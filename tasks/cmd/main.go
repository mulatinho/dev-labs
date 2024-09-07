package main

import (
	"fmt"

	"github.com/mulatinho/golabs/tasks/routes"
)

func main() {
	// Initialize the routes
	r := routes.SetupRouter()

	// Start the server
	fmt.Println(":. starting server on port 8181..")
	r.Run(":8181")
}
