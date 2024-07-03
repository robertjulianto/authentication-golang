package handlers

import (
	"fmt"
	"ims/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type roleHandler struct {
	roleService services.RoleService
}

type RoleHandler interface {
	HandleCreateRole(ctx *gin.Context)
	HandleGetAllRoles(ctx *gin.Context)
	HandleGetRoleByID(ctx *gin.Context)
	HandleUpdateRole(ctx *gin.Context)
	HandleDeleteRole(ctx *gin.Context)
}

func NewRoleHandler(roleService services.RoleService) *roleHandler {
	return &roleHandler{roleService}
}

type handleRoleRequest struct {
	RoleName string `json:"role_name" binding:"required"`
}

func (handler *roleHandler) HandleCreateRole(ctx *gin.Context) {
	var handleRoleRequest handleRoleRequest
	err := ctx.ShouldBindBodyWithJSON(&handleRoleRequest)

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

	err = handler.roleService.CreateRole(handleRoleRequest.RoleName)

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

func (handler *roleHandler) HandleGetAllRoles(ctx *gin.Context) {
	roles := handler.roleService.GetAllRoles()

	ctx.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func (handler *roleHandler) HandleGetRoleByID(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	role := handler.roleService.GetRoleByID(roleID)
	ctx.JSON(http.StatusOK, role)
}

func (handler *roleHandler) HandleUpdateRole(ctx *gin.Context) {
	var handleRoleRequest handleRoleRequest
	err := ctx.ShouldBindBodyWithJSON(&handleRoleRequest)

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

	roleID, _ := strconv.Atoi(ctx.Param("id"))
	spec := services.UpdateRoleSpec{
		ID:          roleID,
		RoleName:    handleRoleRequest.RoleName,
	}

	err = handler.roleService.UpdateRole(spec)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})

}

func (handler *roleHandler) HandleDeleteRole(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	err := handler.roleService.DeleteRoleByID(roleID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}
