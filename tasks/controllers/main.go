package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulatinho/golabs/tasks/models"
)

func GetIndex(ctx *gin.Context) {
	fmt.Println("GET /")
	ctx.HTML(200, "index.html", gin.H{
		"title":   "GoLabs Mulatinho: Home",
		"visible": "d-block",
	})
}

func GetTasks(ctx *gin.Context) {
	fmt.Println("GET /tasks")
	tasks := models.GetAllTasks()
	ctx.HTML(http.StatusOK, "task.html", gin.H{
		"title":   "GoLabs Mulatinho: Tasks",
		"tasks":   tasks,
		"visible": "d-none",
	})
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("name")
	fmt.Println("GET /tasks/:id", id)
	tasks := models.GetTaskById(id)
	ctx.HTML(http.StatusOK, "task.html", gin.H{
		"title":   "GoLabs Mulatinho: Tasks",
		"tasks":   tasks,
		"visible": "d-none",
	})
}

func PostTask(ctx *gin.Context) {
	var body models.Task
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/**")
	r.Static("/static", "./static")

	// Initialize the database
	models.InitDB()

	// Define routes
	r.GET("/", GetIndex)
	r.GET("/tasks", GetTasks)
	r.GET("/tasks/:id", GetTaskById)
	r.POST("/tasks", PostTask)

	return r
}
