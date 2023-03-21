package user

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/configs"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

// 登录方法
func Login(c *gin.Context) {
	type LoginForm struct {
		UserName string `json:"userName" required:"true" label:"用户名"`
		Password string `json:"password" required:"true" label:"密码"`
	}
	form := &LoginForm{}
	c.ShouldBind(form)

	// 判断必填项
	requiredErr := utils.ValRequired(form)
	if requiredErr != nil {
		utils.ResponseFail(c, 400, requiredErr.Error())
		return
	}

	// 查询用户
	user := &structs.User{}
	if err := db.Sql.Where("user_name = ?", form.UserName).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ResponseFail(c, 400, "用户未注册或密码不正确")
		} else {
			logrus.Error(err)
			utils.ResponseFailDefault(c)
		}
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		utils.ResponseFail(c, 400, "用户未注册或密码不正确")
		return
	}

	// 密码正确，生成token
	// token数据结构体
	claim := &structs.Claims{
		UserName: user.UserName,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(configs.TokenExpire).Unix(),
			Issuer:    "laoyutang",
		},
	}
	token, jwtErr := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(os.Getenv("LAOYUTANG_SECRET_KEY")))
	if jwtErr != nil {
		logrus.Error(jwtErr)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, &gin.H{
		"token": token,
	})
}
