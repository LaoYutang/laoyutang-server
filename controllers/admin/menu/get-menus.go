package menu

import (
	"context"

	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/redis/go-redis/v9"
)

func GetMenus() (res map[int]map[string]any, errOut error) {
	ctx := context.Background()
	var (
		err    error
		resStr string
	)
	menus := map[int]map[string]any{}

	resStr, err = db.Redis.Get(ctx, "menus").Result()
	if err == redis.Nil {
		// 数据库读取并格式化
		menuStructs := &[]structs.Menu{}
		if err = db.Sql.Find(menuStructs).Error; err != nil {
			return nil, err
		}

		for _, menu := range *menuStructs {
			menus[menu.Id] = structs.H{
				"pid":  menu.Pid,
				"name": menu.Name,
				"type": menu.Type,
				"sign": menu.Sign,
			}
		}

		db.Redis.Set(ctx, "menus", utils.ToJson(menus), 0)
	} else if err != nil {
		return nil, err
	} else {
		// 解析json数据
		err = utils.ParseJson(resStr, &menus)
		if err != nil {
			return nil, err
		}
	}

	return menus, nil
}
