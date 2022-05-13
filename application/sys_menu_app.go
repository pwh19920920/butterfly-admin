package application

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/pwh19920920/butterfly-admin/common"
	"github.com/pwh19920920/butterfly-admin/config/sequence"
	"github.com/pwh19920920/butterfly-admin/domain/entity"
	"github.com/pwh19920920/butterfly-admin/infrastructure/persistence"
	"github.com/pwh19920920/butterfly-admin/types"
	"github.com/pwh19920920/snowflake"
	"github.com/sirupsen/logrus"
)

type SysMenuApplication struct {
	sequence   *snowflake.Node
	repository *persistence.Repository
}

// Query 分页查询
func (application *SysMenuApplication) Query(request *types.SysMenuQueryRequest) (int64, []entity.SysMenu, error) {
	total, data, err := application.repository.SysMenuRepository.Select(request)

	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Select() happen error for", err)
	}
	return total, data, err
}

func (application *SysMenuApplication) QueryForTree(withOption bool) ([]types.SysMenuTreeResponse, error) {
	allMenus, err := application.repository.SysMenuRepository.SelectAll()
	if err != nil {
		logrus.Error("SysMenuRepository.SelectAll() happen error for", err)
		return nil, err
	}

	// map<menuId, []option>
	menusOptionMap := make(map[int64][]entity.SysMenuOption, 0)
	if withOption {
		menusOptions, err := application.repository.SysMenuOptionRepository.SelectAll()
		if err != nil {
			return nil, err
		}

		for _, option := range menusOptions {
			options, ok := menusOptionMap[option.MenuId]
			if !ok {
				options = make([]entity.SysMenuOption, 0)
			}
			options = append(options, option)
			menusOptionMap[option.MenuId] = options
		}
	}

	rootMenus := make([]types.SysMenuTreeResponse, 0)
	var menuMap = make(map[int64][]types.SysMenuTreeResponse, 0)
	for _, item := range allMenus {
		// 放数据到menuMap
		menu, ok := menuMap[*item.Parent]
		if !ok {
			menu = make([]types.SysMenuTreeResponse, 0)
		}

		menu = append(menu, types.SysMenuTreeResponse{SysMenu: item})
		menuMap[*item.Parent] = menu

		// 得到rootMenus
		if *item.Parent == 0 {
			rootMenus = append(rootMenus, types.SysMenuTreeResponse{SysMenu: item})
		}
	}

	application.recursionAssignment(withOption, menusOptionMap, rootMenus, menuMap)
	return rootMenus, nil
}

func (application *SysMenuApplication) recursionAssignment(withOption bool, menusOptionMap map[int64][]entity.SysMenuOption,
	rootMenus []types.SysMenuTreeResponse, menuMap map[int64][]types.SysMenuTreeResponse) {

	for index, item := range rootMenus {
		menus, _ := menuMap[item.Id]

		// 如果需要带option
		if withOption {
			options, ok := menusOptionMap[item.Id]
			if ok {
				item.Options = options
			}
		}

		item.Children = menus
		rootMenus[index] = item

		// 判断退出条件
		if menus != nil && len(menus) != 0 {
			// 继续赋值
			application.recursionAssignment(withOption, menusOptionMap, menus, menuMap)
		}
	}
}

// Create 创建菜单
func (application *SysMenuApplication) Create(request *types.SysMenuCreateRequest) error {
	menu := request.SysMenu
	menu.Id = sequence.GetSequence().Generate().Int64()

	route, err := application.getRoutePath(menu.Id, *menu.Parent)
	if err != nil {
		return err
	}

	// 对requestOption重新赋值, 程序是先删除, 后insert into on duplicate update
	if request.Options != nil {
		for index, option := range request.Options {
			option.MenuId = menu.Id
			codeKey := fmt.Sprintf("%v-%v-%v-%v", option.MenuId, option.Value, option.Method, option.Path)
			option.Id = sequence.GetSequence().Generate().Int64()
			option.Code = fmt.Sprintf("%x", md5.Sum([]byte(codeKey)))
			option.Deleted = common.DeletedFalse
			request.Options[index] = option
		}
	}

	menu.Route = route

	err = application.repository.SysMenuRepository.Save(&menu, &request.Options)

	// 错误记录
	if err != nil {
		logrus.Error("SysMenuRepository.Save() happen error", err)
	}
	return err
}

