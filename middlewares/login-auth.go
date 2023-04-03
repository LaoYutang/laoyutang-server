package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseFail(c, http.StatusUnauthorized, "登录失效，请重新登录")
			c.Abort()
			return
		}

		// 解析token
		token, parseErr := jwt.ParseWithClaims(authHeader, &structs.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
			return []byte(os.Getenv("LAOYUTANG_SECRET_KEY")), nil
		})
		if parseErr != nil {
			logrus.Error("token parse error", parseErr)
			utils.ResponseFail(c, http.StatusUnauthorized, "登录失效，请重新登录")
			c.Abort()
			return
		}

		// 解析claim数据，放到context中
		claims, ok := token.Claims.(*structs.Claims)
		if !(ok && token.Valid) {
			logrus.Error("token parse error", parseErr)
			utils.ResponseFail(c, http.StatusUnauthorized, "登录失效，请重新登录")
			c.Abort()
			return
		}

		c.Set("UserName", claims.UserName)
		c.Set("UserId", claims.UserId)
		c.Next()
	}
}
