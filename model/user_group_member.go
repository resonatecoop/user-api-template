package model

// import (
// 	"time"

// 	uuid "github.com/google/uuid"
// )

// type UserGroupMember struct {
// 	CreatedAt   time.Time `bun:"default:now()"`
// 	UpdatedAt   time.Time
// 	UserGroupID uuid.UUID `bun:",pk,type:uuid,notnull"`
// 	MemberID    uuid.UUID `bun:",pk,type:uuid,notnull"`
// 	DisplayName string
// 	Tags        []uuid.UUID `bun:",type:uuid[]" pg:",array"`
// }
