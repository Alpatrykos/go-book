package main

import (
    "bytes"
    "fmt"
    "os"
)

var buf bytes.Buffer
//comma dodaje znak przecinka co 3 cyfry w lancuchu liczby rzeczywistej
func comma(num string) string {
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
    return s
}
func main() {
    arg := os.Args[1]
    fmt.Printf("%s\n", comma(arg))
}
