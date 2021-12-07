package model

import (
	uuid "github.com/google/uuid"
)

// ShareTransaction model
type ShareTransaction struct {
	IDRecord
	UserID    uuid.UUID `bun:"type:uuid,notnull"`
	User      *User     `bun:"rel:has-one"`
	InvoiceID string    `bun:",unique,notnull"`
	Quantity  int64     `bun:",notnull,"`
}
