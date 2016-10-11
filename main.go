package main

import (
	// external
	"github.com/alecthomas/kingpin"
)

var (
	action = kingpin.Arg("action", "What action to take [list, scan, remove, upload, download]").Required().String()
	target = kingpin.Arg("target", "What to take action on").String()
)

func main() {
	kingpin.Parse()

	switch *action {
	case "list":
		handleList()
	case "scan":
		handleScan(*target)
	case "remove":
		handleRemove(*target)
	}
}
