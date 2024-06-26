package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adhyaksasb/mycommerce-api/handler"
	"github.com/gin-gonic/gin"
)

var(
	app *gin.Engine
)

func route(r *gin.RouterGroup) {
	// Write Service
	// r.POST("/posts", handler.CreatePosts)
	// r.PUT("/posts/:id", handler.UpdatePost)
	// r.DELETE("/posts/:id", handler.DeletePost)

	// // Read Service
	// r.GET("/posts", handler.IndexPosts)
	// r.GET("/posts/:id", handler.ShowPost)

	r.GET("/ping", handler.Ping)

	r.GET("/err", handler.ErrRouter)


	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"message": "name not found",
			})
		} else {
			c.JSON(200, gin.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	r.GET("/env", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": os.Getenv("DB_URL"),
		})
	})
	r.GET("/user/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": gin.H{
				"id": c.Param("id"),
			},
		})
	})
}

func init() {
	// handler.LoadEnvVariables()
	// handler.ConnectToDB()
	app = gin.New()
	r := app.Group("/api")
	route(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
