package model

import (
	"github.com/go-pg/pg"

	//uuidpkg "github.com/resonatecoop/user-api/pkg/uuid"

	uuid "github.com/google/uuid"

	pbUser "github.com/resonatecoop/user-api-template/proto/user"
)

// Tag provides basic tag structure
type Tag struct {
	ID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Type string    `bun:",notnull"`
	Name string    `bun:",notnull"`
}

// SearchTags find a tag by query string on Name
func SearchTags(query string, tagType string, db *pg.DB) ([]*Tag, error) {
	var tags []*Tag

	err := db.Model(&tags).
		ColumnExpr("tag.id").
		Where("to_tsvector('english'::regconfig, COALESCE(name, '')) @@ (plainto_tsquery('english'::regconfig, ?)) = true", query).
		Where("type = ?", tagType).
		Select()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// GetTags given tag UUID returns a Tag
func GetTags(tagIds []uuid.UUID, db *pg.DB) ([]*pbUser.Tag, error) {
	tags := make([]*pbUser.Tag, len(tagIds))
	if len(tags) > 0 {
		var t []Tag
		err := db.Model(&t).
			Where("id in (?)", pg.In(tagIds)).
			Select()
		if err != nil {
			return nil, err
		}
		for i, tag := range t {
			tags[i] = &pbUser.Tag{Id: tag.ID.String(), Type: tag.Type, Name: tag.Name}
		}
	}
	return tags, nil
}

// GetTagIDs accepts a slice tags
func GetTagIDs(t []*pbUser.Tag, db *pg.Tx) ([]uuid.UUID, error) {
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
			tagID, err := uuid.Parse(tag.Id)
			if err != nil {
				return nil, err
			}
			tagIDs[i] = tagID
		}
	}
	return tagIDs, nil
}
