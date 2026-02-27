package users

import "errors"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	errInvalidEmail      = errors.New("invalid email")
	errUserAlreadyExists = errors.New("user already exists")
)

func (u *User) Validate() error {
	if u.Email == "" {
		return errInvalidEmail
	}
	return nil
}
