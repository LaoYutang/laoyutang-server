package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func Read(c *gin.Context) {
	type Form struct {
		UserName string `form:"userName"`
		structs.Pager
	}
	form := &Form{}
	c.ShouldBind(form)

	// 字段校验
	vadErr := validator.New().Struct(form)
	if vadErr != nil {
		utils.ResponseFail(c, http.StatusBadRequest, vadErr.Error())
		return
	}

	users := &[]structs.UserApi{}
	offset, limit := (form.CurrentPage-1)*form.PageSize, form.PageSize
	if err := db.Sql.Model(&structs.User{}).Where("user_name like ?", form.UserName+"%").Offset(offset).Limit(limit).Order("created_at desc").Find(users).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}
	var total int64
	db.Sql.Model(&structs.User{}).Count(&total)

	utils.ResponseSuccess(c, &structs.H{
		"total": total,
		"list":  users,
	})

}
