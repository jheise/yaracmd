package main

import (
	// standard
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleDownload(target string) {
	endpoint := buildEndpoint() + target

	// create a new request because the default http client in go doesnt delete
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

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
