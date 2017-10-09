package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    fmt.Println(anagram(args[0], args[1]))
}

func anagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    n := len(s1)
    for i := 0; i < n; i++ {
        if s1[i] != s2[n-i-1] {
            return false
        }
    }
    return true
}
