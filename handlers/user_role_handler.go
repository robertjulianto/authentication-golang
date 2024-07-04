package handlers

import (
	"fmt"
	"ims/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userRoleHandler struct {
	userRoleService services.UserRoleService
}

type UserRoleHandler interface {
	HandleCreateUserRole(ctx *gin.Context)
	HandleDeleteUserRole(ctx *gin.Context)
	HandleGetRoleMembers(ctx *gin.Context)
}

func NewUserRoleHandler(userRoleService services.UserRoleService) *userRoleHandler {
	return &userRoleHandler{userRoleService}
}

type userRoleRequest struct {
	RoleID  int   `json:"role_id" binding:"required"`
	UserIDs []int `json:"user_ids" binding:"required"`
}

func (handler *userRoleHandler) HandleCreateUserRole(ctx *gin.Context) {
	var userRoleRequest userRoleRequest
	err := ctx.ShouldBindBodyWithJSON(&userRoleRequest)

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

	err = handler.userRoleService.CreateUserRole(
		userRoleRequest.RoleID,
		userRoleRequest.UserIDs,
	)

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

func (handler *userRoleHandler) HandleDeleteUserRole(ctx *gin.Context) {
	var userRoleRequest userRoleRequest
	err := ctx.ShouldBindBodyWithJSON(&userRoleRequest)

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

	err = handler.userRoleService.DeleteUserRole(
		userRoleRequest.RoleID,
		userRoleRequest.UserIDs,
	)

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

func (handler *userRoleHandler) HandleGetRoleMembers(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("role_id"))
	users := handler.userRoleService.GetRoleMembers(roleID)

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
