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
	"fmt"
	"log"
)

type Task struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// QueryAllTasks retrieves an array of tasks.
//
// Returns:
//   - []Task: An array of task objects
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

// QueryTaskById retrieves a task by its ID.
//
// Parameters:
//   - id: The unique identifier of the task to retrieve.
//
// Returns:
//   - []Task: The task object specified by ID inside of an array of them.
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
