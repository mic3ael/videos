package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mic3ael/pragmaticreviews/controllers"
	"github.com/mic3ael/pragmaticreviews/middlewares"
	"github.com/mic3ael/pragmaticreviews/services"
)

var (
	videoConroller  controllers.VideoController = controllers.New(services.New())
	loginController controllers.LoginController = controllers.NewLoginController(services.NewLoginService(), services.NewJWTService())
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.Default()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
			return
		}

		ctx.JSON(http.StatusUnauthorized, nil)
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoConroller.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoConroller.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message: ": "OK"})
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoConroller.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
