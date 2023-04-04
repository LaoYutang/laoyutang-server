package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func Read(c *gin.Context) {
	menus := &[]structs.Menu{}
	if err := db.Sql.Order("created_at desc").Find(menus).Error; err != nil {
		logrus.Error(err)
		utils.ResponseFailDefault(c)
		return
	}

	utils.ResponseSuccess(c, menus)
}
