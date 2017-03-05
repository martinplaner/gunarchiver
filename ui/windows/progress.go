// +build windows

package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/martinplaner/gunarchiver/ui"
)

type ProgressWindow struct {
	mainWindow      *walk.MainWindow
	progressBar     *walk.ProgressBar
	requestedCancel bool
}

func NewProgressWindow() *ProgressWindow {
	return &ProgressWindow{}
}

func (w *ProgressWindow) Show(progress ui.Progress) error {

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
						},
					},
				},
			},
		},
	}

	_, err := main.Run()
	return err
}

func (w *ProgressWindow) Update(p ui.Progress) {
	w.progressBar.SetValue(p.Percentage)

	pi := w.mainWindow.ProgressIndicator()
	pi.SetTotal(uint32(ui.ProgressMaxValue))
	pi.SetCompleted(uint32(p.Percentage))
}

func (w *ProgressWindow) Close() {
	w.mainWindow.Close()
}

func (w *ProgressWindow) RequestedCancel() bool {
	return w.requestedCancel
}
