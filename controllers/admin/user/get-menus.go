package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
)

func GetMenus(c *gin.Context) {
	// 获取用户角色
	user := &structs.User{}
	db.Sql.Where("user-name = ?", c.MustGet("UserName")).First(user)
	fmt.Println(user)
}
