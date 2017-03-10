// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
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

// RegisterFormat registers an archive format for use by Decode.
// Name is the name of the format, like "zip", "tar" or "rar".
// Matches is a function that determines if the format is capable of reading the archive from reader r.
// Decode is the function that decodes the archive.
func RegisterFormat(name string, matches func(filename string, r io.Reader) bool, decode func(*os.File) (Archive, error)) {
	formats = append(formats, format{name, matches, decode})
}

// Decode decodes an archive in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Decode(file *os.File) (Archive, string, error) {
	format := detect(file)
	if format.decode == nil {
		return nil, "", ErrUnknownFormat
	}
	m, err := format.decode(file)
	return m, format.name, err
}

// detect tries to match the given archive file with a registered format.
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
