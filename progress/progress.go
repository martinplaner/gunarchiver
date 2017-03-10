// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package progress provides data types for sharing progress between components.
package progress

type UpdateCloser interface {
	Update(Progress)
	Close()
}

// Progress is the main data structure for sharing progress.
type Progress struct {
	Percentage  int
	CurrentFile string
}

// Sync receives progress updates through a channel and applies them to an UpdateCloser.
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
