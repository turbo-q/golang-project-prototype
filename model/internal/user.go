package internal

// 持久化结构
// User struct for table user
type User struct {
	ID       string `gorm:"type:varchar(64);not null"`
	Password string `gorm:"type:varchar(64);not null"`
	Account  string `gorm:"type:varchar(64);not null"`
}

func (User) TableName() string {
	return "t_user"
}
