package ginhandler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

// ReturnMessage  Http API return data
type ReturnMessage struct {
	Success      bool        `json:"Success"`
	Data         interface{} `json:"Data"`
	ErrorCode    string      `json:"ErrorCode"`
	ErrorMessage string      `json:"ErrorMessage"`
}

var myjson jsoniter.API
var validate *validator.Validate

func init() {
	myjson = jsoniter.Config{
		EscapeHTML:    true,
		CaseSensitive: true, // 配置大小写敏感
	}.Froze()
	validate = validator.New()
}

// BaseHandler  Command tool for gin handler
type BaseHandler struct {
	Version string
}

// // ReturnMessage  Http API return data
// type ReturnMessage struct {
// 	Success      bool        `json:"Success"`
// 	Data         interface{} `json:"Data"`
// 	ErrorCode    string      `json:"ErrorCode"`
// 	ErrorMessage string      `json:"ErrorMessage"`
// }

// GetVersion  return version
func (h BaseHandler) GetVersion() string {
	return h.Version
}

// Pong default response for request ping
func (h BaseHandler) Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// PrintLog 打印日志
func (h BaseHandler) PrintLog(log string) func(*gin.Context) {

	return func(c *gin.Context) {
		fmt.Println(log)
	}
}

// UnmarshalPost Unmarshal struct from Post Content to data v
func (h BaseHandler) UnmarshalPost(c *gin.Context, v interface{}) error {

	var err error
	// 支持读c.Request.Body多次
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	err = myjson.NewDecoder(bytes.NewReader(bodyBytes)).Decode(v)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err != nil {
		rm := ReturnMessage{
			Success:      false,
			ErrorCode:    HttpErrorCode.InvalidQueryParameter,
			ErrorMessage: fmt.Sprintf("Post Data Should Be JSON Map Format : %s", err.Error()),
		}
		c.AbortWithStatusJSON(http.StatusOK, rm)
		return err
	}
	err = validate.Struct(v)

	if err != nil {
		rm := ReturnMessage{
			Success:      false,
			ErrorCode:    HttpErrorCode.InvalidQueryParameter,
			ErrorMessage: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusOK, rm)
		return err
	}
	return nil
}

//SendSuccess 发送成功数据
func (h BaseHandler) SendSuccess(c *gin.Context, data interface{}) {
	rm := ReturnMessage{
		Success: true,
		Data:    data,
	}
	c.AbortWithStatusJSON(http.StatusOK, rm)
}

//SendFailure 发送失败数据
func (h BaseHandler) SendFailure(c *gin.Context, errcode string, err error) {

	rm := ReturnMessage{
		Success:      false,
		ErrorCode:    errcode,
		ErrorMessage: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusOK, rm)
}

// ValidatePostJSON middleware for check post data  if is map type
func ValidatePostJSON(c *gin.Context) {
	var v map[string]interface{}
	var err error
	if c.Request.Method == http.MethodPost {
		err = c.BindJSON(&v)
		if err != nil {
			rm := ReturnMessage{
				Success:      false,
				ErrorCode:    HttpErrorCode.InvalidQueryParameter,
				ErrorMessage: fmt.Sprintf("Post Data Should Be JSON Map Format : %s", err.Error()),
			}
			c.AbortWithStatusJSON(http.StatusOK, rm)
		}
	}
}
