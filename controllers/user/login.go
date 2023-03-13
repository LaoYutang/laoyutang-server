package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"golang.org/x/crypto/bcrypt"
)

var sql = db.GetDB()

func SignIn(c *gin.Context) {
	type signForm struct {
		UserName        string `json:"userName"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	form := &signForm{}
	c.BindJSON(form)

	// 判断密码是否一致
	if form.Password != form.ConfirmPassword {
		c.JSON(http.StatusOK, &structs.Response{
			Success: false,
			ErrNo:   400,
			Message: "两次输入的密码不一致，请重新输入",
		})
		return
	}

	// 判断用户是否存在
	var count int64
	sql.Model(&structs.User{}).Where("user_name = ?", form.UserName).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, &structs.Response{
			Success: false,
			ErrNo:   400,
			Message: "用户名已存在",
		})
		return
	}

	// 密码加密
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 4)
	fmt.Printf("%v", string(encryptPassword))

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

	result := sql.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusOK, &structs.Response{
			Success: false,
			ErrNo:   500,
			Message: "服务器异常",
		})
		return
	}

	c.JSON(http.StatusOK, &structs.Response{
		Success: true,
		Data: &gin.H{
			"id": user.Id,
		},
	})
}
