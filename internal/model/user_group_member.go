package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserGroupMember struct {
	CreatedAt   time.Time `sql:"default:now()"`
	UpdatedAt   time.Time
	UserGroupID uuid.UUID `sql:",pk,type:uuid,notnull"`
	MemberID    uuid.UUID `sql:",pk,type:uuid,notnull"`
	DisplayName string
	Tags        []uuid.UUID `sql:",type:uuid[]" pg:",array"`
}
