package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kayceeDev/caspa-events/lib"
)

type userService struct {
	userRepository UserRepository
	logger         *lib.CustomLogger
}

var _ UserService = (*userService)(nil)

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateUser(c *gin.Context) error {
	dto := &CreateUserDTO{}
	if err := lib.ReadJSON(c, dto); err != nil {
		lib.JSON(c, "ERROR_4001: Invalid request body", http.StatusBadRequest, err, u.logger)
		return err
	}
	err := u.userRepository.CreateUser(c, dto)
	if err != nil {
		lib.JSON(c, "ERROR_4004: Failed to create user", http.StatusInternalServerError, err, u.logger)
		return err
	}
	lib.JSON(c, "User created successfully", http.StatusCreated, nil, u.logger)
	return nil
}

func (u *userService) GetUserByUUID(c *gin.Context) error {
	userUUID, err := lib.ParseNamedUUIDParam("user_uuid", c)
	if err != nil {
		lib.JSON(c, "ERROR_4002: Invalid ID passed", http.StatusUnauthorized, err, u.logger)
		return err
	}
	uuid, err := uuid.Parse(userUUID.String())
	if err != nil {
		lib.JSON(c, "ERROR_4002: Invalid UUID format", http.StatusBadRequest, err, u.logger)
		return err
	}
	result, err := u.userRepository.GetUserByUUID(c, uuid)
	if err != nil {
		lib.JSON(c, "ERROR_4003: User not found", http.StatusNotFound, err, u.logger)
		return err
	}
	lib.JSON(c, "User retrieved successfully", http.StatusOK, result, u.logger)
	return nil
}
