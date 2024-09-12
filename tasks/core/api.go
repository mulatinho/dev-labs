package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	API_VERSION = "v1"
	API_PREFIX  = "/api/" + API_VERSION
)

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Tasks   []Task `json:"tasks"`
}

func GetTasks(ctx *gin.Context) {
	taskApp.log.Println("GET ", API_PREFIX+"/tasks")
	tasks := QueryAllTasks()
	taskApp.log.Println(tasks)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"tasks":   tasks,
	})
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("name")
	taskApp.log.Println("GET ", API_PREFIX+"/tasks/"+id)
	tasks := QueryTaskById(id)
	ctx.HTML(http.StatusOK, "task.html", gin.H{
		"title":   "GoLabs Mulatinho: Tasks",
		"tasks":   tasks,
		"visible": "d-none",
	})
}

func PostTask(ctx *gin.Context) {
	taskApp.log.Println("POST", API_PREFIX+"/tasks")
	var body []byte
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func SetupAPIs(router *gin.Engine) {

	v1 := router.Group(API_PREFIX)

	v1.GET("/tasks", GetTasks)
	v1.GET("/tasks/:id", GetTaskById)
	v1.POST("/tasks", PostTask)
}
