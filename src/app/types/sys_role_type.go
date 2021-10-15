package types

import (
	"butterfly-admin/src/app/domain/entity"
	"github.com/pwh19920920/butterfly/response"
)

type SysRoleQueryRequest struct {
	response.RequestPaging
}

type SysRolePermissionQueryResponse struct {
	entity.SysPermission
	Options []string `json:"options"`
}

type SysRoleCreateRequest struct {
	entity.SysRole
	Permissions []entity.SysPermission `json:"permissions"`
}
