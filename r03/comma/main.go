package main

import (
    "bytes"
    "fmt"
    "os"
)

const dash = '-'

var buf bytes.Buffer
var fract string
//comma dodaje znak przecinka co 3 cyfry w lancuchu liczby rzeczywistej
func comma(num1 string) string {
    num := num1
    if num[0] == dash {
        buf.WriteByte('-')
        num = num[1:]
    }
    for i := 0; i < len(num); i++ {
        if num[i] == '.' {
            fract = num[i:]
            num = num[:i]
        }
    }
    var nnum bytes.Buffer
    if len(num) <= 3 {
        nnum.WriteString(num)
    }else {
        var temp string
        nnum.WriteString(num[len(num)-3:])
        num = num[:len(num)-3]
        for len(num) > 3 {
            temp = nnum.String()
            nnum.Reset()
            fmt.Fprintf(&nnum, "%s,%s", num[len(num)-3:], temp)
            num = num[:len(num)-3]
            if len(num) < 3 && len(num) > 0 {
                temp = nnum.String()
                nnum.Reset()
                fmt.Fprintf(&nnum, "%s,%s", num, temp)
            }
        }
    }
    buf.WriteString(nnum.String())
    buf.WriteString(fract)
    return buf.String()
}

func main() {
    arg := os.Args[1]
    fmt.Printf("%s\n", comma(arg))
}
