package aiApiKeys

import (
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
	keys := &[]structs.AiApiKeyApi{}
	if err := db.Sql.Model(&structs.AiApiKey{}).Order("id desc").Find(keys).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, keys)
}

type keyForm struct {
	Name   string `json:"name" validate:"required,min=2,max=30"`
	Type   string `json:"type" validate:"required,min=2,max=30,oneof=Openai"`
	ApiKey string `json:"apiKey" validate:"required,min=1,max=100"`
}

func Create(c *gin.Context) {
	form := &keyForm{}
	c.ShouldBindJSON(form)

	if err := validator.New().Struct(form); err != nil {
		utils.ResponseFail(c, http.StatusBadRequest, err.Error())
		return
	}

	key := &structs.AiApiKey{
		Name:   form.Name,
		Type:   form.Type,
		ApiKey: utils.AesEncrypt(form.ApiKey), // apiKey加密存储
		Model: structs.Model{
			CreatedBy: c.MustGet("UserName").(string),
			UpdatedBy: c.MustGet("UserName").(string),
		},
	}

	if err := db.Sql.Create(key).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, nil)
}

func Delete(c *gin.Context) {
	ids := []int{}
	c.ShouldBindJSON(&ids)

	delData := &map[string]interface{}{
		"DeletedAt": time.Now(),
		"DeletedBy": c.MustGet("UserName"),
	}

	if err := db.Sql.Model(&structs.AiApiKey{}).Where(ids).Updates(delData).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, nil)
}
