package dao

import (
	"bluenote/internal/domain"
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail    = errors.New("user already exists")
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
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
	Nickname    string
	Phone       string `gorm:"unique"`
	Birthday    int64
	AboutMe     string
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().Unix()
	u.CreatedTime = now
	u.UpdatedTime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return ErrUserDuplicateEmail
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).First(&user, "email = ?", u.Email).Error
	return user, err
}

func (dao *UserDAO) FindById(ctx context.Context, u domain.User) (domain.User, error) {
	var user domain.User
	err := dao.db.WithContext(ctx).First(&user, "id = ?", u.ID).Error
	return user, err
}

func (dao *UserDAO) EditById(ctx context.Context, u domain.User) error {
	// Convert domain user to DAO user
	user := User{
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		AboutMe:  u.AboutMe,
	}

	err := dao.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", u.ID).
		Updates(&user).Error
	return err
}
