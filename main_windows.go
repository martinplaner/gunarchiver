package main

import (
	"github.com/martinplaner/gunarchiver/ui"
	"github.com/martinplaner/gunarchiver/ui/windows"
)

func init() {
	ui.Default = windows.UserInterface{}
}
