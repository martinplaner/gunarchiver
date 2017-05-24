// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rar provides support for reading and extracting ZIP archives.
package rar

import (
	"io"
	"os"
	"strings"

	"github.com/martinplaner/gunarchiver/archive"
	"github.com/nwaples/rardecode"
)

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type rarArchive struct {
	file     *os.File
	rar      *rardecode.Reader
	numFiles int
}

func (a *rarArchive) Basename() string {
	return archive.Basename(a.file.Name())
}

func (a *rarArchive) NumFiles() int {
	return a.numFiles
}

func (a *rarArchive) Next() (archive.File, error) {
	hdr, err := a.rar.Next()
	if err != nil {
		return nil, err
	}

	return &file{
		rar: a.rar,
		hdr: hdr,
	}, nil
}

func (a *rarArchive) Reset() error {
	a.file.Seek(0, 0)
	r, err := rardecode.NewReader(a.file, "")
	if err != nil {
		return err
	}
	a.rar = r

	return nil
}

type file struct {
	rar *rardecode.Reader
	hdr *rardecode.FileHeader
}

func (f *file) Path() string {
	return f.hdr.Name
}

func (f *file) Mode() os.FileMode {
	return f.hdr.Mode()
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.rar.Read(p)
}

func (f *file) Close() error {
	// TODO: check if anything needs to be closed here
	return nil
}

func matches(filename string, r io.Reader) bool {
	// Check zip suffix
	if strings.HasSuffix(strings.ToLower(filename), ".rar") {
		return true
	}

	_, err := rardecode.NewReader(r, "")
	if err != nil {
		return false
	}

	return true
}

func decode(file *os.File) (archive.Archive, error) {

	r, err := rardecode.NewReader(file, "")
	if err != nil {
		return nil, err
	}

	// TODO: implement numFiles

	return &rarArchive{
		file: file,
		rar:  r,
	}, nil
}

func init() {
	archive.RegisterFormat("rar", matches, decode)
}
