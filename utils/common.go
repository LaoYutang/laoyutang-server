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
		logrus.Error(errors.New("ToJson Error: " + err.Error()))
		return ""
	}
	return string(bytes)
}

// 解析json方法
func ParseJson(jsonStr string, target any) error {
	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		logrus.Error(errors.New("ParseJson Error: " + err.Error()))
		logrus.Error(errors.New("ParseJson String: " + jsonStr))
		return err
	}
	return nil
}

// 切片去重方法
func RemoveDup(arr []any) (res []any) {
	tmp := map[any]bool{}
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
			res = append(res, v)
		}
	}
	return res
}
