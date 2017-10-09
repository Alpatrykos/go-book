package main

import (
    "fmt"
    "crypto/sha256"
    "os"
)

func main() {
    if len(os.Args) == 1 || len(os.Args) > 2 {
        fmt.Printf("\n usage: sha256 arg1")
    }
    arg := os.Args[1]
    c := sha256.Sum256([]byte(arg))
    fmt.Printf("arg: %s sha256: %x\n", arg, c)
}
