package ui

import (
	"context"

	"github.com/martinplaner/gunarchiver/progress"
)

var ProgressMaxValue = 100

var Default UserInterface

type Window interface {
	Close()
}

type ProgressWindow interface {
	Window
	Show() error
	// Update updates the progress.
	// Must be called _after_ calling Show()!
	Update(progress.Progress)
	RequestedCancel() bool
}

type ErrorWindow interface {
	Window
	Show(message string) error
}

type UserInterface interface {
	NewProgressWindow(context.CancelFunc) ProgressWindow
	NewErrorWindow() ErrorWindow
}
