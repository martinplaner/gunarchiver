package progress

type UpdateCloser interface {
	Update(Progress)
	Close()
}

type Progress struct {
	Percentage  int
	CurrentFile string
}

type Sync struct {
	UpdateCloser UpdateCloser
	Progress     chan Progress
}

func (s Sync) Run() {
	for p := range s.Progress {
		s.UpdateCloser.Update(p)
	}
	s.UpdateCloser.Close()
}
