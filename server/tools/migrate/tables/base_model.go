package tables

import "time"

type BaseModel struct {
	Id         uint      `gorm:"column:id;primary_key"`
	Uuid       string    `gorm:"column:uuid;unique;not null"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreateBy   string    `gorm:"column:create_by;not null;default:''"`
}
