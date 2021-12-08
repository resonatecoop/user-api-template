package server

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	uuid "github.com/google/uuid"

	"github.com/resonatecoop/user-api/model"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// // Server implements the UserService
// type Server struct {
// 	db *bun.DB
// }

// // New creates an instance of our server
// func New(db *bun.DB) *Server {
// 	return &Server{db: db}
// }

// AddUser gets a user to the in-memory store.
func (s *Server) AddUserGroup(ctx context.Context, usergroup *pbUser.UserGroupCreateRequest) (*pbUser.UserRequest, error) {

	requiredErr := s.checkRequiredAddUserGroupAttributes(ctx, usergroup)

	if requiredErr != nil {
		return nil, requiredErr
	}

	OwnerUUID, err := uuid.Parse(usergroup.Id)

	if err != nil {
		return nil, errors.New("supplied user_id is not a valid UUID")
	}

	existingGroupCount, _ := s.db.NewSelect().
		Model((*model.UserGroup)(nil)).
		Where("owner_id = ?", OwnerUUID).
		Count(ctx)

	owningUser := new(model.User)

	err = s.db.NewSelect().
		Model(owningUser).
		Where("id = ?", OwnerUUID).
		Scan(ctx)

	if err != nil {
		return nil, errors.New("supplied owner_id could not be found in Users")
	}

	if owningUser.RoleID == int32(model.UserRole) && existingGroupCount > 0 {
		return nil, errors.New("supplied owner_id is a user and already has a user group profile")
	}

	group := new(model.GroupType)

	err = s.db.NewSelect().
		Model(group).
		Where("name = ?", usergroup.GroupType).
		Scan(ctx)

	if err != nil {
		return nil, errors.New("supplied group type is not valid")
	}

	AvatarUUID, err := uuid.Parse(usergroup.Avatar)

	if usergroup.Avatar != "" && err != nil {
		return nil, errors.New("supplied avatar is not a valid UUID")
	}

	BannerUUID, err := uuid.Parse(usergroup.Banner)

	if usergroup.Banner != "" && err != nil {
		return nil, errors.New("supplied banner is not a valid UUID")
	}

	newUserGroup := &model.UserGroup{
		OwnerID:     OwnerUUID,
		TypeID:      group.ID,
		Type:        group,
		DisplayName: usergroup.DisplayName,
		Description: usergroup.Description,
		ShortBio:    usergroup.ShortBio,
		Avatar:      AvatarUUID,
		Banner:      BannerUUID,
		GroupEmail:  usergroup.GroupEmail,
	}

	newUserGroup.CreatedAt = time.Now().UTC()

	_, err = s.db.NewInsert().Model(newUserGroup).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.UserRequest{Id: newUserGroup.ID.String()}, nil
}

