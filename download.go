package main

import (
	// standard
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handleDownload(target string) {
	endpoint := buildEndpoint() + target

	// create a new request because the default http client in go doesnt delete
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Println("Failed to download " + target)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(target, body, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded " + target)

}
