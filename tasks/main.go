package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID   int    `json:"id" `
	Name string `json:"name" validate:"required"`
}

var tasks = []task{
	{ID: 0, Name: "Do some labs"},
	{ID: 1, Name: "Train more Go"},
}

func getTasks(ctx *gin.Context) {
	fmt.Printf("%#v\n", tasks)
	ctx.JSON(http.StatusOK, tasks)
}

func getTaskById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, elem := range tasks {
		if elem.ID == id {
			ctx.IndentedJSON(http.StatusOK, elem)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{})
}

func postTasks(ctx *gin.Context) {
	var body task
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	body.ID = len(tasks)
	tasks = append(tasks, body)
	ctx.JSON(http.StatusOK, gin.H{"count": len(tasks)})
}

func main() {
	app := gin.Default()

	fmt.Println(":. mulato labs: gin crud lab01")

	app.GET("/tasks", getTasks)
	app.GET("/tasks/:id", getTaskById)
	app.POST("/tasks", postTasks)

	app.Run(":8181")
}
