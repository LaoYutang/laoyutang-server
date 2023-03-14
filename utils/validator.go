package utils

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
)

// 判断必填项
func ValRequired(data interface{}) error {
	val := reflect.ValueOf(data).Elem() //获取字段值
	typ := reflect.TypeOf(data).Elem()  //获取字段类型
	// 遍历struct中的字段
	for i := 0; i < typ.NumField(); i++ {
		// 当struct中的tag为 required:"true" 且当前字段值为空值时，输出
		if typ.Field(i).Tag.Get("required") == "true" && val.Field(i).IsZero() {
			return errors.New(typ.Field(i).Tag.Get("label") + "不能为空")
		}
	}
	return nil
}

// 判断必填项
func ValRequiredAndResponse(data interface{}, c *gin.Context) error {
	err := ValRequired(data)
	if err != nil {
		c.JSON(http.StatusOK, &structs.Response{
			Success: false,
			ErrNo:   400,
			Message: err.Error(),
		})
		return err
	}
	return nil
}
