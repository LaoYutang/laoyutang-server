package base

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// 注册方法
func SignIn(c *gin.Context) {
	type signForm struct {
		UserName        string `json:"userName" validate:"required,min=4,max=16"`
		Password        string `json:"password" validate:"required,min=6,max=16"`
		ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password,min=6,max=16"`
		Invitation      string `json:"invitation" validate:"required,eq=142857"`
	}
	form := &signForm{}
	c.ShouldBind(form)

	// 字段校验
	vadErr := validator.New().Struct(form)
	if vadErr != nil {
		utils.ResponseFail(c, 400, vadErr.Error())
		return
	}

	// 判断用户是否存在
	var count int64
	sqlErr1 := db.Sql.Model(&structs.User{}).Where("user_name = ?", form.UserName).Count(&count).Error
	if sqlErr1 != nil {
		logrus.Error(sqlErr1)
		utils.ResponseFailDefault(c)
		return
	} else if count > 0 {
		utils.ResponseFail(c, 400, "用户名已存在")
		return
	}

	// 密码加密
	encryptPassword, encryptErr := bcrypt.GenerateFromPassword([]byte(form.Password), 4)
	if encryptErr != nil {
		logrus.Error(encryptErr)
		utils.ResponseFailDefault(c)
		return
	}

	user := &structs.User{
		UserName: form.UserName,
		Password: string(encryptPassword),
		Model: structs.Model{
			CreatedBy: "sys",
			UpdatedBy: "sys",
		},
	}

	result := db.Sql.Create(user)
	if result.Error != nil {
		logrus.Error(result.Error)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, &gin.H{
		"id": user.Id,
	})
}
