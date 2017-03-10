// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zip provides support for reading and extracting ZIP archives.
package zip

import (
	"archive/zip"

	"bytes"
	"io"
	"os"
	"strings"

	"github.com/martinplaner/gunarchiver/archive"
)

const zipFileSignature = "PK\x03\x04"

type zipArchive struct {
	file     *os.File
	zip      *zip.Reader
	current  int
	numFiles int
}

func (a *zipArchive) Basename() string {
	return archive.Basename(a.file.Name())
}

func (a *zipArchive) NumFiles() int {
	return a.numFiles
}

func (a *zipArchive) Next() (archive.File, error) {
	if a.current > len(a.zip.File)-1 {
		return nil, io.EOF
	}

	f := a.zip.File[a.current]
	a.current++

	return &file{
		zip: f,
	}, nil
}

func (a *zipArchive) Reset() error {
	a.file.Seek(0, 0)
	stat, err := a.file.Stat()
	if err != nil {
		return err
	}

	r, err := zip.NewReader(a.file, stat.Size())
	if err != nil {
		return err
	}

	a.zip = r
	a.current = 0
	return nil
}

type file struct {
	zip    *zip.File
	r      io.ReadCloser
	isOpen bool
}

func (f *file) Path() string {
	return f.zip.Name
}

func (f *file) Mode() os.FileMode {
	return f.zip.Mode()
}

func (f *file) Read(p []byte) (n int, err error) {
	if !f.isOpen {
		r, err := f.zip.Open()
		if err != nil {
			return 0, err
		}
		f.r = r
		f.isOpen = true
	}

	return f.r.Read(p)
}

func (f *file) Close() error {
	err := f.r.Close()
	if err != nil {
		return err
	}

	f.r = nil
	f.isOpen = false
	return nil
}

func matches(filename string, r io.Reader) bool {
	// Check zip suffix
	if strings.HasSuffix(strings.ToLower(filename), ".zip") {
		return true
	}

	// Check zip signature
	buf := make([]byte, 4)
	n, err := r.Read(buf)
	if err != nil || n < len(buf) {
		return false
	}

	return bytes.Equal(buf, []byte(zipFileSignature))
}

func decode(file *os.File) (archive.Archive, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return nil, err
	}

	numFiles := 0
	for range reader.File {
		numFiles++
	}

	return &zipArchive{
		file:     file,
		zip:      reader,
		numFiles: numFiles,
	}, nil
}

func init() {
	archive.RegisterFormat("zip", matches, decode)
}
