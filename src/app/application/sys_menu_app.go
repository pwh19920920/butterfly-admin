package application

import (
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/types"
	"github.com/bwmarrin/snowflake"
)

type SysMenuApplication struct {
	sequence   *snowflake.Node
	repository *persistence.Repository
}

// 分页查询
func (application SysMenuApplication) Query(request types.SysMenuQueryRequest) (int64, []entity.SysMenu, error) {
	return application.repository.SysMenuRepository.Select(request)
}
