package main

import (
	// standard
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// external
	"github.com/jheise/yaramsg"
)

func handleUpload(target string) {
	endpoint := buildEndpoint() + target

	// open file
	f, err := os.Open(target)
	if err != nil {
		panic(err)
	}

	// make reader
	reader := bufio.NewReader(f)

	// create a new request because the default http client in go doesnt delete
	req, err := http.NewRequest("PUT", endpoint, reader)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
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
		fmt.Println("Successfully uploaded " + target)
	} else {
		fmt.Println("Failed to upload " + target)
	}
}
