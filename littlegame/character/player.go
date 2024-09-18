package character

import (
	"fmt"

	"github.com/mulatinho/golabs/littlegame/utils"
)

const (
	PACKAGE_NAME = "character"
	LOOP_MSG     = "loop is running"
)

type Player struct {
	id   int
	name string
}

func CreateNewPlayer() *Player {
	fmt.Println()

	player := &Player{
		id:   utils.GetPlayerID(),
		name: utils.GenerateName(utils.NAME_TYPE_MAN),
	}

	return player
}
