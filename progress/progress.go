package progress

const ProgressChan = "progress"

type Progress struct {
	Percentage  int
	CurrentFile string
	Done        bool
}
