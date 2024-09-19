package engine

import (
	"testing"

	"github.com/mulatinho/golabs/littlegame/character"
	"github.com/mulatinho/golabs/littlegame/utils"
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

func TestTurnOrder(t *testing.T) {
	level := NewLevel(levelId, "The Level One")
	playerOne := character.NewPlayer()
	playerTwo := character.NewPlayer()
	playerThree := character.NewPlayer()
	playerFour := character.NewPlayer()

	turnOne := NewTurn(TURN_MODE_SINGLE_CHALLENGE)
	turnTwo := NewTurn(TURN_MODE_COMBAT)

	level.AddScene(newScene)

	turnOne.AddPlayer(playerOne)
	turnOne.AddPlayer(playerTwo)

	turnTwo.AddPlayer(playerOne)
	turnTwo.AddPlayer(playerTwo)
	turnTwo.AddPlayer(playerThree)
	turnTwo.AddPlayer(playerFour)

	newScene.AddTurn(turnOne)
	newScene.AddTurn(turnTwo)

	turnOne.ApplyRandomOrder()
	turnTwo.ApplyRandomOrderPreBuilt([]int{1})

	mlt.Equals(t, len(turnOne.Players), 2)
	mlt.Equals(t, len(turnTwo.Players), 4)

	utils.LogMessage(utils.LOG_TYPE_DEBUG, turnOne)
	utils.LogMessage(utils.LOG_TYPE_DEBUG, turnTwo)
}
