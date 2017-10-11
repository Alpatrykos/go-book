package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

const str = "  ada      aaa"

func main() {
    fmt.Println(str)
    fmt.Println([]byte(str))
    str := rmSpace([]byte(str))
    fmt.Println(str)
}

func rmSpace(str []byte) []byte {
    s := string(str)
    size := 0
    prev_space := false
    for _, r := range s {
        if unicode.IsSpace(r) {
            if prev_space {
                continue
            } else {
                prev_space = true
            }
        } else {
            prev_space = false
        }
        n := utf8.RuneLen(r)
        copy(str[size:size+n], []byte(string(r)))
        size += n
    }
    return str[:size]
}
