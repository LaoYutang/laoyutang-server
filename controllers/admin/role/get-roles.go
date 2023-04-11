package role

import (
	"context"

	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/redis/go-redis/v9"
)

func GetRoles() (res map[int]map[string]any, errOut error) {
	ctx := context.Background()
	var (
		err    error
		resStr string
	)
	roles := map[int]map[string]any{}

	resStr, err = db.Redis.Get(ctx, "roles").Result()
	if err == redis.Nil {
		// 数据库读取并格式化
		roleStructs := &[]structs.Role{}
		if err = db.Sql.Find(roleStructs).Error; err != nil {
			return nil, err
		}

		for _, role := range *roleStructs {
			roles[role.Id] = structs.H{
				"name":  role.Name,
				"menus": role.Menus,
			}
		}

		db.Redis.Set(ctx, "roles", utils.ToJson(roles), 0)
	} else if err != nil {
		return nil, err
	} else {
		// 解析json数据
		err = utils.ParseJson(resStr, &roles)
		if err != nil {
			return nil, err
		}
	}

	return roles, nil
}
