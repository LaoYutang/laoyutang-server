package structs

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	ErrNo   int         `json:"errNo"`
}
