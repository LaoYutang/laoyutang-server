package structs

type User struct {
	Model
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	DepartmentId   int    `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

func (u User) TableName() string {
	return "user_info"
}
