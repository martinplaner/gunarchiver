// +build windows

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package windows implements the user interface for the windows platform.
package windows

import (
	"github.com/martinplaner/gunarchiver/progress"
	"github.com/martinplaner/gunarchiver/ui"
)

type UserInterface struct{}

func (UserInterface) NewProgressWindow() ui.ProgressWindow {
	return &progressWindow{
		progress: &progress.Progress{},
	}
}

func (UserInterface) NewErrorWindow(message string) ui.ErrorWindow {
	return &errorWindow{
		message: message,
	}
}
