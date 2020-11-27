package types

import (
	"butterfly-admin/src/app/domain/entity"
	"github.com/pwh19920920/butterfly/response"
)

type SysMenuQueryRequest struct {
	response.RequestPaging
}

type SysMenuCreateRequest struct {
	entity.SysMenu
}
