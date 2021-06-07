package model

import (
	"context"

	uuid "github.com/google/uuid"
)

// AccessRole represents access role type
type AccessRole int32

const (
	// SuperAdminRole has all permissions and can assign admins
	SuperAdminRole AccessRole = iota + 1 // 1

	// AdminRole has admin permissions across all tenants, except the ability to assign other Admins
	AdminRole // 2

	// TenantAdminRole has admin permissions over other users in their tenant.
	TenantAdminRole // 3

	// LabelRole is a like an artist user, but can manage their artists content.
	LabelRole // 4

	// ArtistRole is a like a standard user, but can have multiple personas and the ability to upload
	ArtistRole // 5

	// UserRole is a standard user
	UserRole // 6
)

// RBACService represents role-based access control service interface
type RBACService interface {
	EnforceRole(context.Context, AccessRole) bool
	EnforceUser(context.Context, uuid.UUID) bool
	EnforceTenant(context.Context, uuid.UUID) bool
	EnforceTenantAdmin(context.Context, int32) bool
	EnforceTenantAndRole(context.Context, AccessRole, int32) bool
	IsLowerRole(context.Context, AccessRole) bool
}

// Role entity
// type Role struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// Role is a one of roles user can have
type Role struct {
	ID          int32  `bun:"type:,unique"`
	Name        string `bun:"type:varchar(50),unique,notnull"`
	Description string `bun:"type:varchar(200),notnull"`
	IsDefault   bool   `bun:"default:false"`
}
