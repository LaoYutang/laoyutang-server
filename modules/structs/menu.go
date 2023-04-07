package structs

type Menu struct {
	Model
	Pid  int    `json:"pid" gorm:"type:bigint;comment:父级id 0为根节点"`
	Name string `json:"name" gorm:"type:varchar(100);comment:名称"`
	Type int    `json:"type" gorm:"type:tinyint;comment:类型 0菜单 1权限"`
	Sign string `json:"sign" gorm:"type:varchar(100);comment:标识"`
}

func (Menu) TableName() string {
	return "t_menu"
}
