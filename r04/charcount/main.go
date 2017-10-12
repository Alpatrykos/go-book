package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int) //zliczanie wystapien znakow Unicode
    letters := make(map[rune]int)
    digits := make(map[rune]int)
    var utflen [ utf8.UTFMax + 1]int //zliczanie dlugosci kodowan UTF-8
    invalid := 0 // liczba nieprawidlowych znakow UTF-8

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune() //zwraca rune, liczbe bajtow i blad
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        counts[r]++
        utflen[n]++
        if unicode.IsLetter(r) {
            letters[r]++
        }
        if unicode.IsNumber(r) {
            digits[r]++
        }
    }
    fmt.Printf("runa\tliczba wystapien\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Print("\ndlugosc\tliczba wystapien\n")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    if invalid > 0 {
        fmt.Printf("\n%d niewlasciwych znakow UTF-8\n", invalid)
    }
}
