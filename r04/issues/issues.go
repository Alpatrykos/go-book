package main 

import (
    "fmt"
    "log"
    "os"
    "../github"
    "time"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d tematow:\n", result.TotalCount)
    now := time.Now()
    monthAgo := now.AddDate(0, -1, 0)
    yearAgo := now.AddDate(-1, 0, 0)
    fmt.Println(" < Month ago:")
    for _, item := range result.Items {
        if item.CreatedAt.After(monthAgo) {
            fmt.Printf("#%-5d %9.9s %60.60s %15s\n",
            item.Number, item.User.Login, item.Title, item.CreatedAt)
        }
    }
    fmt.Println("\n < Year ago:")
    for _, item := range result.Items {
        if item.CreatedAt.After(yearAgo) {
            fmt.Printf("#%-5d %9.9s %60.60s %15s\n",
            item.Number, item.User.Login, item.Title, item.CreatedAt)
        }
    }
 
    fmt.Println("\n > Year ago:")
    for _, item := range result.Items {
        if item.CreatedAt.Before(yearAgo) {
            fmt.Printf("#%-5d %9.9s %60.60s %15s\n",
            item.Number, item.User.Login, item.Title, item.CreatedAt)
        }
    }

}
