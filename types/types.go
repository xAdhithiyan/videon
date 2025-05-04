package types

type UserStore interface {
	CreateUser()
	GetUserByEmail()
}
