package structs

type AiApiKey struct {
	Model
	Name   string `json:"name" gorm:"type:varchar(100);unique;comment:名称"`
	Type   string `json:"type" gorm:"type:varchar(20);comment:类型 暂时只支持Openai"`
	ApiKey string `json:"apiKey" gorm:"type:varchar(200);unique;comment:ApiKey"`
}

func (AiApiKey) TableName() string {
	return "t_ai_api_keys"
}
