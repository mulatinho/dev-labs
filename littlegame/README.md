# LittleGame Engine

This is a package that aims to be a game engine created only for fun and training, the initial purpose is the ability to create levels, scenes, and for each scene create objectives that players could achieve.

- Levels
- - Scenes
- Scenes
- - Objectives
- - Turns
- - Combats

Players also have three objective modes: single and multiple challenges, combat, each of them can have turns where players and NPCs could execute actions.

- Players 
- - Types
- - Skills
- - Combat
- - Common Actions

```txt
while GameLoop: 
    for each level:
        run the Scene XXX:
            checkTheObjectives()
            while turnsAreRunning():
                if Objective.isCompleted { break }
            endScene()
        
        if level.isCompleted {
            endLevel()
        }
```


## Simple Example of Levels and Scenes creation

```go
func main() {
	levelOne := engine.NewLevel(1, "The Beginning")
	levelTwo := engine.NewLevel(2, "The Combat")

	levelOne.NewScene("Scene 01", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 03", engine.SCENE_DIFFICULT_MEDIUM)
	levelTwo.NewScene("Scene 04", engine.SCENE_DIFFICULT_HARD)

	levels := []engine.Level{*levelOne, *levelTwo}
	engine.GameLoop(levels)
}
```
outputs
```log
:. engine.Level{Id:1, Name:"The Beginning", Scenes:[]*engine.Scene{(*engine.Scene)(0xc00008a050)}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:true}
:.:. LEVEL 1, SCENE FINISHED => 1
:. engine.Level{Id:2, Name:"The Combat", Scenes:[]*engine.Scene{(*engine.Scene)(0xc00008a0a0), (*engine.Scene)(0xc00008a0f0), (*engine.Scene)(0xc00008a140)}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, isCompleted:true}
:.:. LEVEL 2, SCENE FINISHED => 1
:.:. &engine.Scene{Id:2, LevelId:2, Name:"Scene 03", Difficult:2, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:2, LevelId:2, Name:"Scene 03", Difficult:2, Turns:[]*engine.Turn{}, isCompleted:true}
:.:. LEVEL 2, SCENE FINISHED => 2
:.:. &engine.Scene{Id:3, LevelId:2, Name:"Scene 04", Difficult:3, Turns:[]*engine.Turn{}, isCompleted:false}
:.:. &engine.Scene{Id:3, LevelId:2, Name:"Scene 04", Difficult:3, Turns:[]*engine.Turn{}, isCompleted:true}
:.:. LEVEL 2, SCENE FINISHED => 3
```