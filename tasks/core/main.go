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
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	Name    = "golabs/mulatinho"
	Version = "1.0"
	Port    = "8181"
	FQDN    = "http://localhost:" + Port
	Banner  = `
 ________  ________  ___       ________  ________  ________                                     
|\   ____\|\   __  \|\  \     |\   __  \|\   __  \|\   ____\                                    
\ \  \___|\ \  \|\  \ \  \    \ \  \|\  \ \  \|\ /\ \  \___|_                                   
 \ \  \  __\ \  \\\  \ \  \    \ \   __  \ \   __  \ \_____  \                                  
  \ \  \|\  \ \  \\\  \ \  \____\ \  \ \  \ \  \|\  \|____|\  \                                 
   \ \_______\ \_______\ \_______\ \__\ \__\ \_______\____\_\  \                                
    \|_______|\|_______|\|_______|\|__|\|__|\|_______|\_________\                               
                                                     \|_________|                               
                                                                                                
                                                                                                
 _____ ______   ___  ___  ___       ________  _________  ___  ________   ___  ___  ________     
|\   _ \  _   \|\  \|\  \|\  \     |\   __  \|\___   ___\\  \|\   ___  \|\  \|\  \|\   __  \    
\ \  \\\__\ \  \ \  \\\  \ \  \    \ \  \|\  \|___ \  \_\ \  \ \  \\ \  \ \  \\\  \ \  \|\  \   
 \ \  \\|__| \  \ \  \\\  \ \  \    \ \   __  \   \ \  \ \ \  \ \  \\ \  \ \   __  \ \  \\\  \  
  \ \  \    \ \  \ \  \\\  \ \  \____\ \  \ \  \   \ \  \ \ \  \ \  \\ \  \ \  \ \  \ \  \\\  \ 
   \ \__\    \ \__\ \_______\ \_______\ \__\ \__\   \ \__\ \ \__\ \__\\ \__\ \__\ \__\ \_______\
    \|__|     \|__|\|_______|\|_______|\|__|\|__|    \|__|  \|__|\|__| \|__|\|__|\|__|\|_______|

`
	URL_API_TASKS = FQDN + API_PREFIX + "/tasks"
)

// Task represents a task in the task management system.
// It contains all necessary information about the task,
// including its ID, name, and status.
//
// Fields:
//   - ID: A unique identifier for the task.
//   - Name: The name of the task.
//   - Status: The current status of the task (e.g., pending, completed).
type TaskApp struct {
	db  *sql.DB
	log *log.Logger
	gin *gin.Engine
}

var taskApp *TaskApp

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/data.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	
	INSERT OR REPLACE INTO tasks (id, name) VALUES (0, 'This is my first Task example for Tests :)');`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SetupRouter() {
	var lock = &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	fmt.Println("HEEEEEEEEEEEEEEEY ", reflect.ValueOf(taskApp).IsNil())
	if reflect.ValueOf(taskApp).IsNil() {
		taskApp = &TaskApp{
			db:  InitDB(),
			gin: gin.Default(),
			log: log.New(os.Stdout, ":. ", log.LstdFlags|log.Lshortfile),
		}

		SetupViews(taskApp.gin)
		SetupAPIs(taskApp.gin)
	}

	fmt.Println("HOOOOOOOOOOOOOOOY ", reflect.ValueOf(taskApp).IsNil())
}

func Start() {
	gin.SetMode(gin.DebugMode)
	SetupRouter()
	fmt.Println("HOOOOOOOOOOOOOOOY ", reflect.ValueOf(taskApp).IsNil())
	fmt.Println("HUUUUUUUUUUUUUUUY ", reflect.ValueOf(taskApp.log).IsNil())
	taskApp.log.Println(":. starting server on port ", Port)
	taskApp.gin.Run(":" + Port)
}
