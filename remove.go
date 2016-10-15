package main

import (
	// standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// external
	"github.com/jheise/yaramsg"
)

func handleRemove(target string) {
	endpoint := buildEndpoint() + target

	// create a new request because the default http client in go doesnt delete
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusInternalServerError:
		fmt.Println("Unable to delete " + target)
		os.Exit(1)
	case http.StatusNotFound:
		fmt.Println(target + " does not exist on server")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var current yaramsg.RemoveResponse
	json.Unmarshal(body, &current)

	if current.Result {
		fmt.Println("Successfully removed " + target)
	} else {
		fmt.Println("Failed to remove " + target)
	}
}
