package types

import (
	"butterfly-admin/src/app/domain/entity"
	"github.com/pwh19920920/butterfly/response"
)

type SysMenuQueryRequest struct {
	response.RequestPaging
}

type SysMenuTreeResponse struct {
	entity.SysMenu
	Options  []entity.SysMenuOption `json:"options"`
	Children []SysMenuTreeResponse  `json:"children"`
}

type SysMenuCreateRequest struct {
	entity.SysMenu
	Options []entity.SysMenuOption `json:"options"`
}

type SysMenuUpdateRequest struct {
	entity.SysMenu
}
