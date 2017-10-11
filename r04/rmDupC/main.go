package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    fmt.Printf("%s\n", args)
    args = rmDup(args)
    fmt.Println(args)
}

func rmDup(s []string) []string{
    for i := 0; i < len(s)-1; i++ {
        if s[i] == s[i+1] {
            copy(s[i:], s[i+1:])
            s = s[:len(s)-1]
        }
    }
    return s
}
