package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/controllers/admin/user"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

// 功能权限鉴权中间件
// perm 权限字符串Sign
func FuncAuth(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("UserId").(int)
		if userId == 0 {
			utils.ResponseFail(c, http.StatusUnauthorized, "登录失效，请重新登录")
			c.Abort()
			return
		}

		permList, err := user.GetUserPerms(userId)
		if err != nil {
			logrus.Error("GetUserPerms Error:" + err.Error())
			utils.ResponseFailDefault(c)
			c.Abort()
			return
		}

		if !utils.SliceIncludes(permList, perm) {
			utils.ResponseFail(c, http.StatusForbidden, "Forbidden!")
			c.Abort()
			return
		}

		c.Next()
	}
}
