package main

func MakeTestRoom(p *Actor) Scene {
    s := MakeScene("Test Room")
    g := MakeGoblin()
    s.AddActor(p)
    s.AddActor(&g)
    return s
}

