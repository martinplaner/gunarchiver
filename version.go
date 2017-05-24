package main

import (
	"fmt"
	"os"
)

var (
	// Version is the version string in compliance with Semantic Versioning 2.x
	Version = "undefined"
	// BuildDate is the date and time of build
	BuildDate = "N/A"
	// GitCommit is the short commit hash of the git source tree
	GitCommit = "undefined"
	// GitBranch is the current branch name the code is built off
	GitBranch string
	// GitState represents whether there are uncommitted changes (dirty/clean)
	GitState string
	// GitSummary is the output of `output of git describe --tags --dirty --always`
	GitSummary string
)

// PrintVersion prints the current version, comment hash and build date
func PrintVersion() {
	fmt.Printf("%s version %s\nbuild %s\ndate %s\n\n", os.Args[0], Version, GitCommit, BuildDate)
}
