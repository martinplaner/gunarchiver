// +build windows

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type errorWindow struct {
	mainWindow *walk.MainWindow
	message    string
}

func (w *errorWindow) Show() error {

	var dialog *walk.Dialog
	var acceptPB *walk.PushButton

	error := Dialog{
		AssignTo:      &dialog,
		FixedSize:     true,
		MinSize:       Size{400, 150},
		DefaultButton: &acceptPB,
		Title:         "Error",
		Icon:          walk.IconError(),
		Layout:        VBox{},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					Label{
						Text: w.message,
					},
					PushButton{
						Text: "OK",
						OnClicked: func() {
							dialog.Close(0)
						},
					},
				},
			},
		},
	}

	_, err := error.Run(nil)

	return err
}

func (w *errorWindow) Close() {
	w.mainWindow.Close()
}
