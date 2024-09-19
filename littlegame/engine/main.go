package engine

import (
	"fmt"
	"time"
)

const (
	ENGINE_CYCLE_TIME = time.Millisecond * 200
	LOOP_MSG          = "loop is running"
)

func GameLoop(levelList []Level) {
	for levelIndex := 0; levelIndex < len(levelList); levelIndex++ {
		level := levelList[levelIndex]

		loopEvent := 0
		subLoopEvent := 0

		fmt.Printf(":. %#v\n", level)

		if level.isCompleted {
			fmt.Printf(":. LEVEL FINISHED => %#v\n", level.Id)
			break
		}
		for sceneIndex := 0; sceneIndex < len(level.Scenes); sceneIndex++ {
			scene := level.Scenes[sceneIndex]

			for {
				fmt.Printf(":.:. %#v\n", scene)
				if scene.isCompleted {
					fmt.Printf(":.:. LEVEL %#v, SCENE FINISHED => %#v\n", level.Id, scene.Id)
					break
				}

				if subLoopEvent > 5 {
					scene.isCompleted = true
				}
				subLoopEvent++
				time.Sleep(ENGINE_CYCLE_TIME)
			}
		}

		if loopEvent > 2 {
			level.isCompleted = true
		}
		loopEvent++
		time.Sleep(ENGINE_CYCLE_TIME)
	}
}
