package ui

var ProgressMaxValue = 100

type Window interface {
	Close()
}

type Progress struct {
	Percentage  int
	CurrentFile string
}

type ProgressWindow interface {
	Window
	Show(progress Progress) error
	Update(progress Progress)
	RequestedCancel() bool
}
