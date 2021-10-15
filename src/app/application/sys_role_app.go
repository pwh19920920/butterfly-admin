package application

import (
	"butterfly-admin/src/app/config/sequence"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/types"
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"strings"
)

type SysRoleApplication struct {
	sequence   *snowflake.Node
	repository *persistence.Repository
}

// Query 分页查询
func (application *SysRoleApplication) Query(request *types.SysRoleQueryRequest) (int64, []entity.SysRole, error) {
	total, data, err := application.repository.SysRoleRepository.Select(request)
	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Select() happen error for", err)
		return total, nil, err
	}
	return total, data, err
}

func (application *SysRoleApplication) SelectAll() ([]entity.SysRole, error) {
	data, err := application.repository.SysRoleRepository.SelectAll()
	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Select() happen error for", err)
	}
	return data, err
}

// QueryPermissionByRoleId 查询
func (application *SysRoleApplication) QueryPermissionByRoleId(roleId int64) ([]types.SysRolePermissionQueryResponse, error) {
	data, err := application.repository.SysPermissionRepository.SelectByRoleId(roleId)
	// 错误记录
	if err != nil {
		logrus.Error("SysPermissionRepository.SelectByRoleId() happen error for", err)
		return nil, err
	}

	result := make([]types.SysRolePermissionQueryResponse, 0)
	for _, item := range data {
		options := make([]string, 0)
		if item.Option != "" {
			options = strings.Split(item.Option, ",")
		}
		result = append(result, types.SysRolePermissionQueryResponse{SysPermission: item, Options: options})
	}
	return result, nil
}

// Create 创建
func (application *SysRoleApplication) Create(request *types.SysRoleCreateRequest) error {
	role := request.SysRole
	role.Id = sequence.GetSequence().Generate().Int64()

	if request.Permissions != nil {
		for index, permission := range request.Permissions {
			permission.Id = sequence.GetSequence().Generate().Int64()
			permission.RoleId = role.Id
			request.Permissions[index] = permission
		}
	}
	return application.repository.SysRoleRepository.Save(&request.Permissions, &role)
}

// Modify 创建
func (application *SysRoleApplication) Modify(request *types.SysRoleCreateRequest) error {
	role := request.SysRole
	if request.Permissions != nil {
		for index, permission := range request.Permissions {
			permission.Id = sequence.GetSequence().Generate().Int64()
			request.Permissions[index] = permission
		}
	}
	return application.repository.SysRoleRepository.UpdateById(request.Id, &request.Permissions, &role)
}

// Delete 更新
func (application *SysRoleApplication) Delete(request int64) error {
	return application.repository.SysRoleRepository.Delete(request)
}
