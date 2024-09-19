package engine

import (
	"math/rand"
	"time"

	"github.com/mulatinho/golabs/littlegame/character"
	"github.com/mulatinho/golabs/littlegame/utils"
)

type (
	SceneDifficult int
	TurnMode       int
)

const (
	SCENE_DIFFICULT_ZERO SceneDifficult = iota
	SCENE_DIFFICULT_EASY
	SCENE_DIFFICULT_MEDIUM
	SCENE_DIFFICULT_HARD
	SCENE_DIFFICULT_TOTAL
)

const (
	TURN_MODE_SINGLE_CHALLENGE = iota
	TURN_MODE_MULTIPLE_CHALLENGE
	TURN_MODE_COMBAT
	TURN_MODE_TOTAL
)

type Turn struct {
	Id          int
	SceneId     int
	Mode        TurnMode
	isCompleted bool
	timestamp   int64
	Players     []*character.Player
	TurnOrder   []int
}

type Scene struct {
	Id          int
	LevelId     int
	Name        string
	Difficult   SceneDifficult
	Turns       []*Turn
	SceneTurn   *Turn
	isCompleted bool
}

func NewScene(level int, name string) *Scene {
	SceneId := utils.GetRandomNumber(10_000)
	return &Scene{
		Id:          SceneId,
		LevelId:     level,
		Name:        name,
		Difficult:   SceneDifficult(rand.Int() % int(SCENE_DIFFICULT_TOTAL)),
		Turns:       []*Turn{},
		SceneTurn:   &Turn{},
		isCompleted: false,
	}
}

func (s *Scene) GetID() int      { return s.Id }
func (s *Scene) GetLevel() int   { return s.LevelId }
func (s *Scene) GetName() string { return s.Name }

func (s *Scene) AddTurn(turn *Turn) {
	s.SceneTurn = turn
	s.Turns = append(s.Turns, turn)
}

func (s *Scene) NewTurn(mode TurnMode) {
	newTurn := &Turn{
		Id:          len(s.Turns) + 1,
		SceneId:     s.Id,
		Mode:        mode,
		isCompleted: false,
		timestamp:   time.Now().Unix(),
		Players:     []*character.Player{},
		TurnOrder:   []int{},
	}

	s.AddTurn(newTurn)
}

func (s *Scene) DeleteTurnById(turnId int) bool {
	turnFound := false
	for o, turn := range s.Turns {
		if turn.Id == turnId {
			s.Turns = append(s.Turns[:o], s.Turns[o+1:]...)
			turnFound = true
		}
	}

	return turnFound
}

func NewTurn(mode TurnMode) *Turn {
	return &Turn{
		Id:          -1,
		SceneId:     -1,
		Mode:        mode,
		isCompleted: false,
		timestamp:   time.Now().Unix(),
		Players:     []*character.Player{},
	}
}

func (t *Turn) SetId(id int)           { t.Id = id }
func (t *Turn) SetSceneId(sceneId int) { t.SceneId = sceneId }

func (t *Turn) AddPlayer(player *character.Player) {
	t.Players = append(t.Players, player)
}

func (t *Turn) DeletePlayerById(playerId int) bool {
	playerFound := false
	for i, player := range t.Players {
		if player.Id == playerId {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			playerFound = true
		}
	}

	return playerFound
}

func (t *Turn) ApplyRandomOrder() {
	t.TurnOrder = utils.ApplyRandomIndexesToArray(len(t.Players))
}

func (t *Turn) ApplyRandomOrderPreBuilt(prebuilt []int) {
	t.TurnOrder = utils.ApplyRandomIndexesToArrayPreBuilt(len(t.Players), prebuilt)
}
