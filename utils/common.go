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
func RemoveDup[T any](arr []T) (res []T) {
	tmp := map[any]bool{}
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
			res = append(res, v)
		}
	}
	return res
}

// 自定义节点切片去重方法
func RemoveDupCustom[T any](arr []T, getValue func(T) any) (res []T) {
	tmp := map[any]bool{}
	for _, v := range arr {
		val := getValue(v)
		if _, ok := tmp[val]; !ok {
			tmp[val] = true
			res = append(res, v)
		}
	}
	return res
}

type treeNode[T any] struct {
	Id       any            `json:"id"`
	Pid      any            `json:"pid"`
	Children *[]treeNode[T] `json:"children"`
	Data     T              `json:"data"`
}

// list生成树方法
func GenerateTree[T any](list []T, getAttr func(T) (id any, pid any)) (res []treeNode[T]) {
	childrenMap := map[any]*[]treeNode[T]{}

	for _, val := range list {
		id, pid := getAttr(val)
		node := &treeNode[T]{
			Id:       id,
			Pid:      pid,
			Children: &[]treeNode[T]{},
			Data:     val,
		}

		if _, ok := childrenMap[pid]; !ok {
			childrenMap[pid] = &[]treeNode[T]{}
		}
		if _, ok := childrenMap[id]; !ok {
			childrenMap[id] = node.Children
		}

		*childrenMap[pid] = append(*childrenMap[pid], *node)

		if pid == 0 || pid == "0" || pid == nil {
			res = append(res, *node)
		}
	}

	return res
}
