package handlers

import (
	"fmt"
	"ims/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type accessHandler struct {
	accessService services.AccessService
}

type AccessHandler interface {
	HandleCreateAccess(ctx *gin.Context)
	HandleGetAllAccesses(ctx *gin.Context)
	HandleGetAccessByID(ctx *gin.Context)
	HandleGetAccessByName(ctx *gin.Context)
	HandleGetAccessByCode(ctx *gin.Context)
	HandleDeleteAccess(ctx *gin.Context)
}

func NewAccessHandler(accessService services.AccessService) *accessHandler {
	return &accessHandler{accessService}
}

type handleCreateAcessRequest struct {
	AccessName string `json:"access_name" binding:"required"`
	AccessCode string `json:"access_code" binding:"required"`
}

func (handler *accessHandler) HandleCreateAccess(ctx *gin.Context) {
	var handleCreateAcessRequest handleCreateAcessRequest
	err := ctx.ShouldBindBodyWithJSON(&handleCreateAcessRequest)

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

	err = handler.accessService.CreateAccess(handleCreateAcessRequest.AccessName, handleCreateAcessRequest.AccessCode)

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

func (handler *accessHandler) HandleGetAllAccesses(ctx *gin.Context) {
	accesses := handler.accessService.GetAllAccesses()

	ctx.JSON(http.StatusOK, gin.H{
		"accesses": accesses,
	})
}

func (handler *accessHandler) HandleGetAccessByID(ctx *gin.Context) {
	accessID, _ := strconv.Atoi(ctx.Param("id"))
	access := handler.accessService.GetAccessByID(accessID)
	if access == nil {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, access)
}

type handleGetAccessByNameRequest struct {
	AccessName string `json:"access_name" binding:"required"`
}

func (handler *accessHandler) HandleGetAccessByNameRequest(ctx *gin.Context) {
	var handleGetAccessByNameRequest handleGetAccessByNameRequest
	err := ctx.ShouldBindBodyWithJSON(&handleGetAccessByNameRequest)

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

	access := handler.accessService.GetAccessByName(handleGetAccessByNameRequest.AccessName)

	ctx.JSON(http.StatusOK, access)
}

type handleGetAccessByCodeRequest struct {
	AccessCode string `json:"access_code" binding:"required"`
}

func (handler *accessHandler) HandleGetAccessByCodeRequest(ctx *gin.Context) {
	var handleGetAccessByCodeRequest handleGetAccessByCodeRequest
	err := ctx.ShouldBindBodyWithJSON(&handleGetAccessByCodeRequest)

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

	access := handler.accessService.GetAccessByCode(handleGetAccessByCodeRequest.AccessCode)

	ctx.JSON(http.StatusOK, access)
}

func (handler *accessHandler) HandleDeleteAccess(ctx *gin.Context) {
	accessID, _ := strconv.Atoi(ctx.Param("id"))
	err := handler.accessService.DeleteAccessByID(accessID)
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
