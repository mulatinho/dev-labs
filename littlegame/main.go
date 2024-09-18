package main

import (
	"fmt"

	"github.com/mulatinho/golabs/littlegame/character"
	"github.com/mulatinho/golabs/littlegame/engine"
)

const (
	PACKAGE_NAME    = "LittleGame"
	PACKAGE_VERSION = 1.0
)

func main() {
	player := character.CreateNewPlayer()

	fmt.Printf("HELLO %v\n", player)
	engine.GameLoop()
}
