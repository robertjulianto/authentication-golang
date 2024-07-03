package handlers

import (
	"fmt"
	"ims/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService services.UserService
}

type UserHandler interface {
	HandleCreateUser(ctx *gin.Context)
	HandleGetAllUsers(ctx *gin.Context)
	HandleGetUserByID(ctx *gin.Context)
	HandleUpdateUser(ctx *gin.Context)
	HandleDeleteUser(ctx *gin.Context)
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

type handleUserRequest struct {
	UserName    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}

func (handler *userHandler) HandleCreateUser(ctx *gin.Context) {
	var handleUserRequest handleUserRequest
	err := ctx.ShouldBindBodyWithJSON(&handleUserRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	err = handler.userService.CreateUser(
		handleUserRequest.UserName,
		handleUserRequest.Password,
		handleUserRequest.DisplayName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})

}

func (handler *userHandler) HandleGetAllUsers(ctx *gin.Context) {
	users := handler.userService.GetAllUsers()

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (handler *userHandler) HandleGetUserByID(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	user := handler.userService.GetUserByID(userID)
	ctx.JSON(http.StatusOK, user)
}

func (handler *userHandler) HandleUpdateUser(ctx *gin.Context) {
	var handleUserRequest handleUserRequest
	err := ctx.ShouldBindBodyWithJSON(&handleUserRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	userID, _ := strconv.Atoi(ctx.Param("id"))
	spec := services.UpdateUserSpec{
		ID:          userID,
		UserName:    handleUserRequest.UserName,
		Password:    handleUserRequest.Password,
		DisplayName: handleUserRequest.DisplayName,
	}

	err = handler.userService.UpdateUser(spec)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}

func (handler *userHandler) HandleDeleteUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	err := handler.userService.DeleteUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}
