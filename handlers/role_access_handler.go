package handlers

import (
	"fmt"
	"ims/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type roleAccessHandler struct {
	roleAccessService services.RoleAccessService
}

type RoleAccessHandler interface {
	HandleCreateRoleAccess(ctx *gin.Context)
	HandleDeleteRoleAccess(ctx *gin.Context)
	HandleGetAccessesByRole(ctx *gin.Context)
}

func NewRoleAccessHandler(roleAccessService services.RoleAccessService) *roleAccessHandler {
	return &roleAccessHandler{roleAccessService}
}

type roleAccessRequest struct {
	RoleID    int   `json:"role_id" binding:"required"`
	AccessIDs []int `json:"access_ids" binding:"required"`
}

func (handler *roleAccessHandler) HandleCreateRoleAccess(ctx *gin.Context) {
	var roleAccessRequest roleAccessRequest
	err := ctx.ShouldBindBodyWithJSON(&roleAccessRequest)

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

	err = handler.roleAccessService.CreateRoleAccess(
		roleAccessRequest.RoleID,
		roleAccessRequest.AccessIDs,
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

func (handler *roleAccessHandler) HandleDeleteRoleAccess(ctx *gin.Context) {
	var roleAccessRequest roleAccessRequest
	err := ctx.ShouldBindBodyWithJSON(&roleAccessRequest)

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

	err = handler.roleAccessService.DeleteRoleAccess(
		roleAccessRequest.RoleID,
		roleAccessRequest.AccessIDs,
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

func (handler *roleAccessHandler) HandleGetAccessesByRole(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("role_id"))
	accesses := handler.roleAccessService.GetAccessByRoleID(roleID)

	ctx.JSON(http.StatusOK, gin.H{
		"accesses": accesses,
	})
}
