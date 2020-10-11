package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mic3ael/pragmaticreviews/dto"
	"github.com/mic3ael/pragmaticreviews/services"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func NewLoginController(loginService services.LoginService,
	jWtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	fmt.Println("Row", err)
	if err != nil {
		return ""
	}
	fmt.Println("Credentials", credentials)
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
