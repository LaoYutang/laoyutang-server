package menu

import (
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
)

func GetMenusMap() (res map[int]map[string]any, errOut error) {
	getData := func() (menus map[int]map[string]any, err error) {
		menus = map[int]map[string]any{}
		menuStructs := &[]structs.Menu{}
		if errOut = db.Sql.Find(menuStructs).Error; errOut != nil {
			return
		}

		for _, menu := range *menuStructs {
			menus[menu.Id] = structs.H{
				"id":   menu.Id,
				"pid":  menu.Pid,
				"name": menu.Name,
				"type": menu.Type,
				"sign": menu.Sign,
			}
		}

		return
	}

	return utils.UseRedis("menus", getData, 0)
}

// func GetList(c *gin.Context) {
// 	menus, err := GetMenus()
// 	if err != nil {
// 		logrus.Error(err)
// 		utils.ResponseFailDefault(c)
// 		return
// 	}

// 	utils.ResponseSuccess(c, menus)
// }
