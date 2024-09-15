/*
 * The golabs/tasks application.
 *
 * Copyright (C) 2024 Alexandre Mulatinho <alex@mulatinho.net>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either  version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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

// GetTasks retrieves a JSON of []Task's
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

// GetTaskById retrieves a JSON of []Task's containing the specific ID
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
