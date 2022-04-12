package main

import (
    "container/list"
)

func MakeGoblin() Actor {
    a := Actor{
        Id: ActorId(),
        Level: 1,
        Name: "Goblin",
        IsPlayer: false,
        Attitude: *Attitudes.Value("hostile"),
        Items: list.New(),
    }
    c := Combat{
        Owner: &a,
        Engaged: nil,
        Health: 3,
        MaxHealth: 3,
        Attack: 2,
        MinDamage: 1,
        MaxDamage: 2,
        Armor: 0,
        Defense: 8,
    }
    a.Combat = &c
    return a
}

