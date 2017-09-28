package main

import (
    "fmt"
    "time"
    "./popcount"
    "./popcount_64shifts"
    "./popcount_clearLSB"
)

const (
    x = 12315124124
    iter = 100000
)

func main() {
    start := time.Now()
    for i:= 0; i < iter; i++ {
        popcount.PopCount(x)
    }
    fmt.Println("time: %d", time.Since(start).Seconds())

    start = time.Now()
    for i:= 0; i < iter; i++ {
        popcount_64shifts.PopCount(x)
    }
    fmt.Println("time: %d", time.Since(start).Seconds())

    start = time.Now()
    for i:= 0; i < iter; i++ {
        popcount_clearLSB.PopCount(x)
    }
    fmt.Println("time: %d", time.Since(start).Seconds())
}
