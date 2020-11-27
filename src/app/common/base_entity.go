package common

type BaseEntity struct {
	Id        int64      `json:"id,string" gorm:"primary_key"`
	CreatedAt *LocalTime `json:"createdAt" gorm:"index;column:created_at"`
	UpdatedAt *LocalTime `json:"updatedAt" gorm:"column:updated_at"`
}
