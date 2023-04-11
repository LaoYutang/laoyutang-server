package role

import (
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
)

func GetRolesMap() (res map[int]map[string]any, errOut error) {
	getData := func() (roles map[int]map[string]any, err error) {
		roles = map[int]map[string]any{}
		roleStructs := &[]structs.Role{}
		if errOut = db.Sql.Find(roleStructs).Error; errOut != nil {
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

	return utils.UseRedis("roles", getData, 0)
}

// func GetList(c *gin.Context) {
// 	roles, err := GetRoles()
// 	if err != nil {
// 		logrus.Error(err)
// 		utils.ResponseFailDefault(c)
// 		return
// 	}

// 	utils.ResponseSuccess(c, roles)
// }
