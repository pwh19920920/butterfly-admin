package entity

import "butterfly-admin/src/app/common"

type User struct {
	Id        uint64            `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt *common.LocalTime `json:"createdAt" gorm:"index;column:created_at"`
	UpdatedAt *common.LocalTime `json:"updatedAt" gorm:"column:updated_at"`
	Username  string            `json:"updatedAt" gorm:"column:username"`
	Password  string            `json:"updatedAt" gorm:"column:password"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "t_sys_user"
}
