package role

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func Read(c *gin.Context) {
	roles := &[]structs.Role{}
	if err := db.Sql.Order("created_at desc").Find(roles).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, roles)
}

type roleForm struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=2,max=30"`
	Menus string `json:"menus"`
	Desc  string `json:"desc"`
}

func Create(c *gin.Context) {
	form := &roleForm{}
	c.ShouldBind(form)

	if err := validator.New().Struct(form); err != nil {
		utils.ResponseFail(c, http.StatusBadRequest, err.Error())
		return
	}

	role := &structs.Role{
		Name:  form.Name,
		Menus: form.Menus,
		Desc:  form.Desc,
		Model: structs.Model{
			CreatedBy: c.MustGet("UserName").(string),
			UpdatedBy: c.MustGet("UserName").(string),
		},
	}

	if err := db.Sql.Create(role).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	ctx := context.Background()
	db.Redis.Del(ctx, "roles")

	utils.ResponseSuccess(c, nil)
}

func Update(c *gin.Context) {
	form := &roleForm{}
	c.ShouldBind(form)

	if err := validator.New().Struct(form); err != nil {
		utils.ResponseFail(c, http.StatusBadRequest, err.Error())
		return
	}

	role := &map[string]interface{}{
		"UpdatedBy": c.MustGet("UserName").(string),
		"Name":      form.Name,
		"Menus":     form.Menus,
		"Desc":      form.Desc,
	}

	if err := db.Sql.Model(&structs.Role{}).Where(form.Id).Updates(role).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	ctx := context.Background()
	db.Redis.Del(ctx, "roles")

	utils.ResponseSuccess(c, nil)
}

func Delete(c *gin.Context) {
	ids := []int{}
	c.ShouldBindJSON(&ids)

	delData := &map[string]interface{}{
		"DeletedAt": time.Now(),
		"DeletedBy": c.MustGet("UserName"),
	}

	if err := db.Sql.Model(&structs.Role{}).Where(ids).Updates(delData).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}
	ctx := context.Background()
	db.Redis.Del(ctx, "roles")

	utils.ResponseSuccess(c, nil)
}
