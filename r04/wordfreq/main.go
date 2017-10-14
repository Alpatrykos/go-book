package main

import (
    "fmt"
    "bufio"
    "os"
)
// no error handling
func main() {
    wordcounts := make(map[string]int)

    in := bufio.NewScanner(os.Stdin)
    in.Split(bufio.ScanWords)
    for in.Scan() {
        wordcounts[in.Text()]++
    }
    for word, count := range wordcounts {
        fmt.Printf("%10q: %d\n", word, count)
    }
}
