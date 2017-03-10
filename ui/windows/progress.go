// +build windows

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/martinplaner/gunarchiver/progress"
)

type progressWindow struct {
	mainWindow      *walk.MainWindow
	progressBar     *walk.ProgressBar
	currentFile     *walk.Label
	dataBinder      *walk.DataBinder
	requestedCancel bool
	err             error
	progress        *progress.Progress
}

func (w *progressWindow) Show() error {

	var cancelButton *walk.PushButton

	main := MainWindow{
		AssignTo: &w.mainWindow,
		Title:    "gunarchiver",
		Size:     Size{400, 100},
		Layout:   VBox{},
		DataBinder: DataBinder{
			AssignTo:   &w.dataBinder,
			DataSource: w.progress,
		},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{
						Text: "Extracting:",
					},
					Label{
						AssignTo: &w.currentFile,
						Text:     "Starting...",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					ProgressBar{
						AssignTo:           &w.progressBar,
						MaxValue:           100,
						AlwaysConsumeSpace: true,
					},
					PushButton{
						AssignTo: &cancelButton,
						Text:     "Cancel",
						MaxSize:  Size{50, 50},
						OnClicked: func() {
							cancelButton.SetText("Canceling...")
							cancelButton.SetEnabled(false)
							w.requestedCancel = true
						},
					},
				},
			},
		},
	}

	_, err := main.Run()
	return err
}

func (w *progressWindow) Update(p progress.Progress) {
	*w.progress = p

	if w.progressBar != nil {
		w.progressBar.SetValue(p.Percentage)
	}

	if w.currentFile != nil {
		w.currentFile.SetText(p.CurrentFile)
	}

	if w.mainWindow != nil {
		pi := w.mainWindow.ProgressIndicator()
		pi.SetTotal(100)
		pi.SetCompleted(uint32(p.Percentage))
	}
}

func (w *progressWindow) Close() {
	for w.mainWindow == nil || w.err != nil {
	}
	w.mainWindow.Close()
}

func (w *progressWindow) RequestedCancel() bool {
	return w.requestedCancel
}
