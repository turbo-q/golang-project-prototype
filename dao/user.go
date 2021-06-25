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

var User = userDao{db.GetDBDefault().Model(model.User{})}

type userDao struct {
	M *gorm.DB
}

//	通用的db
func (d *userDao) Fields(fields ...string) *gorm.DB {
	return d.M.Select(fields)
}

// 如果需要select 则使用Fields更加通用
func (d *userDao) FindById(id string) (user model.User, g *gorm.DB) {
	g = d.M.Where("id=?", id).
		First(&user)
	return
}
