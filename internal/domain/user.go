package domain

// User 领域对象
type User struct {
	ID          int64
	Email       string
	Password    string
	CreatedTime int64
	Nickname    string
	Phone       string
	Birthday    int64
	AboutMe     string
}
