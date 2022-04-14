package main

import (
    "container/list"
    "fmt"
)

type Game struct {
    Seed int64
    Player *Actor
    CurrentScene *Scene
    Scenes *list.List
}

func MakeGame(Seed int64, name string) Game {
    ls, p := list.New(), MakePlayer("Sam")
    s := MakeTestRoom(&p)
    ls.PushFront(&s)
    return Game{
        Seed: Seed,
        Player: &p,
        CurrentScene: &s,
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
        fmt.Printf("Input Received: %s", noun)

        // Handle Input:
        q := false
        switch verb {
            case "quit":
                q = true
            case "describe":
                switch noun {
                    case "self":
                        fmt.Println(g.Player.Repr())
                    case "scene":
                        fmt.Println(g.CurrentScene.Repr())
                    default:
                        fmt.Println("Invalid noun.")
                }
            case "fight": {
                // TODO: Some kind of target menu (in progress)
            }
            case "engage": {
                // TODO: Some kind of target menu (in progress)
                if a := g.CurrentScene.HostileActors(); a != nil {
                    fmt.Println("Valid targets:")
                    for i := range a {
                        fmt.Printf("(%s) %s, ", Az.Label(i), a[i].Name)
                    }
                    fmt.Print("\n")
                    // Get input:
                    fmt.Println("\nInput Target:")
                    fmt.Print("> ")
                    fmt.Scanln(&noun)
                    fmt.Printf("Input Received: %s\n", noun)
                    g.Player.Combat.Engaged = a[*Az.Value(noun)]
                }
                fmt.Println("No valid targets.")
            }
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

