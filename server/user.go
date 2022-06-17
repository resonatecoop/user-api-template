package server

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	uuidpkg "github.com/resonatecoop/user-api-template/pkg/uuid"

	"github.com/resonatecoop/user-api-template/model"
	pbUser "github.com/resonatecoop/user-api-template/proto/user"
)

// AddUser adds a user to the DB
func (s *Server) AddUser(ctx context.Context, user *pbUser.UserAddRequest) (*pbUser.UserRequest, error) {
	err := checkRequiredAddAttributes(user)
	if err != nil {
		return nil, err
	}

	err = s.db.NewSelect().Model(&model.User{}).
		Where("username = ?", strings.ToLower(user.Username)).
		Scan(ctx)

	if err == nil {
		return nil, errors.New("Email is taken")
	}

	var thisRole int32

	// if requested role is not admin, grant it
	if user.RoleId != nil && *user.RoleId >= int32(model.LabelRole) {
		thisRole = *user.RoleId
	} else {

		defaultRole := new(model.Role)

		err = s.db.NewSelect().
			Model(defaultRole).
			Where("is_default = TRUE").
			Scan(ctx)

		if err != nil {
			return nil, err
		}

		thisRole = int32(defaultRole.ID)
	}

	// defaults to User Role, must update with greater privileges to change role
	newUser := &model.User{
		Username:               strings.ToLower(user.Username),
		RoleID:                 thisRole,
		FullName:               user.FullName,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Member:                 user.Member,
		Country:                user.Country,
		NewsletterNotification: user.NewsletterNotification,
	}

	newUser.ID = uuid.Must(uuid.NewRandom())

	_, err = s.db.NewInsert().
		Column(
			"id",
			"username",
			"full_name",
			"first_name",
			"last_name",
			"role_id",
			"tenant_id",
			"member",
			"country",
			"newsletter_notification",
		).
		Model(newUser).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	res := &pbUser.UserRequest{Id: newUser.ID.String()}

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUser Gets a user from the DB
func (s *Server) GetUser(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserPublicResponse, error) {

	u := new(model.User)

	err := s.db.NewSelect().Model(u).
		Column("user.*").
		Where("id = ?", user.Id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.UserPublicResponse{
		Id:             u.ID.String(),
		Username:       u.Username,
		FullName:       u.FullName,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Member:         u.Member,
		Country:        u.Country,
		FollowedGroups: uuidpkg.ConvertUUIDToStrArray(u.FollowedGroups),
	}, nil
}

// GetUserRestricted intended for privileged roles only supplies more detailed, private info about user.
func (s *Server) GetUserRestricted(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserPrivateResponse, error) {

	u := new(model.User)

	err := s.db.NewSelect().Model(u).
		Column("user.*").
		Where("id = ?", user.Id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.UserPrivateResponse{
		Id:                     u.ID.String(),
		Username:               u.Username,
		FullName:               u.FullName,
		FirstName:              u.FirstName,
		LastName:               u.LastName,
		Member:                 u.Member,
		RoleId:                 u.RoleID,
		TenantId:               u.TenantID,
		FollowedGroups:         uuidpkg.ConvertUUIDToStrArray(u.FollowedGroups),
		NewsletterNotification: u.NewsletterNotification,
	}, nil
}

// DeleteUser Deletes a user from the DB
func (s *Server) DeleteUser(ctx context.Context, user *pbUser.UserRequest) (*pbUser.Empty, error) {
	u := new(model.User)

	_, err := s.db.NewDelete().
		Model(u).
		Where("id = ?", user.Id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// UpdateUser updates a users basic attributes
func (s *Server) UpdateUser(ctx context.Context, UserUpdateRequest *pbUser.UserUpdateRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if UserUpdateRequest.Username != nil {
		updatedUserValues["username"] = *UserUpdateRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*UserUpdateRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}

	if UserUpdateRequest.RoleId != nil && *UserUpdateRequest.RoleId >= int32(model.LabelRole) {
		updatedUserValues["role_id"] = *UserUpdateRequest.RoleId
	}
	if UserUpdateRequest.FirstName != nil {
		updatedUserValues["first_name"] = *UserUpdateRequest.FirstName
	}
	if UserUpdateRequest.LastName != nil {
		updatedUserValues["last_name"] = *UserUpdateRequest.LastName
	}
	if UserUpdateRequest.FullName != nil {
		updatedUserValues["full_name"] = *UserUpdateRequest.FullName
	}
	if UserUpdateRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *UserUpdateRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", UserUpdateRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// UpdateUserRestricted updates a users more restricted attributes
func (s *Server) UpdateUserRestricted(ctx context.Context, UserUpdateRestrictedRequest *pbUser.UserUpdateRestrictedRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if UserUpdateRestrictedRequest.Username != nil {
		updatedUserValues["username"] = *UserUpdateRestrictedRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*UserUpdateRestrictedRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}
	if UserUpdateRestrictedRequest.FirstName != nil {
		updatedUserValues["first_name"] = *UserUpdateRestrictedRequest.FirstName
	}
	if UserUpdateRestrictedRequest.LastName != nil {
		updatedUserValues["last_name"] = *UserUpdateRestrictedRequest.LastName
	}
	if UserUpdateRestrictedRequest.FullName != nil {
		updatedUserValues["full_name"] = *UserUpdateRestrictedRequest.FullName
	}
	if UserUpdateRestrictedRequest.Member != nil {
		updatedUserValues["member"] = *UserUpdateRestrictedRequest.Member
	}
	if UserUpdateRestrictedRequest.RoleId != nil {
		updatedUserValues["role_id"] = *UserUpdateRestrictedRequest.RoleId
	}
	if UserUpdateRestrictedRequest.TenantId != nil {
		updatedUserValues["tenant_id"] = *UserUpdateRestrictedRequest.TenantId
	}
	if UserUpdateRestrictedRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *UserUpdateRestrictedRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", UserUpdateRestrictedRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// ListUsers lists all users in the store.
func (s *Server) ListUsers(ctx context.Context, Empty *pbUser.Empty) (*pbUser.UserListResponse, error) {

	var users []model.User
	var results pbUser.UserListResponse

	err := s.db.NewSelect().
		Model(&users).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		var result pbUser.UserPrivateResponse
		result.Id = uuid.UUID.String(user.ID)
		result.Username = user.Username
		result.FullName = user.FullName
		result.FirstName = user.FirstName
		result.LastName = user.LastName
		result.Member = user.Member
		result.RoleId = user.RoleID
		result.TenantId = user.TenantID
		result.NewsletterNotification = user.NewsletterNotification
		result.FollowedGroups = uuidpkg.ConvertUUIDToStrArray(user.FollowedGroups)
		results.User = append(results.User, &result)
	}

	return &results, nil
}

func checkRequiredAddAttributes(user *pbUser.UserAddRequest) error {
	if user.Username == "" {
		argument := "username"
		return fmt.Errorf("argument %v is required", argument)
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Username) {
		return errors.New("username must be a valid email")
	}
	return nil
}

func getUserGroupResponse(ownerOfGroup []model.UserGroup) []*pbUser.RelatedUserGroup {
	groups := make([]*pbUser.RelatedUserGroup, len(ownerOfGroup))
	for i, group := range ownerOfGroup {
		groups[i] = &pbUser.RelatedUserGroup{Id: group.ID.String(), DisplayName: group.DisplayName}
	}
	return groups
}

func (s *Server) DerefString(str *string) string {
	if str != nil {
		return *str
	}

	return ""
}
