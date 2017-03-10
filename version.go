package main

import (
	"fmt"
	"os"
)

var (
	Version    string = "undefined"
	BuildDate  string = "N/A"
	GitCommit  string = "undefined"
	GitBranch  string = ""
	GitState   string = ""
	GitSummary string = ""
)

func PrintVersion() {
	fmt.Printf("%s version %s\nbuild %s\ndate %s\n\n", os.Args[0], Version, GitCommit, BuildDate)
}
