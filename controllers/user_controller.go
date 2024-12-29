package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/romanmufid16/go-auth-learn/dto"
	"github.com/romanmufid16/go-auth-learn/service"
	"github.com/romanmufid16/go-auth-learn/utils"
	"net/http"
	"strconv"
)

var validate = validator.New()

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (h *UserController) RegisterController(c *gin.Context) {
	var userDto dto.RegisterUser

	errDto := c.ShouldBindJSON(&userDto)
	if errDto != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errDto.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errValidation := validate.Struct(userDto)
	if errValidation != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errValidation.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	result, err := h.UserService.Register(&userDto)
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Invalid Request", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := utils.BuildResponse(true, "User Registered Successfully", result)
	c.JSON(http.StatusCreated, response)

}

func (h *UserController) LoginController(c *gin.Context) {
	var userDto dto.LoginUser
	errDto := c.ShouldBindJSON(&userDto)
	if errDto != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errDto.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errValidation := validate.Struct(userDto)
	if errValidation != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errValidation.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	result, err := h.UserService.Login(&userDto)
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Invalid Request", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := utils.BuildResponse(true, "Login Successful", result)
	c.JSON(http.StatusOK, response)
}

func (h *UserController) GetUserInfo(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		errorResponse := utils.BuildErrorResponse("Unauthorized", "User ID not found in context", utils.EmptyObj{})
		c.JSON(http.StatusUnauthorized, errorResponse)
		return
	}

	// Lakukan type assertion untuk memastikan id adalah uint
	userID, ok := id.(uint)
	if !ok {
		errorResponse := utils.BuildErrorResponse("Invalid Request", "Invalid user ID type", utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Panggil UserService untuk mendapatkan data user berdasarkan userID
	result, err := h.UserService.GetUser(userID)
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Invalid Request", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Response sukses jika user ditemukan
	response := utils.BuildResponse(true, "Retrieved user data successfully", result)
	c.JSON(http.StatusOK, response)
}

func (h *UserController) UpdateUser(c *gin.Context) {
	id, _ := c.Get("user_id")

	var userDto dto.UpdateUser
	errDto := c.ShouldBindJSON(&userDto)
	if errDto != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errDto.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errValidation := validate.Struct(userDto)
	if errValidation != nil {
		errorResponse := utils.BuildErrorResponse("Invalid input", errValidation.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userId, ok := id.(uint)
	if !ok {
		errorResponse := utils.BuildErrorResponse("Invalid Request", "Invalid user ID type", utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	result, err := h.UserService.UpdateUser(userId, &userDto)
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Invalid Request", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := utils.BuildResponse(true, "Update user success", result)

	c.JSON(http.StatusOK, response)

}

func (h *UserController) DeleteUser(c *gin.Context) {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")

	// Konversi ID dari string ke uint
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Invalid Request", "ID must be a number", utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Panggil service untuk menghapus user
	success, err := h.UserService.DeleteUser(uint(id))
	if err != nil {
		errorResponse := utils.BuildErrorResponse("Failed to delete user", err.Error(), utils.EmptyObj{})
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// Kirimkan response berhasil
	response := utils.BuildResponse(success, "User deleted successfully", utils.EmptyObj{})
	c.JSON(http.StatusOK, response)
}
