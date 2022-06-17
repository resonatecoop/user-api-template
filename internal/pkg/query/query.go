package query

import (
	"fmt"

	"github.com/resonatecoop/user-api-template/model"
)

// ForTenant returns query for filtering rows by tenant_id
func ForTenant(u *model.AuthUser, tenantId int32) string {
	switch u.Role {
	case model.SuperAdminRole, model.AdminRole:
		if tenantId != 0 {
			return fmt.Sprintf("tenant_id = %v", tenantId)
		}
		return ""
	default:
		return fmt.Sprintf("tenant_id = %v", u.TenantID)

	}
}
