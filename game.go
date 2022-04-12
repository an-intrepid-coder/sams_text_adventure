package main

import (
    "container/list"
    "fmt"
)

type Game struct {
    seed int64
    player *Actor
    currentScene *Scene
    Scenes *list.List
}

func MakeGame(seed int64, name string) Game {
    ls, p := list.New(), MakePlayer("Sam")
    s := MakeTestRoom(&p)
    ls.PushFront(&s)
    return Game{
        seed: seed,
        player: &p,
        currentScene: &s,
        Scenes: ls,
    }
}

func (g *Game) GameLoop() {
    for {
        // Get input:
        fmt.Println("\nInput Command:")
        fmt.Print("> ")
        var verb, noun string
        fmt.Scanln(&verb, &noun)
        fmt.Println(fmt.Sprintf("Input Received: %s %s", verb, noun))

        // Handle Input:
        q := false
        switch verb {
            case "quit":
                q = true
            case "describe":
                switch noun {
                    case "self":
                        fmt.Println(g.player.Repr())
                    case "scene":
                        fmt.Println(g.currentScene.Repr())
                    default:
                        fmt.Println("Invalid noun.")
                }
            case "fight":   
                // TODO: Some kind of target menu
            case "engage":
                // TODO: Some kind of target menu
            case "disengage":
                // TODO
            default:
                fmt.Println("Invalid verb.")
        }

        if q {
            break
        }
    }
}

