package model

import uuid "github.com/satori/go.uuid"

// Tenant table
type Tenant struct {
	Id     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name   string    `json:"name"`
	Active bool      `json:"active"`
}
