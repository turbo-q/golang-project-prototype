package model

import "errors"

// http error
var (
	ErrRequest    = errors.New("server:请求数据失败")
	ErrJSONDecode = errors.New("json:decode fail")
)

// model error,自定义错误
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
