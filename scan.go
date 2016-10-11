package main

import (
	// standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// external
	"github.com/jheise/yaramsg"
)

func handleScan(target string) {
	endpoint := fmt.Sprintf("http://localhost:9999/scanner/v1/%s/scan/", target)
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Got back " + string(body))

	var current yaramsg.ScanResponse
	json.Unmarshal(body, &current)

	fmt.Printf("Scan Results - %s:\n", target)
	for _, scan := range current.Matches {
		fmt.Printf("    Rule: %s\n", scan.Rule)
		fmt.Printf("\t- Namespace: %s\n", scan.Namespace)
		fmt.Printf("\t- Tags: %s\n", strings.Join(scan.Tags, ", "))
	}
}
