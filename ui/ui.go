// Package ui is the abstraction layer for the user interface and defines all necessary interfaces
// that need to be implemented by concrete UI implementations.
package ui

import "github.com/martinplaner/gunarchiver/progress"

// Window is the common interface for all windows.
type Window interface {
	Show() error
	Close()
}

// ProgressWindow is the window that is shown during the extraction process.
type ProgressWindow interface {
	Window
	// Update updates the progress.
	Update(progress.Progress)
	RequestedCancel() bool
}

// ErrorWindow is the windows that is shown after an error occurred.
type ErrorWindow interface {
	Window
}

// UserInterface is the factory for creating new windows.
type UserInterface interface {
	NewProgressWindow() ProgressWindow
	NewErrorWindow(message string) ErrorWindow
}
