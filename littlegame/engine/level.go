package engine

import "fmt"

type Level struct {
	Id          int
	Name        string
	Scenes      []*Scene
	isCompleted bool
}

func NewLevel(level int, name string) *Level {
	return &Level{
		Id:          level,
		Name:        name,
		Scenes:      []*Scene{},
		isCompleted: false,
	}
}

func (l *Level) GetID() int          { return l.Id }
func (l *Level) GetName() string     { return l.Name }
func (l *Level) GetScenes() []*Scene { return l.Scenes }

func (l *Level) AddScene(scene *Scene) error {
	if scene == nil {
		return fmt.Errorf("no such Scene in Level")
	}
	l.Scenes = append(l.Scenes, scene)
	return nil
}

func (l *Level) NewScene(name string, difficult SceneDifficult) {
	newScene := &Scene{
		Id:          len(l.Scenes) + 1,
		LevelId:     l.Id,
		Name:        name,
		Difficult:   difficult,
		Turns:       []*Turn{},
		SceneTurn:   &Turn{},
		isCompleted: false,
	}

	l.AddScene(newScene)
}

func (l *Level) DeleteSceneById(sceneId int) bool {
	sceneFound := false
	for s, scene := range l.Scenes {
		if scene.Id == sceneId {
			l.Scenes = append(l.Scenes[:s], l.Scenes[s+1:]...)
			sceneFound = true
		}
	}

	return sceneFound
}
