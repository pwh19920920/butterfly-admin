package types

import (
	"github.com/pwh19920920/butterfly-admin/common"
	"github.com/pwh19920920/butterfly/response"
)

type SysUserQueryRequest struct {
	response.RequestPaging
	Name     string `json:"name"`     // 名称
	Username string `json:"username"` // 用户
}

type SysUserQueryResponse struct {
	common.BaseEntity
	Name     string   `json:"name"`     // 名称
	Avatar   string   `json:"avatar"`   // 头像
	Username string   `json:"username"` // 用户
	Roles    string   `json:"roles"`    // 角色串
	RoleList []string `json:"roleList"` // 角色列表
}
