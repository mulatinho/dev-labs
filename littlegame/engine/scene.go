package engine

import (
	"math/rand"
)

const (
	SCENE_DIFFICULT_ZERO int = iota
	SCENE_DIFFICULT_EASY
	SCENE_DIFFICULT_MEDIUM
	SCENE_DIFFICULT_HARD
	SCENE_DIFFICULT_TOTAL
)

type Scene struct {
	Id        int
	LevelId   int
	Name      string
	Difficult int
}

type Scenes []*Scene

func New(level int, name string) *Scene {
	SceneId := rand.Intn(10_000) // this should get the last index of scenes
	return &Scene{
		Id:        SceneId,
		LevelId:   level,
		Name:      name,
		Difficult: rand.Intn(SCENE_DIFFICULT_TOTAL),
	}
}

func (s *Scene) GetID() int      { return s.Id }
func (s *Scene) GetLevel() int   { return s.LevelId }
func (s *Scene) GetName() string { return s.Name }
