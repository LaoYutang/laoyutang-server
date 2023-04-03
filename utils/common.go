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

// 自定义节点切片去重方法
func RemoveDupCustom(arr []any, getValue func(any) any) (res []any) {
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

type treeNode struct {
	id       any
	pid      any
	children *[]treeNode
	data     any
}

// list生成树方法
func GenerateTree(list []any, getAttr func(any) (id any, pid any)) (res []treeNode) {
	childrenMap := map[any]*[]treeNode{}

	for _, val := range list {
		id, pid := getAttr(val)
		node := &treeNode{
			id:       id,
			pid:      pid,
			children: &[]treeNode{},
			data:     val,
		}

		if _, ok := childrenMap[pid]; !ok {
			childrenMap[pid] = &[]treeNode{}
		}
		if _, ok := childrenMap[id]; !ok {
			childrenMap[id] = node.children
		}

		*childrenMap[pid] = append(*childrenMap[pid], *node)

		if pid == 0 || pid == "0" || pid == nil {
			res = append(res, *node)
		}
	}

	return res
}
