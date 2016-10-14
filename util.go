package main

import (
	"fmt"
)

func buildEndpoint() string {
	return fmt.Sprintf("http://%s:%s/scanner/v1/files/", *host, *port)
}
