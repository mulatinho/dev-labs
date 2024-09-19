package main

import (
	"time"

	"github.com/mulatinho/golabs/littlegame/character"
	"github.com/mulatinho/golabs/littlegame/engine"
)

const (
	PACKAGE_NAME    = "LittleGame"
	PACKAGE_VERSION = 1.0
)

func ExampleGameTwo() {
	playerOne := character.NewPlayer()
	playerTwo := character.NewPlayer()
	playerThree := character.NewPlayer()

	levelOne := engine.NewLevel(1, "The Beginning")
	sceneOne := engine.NewScene(levelOne.Id, "Me vs You")
	turnOne := engine.NewTurn(engine.TURN_MODE_SINGLE_CHALLENGE)
	turnOne.AddPlayer(playerOne)
	turnOne.AddPlayer(playerTwo)
	sceneOne.AddTurn(turnOne)
	levelOne.AddScene(sceneOne)

	levelTwo := engine.NewLevel(2, "The Combat")
	sceneTwo := engine.NewScene(levelTwo.Id, "Against Two Junkies")
	turnTwo := engine.NewTurn(engine.TURN_MODE_COMBAT)
	turnTwo.AddPlayer(playerOne)
	turnTwo.AddPlayer(playerTwo)
	turnTwo.AddPlayer(playerThree)
	sceneTwo.AddTurn(turnTwo)
	levelTwo.AddScene(sceneTwo)

	levels := []*engine.Level{levelOne, levelTwo}
	gameParams := &engine.GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 250,
	}
	engine.GameLoop(gameParams)
}

func ExampleGameOne() {
	levelOne := engine.NewLevel(1, "The Beginning")
	levelTwo := engine.NewLevel(2, "The Combat")

	levelOne.NewScene("Scene 01", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", engine.SCENE_DIFFICULT_EASY)

	levels := []*engine.Level{levelOne, levelTwo}
	gameParams := &engine.GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 250,
	}
	engine.GameLoop(gameParams)
}

func main() {
	ExampleGameOne()
	ExampleGameTwo()
}
