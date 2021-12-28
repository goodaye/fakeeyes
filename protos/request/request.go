package request

// 分页请求
type PageRequest struct {
	PageSize   int `json:"page_size,omitempty" validate:"gt=0"`
	PageNumber int `json:"page_number,omitempty" validate:"gt=0"`
}

// DefaultPageRequest 默认请求分页大小
var DefaultPageRequest = PageRequest{
	PageSize:   50,
	PageNumber: 1,
}