// UpdateUser updates a users basic attributes
func (s *Server) UpdateUserGroup(ctx context.Context, UserGroupUpdateRequest *pbUser.UserGroupUpdateRequest) (*pbUser.Empty, error) {

	var updatedUserGroupValues = make(map[string]interface{})

	if UserGroupUpdateRequest.GroupEmail != nil {
		updatedUserGroupValues["group_email_address"] = *UserGroupUpdateRequest.GroupEmail
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*UserGroupUpdateRequest.GroupEmail) {
			return nil, errors.New("group email address must be a valid email")
		}
	}
	if UserGroupUpdateRequest.DisplayName != nil {
		updatedUserGroupValues["display_name"] = *UserGroupUpdateRequest.DisplayName
	}
	if UserGroupUpdateRequest.Description != nil {
		updatedUserGroupValues["description"] = *UserGroupUpdateRequest.Description
	}
	if UserGroupUpdateRequest.ShortBio != nil {
		updatedUserGroupValues["short_bio"] = *UserGroupUpdateRequest.ShortBio
	}
	if UserGroupUpdateRequest.GroupType != nil {
		group := new(model.GroupType)

		err := s.db.NewSelect().
			Model(group).
			Where("name = ?", UserGroupUpdateRequest.GroupType).
			Scan(ctx)

		if err != nil {
			return nil, errors.New("supplied group type is not valid")
		}

		updatedUserGroupValues["type_id"] = group.ID
	}
	if UserGroupUpdateRequest.OwnerId != nil {
		updatedUserGroupValues["owner_id"] = *UserGroupUpdateRequest.OwnerId
	}
	if UserGroupUpdateRequest.Avatar != nil {
		updatedUserGroupValues["avatar"] = *UserGroupUpdateRequest.Avatar
	}
	if UserGroupUpdateRequest.Banner != nil {
		updatedUserGroupValues["banner"] = *UserGroupUpdateRequest.Banner
	}

	updatedUserGroupValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserGroupValues).TableExpr("user_groups").Where("id = ?", UserGroupUpdateRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// DeleteUser Deletes a user from the DB
func (s *Server) DeleteUserGroup(ctx context.Context, usergroup *pbUser.UserGroupRequest) (*pbUser.Empty, error) {
	u := new(model.UserGroup)

	_, err := s.db.NewDelete().
		Model(u).
		Where("id = ?", usergroup.Id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// GetUserGroup returns details of single user group
func (s *Server) GetUserGroup(ctx context.Context, usergrouprequest *pbUser.UserGroupRequest) (*pbUser.UserGroupPublicResponse, error) {

	usergroup := new(model.UserGroup)

	err := s.db.NewSelect().
		Model(usergroup).
		Where("id = ?", usergrouprequest.Id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	group := new(model.GroupType)

	err = s.db.NewSelect().
		Model(group).
		Where("id = ?", usergroup.TypeID).
		Scan(ctx)

	if err != nil {
		return nil, errors.New("supplied group type is not valid")
	}

	return &pbUser.UserGroupPublicResponse{
		DisplayName: usergroup.DisplayName,
		GroupType:   group.Name,
		ShortBio:    usergroup.ShortBio,
		Description: usergroup.Description,
		Avatar:      uuid.UUID.String(usergroup.Avatar),
		Banner:      uuid.UUID.String(usergroup.Banner),
		GroupEmail:  usergroup.GroupEmail,
	}, nil
}

// ListUsersUserGroups lists all the User Groups owned by the supplied User Id
func (s *Server) ListUsersUserGroups(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserGroupListResponse, error) {

	var usergroups []model.UserGroup
	var results pbUser.UserGroupListResponse

	err := s.db.NewSelect().
		Model(&usergroups).
		Where("owner_id = ?", user.Id).
		Order("created_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for _, usergroup := range usergroups {

		group := new(model.GroupType)

		err = s.db.NewSelect().
			Model(group).
			Where("id = ?", usergroup.TypeID).
			Scan(ctx)

		var result pbUser.UserGroupPrivateResponse
		result.Id = uuid.UUID.String(usergroup.ID)
		result.DisplayName = usergroup.DisplayName
		result.GroupType = group.Name
		result.ShortBio = usergroup.ShortBio
		result.Description = usergroup.Description
		result.Avatar = uuid.UUID.String(usergroup.Avatar)
		result.Banner = uuid.UUID.String(usergroup.Banner)
		result.GroupEmail = usergroup.GroupEmail

		results.Usergroup = append(results.Usergroup, &result)
	}

	return &results, nil
}

func (s *Server) checkRequiredAddUserGroupAttributes(ctx context.Context, usergroup *pbUser.UserGroupCreateRequest) error {
	if usergroup.Id == "" || usergroup.Id == uuid.Nil.String() || usergroup.DisplayName == "" {
		var argument string
		switch {
		case usergroup.Id == "":
			argument = "owner_id"
		case usergroup.Id == uuid.Nil.String():
			argument = "owner_id"
		case usergroup.DisplayName == "":
			argument = "display_name"
		}
		return fmt.Errorf("argument %v is required", argument)
	}

	if usergroup.GroupEmail != "" {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(usergroup.GroupEmail) {
			return errors.New("group email address must be a valid email")
		}
	}

	err := s.db.NewSelect().
		Model(new(model.User)).
		Where("id = ?", usergroup.Id).
		Scan(ctx)

	if err != nil {
		return errors.New("supplied owner_id does not exist")
	}

	err = s.db.NewSelect().
		Model(new(model.GroupType)).
		Where("name = ?", usergroup.GroupType).
		Scan(ctx)

	if err != nil {
		return errors.New("supplied group type does not exist")
	}

	return nil
}
