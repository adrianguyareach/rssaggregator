package users

type RegisterUserUseCase struct {
	userCreator       userCreator
	userByEmailFinder userByEmailFinder
}

func NewRegisterUserUseCase(userCreator userCreator, userByEmailFinder userByEmailFinder) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userCreator:       userCreator,
		userByEmailFinder: userByEmailFinder,
	}
}

func (u *RegisterUserUseCase) Register(name, email, password string) (*User, error) {

	existingUser, _ := u.userByEmailFinder.GetUserByEmail(email)
	if existingUser != nil {
		return nil, errUserAlreadyExists
	}

	user := &User{
		FirstName: name,
		Email:     email,
		Password:  password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := u.userCreator.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil

}
