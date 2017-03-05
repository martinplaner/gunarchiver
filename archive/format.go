// Inspired by https://golang.org/src/image/format.go, licensed under:
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package archive

import (
	"errors"
	"io"

	"os"
)

// ErrUnknownFormat indicates an unknown archive format.
var ErrUnknownFormat = errors.New("archive: unknown format")

type format struct {
	name    string
	matches func(filename string, r io.Reader) bool
	decode  func(*os.File) (Archive, error)
}

var formats []format

func RegisterFormat(name string, matches func(filename string, r io.Reader) bool, decode func(*os.File) (Archive, error)) {
	formats = append(formats, format{name, matches, decode})
}

func Decode(file *os.File) (Archive, string, error) {
	format := detect(file)
	if format.decode == nil {
		return nil, "", ErrUnknownFormat
	}
	m, err := format.decode(file)
	return m, format.name, err
}

func detect(file *os.File) format {
	for _, format := range formats {
		matches := format.matches(file.Name(), file)
		file.Seek(0, 0)
		if matches {
			return format
		}
	}
	return format{}
}
