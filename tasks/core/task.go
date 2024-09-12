package core

import (
	"fmt"
	"log"
)

type Task struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func QueryAllTasks() []Task {
	rows, err := taskApp.db.Query("SELECT id, name FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	taskApp.log.Println(tasks)
	return tasks
}

func QueryTaskById(id string) []Task {
	var tasks []Task
	rows, err := taskApp.db.Query("SELECT id, name FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(task)
		tasks = append(tasks, task)
	}

	return tasks
}
