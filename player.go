package main

import (
    "container/list"
)

func MakePlayer(name string) Actor {
    a := Actor{
        Id: ActorId(),
        Level: 1,
        Name: name,
        IsPlayer: true,
        Attitude: *Attitudes.Value("ally"),
        Items: list.New(),
    }
    c := Combat{
        Owner: &a,
        Engaged: nil,
        Health: 10,
        MaxHealth: 10,
        Attack: 6,
        MinDamage: 1,
        MaxDamage: 2,
        Armor: 0,
        Defense: 10,
    }
    a.Combat = &c
    return a
}

