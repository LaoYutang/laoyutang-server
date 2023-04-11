package role

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func GetRolesMap() (res map[int]map[string]any, errOut error) {
	getData := func() (roles map[int]map[string]any, err error) {
		roles = map[int]map[string]any{}
		roleStructs := &[]structs.Role{}
		if err = db.Sql.Find(roleStructs).Error; err != nil {
			return
		}

		for _, menu := range *roleStructs {
			roles[menu.Id] = structs.H{
				"id":    menu.Id,
				"name":  menu.Name,
				"menus": menu.Menus,
			}
		}

		return
	}

	return utils.UseRedis("rolesMap", getData, 0)
}

func GetList(c *gin.Context) {
	getData := func() (roles []map[string]any, err error) {
		roleStructs := &[]structs.Role{}
		if err = db.Sql.Find(roleStructs).Error; err != nil {
			return
		}

		roles = make([]map[string]any, len(*roleStructs))
		for index, role := range *roleStructs {
			roles[index] = structs.H{
				"id":    role.Id,
				"name":  role.Name,
				"menus": role.Menus,
			}
		}

		return
	}

	res, err := utils.UseRedis("rolesList", getData, 0)
	if err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, res)
}

func DelRolesCache() {
	ctx := context.Background()
	db.Redis.Del(ctx, "rolesMap")
	db.Redis.Del(ctx, "rolesList")
}
