package main

import (
    "os"
    "log"
)

// NOTE: This is rudimentary
func LogStrings(s []string) {
    f, err := os.OpenFile("game_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    for i := 0; i < len(s); i++ {
        if _, err := f.Write([]byte(s[i])); err != nil {
            f.Close()
            log.Fatal(err)
        }
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

