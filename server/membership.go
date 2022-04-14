package server

import (
	"context"

	"github.com/resonatecoop/user-api/model"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// GetUserMembership
func (s *Server) GetUserMembership(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserMembershipResponse, error) {
	var memberships []model.UserMembership
	var result pbUser.UserMembershipResponse

	// it should not be possible to have two active memberships at the same time
	err := s.db.NewSelect().
		Model(&memberships).
		Where("user_id = ?", user.Id).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	if len(memberships) == 0 {
		return &result, nil
	}

	membership := memberships[0]

	result.Id = membership.ID.String()
	result.Start = membership.Start.UTC().String()
	result.End = membership.End.UTC().String()
	result.MembershipClass = membership.MembershipClass.Name

	return &result, nil
}

// TODO CancelUserMembership
