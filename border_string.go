package main

import (
    "fmt"
    "log"
    "strings"
)

var BorderStringSizes = MakeEnum([]string{
    "small",
    "medium",
    "large",
})

func BorderString(size int) string {
    if !BorderStringSizes.ContainsValue(size) {
        log.Fatal(fmt.Sprintf("Invalid size: %d", size))
    }
    a := []string{}
    for i := 0; i <= size; i++ {
        a = append(a, "-~=")
    }
    return strings.Join(a, "-")
}

