package base

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/laoyutang/laoyutang-server/modules/configs"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 登录方法
func Login(c *gin.Context) {
	type LoginForm struct {
		UserName string `json:"userName" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	form := &LoginForm{}
	c.ShouldBind(form)

	// 字段校验
	vadErr := validator.New().Struct(form)
	if vadErr != nil {
		utils.ResponseFail(c, 400, vadErr.Error())
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
