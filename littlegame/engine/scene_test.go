package engine

import (
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

var (
	sceneName string = "The Beginning"
	newScene  *Scene = NewScene(levelId, sceneName)
)

func TestNewScene(t *testing.T) {
	mlt.Equals(t, newScene.Name, sceneName)
}

func TestSceneHelpers(t *testing.T) {
	mlt.Equals(t, newScene.GetName(), sceneName)
	mlt.Equals(t, newScene.GetLevel(), levelId)
}

func TestScene(t *testing.T) {
	level := NewLevel(levelId, "The Level One")

	level.AddScene(newScene)

	mlt.Equals(t, len(level.Scenes), 1)
	mlt.Equals(t, level.Scenes[0].LevelId, level.Id)

	newScene.NewTurn(TURN_MODE_SINGLE_CHALLENGE)
	newScene.NewTurn(TURN_MODE_COMBAT)

	mlt.Equals(t, len(newScene.Turns), 2)
}
