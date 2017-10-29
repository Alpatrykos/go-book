// exercise 4.11
// create a tool that allows the user to create, read, update and delete issues at github
// (basically crud) application
// that uses github api to operate. evoke vim editor when there is a neccessity to write a considerable
// amount of text.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const ApiURL = "https://api.github.com/repos/"

type Issues struct {
	TotalCount int `json:total_count`
	Items      []*Issue
}

type Issue struct {
	Number    int
	Htmlurl   string `json:html_url`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:created_at`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:html_url`
}

// get github authentication token (alternatively a config file can be used)
// note that the oauth access token is a environment variable

func main() {
	// authtoken := os.Getenv("GITHUB_OAUTH_TOKEN")
	// if authtoken == "" {
	// 	fmt.Println("github_oath_token enviranment variable is not set.")
	// 	os.Exit(1)
	// }
	/*
	   usage:
	   : repo - list all issues of the repo repository
	   : -s user repo number - show issue
	   : -c user repo - create an issue
	   : -u user repo number - update an issue
	   : -d user repo number - delete an issue

	*/

	args := os.Args[1:]
	//todo setup creating an issue in your github repo
	if args[0] == "-s" {
		issue, err := getIssue(args[1], args[2], args[3])
		if err != nil {
			fmt.Printf("Error occured: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d : %s\n", issue.Number, issue.Title)
		fmt.Printf("%s\n", issue.Body)
	} else if args[0] == "-c" {
		issue, err := createIssue()
		postIssue(issue)
	}
}

func getIssue(user, repo, number string) (*Issue, error) {
	t := user + "/" + repo + "/issues/" + number
	// fmt.Println(ApiURL + t)
	resp, err := http.Get(ApiURL + t)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("uzyskanie sprawy nie powiodlo sie: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func createIssue() Issue, error {
  var result Issue
  
}

//
// func createIssue(terms []string, issue *Issue) error {
// 	resp, err := http.PostForm(url, data) //change url and data later
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("utworzenie sprawy nie powiodlo sie: %s", resp.Status)
// 	}
//
// }
