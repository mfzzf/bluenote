package domain

import "time"

// User 领域对象
type User struct {
	Email       string
	Password    string
	CreatedTime time.Time
}
