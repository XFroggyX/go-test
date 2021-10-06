package controller

import (
	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jWtService   JWTService
}

func LoginHandler(loginService LoginService,
	jWtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
