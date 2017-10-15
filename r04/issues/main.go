package main 

import (
    "fmt"
    "log"
    "os"
    "./github"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printff("%d tematow:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s\n",
        item.Number, item.User.Login, item.Title)
    }
}
