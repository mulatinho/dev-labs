package models

import (
	"fmt"
	"log"
)

type Task struct {
	id   int    //`json:"id" `
	name string //`json:"name" validate:"required"`
}

func GetAllTasks() []Task {
	rows, err := db.Query("SELECT id, name FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.id, &task.name)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func GetTaskById(id string) []Task {
	var tasks []Task

	rows, err := db.Query("SELECT id, name FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.id, &task.name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(task)
		tasks = append(tasks, task)
	}

	return tasks
}
