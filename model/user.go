package model

import (
	"golang-project-prototype/model/internal"
)

// User struct for table user
type User internal.User

// **********************************
// api
// **********************************
type UserApiUpdateReq struct {
	UserID       string `form:"user_id" valid:"Required"`
	UserPassword string `form:"password" valid:"Required"`
}

// **********************************
// service
// **********************************
// service查询列表
type UserServiceGetListReq struct {
	Page int
	Size int
}

// service查询列表结果
type UserServiceGetListRes struct {
	List []UserListItem `json:"list"`
	// 可以对internal.User 进行扩展
	internal.User
	Password string `json:"-"`
}

// **********************************
// 通用数据结构
// **********************************
type UserListItem struct {
}
