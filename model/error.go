package model

import "errors"

// http error
var (
	ErrRequest    = errors.New("server:请求数据失败")
	ErrJSONDecode = errors.New("json:decode fail")
)
