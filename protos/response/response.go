package response

import (
	"time"

	"github.com/goodaye/fakeeyes/protos/db"
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
	db.User
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}
