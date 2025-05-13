package types

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(email string) (*User, error)
}
type VideoFuns interface {
	ParseData(msg []byte) (int, bool)
}

type User struct {
	Id       int
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Cookie struct {
	Name  string
	Value string
}
