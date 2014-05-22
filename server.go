package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/martini"
)

func NewServer() *martini.ClassicMartini {
	m := martini.Classic()
	m.Post("/github", PostGitHub)
	return m
}

func PostGitHub(w http.ResponseWriter, r *http.Request) (int, string) {
	event := r.Header.Get("X-GitHub-Event")

	switch event {
	case "deployment":
		return handleDeployment(w, r)
	case "deployment_status":
		return handleDeploymentStatus(w, r)
	default:
		return 400, "Bad Request"
	}
}

// handlDeployment handles the `deployment` event from GitHub.
func handleDeployment(w http.ResponseWriter, r *http.Request) (int, string) {
	var p GitHubDeployment

	err := parsePayload(r, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	return 200, "{}"
}

// handlDeployment handles the `deployment_status` event from GitHub.
func handleDeploymentStatus(w http.ResponseWriter, r *http.Request) (int, string) {
	var p GitHubDeploymentStatus

	err := parsePayload(r, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	return 200, "{}"
}

func parsePayload(r *http.Request, v interface{}) error {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, v)
	if err != nil {
		return err
	}

	return nil
}
