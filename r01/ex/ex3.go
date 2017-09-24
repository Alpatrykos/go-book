package main

import (
    "fmt"
    "os"
    "time"
    "strings"
)


func echo1() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

func echo2() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func echo3() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
    start := time.Now()
    echo1()
    fmt.Println("time: ", time.Since(start).Seconds())

    start = time.Now()
    echo2()
    fmt.Println("time: ", time.Since(start).Seconds())

    start = time.Now()
    echo3()
    fmt.Println("time: ", time.Since(start).Seconds())
}
