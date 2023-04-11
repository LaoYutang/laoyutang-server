package menu

import (
	"fmt"
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
	menus := &[]structs.Menu{}
	if err := db.Sql.Order("id desc").Find(menus).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, menus)
}

type menuForm struct {
	Id   int    `json:"id"`
	Pid  int    `json:"pid" validate:"min=0"`
	Name string `json:"name" validate:"required,min=2,max=30"`
	Type int    `json:"type" validate:"min=0,max=1"`
	Sign string `json:"sign" validate:"required,min=2,max=50"`
}

func Create(c *gin.Context) {
	form := &menuForm{}
	c.ShouldBindJSON(form)

	if err := validator.New().Struct(form); err != nil {
		utils.ResponseFail(c, http.StatusBadRequest, err.Error())
		return
	}

	menu := &structs.Menu{
		Pid:  form.Pid,
		Name: form.Name,
		Type: form.Type,
		Sign: form.Sign,
		Model: structs.Model{
			CreatedBy: c.MustGet("UserName").(string),
			UpdatedBy: c.MustGet("UserName").(string),
		},
	}

	if err := db.Sql.Create(menu).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	DelMenusCache()

	utils.ResponseSuccess(c, nil)
}

func Update(c *gin.Context) {
	form := &menuForm{}
	c.ShouldBind(form)

	if err := validator.New().Struct(form); err != nil {
		utils.ResponseFail(c, http.StatusBadRequest, err.Error())
		return
	}

	menu := &map[string]interface{}{
		"UpdatedBy": c.MustGet("UserName").(string),
		"Pid":       form.Pid,
		"Name":      form.Name,
		"Type":      form.Type,
		"Sign":      form.Sign,
	}

	if err := db.Sql.Model(&structs.Menu{}).Where(form.Id).Updates(menu).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	DelMenusCache()

	utils.ResponseSuccess(c, nil)
}

func Delete(c *gin.Context) {
	ids := []int{}
	c.ShouldBindJSON(&ids)

	for _, id := range ids {
		roles := &[]structs.Role{}
		if err := db.Sql.Where("menus REGEXP ?", fmt.Sprintf("(^|,)%d(,|$)", id)).Find(roles).Error; err != nil {
			logrus.Error(err)
			utils.ResponseFailDefault(c)
			return
		}

		if len(*roles) > 0 {
			utils.ResponseFail(c, http.StatusBadRequest, "选择的菜单权限已经被角色使用！")
			return
		}
	}

	delData := &map[string]interface{}{
		"DeletedAt": time.Now(),
		"DeletedBy": c.MustGet("UserName"),
	}

	if err := db.Sql.Model(&structs.Menu{}).Where(ids).Updates(delData).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	DelMenusCache()

	utils.ResponseSuccess(c, nil)
}