func (application *SysMenuApplication) getRoutePath(currentId, parentId int64) (string, error) {
	// 顶级root, 默认值为/
	if parentId == 0 {
		return fmt.Sprintf("/%v", currentId), nil
	}

	parent, err := application.repository.SysMenuRepository.GetById(parentId)
	if err != nil || parent == nil {
		return "", errors.New("父级菜单不存在")
	}

	return fmt.Sprintf("%v/%v", parent.Route, currentId), nil
}

// Modify 更新
func (application *SysMenuApplication) Modify(request *types.SysMenuCreateRequest) error {
	// 判断老菜单是否存在
	oldMenu, err := application.repository.SysMenuRepository.GetById(request.Id)
	if err != nil || oldMenu == nil {
		return errors.New("菜单不存在")
	}

	// 对requestOption重新赋值, 程序是先删除, 后insert into on duplicate update
	if request.Options != nil {
		for index, option := range request.Options {
			option.MenuId = request.Id
			codeKey := fmt.Sprintf("%v-%v-%v-%v", option.MenuId, option.Value, option.Method, option.Path)
			option.Id = sequence.GetSequence().Generate().Int64()
			option.Code = fmt.Sprintf("%x", md5.Sum([]byte(codeKey)))
			option.Deleted = common.DeletedFalse
			request.Options[index] = option
		}
	}

	newRoute, err := application.getRoutePath(request.Id, *request.Parent)
	if err != nil {
		return err
	}

	request.Route = newRoute

	// 变更菜单路径, 导致自己以及所有子得route字段都会出错, 因此需要批量替换
	if oldMenu.Parent != request.Parent {
		oldRoute := oldMenu.Route
		return application.repository.SysMenuRepository.UpdateEntityAndChildRouteById(request.Id, oldRoute, &request.SysMenu, &request.Options)
	}

	// route没变化
	return application.repository.SysMenuRepository.UpdateById(request.Id, &request.SysMenu, &request.Options)
}

// Delete 更新
func (application *SysMenuApplication) Delete(request int64) error {
	count, err := application.repository.SysMenuRepository.CountByParent(request)
	if err != nil {
		logrus.Error("SysMenuRepository.CountByParent() happen error", err)
		return err
	}

	if count != 0 {
		return errors.New("子菜单数量不为空, 不可删除")
	}

	count, err = application.repository.SysPermissionRepository.CountByMenuId(request)
	if err != nil {
		logrus.Error("SysPermissionRepository.CountByMenuId() happen error", err)
		return err
	}

	if count != 0 {
		return errors.New("相关权限不为空, 不可删除")
	}
	return application.repository.SysMenuRepository.Delete(request)
}

// QueryOptionByMenuId 查询菜单下的所有操作
func (application *SysMenuApplication) QueryOptionByMenuId(menuId int64) ([]entity.SysMenuOption, error) {
	return application.repository.SysMenuOptionRepository.SelectByMenuId(menuId)
}

// Refresh 刷新
func (application *SysMenuApplication) Refresh() error {
	options, err := application.repository.SysMenuOptionRepository.SelectAll()
	if err != nil {
		logrus.Error("SysMenuOptionRepository.SelectAll() happen error", err)
		return err
	}

	for index, option := range options {
		codeKey := fmt.Sprintf("%v-%v-%v-%v", option.MenuId, option.Value, option.Method, option.Path)
		option.Code = fmt.Sprintf("%x", md5.Sum([]byte(codeKey)))
		options[index] = option
	}
	return application.repository.SysMenuOptionRepository.BatchUpdate(options)
}
