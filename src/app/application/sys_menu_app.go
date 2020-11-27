package application

import (
	"butterfly-admin/src/app/config/sequence"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/types"
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

type SysMenuApplication struct {
	sequence   *snowflake.Node
	repository *persistence.Repository
}

// 分页查询
func (application *SysMenuApplication) Query(request *types.SysMenuQueryRequest) (int64, []entity.SysMenu, error) {
	total, data, err := application.repository.SysMenuRepository.Select(request)

	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Select() happen error for", err)
	}
	return total, data, err
}

// 创建菜单
func (application *SysMenuApplication) Create(request *types.SysMenuCreateRequest) error {
	menu := request.SysMenu
	menu.Id = sequence.GetSequence().Generate().Int64()
	err := application.repository.SysMenuRepository.Save(&menu)

	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Save() happen error", err)
	}
	return err
}
