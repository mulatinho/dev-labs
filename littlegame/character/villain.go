package character

import (
	"github.com/mulatinho/golabs/littlegame/utils"
)

type Villain struct {
	id     int
	name   string
	health int
}

func CreateNewVillain() *Villain {
	villain := &Villain{
		id:   utils.GetPlayerID(),
		name: utils.GenerateName(utils.NAME_TYPE_MAN),
	}

	return villain
}
