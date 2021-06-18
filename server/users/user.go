package userserver

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	uuid "github.com/google/uuid"
	uuidpkg "github.com/resonatecoop/user-api/pkg/uuid"

	"github.com/uptrace/bun"

	"github.com/resonatecoop/user-api/model"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// Server implements the UserService
type Server struct {
	db  *bun.DB
	sec Securer
}

// New creates an instance of our server
func New(db *bun.DB, sec Securer) *Server {
	return &Server{db: db, sec: sec}
}

// Securer represents password securing service
type Securer interface {
	Hash(string) string
	Password(string, ...string) bool
}

// AddUser gets a user to the in-memory store.
func (s *Server) AddUser(ctx context.Context, user *pbUser.AddUserRequest) (*pbUser.Empty, error) {

	requiredErr := checkRequiredAddAttributes(user)
	if requiredErr != nil {
		return nil, requiredErr
	}

	newUser := &model.User{
		Username:               user.Username,
		FullName:               user.FullName,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Member:                 user.Member,
		NewsletterNotification: user.NewsletterNotification,
	}
	_, err := s.db.NewInsert().Model(newUser).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
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
		FullName:       u.FullName,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Member:         u.Member,
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
func (s *Server) UpdateUser(ctx context.Context, updateUserRequest *pbUser.UpdateUserRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if updateUserRequest.Username != nil {
		updatedUserValues["username"] = *updateUserRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*updateUserRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}
	if updateUserRequest.FirstName != nil {
		updatedUserValues["first_name"] = *updateUserRequest.FirstName
	}
	if updateUserRequest.LastName != nil {
		updatedUserValues["last_name"] = *updateUserRequest.LastName
	}
	if updateUserRequest.FullName != nil {
		updatedUserValues["full_name"] = *updateUserRequest.FullName
	}
	if updateUserRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *updateUserRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", updateUserRequest.Id).Exec(ctx)

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
func (s *Server) UpdateUserRestricted(ctx context.Context, updateUserRestrictedRequest *pbUser.UpdateUserRestrictedRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if updateUserRestrictedRequest.Username != nil {
		updatedUserValues["username"] = *updateUserRestrictedRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*updateUserRestrictedRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}
	if updateUserRestrictedRequest.FirstName != nil {
		updatedUserValues["first_name"] = *updateUserRestrictedRequest.FirstName
	}
	if updateUserRestrictedRequest.LastName != nil {
		updatedUserValues["last_name"] = *updateUserRestrictedRequest.LastName
	}
	if updateUserRestrictedRequest.FullName != nil {
		updatedUserValues["full_name"] = *updateUserRestrictedRequest.FullName
	}
	if updateUserRestrictedRequest.Member != nil {
		updatedUserValues["member"] = *updateUserRestrictedRequest.Member
	}
	if updateUserRestrictedRequest.RoleId != nil {
		updatedUserValues["role_id"] = *updateUserRestrictedRequest.RoleId
	}
	if updateUserRestrictedRequest.TenantId != nil {
		updatedUserValues["tenant_id"] = *updateUserRestrictedRequest.TenantId
	}
	if updateUserRestrictedRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *updateUserRestrictedRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", updateUserRestrictedRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// ResetUserPassword reset's a user's password
func (s *Server) ResetUserPassword(ctx context.Context, ResetUserPasswordRequest *pbUser.ResetUserPasswordRequest) (*pbUser.Empty, error) {
	err := checkRequiredResetPasswordAttributes(ResetUserPasswordRequest, s)

	if err != nil {
		return nil, err
	}

	hashedPassword := s.sec.Hash(ResetUserPasswordRequest.Password)

	u := new(model.User)

	_, err = s.db.NewUpdate().
		Model(u).
		Set("updated_at = ?", time.Now().UTC()).
		Set("password = ?", hashedPassword).
		Where("username = ?", ResetUserPasswordRequest.Username).
		Exec(ctx)

	if err != nil {
		return nil, err
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
		result.NewsletterNotification = user.NewsletterNotification
		result.FollowedGroups = uuidpkg.ConvertUUIDToStrArray(user.FollowedGroups)
		results.User = append(results.User, &result)
	}

	return &results, nil
}

func checkRequiredAddAttributes(user *pbUser.AddUserRequest) error {
	if user.Username == "" || user.FullName == "" {
		var argument string
		switch {
		case user.Username == "":
			argument = "username"
		case user.FullName == "":
			argument = "full_name"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Username) {
		return errors.New("username must be a valid email")
	}
	return nil
}

// func checkRequiredUpdateAttributes(user *pbUser.UpdateUserRequest) error {
// 	if user.Username == "" || user.FullName == "" {
// 		var argument string
// 		switch {
// 		case user.Username == "":
// 			argument = "username"
// 		case user.FullName == "":
// 			argument = "full_name"
// 		}
// 		return fmt.Errorf("argument %v is required", argument)
// 	}
// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	if !re.MatchString(user.Username) {
// 		return errors.New("username must be a valid email")
// 	}
// 	return nil
// }

// func checkRequiredRestrictedUpdateAttributes(user *pbUser.UpdateUserRestrictedRequest) error {
// 	if user.Username == "" || user.FullName == "" {
// 		var argument string
// 		switch {
// 		case user.Username == "":
// 			argument = "username"
// 		case user.FullName == "":
// 			argument = "full_name"
// 		}
// 		return fmt.Errorf("argument %v is required", argument)
// 	}
// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	if !re.MatchString(user.Username) {
// 		return errors.New("username must be a valid email")
// 	}
// 	return nil
// }

func checkRequiredResetPasswordAttributes(user *pbUser.ResetUserPasswordRequest, s *Server) error {
	if user.Username == "" || user.Password == "" {
		var argument string
		switch {
		case user.Username == "":
			argument = "email"
		case user.Password == "":
			argument = "Password"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Username) {
		return errors.New("username must be a valid email")
	}
	if !s.sec.Password(user.Password) {
		return errors.New("password is not strong enough")
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
