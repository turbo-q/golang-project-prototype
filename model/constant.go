// 全局常量定义
package model

// 状态码
const (
	SUCCESS_CODE      = 10000 + iota
	ERROR_CODE        // 未知错误
	PARAMS_ERROR_CODE // 参数错误
	TOKEN_ERROR_CODE  // token认证错误
	NOACCESS_CODE     // 未验证
	DBERROR_CODE      // DB 错误
)

// client flag
const (
	CLIENT_DEFAULT = iota
)
