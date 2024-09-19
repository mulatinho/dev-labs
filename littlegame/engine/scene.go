package engine

import (
	"math/rand"
	"time"

	"github.com/mulatinho/golabs/littlegame/character"
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
	Villains    []*character.Villain
}

type Scene struct {
	Id          int
	LevelId     int
	Name        string
	Difficult   SceneDifficult
	Turns       []*Turn
	isCompleted bool
}

func NewScene(level int, name string) *Scene {
	SceneId := rand.Intn(10_000) // this should get the last index of scenes
	return &Scene{
		Id:        SceneId,
		LevelId:   level,
		Name:      name,
		Difficult: SceneDifficult(rand.Int() % int(SCENE_DIFFICULT_TOTAL)),
		Turns:     []*Turn{},
	}
}

func (s *Scene) GetID() int      { return s.Id }
func (s *Scene) GetLevel() int   { return s.LevelId }
func (s *Scene) GetName() string { return s.Name }

func (s *Scene) AddTurn(turn *Turn) {
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
		Villains:    []*character.Villain{},
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
