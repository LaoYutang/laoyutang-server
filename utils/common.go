package utils

import "encoding/json"

func ToJson(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
