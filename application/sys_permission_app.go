package application

import (
	"github.com/pwh19920920/butterfly-admin/domain/entity"
	"github.com/pwh19920920/butterfly-admin/infrastructure/persistence"
	"github.com/pwh19920920/snowflake"
	"github.com/sirupsen/logrus"
)

type SysPermissionApplication struct {
	sequence   *snowflake.Node
	repository *persistence.Repository
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
