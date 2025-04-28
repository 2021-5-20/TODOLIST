package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/2021-5-20/t_d_list/handlers"
	"www.github.com/2021-5-20/t_d_list/middleware"
)

func InitRoutes(r *gin.Engine) {
	r.StaticFile("/favicon.ico", "./frontend/favicon.ico")
	r.POST("/api/register", handlers.HandlePostFormStruct)
	r.POST("/api/login", handlers.LoginHandler)
	r.LoadHTMLFiles("frontend/login.html", "frontend/memo.html", "frontend/register.html")
	r.Static("/static", "./frontend/js")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/memo", func(c *gin.Context) {
		c.HTML(200, "memo.html", nil)
	})
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/memo", handlers.GetAllMemoHandler)
		auth.GET("/memo/:id", handlers.GetMemoHandler)
		auth.POST("/memo", handlers.CreateMemoHandler)
		auth.PUT("/memo/:id", handlers.UpdateMemoHandler)
		auth.DELETE("/memo/:id", handlers.DeleteMemoHandler)
		auth.GET("/profile", handlers.ProfileHandler)
	}
}
