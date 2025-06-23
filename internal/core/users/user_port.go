package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kayceeDev/caspa-events/internal/shared/domains"
)

type UserRepository interface {
	CreateUser(*gin.Context, *CreateUserDTO ) error
	GetUserByUUID(*gin.Context, uuid.UUID) (*domains.UserResponse, error)
}

type UserService interface {
	CreateUser(*gin.Context) error
	GetUserByUUID(*gin.Context) error
}
type UserPort interface { 
	UserRepository
}