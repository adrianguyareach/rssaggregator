package users

import (
	"github.com/gin-gonic/gin"
)

type RegisterUserHandler struct {
	registerUserUseCase *RegisterUserUseCase
}

func NewRegisterUserHandler(u *RegisterUserUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{
		registerUserUseCase: u,
	}
}

type registerUserRequest struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *RegisterUserHandler) RegisterUser(c *gin.Context) {}
