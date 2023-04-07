package structs

type Role struct {
	Model
	Name  string `json:"name" gorm:"type:varchar(100);comment:角色名称"`
	Menus string `json:"menus" gorm:"type:varchar(500);comment:授权的菜单和权限"`
	Desc  string `json:"desc" gorm:"type:varchar(100);comment:描述"`
}

func (Role) TableName() string {
	return "t_role"
}
