package engine

import (
	"math/rand"
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

var (
	levelId   int    = rand.Int() % 20
	levelName string = "The Beginning"
	newLevel  *Level = NewLevel(levelId, levelName)
)

func TestNewLevel(t *testing.T) {
	mlt.Equals(t, newLevel.Id, levelId)
	mlt.Equals(t, newLevel.Name, levelName)
}

func TestLevelHelpers(t *testing.T) {
	mlt.Equals(t, newLevel.GetID(), levelId)
	mlt.Equals(t, newLevel.GetName(), levelName)
	mlt.Equals(t, newLevel.GetScenes(), []*Scene{})
}

func TestLevelScenesHelpers(t *testing.T) {
	scene := &Scene{
		Id:        rand.Int() % 20,
		LevelId:   levelId,
		Name:      "Scene Name",
		Difficult: SCENE_DIFFICULT_EASY,
	}

	newLevel.AddScene(scene)

	mlt.Equals(t, len(newLevel.Scenes), 1)
	mlt.Equals(t, newLevel.Scenes[0].LevelId, levelId)
	mlt.Assert(t, newLevel.Scenes[0].Difficult < SCENE_DIFFICULT_TOTAL)

	mlt.Equals(t, newLevel.DeleteSceneById(scene.Id), true)
	mlt.Equals(t, len(newLevel.Scenes), 0)
	mlt.Equals(t, newLevel.Scenes, []*Scene{})
}
