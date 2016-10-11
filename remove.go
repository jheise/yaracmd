package main

import (
	// standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// external
	"github.com/jheise/yaramsg"
)

func handleRemove(target string) {
	endpoint := fmt.Sprintf("http://localhost:9999/scanner/v1/%s/remove/", target)
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

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
