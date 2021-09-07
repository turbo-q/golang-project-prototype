package dao

import (
	"golang-project-prototype/library/util/db"
	"golang-project-prototype/model"

	"gorm.io/gorm"
)

/*
问题：
	1. 每次都要select *
	2. 每次都要写model
	3. dao层应该只写通用的增删查改
*/

var User = userDao{}

type userDao struct{}

//	通用的db
func (d *userDao) Fields(fields ...string) *gorm.DB {
	return db.GetDBDefault().Select(fields)
}

// 如果需要select 则使用Fields更加通用
func (d *userDao) FindById(id string) (user model.User, g *gorm.DB) {
	g = db.GetDBDefault().Where("id=?", id).
		Find(&user)
	return
}
