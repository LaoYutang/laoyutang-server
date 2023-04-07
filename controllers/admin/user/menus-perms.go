package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func GetMenusAndPerms(c *gin.Context) {
	// 获取用户角色
	user := &structs.User{}
	if err := db.Sql.Where(c.MustGet("UserId")).First(user).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	userRoles := strings.Split(user.RoleIds, ",")
	if len(userRoles) == 0 {
		logrus.Error(fmt.Sprintf("用户角色为空, %v ", c.MustGet("UserId")))
		utils.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("用户角色为空, %v ", c.MustGet("UserId")))
		return
	}

	roles := &[]structs.Role{}
	if err := db.Sql.Where(userRoles).Find(roles).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	menuIds := []string{}
	for _, role := range *roles {
		menuIds = append(menuIds, strings.Split(role.Menus, ",")...)
	}
	menuIds = utils.RemoveDup(menuIds)

	menus := &[]structs.Menu{}
	if err := db.Sql.Where("Type = 0").Find(menus, menuIds).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	tree := utils.GenerateTree(*menus, func(node structs.Menu) (any, any) {
		return node.Id, node.Pid
	})

	permList := &[]map[string]interface{}{}
	if err := db.Sql.Model(&structs.Menu{}).Select("Sign").Where("Type = 1").Find(permList, menuIds).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}
	perms := []string{}
	for _, perm := range *permList {
		perms = append(perms, perm["sign"].(string))
	}

	utils.ResponseSuccess(c, &structs.H{
		"menus": &tree,
		"perms": &perms,
	})
}

// 从数据库中获取用户权限列表方法
func GetUserPermsFromSql(userId int) (list []string, errOut error) {
	user := &structs.User{}
	if err := db.Sql.Where(userId).First(user).Error; err != nil {
		return nil, errors.New("查询角色失败")
	}

	userRoles := strings.Split(user.RoleIds, ",")
	if len(userRoles) == 0 {
		return nil, errors.New("用户角色为空")
	}

	roles := &[]structs.Role{}
	if err := db.Sql.Where(userRoles).Find(roles).Error; err != nil {
		return nil, errors.New("查询角色失败")
	}

	menuIds := []string{}
	for _, role := range *roles {
		menuIds = append(menuIds, strings.Split(role.Menus, ",")...)
	}
	menuIds = utils.RemoveDup(menuIds)

	permList := &[]map[string]interface{}{}
	if err := db.Sql.Model(&structs.Menu{}).Select("Sign").Where("Type = 1").Find(permList, menuIds).Error; err != nil {
		return nil, errors.New("查询权限失败")
	}
	for _, perm := range *permList {
		list = append(list, perm["sign"].(string))
	}
	return
}

// 获取用户权限缓存
func GetUserPerms(userId int) ([]string, error) {
	var err error
	var list []string
	ctx := context.Background()
	resStr, err := db.Redis.Get(ctx, "user_perm").Result()
	if err == redis.Nil {
		resStr = "{}"
	} else if err != nil {
		return nil, err
	}

	res := map[int][]string{}
	err = utils.ParseJson(resStr, &res)
	if err != nil {
		return nil, err
	}

	ok := false
	if list, ok = res[userId]; !ok {
		list, err = GetUserPermsFromSql(userId)
		if err != nil {
			return nil, err
		}
		res[userId] = list
		db.Redis.Set(ctx, "user_perm", utils.ToJson(res), 0)
	}

	return list, nil
}
