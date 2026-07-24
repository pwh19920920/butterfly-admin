package application

import (
	"github.com/pwh19920920/butterfly-admin/internal/domain/entity"
	"github.com/sirupsen/logrus"
)

type SysPermissionApplication struct {
	baseApp
}

// Query 分页查询
func (application *SysPermissionApplication) Query(roleId int64) ([]entity.SysPermission, error) {
	data, err := application.repository.SysPermissionRepository.SelectByRoleId(roleId)

	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Select() happen error for", err)
	}
	return data, err
}
