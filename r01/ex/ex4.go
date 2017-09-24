package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string] int)
    files := os.Args[1:]
    filenames := ""
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            filenames = ""
            for _, arg := range files {
                f, _ := os.Open(arg)
                input := bufio.NewScanner(f)
                for input.Scan() {
                    if line == input.Text() {
                        filenames += " " + arg[9:]
                        break
                    }
                }
            }
            fmt.Printf("%d\t%s.\t\t Found in following files: %s\n", n, line, filenames)
        }
    }
}

func countLines(f *os.File, counts map[string] int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}
