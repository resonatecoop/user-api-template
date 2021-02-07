package model

import (
	pb "github.com/merefield/grpc-user-api/proto"

	"github.com/go-pg/pg"
	"github.com/twitchtv/twirp"

	uuidpkg "github.com/merefield/grpc-user-api/pkg/uuid"

	uuid "github.com/satori/go.uuid"

	errorpkg "github.com/merefield/grpc-user-api/pkg/error"
)

// Tag provides basic tag structure
type Tag struct {
	ID   uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
	Type string    `sql:",notnull"`
	Name string    `sql:",notnull"`
}

// SearchTags find a tag by query string on Name
func SearchTags(query string, tagType string, db *pg.DB) ([]*Tag, twirp.Error) {
	var tags []*Tag

	pgerr := db.Model(&tags).
		ColumnExpr("tag.id").
		Where("to_tsvector('english'::regconfig, COALESCE(name, '')) @@ (plainto_tsquery('english'::regconfig, ?)) = true", query).
		Where("type = ?", tagType).
		Select()
	if pgerr != nil {
		return nil, errorpkg.CheckError(pgerr, "tag")
	}
	return tags, nil
}

// GetTags given tag UUID returns a Tag
func GetTags(tagIds []uuid.UUID, db *pg.DB) ([]*pb.Tag, twirp.Error) {
	tags := make([]*pb.Tag, len(tagIds))
	if len(tags) > 0 {
		var t []Tag
		pgerr := db.Model(&t).
			Where("id in (?)", pg.In(tagIds)).
			Select()
		if pgerr != nil {
			return nil, errorpkg.CheckError(pgerr, "tag")
		}
		for i, tag := range t {
			tags[i] = &pb.Tag{Id: tag.ID.String(), Type: tag.Type, Name: tag.Name}
		}
	}
	return tags, nil
}

// GetTagIDs accepts a slice tags
func GetTagIDs(t []*pb.Tag, db *pg.Tx) ([]uuid.UUID, error) {
	tags := make([]*Tag, len(t))
	tagIDs := make([]uuid.UUID, len(t))
	for i, tag := range t {
		if tag.Id == "" { // new tag to create and add
			tags[i] = &Tag{Type: tag.Type, Name: tag.Name}
			_, pgerr := db.Model(tags[i]).
				Where("type = ?", tags[i].Type).
				Where("lower(name) = lower(?)", tags[i].Name).
				Returning("*").
				SelectOrInsert()
			if pgerr != nil {
				return nil, pgerr
			}
			tagIDs[i] = tags[i].ID
			tag.Id = tags[i].ID.String()
		} else {
			tagID, twerr := uuidpkg.GetUUIDFromString(tag.Id)
			if twerr != nil {
				return nil, twerr.(error)
			}
			tagIDs[i] = tagID
		}
	}
	return tagIDs, nil
}
