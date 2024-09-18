package engine

type Level struct {
	id     int
	name   string
	scenes []*Scene
}

func (l *Level) New(level int, name string) *Level {
	return &Level{
		id:     level,
		name:   name,
		scenes: []*Scene{},
	}
}

func (l *Level) GetID() int          { return l.id }
func (l *Level) GetName() string     { return l.name }
func (l *Level) GetScenes() []*Scene { return l.scenes }
