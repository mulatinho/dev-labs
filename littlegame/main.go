package main

import (
	"github.com/mulatinho/golabs/littlegame/engine"
)

const (
	PACKAGE_NAME    = "LittleGame"
	PACKAGE_VERSION = 1.0
)

func main() {
	levelOne := engine.NewLevel(1, "The Beginning")
	levelTwo := engine.NewLevel(2, "The Combat")

	levelOne.NewScene("Scene 01", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 03", engine.SCENE_DIFFICULT_MEDIUM)
	levelTwo.NewScene("Scene 04", engine.SCENE_DIFFICULT_HARD)

	levels := []engine.Level{*levelOne, *levelTwo}
	engine.GameLoop(levels)
}
