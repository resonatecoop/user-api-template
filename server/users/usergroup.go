package userserver

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

// AddUser gets a user to the in-memory store.
func (s *Server) AddUserGroup(ctx context.Context, usergroup *pbUser.UserGroupCreateRequest) (*pbUser.Empty, error) {

	requiredErr := s.checkRequiredAddUserGroupAttributes(ctx, usergroup)

	if requiredErr != nil {
		return nil, requiredErr
	}

	OwnerUUID, err := uuid.Parse(usergroup.OwnerId)

	if err != nil {
		return nil, fmt.Errorf("supplied user_id is not a valid UUID")
	}

	TypeUUID, err := uuid.Parse(usergroup.TypeId)

	if err != nil {
		return nil, fmt.Errorf("supplied type is not a valid UUID")
	}

	AvatarUUID, err := uuid.Parse(usergroup.Avatar)

	if usergroup.Avatar != "" && err != nil {
		return nil, fmt.Errorf("supplied avatar is not a valid UUID")
	}

	BannerUUID, err := uuid.Parse(usergroup.Banner)

	if usergroup.Banner != "" && err != nil {
		return nil, fmt.Errorf("supplied banner is not a valid UUID")
	}

	newusergroup := &model.UserGroup{
		OwnerID:     OwnerUUID,
		TypeID:      TypeUUID,
		DisplayName: usergroup.DisplayName,
		Description: usergroup.Description,
		ShortBio:    usergroup.ShortBio,
		Avatar:      AvatarUUID,
		Banner:      BannerUUID,
		GroupEmail:  usergroup.GroupEmail,
	}

	newusergroup.CreatedAt = time.Now().UTC()

	_, err = s.db.NewInsert().Model(newusergroup).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
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
	if UserGroupUpdateRequest.TypeId != nil {
		updatedUserGroupValues["type_id"] = *UserGroupUpdateRequest.TypeId
	}
	if UserGroupUpdateRequest.OwnerId != nil {
		updatedUserGroupValues["owner_id"] = *UserGroupUpdateRequest.OwnerId
	}
	if UserGroupUpdateRequest.Avatar != nil {
		updatedUserGroupValues["avatar"] = *UserGroupUpdateRequest.Avatar
	}
	if UserGroupUpdateRequest.Avatar != nil {
		updatedUserGroupValues["banner"] = *UserGroupUpdateRequest.Banner
	}

	updatedUserGroupValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserGroupValues).TableExpr("users").Where("id = ?", UserGroupUpdateRequest.Id).Exec(ctx)

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

	_, err := s.db.NewUpdate().
		Model(u).
		Set("deleted_at = ?", time.Now().UTC()).
		Where("id = ?", usergroup.Id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// ListUsers lists all users in the store.
func (s *Server) ListUsersGroups(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserGroupListResponse, error) {

	var usergroups []model.UserGroup
	var results pbUser.UserGroupListResponse

	err := s.db.NewSelect().
		Model(&usergroups).
		Where("owner_id = ?", user.Id).
		Where("deleted_at IS NOT NULL").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for _, usergroup := range usergroups {
		var result pbUser.UserGroupPrivateResponse
		result.Id = uuid.UUID.String(usergroup.ID)
		result.DisplayName = usergroup.DisplayName
		result.TypeId = uuid.UUID.String(usergroup.TypeID)
		result.ShortBio = usergroup.ShortBio
		result.Description = usergroup.Description
		result.Avatar = uuid.UUID.String(usergroup.Avatar)
		result.Banner = uuid.UUID.String(usergroup.Banner)
		result.GroupEmail = usergroup.GroupEmail

		//result.FollowedGroups = uuidpkg.ConvertUUIDToStrArray(user.FollowedGroups)
		results.Usergroup = append(results.Usergroup, &result)
	}

	return &results, nil
}

func (s *Server) checkRequiredAddUserGroupAttributes(ctx context.Context, usergroup *pbUser.UserGroupCreateRequest) error {
	if usergroup.OwnerId == "" || usergroup.OwnerId == uuid.Nil.String() || usergroup.TypeId == uuid.Nil.String() || usergroup.TypeId == "" || usergroup.DisplayName == "" {
		var argument string
		switch {
		case usergroup.OwnerId == "":
			argument = "owner_id"
		case usergroup.OwnerId == uuid.Nil.String():
			argument = "owner_id"
		case usergroup.TypeId == "":
			argument = "type"
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
		Where("id = ?", usergroup.OwnerId).
		Scan(ctx)

	if err != nil {
		return errors.New("supplied owner_id does not exist")
	}

	err = s.db.NewSelect().
		Model(new(model.GroupTaxonomy)).
		Where("id = ?", usergroup.TypeId).
		Scan(ctx)

	if err != nil {
		return errors.New("supplied group type does not exist")
	}

	return nil
}
