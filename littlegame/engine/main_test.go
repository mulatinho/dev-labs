package engine

import (
	"testing"
	"time"

	"github.com/mulatinho/golabs/littlegame/character"
	"github.com/mulatinho/golabs/littlegame/utils"
	"github.com/mulatinho/golabs/mlt"
)

func TestGameLoop(t *testing.T) {
	levelOne := NewLevel(1, "The Beginning")
	levelTwo := NewLevel(2, "The Combat")

	playerOne := character.NewPlayer()
	playerTwo := character.NewPlayer()
	playerThree := character.NewPlayer()
	playerFour := character.NewPlayer()

	levelOne.NewScene("Scene 01", SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 03", SCENE_DIFFICULT_MEDIUM)

	levelOne.Scenes[0].NewTurn(TURN_MODE_COMBAT)
	levelOne.Scenes[0].Turns[0].AddPlayer(playerOne)
	levelOne.Scenes[0].Turns[0].AddPlayer(playerTwo)

	levelTwo.Scenes[1].NewTurn(TURN_MODE_COMBAT)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerOne)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerTwo)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerThree)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerFour)

	levels := []*Level{levelOne, levelTwo}
	gameParams := &GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 150,
	}

	timeStart := time.Now()
	GameLoop(gameParams)
	timeDuration := time.Since(timeStart)
	mlt.Assert(t, timeDuration < (1_000_000_000*15))
	utils.LogMessage(utils.LOG_TYPE_DEBUG, float64(timeDuration/1_000_000_000))
}
