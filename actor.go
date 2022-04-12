package main

/* An Actor in this case is any object that exists in the game world and
   which can interact with things in some way. It may include some game
   objects, but not typically inventory items. Actors can have Behavior
   properties.  */

import (
    "container/list"
    "fmt"
    "math/rand"
    "strings"
)

var numActors = MakeCounter(0, 1)

func ActorId() int {
    return numActors.GetAndIncrement()
}

type Actor struct {
    Id int
    Level int
    Name string
    IsPlayer bool
    // Their attitude towards the player:
    Attitude int
    // The combat stats bundle/component:
    Combat *Combat
    Items *list.List
    // TODO: Behavior
}

// For now, only Actors with Health are alive:
func (a *Actor) IsAlive() bool {
    if a.Combat != nil {
        return a.Combat.Health > 0
    }
    return false
}

func (a *Actor) Repr() string {
    min, max := a.TotalDamageRange()
    s := []string{
        "Actor Stats:",
        BorderString(0),
        fmt.Sprintf("ID: %d", a.Id),
        fmt.Sprintf("Level: %d", a.Id),
        fmt.Sprintf("Name: %s", a.Name),
        fmt.Sprintf("Is Player: %t", a.IsPlayer),
        fmt.Sprintf("Attitude: %s", Attitudes.Label(a.Attitude)),
    }
    if a.Combat != nil {
        b := []string{
            fmt.Sprintf("Health: %d / %d", a.Combat.Health, a.Combat.MaxHealth),
            fmt.Sprintf("Base Attack: %d", a.TotalAttack()),
            fmt.Sprintf("Base Damage: %d-%d", min, max),
            fmt.Sprintf("Base Armor: %d", a.TotalArmor()),
            fmt.Sprintf("Base Defense: %d", a.TotalDefense()),
        }
        if a.Combat.Engaged != nil {
            b = append(b, fmt.Sprintf("Engaged With: %s", a.Combat.Engaged.Name))
        }
        s = append(s, b...)
    }
    for e := a.Items.Front(); e != nil; e = e.Next() {
        i := e.Value.(*Item)
        s = append(s, i.Repr())
    }
    return strings.Join(s, "\n")
}

// NOTE: In progress -- will take inventory in to account later.
func (a *Actor) TotalAttack() int {
    return a.Combat.Attack
}

// NOTE: In progress -- will take inventory in to account later.
func (a *Actor) TotalDefense() int {
    return a.Combat.Defense
}

// NOTE: In progress -- will take inventory in to account later.
func (a *Actor) TotalDamage() int {
    max, min := a.TotalDamageRange()
    return rand.Intn(max - min) + min
}

// NOTE: In progress -- will take inventory in to account later.
func (a *Actor) TotalArmor() int {
    return a.Combat.Armor
}

// NOTE: In progress -- will take inventory in to account later.
func (a *Actor) TotalDamageRange() (int, int) {
    return a.Combat.MinDamage, a.Combat.MaxDamage
}

type FightResult struct {
    Attacker *Actor
    Defender *Actor
    DamageDealt int
    DefenderLives bool
    Repr string
}

func (a *Actor) Fight(other *Actor) FightResult {
    // NOTE: Combat essentially works on a D20 system for now.
    b := []string{
        fmt.Sprintf("%s attacks %s!", a.Name, other.Name),
        BorderString(1),
    }
    r := FightResult{
        Attacker: a,
        Defender: other,
    }
    attackRoll := RollD20() + a.TotalAttack()
    b = append(b, fmt.Sprintf("%s rolls a %d (inc. modifiers).", a.Name, attackRoll))
    if attackRoll >= other.TotalDefense() {
        b = append(b, fmt.Sprintf("%s hits!", a.Name))
        dmg := a.TotalDamage() - other.Combat.Armor
        b = append(b, fmt.Sprintf("%s does %d damage to %s!", a.Name, dmg, other.Name))
        other.Combat.Health -= dmg
        if !other.IsAlive() {
            b = append(b, fmt.Sprintf("%s kills %s!", a.Name, other.Name))
            r.DefenderLives = false
        } else {
            r.DefenderLives = true
        }
    }
    r.Repr = strings.Join(b, "\n")
    return r
}

