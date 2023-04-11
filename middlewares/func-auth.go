package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/controllers/admin/menu"
	"github.com/laoyutang/laoyutang-server/controllers/admin/role"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

// 功能权限鉴权中间件
// perm 权限字符串Sign
func FuncAuth(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			roles, menus map[int]map[string]any
			err          error
		)

		// 获取角色信息
		if roles, err = role.GetRolesMap(); err != nil {
			logrus.Error(err)
			utils.ResponseFailDefault(c)
			c.Abort()
			return
		}

		// 获取菜单信息
		if menus, err = menu.GetMenusMap(); err != nil {
			logrus.Error(err)
			utils.ResponseFailDefault(c)
			c.Abort()
			return
		}

		for _, roleId := range strings.Split(c.MustGet("UserRoles").(string), ",") {
			roleIdInt, err := strconv.Atoi(roleId)
			if err != nil {
				logrus.Error(err)
				utils.ResponseFailDefault(c)
				c.Abort()
				return
			}
			for _, menuId := range strings.Split(roles[roleIdInt]["menus"].(string), ",") {
				if menuId == "0" {
					continue
				}
				menuIdInt, err := strconv.Atoi(menuId)
				if err != nil {
					logrus.Error(err)
					utils.ResponseFailDefault(c)
					c.Abort()
					return
				}

				if menus[menuIdInt]["sign"].(string) == perm {
					c.Next()
					return
				}
			}
		}

		utils.ResponseFail(c, http.StatusForbidden, "Forbidden!")
		c.Abort()
	}
}
