package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// User 直接对应数据库表结构 //entity || model || Persistent Object
type User struct {
	Id          int64  `gorm:"primary_key,AUTO_INCREMENT"`
	Email       string `gorm:"unique"`
	Password    string
	CreatedTime int64
	UpdatedTime int64
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().Unix()
	u.CreatedTime = now
	u.UpdatedTime = now
	return dao.db.WithContext(ctx).Create(&u).Error
}
