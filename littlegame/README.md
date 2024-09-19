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

from `main.go`

```go
func ExampleGameTwo() {
	playerOne := character.NewPlayer()
	playerTwo := character.NewPlayer()
	playerThree := character.NewPlayer()

	levelOne := engine.NewLevel(1, "The Beginning")
	sceneOne := engine.NewScene(levelOne.Id, "Me vs You")
	turnOne := engine.NewTurn(engine.TURN_MODE_SINGLE_CHALLENGE)
	turnOne.AddPlayer(playerOne)
	turnOne.AddPlayer(playerTwo)
	sceneOne.AddTurn(turnOne)
	levelOne.AddScene(sceneOne)

	levelTwo := engine.NewLevel(2, "The Combat")
	sceneTwo := engine.NewScene(levelTwo.Id, "Against Two Junkies")
	turnTwo := engine.NewTurn(engine.TURN_MODE_COMBAT)
	turnTwo.AddPlayer(playerOne)
	turnTwo.AddPlayer(playerTwo)
	turnTwo.AddPlayer(playerThree)
	sceneTwo.AddTurn(turnTwo)
	levelTwo.AddScene(sceneTwo)

	levels := []*engine.Level{levelOne, levelTwo}
	gameParams := &engine.GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 250,
	}
	engine.GameLoop(gameParams)
}

func ExampleGameOne() {
	levelOne := engine.NewLevel(1, "The Beginning")
	levelTwo := engine.NewLevel(2, "The Combat")

	levelOne.NewScene("Scene 01", engine.SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", engine.SCENE_DIFFICULT_EASY)

	levels := []*engine.Level{levelOne, levelTwo}
	gameParams := &engine.GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 250,
	}
	engine.GameLoop(gameParams)
}
```

#### TestGameLoop function, needs to be refactored :)

from `engine/main_test.go`
```go
func TestGameLoop(t *testing.T) {
	levelOne := NewLevel(1, "The Beginning")
	levelTwo := NewLevel(2, "The Combat")

	playerOne := character.NewPlayer()
	playerTwo := character.NewPlayer()
	playerThree := character.NewPlayer()
	playerFour := character.NewPlayer()

	levelOne.NewScene("Scene 01", SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 02", SCENE_DIFFICULT_EASY)
	levelTwo.NewScene("Scene 03", SCENE_DIFFICULT_MEDIUM)

	levelOne.Scenes[0].NewTurn(TURN_MODE_COMBAT)
	levelOne.Scenes[0].Turns[0].AddPlayer(playerOne)
	levelOne.Scenes[0].Turns[0].AddPlayer(playerTwo)

	levelTwo.Scenes[1].NewTurn(TURN_MODE_COMBAT)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerOne)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerTwo)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerThree)
	levelTwo.Scenes[1].Turns[0].AddPlayer(playerFour)

	levels := []*Level{levelOne, levelTwo}
	gameParams := &GameLoopParameters{
		LevelList: levels,
		LoopTime:  time.Millisecond * 150,
	}

	timeStart := time.Now()
	GameLoop(gameParams)
	timeDuration := time.Since(timeStart)
	mlt.Assert(t, timeDuration < (1_000_000_000*15))
	utils.LogMessage(utils.LOG_TYPE_DEBUG, float64(timeDuration/1_000_000_000))
}
```
### outputs
```log
=== RUN   TestGameLoop
DEBUG main.go:27 engine.Level{Id:1, Name:"The Beginning", Scenes:[]*engine.Scene{(*engine.Scene)(0xc00008a190)}, isCompleted:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{0, 1}}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{0, 1}}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{1, 0}}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{1, 0}}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{0, 1}}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:1, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0)}, TurnOrder:[]int{0, 1}}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:1, Name:"Scene 01", Difficult:1, Turns:[]*engine.Turn{(*engine.Turn)(0xc000014480)}, SceneTurn:(*engine.Turn)(0xc000014480), isCompleted:true}
DEBUG main.go:27 engine.Level{Id:2, Name:"The Combat", Scenes:[]*engine.Scene{(*engine.Scene)(0xc00008a1e0), (*engine.Scene)(0xc00008a230)}, isCompleted:false}
DEBUG main.go:38 engine.Scene{Id:1, LevelId:2, Name:"Scene 02", Difficult:1, Turns:[]*engine.Turn{}, SceneTurn:(*engine.Turn)(0xc0000143c0), isCompleted:false}
DEBUG main.go:38 engine.Scene{Id:2, LevelId:2, Name:"Scene 03", Difficult:2, Turns:[]*engine.Turn{(*engine.Turn)(0xc0000144e0)}, SceneTurn:(*engine.Turn)(0xc0000144e0), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:2, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0), (*character.Player)(0xc000096720), (*character.Player)(0xc000096750)}, TurnOrder:[]int{0, 2, 1, 3}}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:5400, Name:"Joseph Landry", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:9895, Name:"Albert Mayo", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:2, LevelId:2, Name:"Scene 03", Difficult:2, Turns:[]*engine.Turn{(*engine.Turn)(0xc0000144e0)}, SceneTurn:(*engine.Turn)(0xc0000144e0), isCompleted:false}
DEBUG main.go:49 engine.Turn{Id:1, SceneId:2, Mode:2, isCompleted:false, timestamp:1726757634, Players:[]*character.Player{(*character.Player)(0xc0000966c0), (*character.Player)(0xc0000966f0), (*character.Player)(0xc000096720), (*character.Player)(0xc000096750)}, TurnOrder:[]int{1, 3, 0, 2}}
DEBUG main.go:53 character.Player{Id:7218, Name:"Layne Herring", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:9895, Name:"Albert Mayo", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:7119, Name:"Dion Warner", Health:14, isDead:false}
DEBUG main.go:53 character.Player{Id:5400, Name:"Joseph Landry", Health:14, isDead:false}
DEBUG main.go:38 engine.Scene{Id:2, LevelId:2, Name:"Scene 03", Difficult:2, Turns:[]*engine.Turn{(*engine.Turn)(0xc0000144e0)}, SceneTurn:(*engine.Turn)(0xc0000144e0), isCompleted:true}
DEBUG main_test.go:45[m 2
--- PASS: TestGameLoop (2.71s)
PASS
ok  	github.com/mulatinho/golabs/littlegame/engine	2.714s
```