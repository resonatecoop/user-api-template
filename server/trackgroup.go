package server

import (
	"context"
	"errors"
	"fmt"

	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// GetTrackGroup
func (s *Server) GetTrackGroup(ctx context.Context, trackgroup *pbUser.TrackGroupRequest) (*pbUser.TrackGroupResponse, error) {
	return &pbUser.TrackGroupResponse{}, nil
}

// CreateTrackGroup
func (s *Server) CreateTrackGroup(ctx context.Context, trackgroup *pbUser.TrackGroupCreateRequest) (*pbUser.TrackGroupResponse, error) {
	return &pbUser.TrackGroupResponse{}, nil
}

// UpdateTrackGroup
func (s *Server) UpdateTrackGroup(ctx context.Context, trackgroup *pbUser.TrackGroupUpdateRequest) (*pbUser.Empty, error) {
	return &pbUser.Empty{}, nil
}

// DeleteTrackGroup
func (s *Server) DeleteTrackGroup(ctx context.Context, trackGroup *pbUser.TrackGroupRequest) (*pbUser.Empty, error) {
	return &pbUser.Empty{}, nil
}

// AddTracksToTrackGroup
func (s *Server) AddTracksToTrackGroup(ctx context.Context, tracksToTrackGroup *pbUser.TracksToTrackGroup) (*pbUser.Empty, error) {
	return &pbUser.Empty{}, nil
}

// RemoveTracksFromTrackGroup
func (s *Server) RemoveTracksFromTrackGroup(ctx context.Context, tracksToTrackGroup *pbUser.TracksToTrackGroup) (*pbUser.Empty, error) {
	return &pbUser.Empty{}, nil
}

func checkTrackGroupRequiredAttributes(trackGroup *pbUser.TrackGroupCreateRequest) error {
	if trackGroup.Title == "" || (trackGroup.ReleaseDate == nil) || trackGroup.Type == "" || len(trackGroup.Cover) == 0 || trackGroup.CreatorId == "" {
		var argument string
		switch {
		case trackGroup.Title == "":
			argument = "title"
		case trackGroup.ReleaseDate == nil:
			argument = "release_date"
		case trackGroup.Type == "":
			argument = "type"
		case len(trackGroup.Cover) == 0:
			argument = "cover"
		case trackGroup.CreatorId == "":
			argument = "creator_id"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	// A playlist does not have necessarily a owner user group (with id UserGroupId)
	// if it is a private user playlist
	// But other types of track groups (lp, ep, single) have to belong to a user group
	if trackGroup.Type != "playlist" && trackGroup.UserGroupId == "" {
		return errors.New("user_group_id is required")
	}
	return nil
}
