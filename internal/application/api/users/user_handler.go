package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kayceeDev/caspa-events/internal/core/users"
)

type UserHandler struct {
	user users.UserService
}

func NewUserHandler(userService users.UserService) *UserHandler {
	return &UserHandler{
		user: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	h.user.CreateUser(c)
}
