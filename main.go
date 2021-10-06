package main

import (
	"github.com/gin-gonic/gin"
	controller "go_projct/app/controller"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./templates")
	router.LoadHTMLGlob("templates/html/*")

	var loginService controller.LoginService = controller.StaticLoginService()
	var jwtService controller.JWTService = controller.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"title": "Login Page",
			},
		)
	})

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	err := router.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
