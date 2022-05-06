package server

import (
	"context"
	"errors"
	"time"

	uuid "github.com/google/uuid"
	"github.com/resonatecoop/user-api/model"
	uuidpkg "github.com/resonatecoop/user-api/pkg/uuid"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// AddUserUploadSubmission
func (s *Server) AddUploadSubmission(
	ctx context.Context,
	uploadSubmission *pbUser.UploadSubmissionAddRequest,
) (
	*pbUser.UploadSubmissionRequest,
	error,
) {
	UserUUID, err := uuid.Parse(uploadSubmission.Id)

	if err != nil {
		return nil, errors.New("supplied user_id is not a valid UUID")
	}

	newUploadSubmission := &model.UploadSubmission{
		Name:        uploadSubmission.Name,
		Description: uploadSubmission.Description,
		Files:       uuidpkg.ConvertStrToUUIDArray(uploadSubmission.Files),
		UserID:      UserUUID,
	}

	newUploadSubmission.ID = uuid.Must(uuid.NewRandom())
	newUploadSubmission.CreatedAt = time.Now().UTC()

	_, err = s.db.NewInsert().
		Column(
			"id",
			"user_id",
			"name",
			"description",
			"files",
		).
		Model(newUploadSubmission).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	res := &pbUser.UploadSubmissionRequest{
		Id: newUploadSubmission.ID.String(),
	}

	return res, nil
}

// UpdateUserUploadSubmission
func (s *Server) UpdateUploadSubmission(
	ctx context.Context,
	UploadSubmissionUpdateRequest *pbUser.UploadSubmissionUpdateRequest,
) (
	*pbUser.Empty,
	error,
) {
	uploadSubmission := new(model.UploadSubmission)

	err := s.db.NewSelect().
		Model(uploadSubmission).
		Where("id = ?", UploadSubmissionUpdateRequest.Id).
		Limit(1).
		Scan(ctx)

	// Not found
	if err != nil {
		return nil, errors.New("upload submission not found")
	}

	if UploadSubmissionUpdateRequest.Files != nil {
		uploadSubmission.Files = uuidpkg.ConvertStrToUUIDArray(UploadSubmissionUpdateRequest.Files)
	}

	rows, err := s.db.NewUpdate().
		Column(
			"id",
			"name",
			"description",
			"files",
		).
		Model(uploadSubmission).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// DeleteUploadSubmission Deletes an upload submission
func (s *Server) DeleteUploadSubmission(ctx context.Context, uploadSubmission *pbUser.UploadSubmissionRequest) (*pbUser.Empty, error) {
	u := new(model.UploadSubmission)

	_, err := s.db.NewDelete().
		Model(u).
		Where("id = ?", uploadSubmission.Id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}
