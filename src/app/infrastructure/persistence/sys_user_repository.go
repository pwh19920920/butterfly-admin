package persistence

import (
	"butterfly-admin/src/app/common"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/types"
	"gorm.io/gorm"
)

type SysUserRepositoryImpl struct {
	db *gorm.DB
}

func NewSysUserRepositoryImpl(db *gorm.DB) *SysUserRepositoryImpl {
	return &SysUserRepositoryImpl{db: db}
}

func (repo *SysUserRepositoryImpl) GetByUsername(username string) (*entity.SysUser, error) {
	var user entity.SysUser
	err := repo.
		db.Model(&entity.SysUser{}).
		Where("username = ?", username).
		Not(&entity.SysUser{BaseEntity: common.BaseEntity{Deleted: common.DeletedTrue}}).
		First(&user).Error
	return &user, err
}

func (repo *SysUserRepositoryImpl) GetById(id int64) (*entity.SysUser, error) {
	var user entity.SysUser
	err := repo.db.
		Model(&entity.SysUser{}).
		Where("id = ?", id).
		Not(&entity.SysUser{BaseEntity: common.BaseEntity{Deleted: common.DeletedTrue}}).
		First(&user).Error
	return &user, err
}

func (repo *SysUserRepositoryImpl) Select(req *types.SysUserQueryRequest) (int64, []entity.SysUser, error) {
	var count int64 = 0
	whereCase := &entity.SysUser{
		Name:     req.Name,
		Username: req.Username,
	}
	notCase := &entity.SysUser{BaseEntity: common.BaseEntity{Deleted: common.DeletedTrue}}
	repo.db.Model(&entity.SysUser{}).Where(whereCase).Count(&count)

	var data []entity.SysUser
	err := repo.db.Model(&entity.SysUser{}).
		Where(whereCase).
		Not(notCase).
		Limit(req.PageSize).Offset(req.Offset()).Find(&data).Error
	return count, data, err
}

// Create 创建
func (repo *SysUserRepositoryImpl) Create(user *entity.SysUser) error {
	return repo.db.Model(&entity.SysUser{}).Create(&user).Error
}

// Modify 更新
func (repo *SysUserRepositoryImpl) Modify(user *entity.SysUser) error {
	return repo.db.
		Model(&entity.SysUser{}).
		Where(&entity.SysUser{BaseEntity: common.BaseEntity{Id: user.Id}}).
		Updates(&user).Error
}
