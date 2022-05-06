package server

import (
	"context"
	"fmt"

	"github.com/resonatecoop/user-api/model"

	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// func (s *Server) GetTracks(ctx context.Context, req *pbUser.TracksList) (*pbUser.TracksList, error) {
// 	trackIds := make([]uuid.UUID, len(req.Tracks))
// 	for i, track := range req.Tracks {
// 		id, err := uuidpkg.GetUuidFromString(track.Id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		trackIds[i] = id
// 	}
// 	tracksResponse, twerr := model.GetTracks(trackIds, s.db, true, ctx)
// 	if twerr != nil {
// 		return nil, twerr
// 	}
// 	return &pb.TracksList{
// 		Tracks: tracksResponse,
// 	}, nil
// }
//
// func (s *Server) SearchTracks(ctx context.Context, q *pbUser.Query) (*pbUser.SearchResults, error) {
// 	if len(q.Query) < 3 {
// 		return nil, errors.New("query must be a valid search query")
// 	}
//
// 	searchResults, twerr := model.SearchTracks(q.Query, s.db)
// 	if twerr != nil {
// 		return nil, twerr
// 	}
// 	return searchResults, nil
// }

// CreateTrack
func (s *Server) CreateTrack(ctx context.Context, track *pbUser.Track) (*pbUser.Track, error) {
	// Track is created then added to a TrackGroup on track group creation
	err := checkRequiredAttributes(track)
	if err != nil {
		return nil, err
	}

	t := &model.Track{
		Title:       track.Title,
		Status:      track.Status,
		Enabled:     track.Enabled,
		TrackNumber: track.TrackNumber,
		Duration:    track.Duration,
	}

	_, err = s.db.NewInsert().
		Column(
			"id",
			"title",
			"status",
			"duration",
			"track_number",
		).
		Model(t).
		Exec(ctx)

	res := &pbUser.Track{Id: t.ID.String()}

	return res, nil
}

// GetTrack
func (s *Server) GetTrack(ctx context.Context, track *pbUser.Track) (*pbUser.Track, error) {
	// t, err := getTrackModel(track)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// pgerr := s.db.Model(t).
	// 		Column("track.*").
	// 		WherePK().
	// 		Select()
	// if pgerr != nil {
	// 	return nil, errorpkg.CheckError(pgerr, "track")
	// }
	// track.UserGroupId = t.UserGroupId.String()
	// track.CreatorId = t.CreatorId.String()
	// track.TrackServerId = t.TrackServerId.String()
	// track.Title = t.Title
	// track.Status = t.Status
	// track.Enabled = t.Enabled
	// track.TrackNumber = t.TrackNumber
	// track.Duration = t.Duration
	//
	// // Get tags
	// tags, twerr := model.GetTags(t.Tags, s.db)
	// if twerr != nil {
	// 	return nil, twerr
	// }
	// track.Tags = tags
	//
	// // Get artists (id, name, avatar)
	// artists, pgerr := model.GetRelatedUserGroups(t.Artists, s.db)
	// if pgerr != nil {
	// 	return nil, errorpkg.CheckError(pgerr, "user_group")
	// }
	// track.Artists = artists
	//
	// // Get track_groups (id, title, cover) that are not playlists (i.e. LP, EP or Single)
	// trackGroups, twerr := model.GetTrackGroupsFromIds(t.TrackGroups, s.db, []string{"lp", "ep", "single"})
	// if twerr != nil {
	// 	return nil, twerr
	// }
	// track.TrackGroups = trackGroups

	return &pbUser.Track{}, nil
}

// UpdateTrack
func (s *Server) UpdateTrack(ctx context.Context, track *pbUser.Track) (*pbUser.Empty, error) {
	err := checkRequiredAttributes(track)
	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// DeleteTrack
func (s *Server) DeleteTrack(ctx context.Context, track *pbUser.Track) (*pbUser.Empty, error) {
	return &pbUser.Empty{}, nil
}

func checkRequiredAttributes(track *pbUser.Track) error {
	if track.Title == "" || track.Status == "" || track.TrackNumber == 0 || track.CreatorId == "" || track.UserGroupId == "" { // track.Artists?
		var argument string
		switch {
		case track.Title == "":
			argument = "title"
		case track.Status == "":
			argument = "status"
		case track.CreatorId == "":
			argument = "creator_id"
		case track.UserGroupId == "":
			argument = "user_group_id"
		case track.TrackNumber == 0:
			argument = "track_number"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	return nil
}
