package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"gorm.io/gorm"
)

var sql = db.GetDB()

func Test(c *gin.Context) {
	user := &structs.User{}

	result := sql.Where("id = ?", 1).First(user)
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println(">>> no record!")
		return
	}

	res := &structs.Response{
		Success: true,
		Data: &gin.H{
			"test": c.Request.Method,
		},
	}
	c.JSON(http.StatusOK, res)
}
