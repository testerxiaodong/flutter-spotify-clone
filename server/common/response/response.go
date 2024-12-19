package response

import (
	"net/http"
	"server/common/xerror"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Response 通用返回对象
type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功时调用
//
//	@Description: 成功返回结构
//	@param data
//	@return *Response
func Success(data interface{}) *Response {
	return &Response{200, "OK", data}
}

// ErrResponse 错误返回对象
type ErrResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// Error 失败时调用
//
//	@Description: 错误返回结构
//	@param errCode
//	@param errMsg
//	@return *ErrResponse
func Error(errCode uint32, errMsg string) *ErrResponse {
	return &ErrResponse{errCode, errMsg}
}

func Result(c *gin.Context, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		c.JSON(http.StatusOK, r)
	} else {
		// 默认返回系统错误
		errCode := xerror.SERVER_ERROR
		errMsg := xerror.GetErrMsg(errCode)
		causeErr := errors.Cause(err)
		// err类型
		if e, ok := causeErr.(*xerror.BisErr); ok {
			// 自定义错误类型
			// 自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		}
		// 返回错误信息
		c.JSON(http.StatusBadRequest, Error(errCode, errMsg))
	}
}
