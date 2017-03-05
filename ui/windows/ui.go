// +build windows

package windows

import (
	"context"

	"github.com/martinplaner/gunarchiver/ui"
)

type UserInterface struct{}

func (UserInterface) NewProgressWindow(cancel context.CancelFunc) ui.ProgressWindow {
	return &progressWindow{
		cancel: cancel,
	}
}

func (UserInterface) NewErrorWindow() ui.ErrorWindow {
	return &errorWindow{}
}
