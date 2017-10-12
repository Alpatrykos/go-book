package main

import (
    "fmt"
    "unicode/utf8"
    "os"
)

func main() {
    arg := os.Args[1]
    reverse([]byte(arg))
}
        //TODO fix this; make reverse work with utf-8 encoding
func reverse(s []byte) {
    str := s[:]
    for len(str) > 0 {
        r1, size1 := utf8.DecodeRune(str)
        fmt.Printf("%c %v\n", r1, size1)
        fmt.Println("------")
        r2, size2 := utf8.DecodeLastRune(str)
        fmt.Printf("%c %v\n", r2, size2)
        fmt.Printf("\n\n")
        


        str = str[size1:len(str)-size2]


    }
}
