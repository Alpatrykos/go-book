// exercise 4.11
// create a tool that allows the user to create, read, update and delete issues at github
// (basically crud) application
// that uses github api to operate. evoke vim editor when there is a neccessity to write a considerable
// amount of text.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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
		if err != nil {
			fmt.Printf("Error occured: %s", err)
		}
		err = postIssue(args[1], args[2], *issue)
		if err != nil {
			fmt.Printf("Error occured: %s\n", err)
		} else {
			fmt.Printf("Successfully posted an issue:\nNumber:%d\tTitle: %s\nBody:\n%s",
				issue.Number, issue.Title, issue.Body)
		}
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

// already handles setting Title and Body
func createIssue() (*Issue, error) {
	var result Issue
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Title:")
	result.Title, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Czytanie ze standardowego wejscia nie powiodlo sie: %s", err)
	}
	result.Body, err = evokeEditor()
	if err != nil {
		panic(err)
	}
	return &result, nil
}

func postIssue(user, repo string, issue Issue) error {
	// TODO: api gives 404Error, app need pull level privileges to post issue
	var err error
	var req []byte
	var resp *http.Response
	t := ApiURL + user + "/" + repo + "/issues/"
	fmt.Println(t)
	req, err = json.Marshal(issue)
	if err != nil {
		return fmt.Errorf("Wystapil blad podczas zamiany struktury na json: %s", err)
	}
	resp, err = http.Post(t, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return fmt.Errorf("Wystapil blad podczas wysylania danych: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("Postowanie nie powiodlo sie: %s",
			resp.Status)
	}
	resp.Body.Close()
	return nil
}

// evoke vim editor to be able to input multi lines of text
func evokeEditor() (string, error) {
	vi := "vim"
	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "tempFilePrefix")
	if tmpFileErr != nil {
		return "", fmt.Errorf("Blad podczas tworzenia pliku tymczasowego: %s", tmpFileErr)
	}
	path, err := exec.LookPath(vi)
	if err != nil {
		return "", fmt.Errorf("Blad podczas szukania %s:%s", vi, err)
	}
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return "", fmt.Errorf("Start sie nie udal: %s", err)
	}
	// fmt.Printf("Oczekiwanie na zakonczenie komendy.\n")
	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("Komenda zakonczona z bledem: %s", err)
	}
	var result []byte
	result, err = ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("Odczytanie tymczasowego pliku nie powiodlo sie: %s", err)
	}
	if len(result) < 1 {
		return "", fmt.Errorf("Pusty plik.")
	}
	err = os.RemoveAll(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("Usuniecie pliku tymczasowego nie powiodlo sie: %s", err)
	}

	return string(result), nil
}
