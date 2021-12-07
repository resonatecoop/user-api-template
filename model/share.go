package model

import (
	uuid "github.com/google/uuid"
)

// Credit
type Share struct {
	IDRecord
	UserID    uuid.UUID `bun:"type:uuid,notnull"`
	InvoiceID string    `bun:",unique,notnull"`
	Quantity  int64     `bun:",notnull,"`
}
