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

func handleList() {
	endpoint := buildEndpoint()
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var current yaramsg.ListResponse
	json.Unmarshal(body, &current)

	fmt.Println("Current binaries:")
	for _, bin := range current.Files {
		fmt.Printf("\t- %s\n", bin)
	}
}
