package main

import (
    "bytes"
    "fmt"
    "os"
)

var buf bytes.Buffer
//comma dodaje znak przecinka co 3 cyfry w lancuchu liczby rzeczywistej
func comma(num string) string {
    var minus string
    var fract string
    if num[0] == '-' {
        minus = string('-')
        num = num[1:]
    }
    for i := 0; i < len(num); i++ {
        if num[i] == '.' {
            fract = num[i:]
            num = num[:i]
        }
    }
    n := len(num)
    var s string
    if n <= 3 {
        return num
    }
    buf.WriteString(num)
    for buf.Len() > 3 {
        s = string(',') + buf.String()[buf.Len()-3:] + s
        buf.Truncate(buf.Len() - 3)
    }
    if buf.Len() > 0 {
        s = buf.String() + s
    }
    s = minus + s + fract
    return s
}
func main() {
    arg := os.Args[1]
    fmt.Printf("%s\n", comma(arg))
}
