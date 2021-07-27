package model

import "errors"

// error
var (
	ErrRequest    = errors.New("server:请求数据失败")    // http request fail
	ErrJSONEncode = errors.New("json:encode fail") // json marshal fail
	ErrJSONDecode = errors.New("json:decode fail") // json Unmarshal fail
	ErrUnknown    = errors.New("未知错误")
)

// model error,自定义错误类型
type ModelError struct {
	Code int    // 状态码
	Msg  string // 错误消息
	Err  error  // error
}

func (m ModelError) Error() string {
	if m.Err != nil {
		return m.Msg + m.Err.Error()
	} else {
		return m.Msg
	}
}

// db error
type DBError struct {
	Msg string
}

func (d DBError) Error() string {
	return "db:" + d.Msg
}
