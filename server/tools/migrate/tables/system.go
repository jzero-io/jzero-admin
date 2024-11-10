package tables

type SystemUser struct {
	BaseModel

	Username string `gorm:"column:username;type:varchar(30);unique;not null"`
	Password string `gorm:"column:password;type:varchar(100);not null"`
	Nickname string `gorm:"column:nickname;type:varchar(30);not null"`
	Gender   string `gorm:"column:gender;type:varchar(1);not null"`
	Phone    string `gorm:"column:phone;type:varchar(20)"`
	Email    string `gorm:"column:email;type:varchar(100)"`
	Status   string `gorm:"column:status;type:varchar(1);not null"`
}

func (SystemUser) TableName() string {
	return "system_user"
}

type SystemRole struct {
	BaseModel

	Name   string `gorm:"column:name;type:varchar(50);not null"`
	Status string `gorm:"column:status;type:varchar(1);not null"`
	Code   string `gorm:"column:code;type:varchar(255);not null"`
	Desc   string `gorm:"column:desc;not null"`
}

func (SystemRole) TableName() string {
	return "system_role"
}

type SystemUserRole struct {
	BaseModel

	UserId int `gorm:"column:user_id;not null"`
	RoleId int `gorm:"column:role_id;not null"`
}

func (SystemUserRole) TableName() string {
	return "system_user_role"
}

type SystemMenu struct {
	BaseModel

	Status          string `gorm:"column:status;type:varchar(1);not null"`
	ParentId        int    `gorm:"column:parent_id;not null"`
	MenuType        string `gorm:"column:menu_type;type:varchar(1);not null"`
	MenuName        string `gorm:"column:menu_name;type:varchar(50);not null"`
	HideInMenu      bool   `gorm:"column:hide_in_menu;type:tinyint(1);not null"`
	ActiveMenu      string `gorm:"column:active_menu;type:varchar(50)"`
	Order           int    `gorm:"column:order;not null"`
	RouteName       string `gorm:"column:route_name;type:varchar(255);not null"`
	RoutePath       string `gorm:"column:route_path;type:varchar(255);not null"`
	Component       string `gorm:"column:component;type:varchar(255);not null"`
	Icon            string `gorm:"column:icon;type:varchar(255);not null"`
	IconType        string `gorm:"column:icon_type;type:varchar(1);not null"`
	I18nKey         string `gorm:"column:i18n_key;type:varchar(255);not null"`
	KeepAlive       bool   `gorm:"column:keep_alive;not null"`
	Href            string `gorm:"column:href"`
	MutiTab         bool   `gorm:"column:multi_tab"`
	FixedIndexInTab int    `gorm:"column:fixed_index_in_tab;not null"`
	Query           string `gorm:"column:query"`
	Buttons         string `gorm:"column:buttons"`
	Constant        bool   `gorm:"column:constant;not null"`
}

func (SystemMenu) TableName() string {
	return "system_menu"
}

type SystemRoleMenu struct {
	BaseModel

	RoleId int `gorm:"column:role_id;not null"`
	MenuId int `gorm:"column:menu_id;not null"`
}

func (SystemRoleMenu) TableName() string {
	return "system_role_menu"
}

type SystemEmail struct {
	BaseModel

	From      string `gorm:"column:from;not null"`
	Host      string `gorm:"column:host;not null"`
	Port      int    `gorm:"column:port;not null"`
	Username  string `gorm:"column:username;not null"`
	Password  string `gorm:"column:password;not null"`
	EnableSsl bool   `gorm:"column:enable_ssl;type:tinyint(1);not null"`
	IsVerify  bool   `gorm:"column:is_verify;type:tinyint(1);not null"`
}

func (SystemEmail) TableName() string {
	return "system_email"
}
