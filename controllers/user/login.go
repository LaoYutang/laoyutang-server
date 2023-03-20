package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// 注册方法
func SignIn(c *gin.Context) {
	type signForm struct {
		UserName        string `json:"userName" required:"true" label:"用户名"`
		Password        string `json:"password" required:"true" label:"密码"`
		ConfirmPassword string `json:"confirmPassword" required:"true" label:"确认密码"`
	}
	form := &signForm{}
	c.ShouldBind(form)

	// 判断必填项
	requiredErr := utils.ValRequired(form)
	if requiredErr != nil {
		utils.ResponseFail(c, 400, requiredErr.Error())
		return
	}

	// 判断密码是否一致
	if form.Password != form.ConfirmPassword {
		utils.ResponseFail(c, 400, "两次输入的密码不一致，请重新输入")
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
			CreatedAt: time.Now(),
			CreatedBy: "sys",
			UpdatedAt: time.Now(),
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
