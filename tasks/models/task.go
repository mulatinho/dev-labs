package models

import "log"

type Task struct {
	id   int    `json:"id" `
	name string `json:"name" validate:"required"`
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

func GetTaskById(id int) Task {
	var task Task

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
	}

	return task
}
