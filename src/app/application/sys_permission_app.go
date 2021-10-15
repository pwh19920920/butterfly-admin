package application

import (
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/infrastructure/persistence"
	"github.com/bwmarrin/snowflake"
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