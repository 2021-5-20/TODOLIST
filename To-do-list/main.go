package main

import (
	"github.com/gin-gonic/gin"
	"www.github.com/2021-5-20/t_d_list/config"
	"www.github.com/2021-5-20/t_d_list/routes"
)

func main() {
	config.InitDB()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 或者具体 origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	routes.InitRoutes(r)
	r.Run(":3000") // listen and serve on  (for windows "localhost:3000")
}
