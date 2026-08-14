package api

import (
	"go-landing-page/internal/repositories"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) UserHandler {
	return UserHandler{userRepository: userRepository}
}

func (h *UserHandler) Create(echo echo.Context) error {
	return nil
}
