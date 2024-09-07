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
	tasks := models.GetAllTasks()
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
