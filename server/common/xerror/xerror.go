package xerror

import (
	"fmt"
)

// 一般业务错误
var errMsgMap map[uint32]string

// 数据库底层错误
var sysErrMsgMap map[uint32]string

// BisErr 自定义业务错误
type BisErr struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 获取错误码
func (e *BisErr) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 获取错误信息
func (e *BisErr) GetErrMsg() string {
	return e.errMsg
}

// Error 实现error接口
func (e *BisErr) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *BisErr {
	return &BisErr{errCode: errCode, errMsg: errMsg}
}

// NewErrCode
//
//	@Description: 使用ErrCode 创建 bisErr
//	@param errCode
//	@return *BisErr
func NewErrCode(errCode uint32) *BisErr {
	return &BisErr{errCode: errCode, errMsg: GetErrMsg(errCode)}
}

// NewErrMsg
//
//	@Description: 创建通用 BisErr（SERVER_ERROR）
//	@param errMsg 错误消息
//	@return *BisErr
func NewErrMsg(errMsg string) *BisErr {
	return &BisErr{errCode: SERVER_ERROR, errMsg: errMsg}
}

// 错误枚举
const (
	// OK 成功返回
	OK uint32 = 200
	// SERVER_ERROR 系统一般错误
	SERVER_ERROR uint32 = 10010
	// REQUEST_PARAM_ERROR 请求参数错误
	REQUEST_PARAM_ERROR uint32 = 10020
	// TOKEN_EXPIRE_ERROR Token过期
	TOKEN_EXPIRE_ERROR uint32 = 10030

	// 数据库错误
	DB_DELETE_ERR uint32 = 10051
	DB_INSERT_ERR uint32 = 10052
	DB_UPDATE_ERR uint32 = 10053
	DB_SEARCH_ERR uint32 = 10054

	// AUTH_CHECK_FAILURE 未授权
	AUTH_CHECK_FAILURE = 401
)

func init() {
	errMsgMap = make(map[uint32]string)
	errMsgMap[OK] = "success"
	errMsgMap[SERVER_ERROR] = "服务器开小差啦,稍后再来试一试"
	errMsgMap[REQUEST_PARAM_ERROR] = "参数错误"
	errMsgMap[TOKEN_EXPIRE_ERROR] = "token过期，请重新登陆"
	errMsgMap[AUTH_CHECK_FAILURE] = "用户未授权"

	sysErrMsgMap = make(map[uint32]string)
	sysErrMsgMap[DB_DELETE_ERR] = "DB删除失败"
	sysErrMsgMap[DB_INSERT_ERR] = "DB插入失败"
	sysErrMsgMap[DB_UPDATE_ERR] = "DB更新失败"
	sysErrMsgMap[DB_SEARCH_ERR] = "DB查询错误"
}

// GetErrMsg
//
//	@Description: 获取默认错误码信息
//	@param code
//	@return string
func GetErrMsg(code uint32) string {
	if msg, ok := errMsgMap[code]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

// IsBisCodeErr
//
//	@Description: 是否业务错误码错误
//	@param code
//	@return bool
func IsBisCodeErr(code uint32) bool {
	if _, ok := errMsgMap[code]; ok {
		return true
	} else {
		return false
	}
}

// IsSysCodeErr
//
//	@Description: 是否系统底层错误码错误
//	@param code
//	@return bool
func IsSysCodeErr(code uint32) bool {
	if _, ok := sysErrMsgMap[code]; ok {
		return true
	} else {
		return false
	}
}
