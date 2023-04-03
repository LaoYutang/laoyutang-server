package structs

type Pager struct {
	PageSize    int `json:"pageSize" form:"pageSize" validate:"required"`
	CurrentPage int `json:"currentPage" form:"currentPage" validate:"required"`
}
