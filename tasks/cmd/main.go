package cmd

import (
	"fmt"

	"github.com/mulatinho/golabs/tasks/controllers"
)

func Start() {
	// Initialize the routes
	r := controllers.SetupRouter()

	// Start the server
	fmt.Println(":. starting server on port 8181..")
	r.Run(":8181")
}
