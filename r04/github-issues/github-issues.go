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

const apiurl = "https://api.github.com/repos"

type issues struct {
    totalcount int `json:total_count`
    items []*issue
}

type issue struct {
    number int
    htmlurl string `json:html_url`
    title string
    state string
    user *user
    createdat time.time `json:created_at`
    body string
}

type user struct {
    login string
    htmlurl string `json:html_url`
}

// get github authentication token (alternatively a config file can be used)
// note that the oauth access token is a environment variable
const authtoken = os.getenv("github_oauth_token")

func main() {
    if authtoken == "" {
        fmt.println("github_oath_token enviranment variable is not set.")
        os.exit(1)
    }
    /*
    usage:
    : repo - list all issues of the repo repository
    : -s repo number - show issue
    : -c repo - create an issue
    : -u repo number - update an issue
    : -d repo number - delete an issue

    */
    args := os.args[1:]
    }
    //todo setup creating an issue in your github repo

}

func getIssue(terms []string) (*issue, error) {
    q := url.queryescape(strings.join(terms, "/"))
    resp, err := http.get(apiurl + q)
    if err != nil {
        return nil, err
    }

    if resp.statuscode != http.statusok {
        resp.body.close()
        return nil, fmt.errorf("uzyskanie sprawy nie powiodlo sie: %s", resp.status)
    }

    var result issue
    if err := json.newdecoder(resp.body).decode(&result); err != nil {
        resp.body.close()
        return nil. err
    }
    resp.body.close()
    return &result, nil
}


