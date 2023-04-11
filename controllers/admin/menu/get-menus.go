package menu

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func GetMenusMap() (res map[int]map[string]any, errOut error) {
	getData := func() (menus map[int]map[string]any, err error) {
		menus = map[int]map[string]any{}
		menuStructs := &[]structs.Menu{}
		if err = db.Sql.Find(menuStructs).Error; err != nil {
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

	return utils.UseRedis("menusMap", getData, 0)
}

func GetList(c *gin.Context) {
	getData := func() (menus []map[string]any, err error) {
		menuStructs := &[]structs.Menu{}
		if err = db.Sql.Find(menuStructs).Error; err != nil {
			return
		}

		menus = make([]map[string]any, len(*menuStructs))
		for index, menu := range *menuStructs {
			menus[index] = structs.H{
				"id":   menu.Id,
				"pid":  menu.Pid,
				"name": menu.Name,
				"type": menu.Type,
				"sign": menu.Sign,
			}
		}

		return
	}

	res, err := utils.UseRedis("menusList", getData, 0)
	if err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, res)
}

func DelMenusCache() {
	ctx := context.Background()
	db.Redis.Del(ctx, "menusMap")
	db.Redis.Del(ctx, "menusList")
}
