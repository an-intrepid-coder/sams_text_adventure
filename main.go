package main

import (
    "math/rand"
    "time"
)

// TODO: Command history, like bash.

func main() {
    // Test game:
    seed := time.Now().UnixNano()
    rand.Seed(seed)

    g := MakeGame(seed, "Sam")
    g.GameLoop()
}

