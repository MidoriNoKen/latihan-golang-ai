package controller

import (
	"net/http"
	"strconv"

	"github.com/MidoriNoKen/latihan-golang-ai/domain"
	"github.com/MidoriNoKen/latihan-golang-ai/pkg/response"
	"github.com/gin-gonic/gin"
)

// UserController handles user incoming HTTP requests
type UserController struct {
	userService domain.UserService
}

// NewUserController instantiates a new UserController with its dependencies
func NewUserController(userService domain.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterInput represents the request payload structure for user registration
type RegisterInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// Register handles user registration POST /users
func (ctrl *UserController) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.JSONError(c, http.StatusBadRequest, "Invalid input request payload", err.Error())
		return
	}

	user, err := ctrl.userService.Register(c.Request.Context(), input.Name, input.Email)
	if err != nil {
		response.JSONError(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	response.JSONSuccess(c, http.StatusCreated, "User registered successfully", user)
}

// GetAll handles fetching all users GET /users
func (ctrl *UserController) GetAll(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		response.JSONError(c, http.StatusInternalServerError, "Failed to fetch users list", err.Error())
		return
	}

	response.JSONSuccess(c, http.StatusOK, "Users list retrieved successfully", users)
}

// GetByID handles fetching a user by their ID GET /users/:id
func (ctrl *UserController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.JSONError(c, http.StatusBadRequest, "Invalid user ID format", "ID must be a positive integer")
		return
	}

	user, err := ctrl.userService.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		response.JSONError(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	response.JSONSuccess(c, http.StatusOK, "User detail retrieved successfully", user)
}
