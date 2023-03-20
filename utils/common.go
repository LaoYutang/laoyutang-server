package utils

import (
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
)

// 转json字符串方法
func ToJson(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		logrus.Error(errors.New("ToJson Error: \n" + err.Error()))
	}
	return string(bytes)
}
