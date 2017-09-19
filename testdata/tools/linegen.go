package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
    "strconv"
    "io/ioutil"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    arg := os.Args[1]
    data := ""
    lines := []string{
        "Thou shall not pass",
        "Fly you fools",
        "You have my sword",
        "And my bow",
        "And my axe",
    }
    n, _ := strconv.Atoi(arg);
    i := n
    for i > 0 {
        data += lines[rand.Intn(len(lines))] + "\n"
        i--
    }
    data1 := []byte(data)
    filename := fmt.Sprintf("%dlines", n)
    ioutil.WriteFile(filename, data1, 0644)
}
