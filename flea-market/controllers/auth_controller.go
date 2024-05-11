package controllers

import (
	"flea-market/dto"
	"flea-market/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var newUserInput dto.SignUpUserInput
	if err := ctx.ShouldBindJSON(&newUserInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.SignUp(newUserInput.Email, newUserInput.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
