package main

/* The world will be represented as a graph of nodes, and
   Scenes are the nodes.  */

import (
    "container/list"
    "fmt"
    "strings"
)

var numScenes = MakeCounter(0, 1)

func SceneId() int {
    return numScenes.GetAndIncrement()
}

type Scene struct {
    Id int
    Name string
    Actors *list.List
}

func (s *Scene) AddActor(a *Actor) {
    s.Actors.PushFront(a)
}

/* Returns a slice containing pointers to all actors which can be engaged
   or fought.  */
func (s *Scene) HostileActors() []*Actor {
    a := []*Actor{}
    for e := s.Actors.Front(); e != nil; e = e.Next() {
        b := e.Value.(*Actor)
        if Attitudes.Label(b.Attitude) == "hostile" {
            a = append(a, b)
        }
    }
    if len(a) > 0 { 
        return a
    }
    return nil
}

// Returns true/false on success/failure.
func (s *Scene) RemoveActor(a *Actor) bool {
    for e := s.Actors.Front(); e != nil; e = e.Next() {
        b := e.Value.(*Actor)
        if a == b {
            s.Actors.Remove(e)
            return true
        }
    }
    return false
}

/* Removes actors with less than 1 health.
   Returns number of removed Actors.  */
func (s *Scene) RemoveDeadActors() int {
    i := 0
    for e := s.Actors.Front(); e != nil; e = e.Next() {
        a := e.Value.(*Actor)
        if !a.IsAlive() {
            i++
            s.Actors.Remove(e)
        }
    }
    return i
}

func MakeScene(name string) Scene {
    s := Scene{
        Id: SceneId(),
        Name: name,
        Actors: list.New(),
    }
    return s
}

func (s *Scene) Repr() string {
    a := []string{
        "Scene Stats:",
        BorderString(2),
        fmt.Sprintf("ID: %d", s.Id),
        fmt.Sprintf("Name: %s", s.Name),
        "Describing Actors:",
        BorderString(1),
    }
    for e := s.Actors.Front(); e != nil; e = e.Next() {
        b := e.Value.(*Actor)
        a = append(a, b.Repr())
    }
    return strings.Join(a, "\n")
}

