package response

import (
	"time"
)

// ReturnMessage  Http API return data
type ReturnMessage struct {
	Success      bool        `json:"Success"`
	Data         interface{} `json:"Data"`
	ErrorCode    string      `json:"ErrorCode"`
	ErrorMessage string      `json:"ErrorMessage"`
}

//PageResponse 分页返回
type PageResponse struct {
	// 总数
	Count      int64 `json:"count"`
	PageSize   int   `json:"page_size"`
	PageNumber int   `json:"page_number"`
	// 分页总数
	PageCount int `json:"page_count"`
}

type UserLogin struct {
	Name       string    `json:"name" xorm:"not null unique VARCHAR(255) comment('用户名')"`
	LastLogin  time.Time `json:"last_login" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}
