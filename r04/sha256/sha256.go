package main

import (
    "fmt"
    "crypto/sha256"
    "crypto/sha512"
    "os"
)

func main() {
    flag := "256"
    args := os.Args[1:]
    if len(args) < 1 || len(args) > 2 {
        fmt.Printf("\nusage: sha256 [-384,-512] arg\n Display sha256 (optionally:sha384 or sha512).\n")
    }else if len(args) == 2 {
        flag = args[0][1:]
        if args[0] == "-384" {
            c := sha512.Sum384([]byte(args[1]))
            fmt.Printf("sha%s: %x\n", flag, c)
        }else if args[0] == "-512" {
            c := sha512.Sum512([]byte(args[1]))
            fmt.Printf("sha%s: %x\n", flag, c)
        }else {
            fmt.Printf("incorect flag supplied.\n")
        }
    }else {
        c := sha256.Sum256([]byte(args[0]))
        fmt.Printf("sha%s: %x\n", flag, c)
    }
}
