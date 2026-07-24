package application

import (
	"context"

	"github.com/pwh19920920/butterfly-admin/internal/domain/entity"
	"github.com/pwh19920920/butterfly/pkg/logger"
)

type SysPermissionApplication struct {
	baseApp
}

// Query 分页查询
func (application *SysPermissionApplication) Query(ctx context.Context, roleId int64) ([]entity.SysPermission, error) {
	data, err := application.repository.SysPermissionRepository.SelectByRoleId(roleId)

	// 错误记录
	if err != nil {
		logger.Error(ctx, "SysMenuRepository.Select() happen error for", err)
	}
	return data, err
}
