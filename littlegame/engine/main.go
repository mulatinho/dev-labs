package engine

import (
	"time"

	"github.com/mulatinho/golabs/littlegame/utils"
)

const (
	ENGINE_CYCLE_TIME = time.Millisecond * 200
	LOOP_MSG          = "loop is running"
)

type GameLoopParameters struct {
	LevelList []*Level
	LoopTime  time.Duration
}

func GameLoop(GameLoopParams *GameLoopParameters) {
	for levelIndex := 0; levelIndex < len(GameLoopParams.LevelList); levelIndex++ {
		level := GameLoopParams.LevelList[levelIndex]

		loopEvent := 0
		subLoopEvent := 0
		turnLoopEvent := 0

		utils.LogMessage(utils.LOG_TYPE_DEBUG, level)

		if level.isCompleted {
			break
		}
		for sceneIndex := 0; sceneIndex < len(level.Scenes); sceneIndex++ {
			scene := level.Scenes[sceneIndex]
			sceneMagicNumber := utils.GetRandomNumber(5)
			turnMagicNumber := utils.GetRandomNumber(10)

			for {
				utils.LogMessage(utils.LOG_TYPE_DEBUG, scene)

				if scene.isCompleted {
					break
				} else if len(scene.Turns) == 0 {
					break
				}

				for turnIndex := 0; turnIndex < len(scene.Turns); turnIndex++ {
					turnAction := scene.Turns[turnIndex]
					turnAction.ApplyRandomOrder()
					utils.LogMessage(utils.LOG_TYPE_DEBUG, turnAction)

					for orderIndex := 0; orderIndex < len(turnAction.TurnOrder); orderIndex++ {
						playerIndex := turnAction.TurnOrder[orderIndex]
						utils.LogMessage(utils.LOG_TYPE_DEBUG, turnAction.Players[playerIndex])
					}

					if turnLoopEvent > turnMagicNumber {
						turnAction.isCompleted = true
						turnLoopEvent = 0
					}

					turnLoopEvent++
					time.Sleep(GameLoopParams.LoopTime)
				}

				if subLoopEvent > sceneMagicNumber {
					scene.isCompleted = true
					subLoopEvent = 0
				}

				subLoopEvent++
				time.Sleep(GameLoopParams.LoopTime)
			}
		}

		if loopEvent > 2 {
			level.isCompleted = true
			loopEvent = 0
		}
		loopEvent++
		time.Sleep(GameLoopParams.LoopTime)
	}
}
