// Exercise 4.11 
// Create a tool that allows the user to create, read, update and delete issues at github 
// (basically CRUD) application 
// that uses github api to operate. Evoke Vim editor when there is a neccessity to write a considerable
// amount of text.
package main

import (
    "fmt"
    "os"
)

// get github authentication token (alternatively a config file can be used)
// Note that the OAuth access token is a environment variable
const authToken = os.Getenv("GITHUB_OAUTH_TOKEN")

func main() {
    if authToken == "" {
        fmt.Println("GITHUB_OATH_TOKEN enviranment variable holding github OAuth access token is not set.")
        os.Exit(1)
    }
    //TODO setup creating an issue in your github repo

}
