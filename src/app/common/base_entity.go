package common

type BaseEntity struct {
	Id        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt *LocalTime `json:"createdAt" gorm:"index;column:created_at"`
	UpdatedAt *LocalTime `json:"updatedAt" gorm:"column:updated_at"`
}
