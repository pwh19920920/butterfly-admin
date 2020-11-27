package persistence

import (
	"butterfly-admin/src/app/common"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/types"
	"gorm.io/gorm"
)

type SysMenuRepositoryImpl struct {
	db *gorm.DB
}

func NewSysMenuRepositoryImpl(db *gorm.DB) *SysMenuRepositoryImpl {
	return &SysMenuRepositoryImpl{db: db}
}

// 保存
func (s *SysMenuRepositoryImpl) Save(menu entity.SysMenu) error {
	return s.db.Model(&entity.SysMenu{}).Create(&menu).Error
}

// 更新
func (s *SysMenuRepositoryImpl) UpdateById(id uint64, menu entity.SysMenu) error {
	return s.db.Model(&entity.SysMenu{}).
		Where(&entity.SysMenu{BaseEntity: common.BaseEntity{Id: id}}).Updates(&menu).Error
}

// 删除
func (s *SysMenuRepositoryImpl) Delete(id uint64) error {
	return s.db.Model(&entity.SysMenu{}).
		Where(&entity.SysMenu{BaseEntity: common.BaseEntity{Id: id}}).
		Updates(&entity.SysMenu{Deleted: 1}).Error
}

// 查询分页
func (s *SysMenuRepositoryImpl) Select(req types.SysMenuQueryRequest) (int64, []entity.SysMenu, error) {
	var count int64 = 0
	s.db.Model(&entity.SysMenu{}).Count(&count)

	var data []entity.SysMenu
	err := s.db.Model(&entity.SysMenu{}).Limit(req.PageSize).Offset(req.Offset()).Find(&data).Error
	return count, data, err
}

// 查询全部
func (s *SysMenuRepositoryImpl) SelectAll() ([]entity.SysMenu, error) {
	var data []entity.SysMenu
	err := s.db.Model(&entity.SysMenu{}).Find(&data).Error
	return data, err
}
