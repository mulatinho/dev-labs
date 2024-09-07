package routes

import (
	"github.com/mulatinho/golabs/tasks/controllers"
	"github.com/mulatinho/golabs/tasks/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	r.Static("/static", "./static")

	// Initialize the database
	models.InitDB()

	// Define routes
	r.GET("/", controllers.GetIndex)
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.PostTask)

	return r
}
