// +build windows

package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type errorWindow struct {
	mainWindow *walk.MainWindow
}

func (w *errorWindow) Show(message string) error {

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
						Text: message,
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
