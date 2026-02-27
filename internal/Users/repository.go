package users

type userCreator interface {
	CreateUser(user *User) error
}

type userByEmailFinder interface {
	GetUserByEmail(email string) (*User, error)
}
