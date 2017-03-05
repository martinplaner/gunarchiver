// +build windows

package windows

import (
	"context"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/martinplaner/gunarchiver/progress"
	"github.com/martinplaner/gunarchiver/ui"
)

type progressWindow struct {
	mainWindow      *walk.MainWindow
	progressBar     *walk.ProgressBar
	requestedCancel bool
	cancel          context.CancelFunc
	err             error
}

func (w *progressWindow) Show() error {

	var cancelButton *walk.PushButton

	main := MainWindow{
		AssignTo: &w.mainWindow,
		Title:    "gunarchiver",
		Size:     Size{400, 100},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ProgressBar{
						AssignTo: &w.progressBar,
						MaxValue: ui.ProgressMaxValue,
					},
					PushButton{
						AssignTo: &cancelButton,
						Text:     "Cancel",
						OnClicked: func() {
							cancelButton.SetText("Canceling...")
							cancelButton.SetEnabled(false)
							w.requestedCancel = true
							w.cancel()
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
	if w.progressBar != nil {
		w.progressBar.SetValue(p.Percentage)
	}

	if w.mainWindow != nil {
		pi := w.mainWindow.ProgressIndicator()
		pi.SetTotal(uint32(ui.ProgressMaxValue))
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
