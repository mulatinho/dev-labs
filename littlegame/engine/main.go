package engine

import (
	"fmt"
	"time"
)

const (
	ENGINE_CYCLE_TIME = time.Millisecond * 500
	LOOP_MSG          = "loop is running"
)

func GameLoop() {
	for {
		fmt.Println(LOOP_MSG)
		time.Sleep(ENGINE_CYCLE_TIME)
	}
}
