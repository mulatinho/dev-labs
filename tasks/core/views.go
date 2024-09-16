package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PageViewIndex(ctx *gin.Context) {
	taskApp.log.Println("GET /")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "GoLabs Mulatinho: Home",
		"visible": "d-none",
	})
}

func PageViewTasks(ctx *gin.Context) {
	taskApp.log.Println("GET /tasks")
	restCall := RestCall("GET", URL_API_TASKS, nil)
	if restCall.RestError != nil {
		taskApp.log.Println("call to /tasks failed.")
	}

	var data APIResponse
	json.Unmarshal(restCall.RestBody, &data)
	taskApp.log.Println(string(restCall.RestBody))
	ctx.HTML(http.StatusOK, "task.html", gin.H{
		"title":   "GoLabs Mulatinho: Tasks",
		"tasks":   data.Tasks,
		"visible": "d-none",
	})
	taskApp.log.Println(data.Tasks)
}

func PageViewTaskById(ctx *gin.Context) {
	var tasks []Task
	id := ctx.Param("name")
	fmt.Println("GET /tasks/:id", id)
	restCall := RestCall("GET", URL_API_TASKS+id, nil)
	if restCall.RestError != nil {
		taskApp.log.Println("call to /tasks failed.")
	}

	json.Unmarshal(restCall.RestBody, &tasks)
	ctx.HTML(http.StatusOK, "task.html", gin.H{
		"title":   "GoLabs Mulatinho: Task Info",
		"tasks":   tasks,
		"visible": "d-none",
	})
}

func SetupViews(router *gin.Engine) {
	possibleDirs := []string{"/../core", "/core", "./core"}
	var coreDir, templateDir, staticDir string

	for _, checkDir := range possibleDirs {
		coreDir = os.Getenv("PWD") + checkDir
		templateDir = coreDir + "/templates/**"
		staticDir = coreDir + "/static"
		if _, err := os.Stat(coreDir); err == nil {
			break
		}
	}

	fmt.Println(templateDir, staticDir)
	router.LoadHTMLGlob(templateDir)
	router.Static("/static", staticDir)

	router.GET("/", PageViewIndex)
	router.GET("/tasks", PageViewTasks)
	router.GET("/tasks/:id", PageViewTaskById)
}
