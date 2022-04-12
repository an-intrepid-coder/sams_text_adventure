package main

import (
    "math/rand"
)

func RollD20() int {
    return rand.Intn(20) + 1
}

