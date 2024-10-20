package tables

import "time"

type BaseModel struct {
	Id         uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreateBy   int       `gorm:"column:create_by"`
	UpdateBy   int       `gorm:"column:update_by"`
}
