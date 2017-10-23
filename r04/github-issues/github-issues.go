// exercise 4.11 
// create a tool that allows the user to create, read, update and delete issues at github 
// (basically crud) application 
// that uses github api to operate. evoke vim editor when there is a neccessity to write a considerable
// amount of text.
package main

import (
    "fmt"
    "os"
    "time"
)

const ApiURL = "https://api.github.com/repos"

type Issues struct {
    TotalCount int `json:total_count`
    Items []*issue
}

type Issue struct {
    Number int
    Htmlurl string `json:html_url`
    Title string
    State string
    User *user
    CreatedAt time.time `json:created_at`
    Body string
}

type User struct {
    Login string
    HTMLURL string `json:html_url`
}

// get github authentication token (alternatively a config file can be used)
// note that the oauth access token is a environment variable
const authtoken = os.getenv("github_oauth_token")

func main() {
    if authtoken == "" {
        fmt.Println("github_oath_token enviranment variable is not set.")
        os.Exit(1)
    }
    /*
    usage:
    : repo - list all issues of the repo repository
    : -s repo number - show issue
    : -c repo - create an issue
    : -u repo number - update an issue
    : -d repo number - delete an issue

    */
    args := os.Args[1:]
    //todo setup creating an issue in your github repo

}

func getIssue(terms []string) (*Issue, error) {
    q := url.QueryEscape(strings.join(terms, "/"))
    resp, err := http.Get(apiurl + q)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("uzyskanie sprawy nie powiodlo sie: %s", resp.Status)
    }

    var result issue
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil. err
    }
    resp.Body.Close()
    return &result, nil
}

func createIssue(terms []string, issue *Issue) error {
    resp, err := http.PostForm(url, data) //change url and data later
    if err != nil {
        return err
    }
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("utworzenie sprawy nie powiodlo sie: %s", resp.Status)
    }

