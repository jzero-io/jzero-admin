package tables

type CasbinRule struct {
	Id    int    `gorm:"primary_key;AUTO_INCREMENT"`
	Ptype string `gorm:"column:ptype;type:varchar(100)"`
	V0    string `gorm:"column:v0;type:varchar(100)"`
	V1    string `gorm:"column:v1;type:varchar(100)"`
	V2    string `gorm:"column:v2;type:varchar(100)"`
	V3    string `gorm:"column:v3;type:varchar(100)"`
	V4    string `gorm:"column:v4;type:varchar(100)"`
	V5    string `gorm:"column:v5;type:varchar(100)"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
