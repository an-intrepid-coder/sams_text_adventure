package main

/* Items are kind of like Actors, but have less agency in the game
   world and are often consumables. Unlike Actors, they should never
   have a Behavior property.  */

import (
    "fmt"
    "strings"
)

var numItems = MakeCounter(0, 1)

func ItemId() int {
    return numItems.GetAndIncrement()
}

type Item struct {
    Id int
    Name string
}

func MakeItem(name string) Item {
    i := Item{
        Id: ItemId(),
        Name: name,
    }
    return i
}

func (i *Item) Repr() string {
    a := []string{
        "Item Stats:",
        fmt.Sprintf("ID: %d", i.Id),
    }
    return strings.Join(a, "\n")
}
