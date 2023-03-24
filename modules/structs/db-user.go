package structs

type User struct {
	Model
	UserName string `json:"userName" gorm:"type:varchar(100);comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(100);comment:密码加密串"`
	RoleIds  string `json:"roleIds" gorm:"type:varchar(100);comment:角色id集 使用,分隔"`
	Email    string `json:"email" gorm:"type:varchar(100)"`
	Phone    string `json:"phone" gorm:"type:varchar(100)"`
}

func (u User) TableName() string {
	return "user_info"
}
