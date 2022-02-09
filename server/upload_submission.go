package server

import (
	"context"

	"github.com/resonatecoop/user-api/model"
	uuidpkg "github.com/resonatecoop/user-api/pkg/uuid"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// AddUserUploadSubmission
func (s *Server) AddUserUploadSubmission(
	ctx context.Context,
	uploadSubmission *pbUser.UserAddUploadSubmissionRequest,
) (
	*pbUser.UserUploadSubmissionRequest,
	error,
) {
	// TODO find usergroup
	newUploadSubmission := &model.UserUploadSubmission{
		Name:        uploadSubmission.Name,
		Description: uploadSubmission.Description,
		Files:       uuidpkg.ConvertStrToUUIDArray(uploadSubmission.Files),
	}

	_, err := s.db.NewInsert().
		Column(
			"id",
			"name",
			"description",
		).
		Model(newUploadSubmission).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	res := &pbUser.UserUploadSubmissionRequest{
		Id: newUploadSubmission.ID.String(),
	}

	return res, nil
}
