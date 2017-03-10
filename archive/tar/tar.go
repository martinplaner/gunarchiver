// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zip provides support for reading and extracting ZIP archives.
package tar

import (
	"io"
	"os"
	"strings"

	"archive/tar"
	"compress/gzip"

	"github.com/martinplaner/gunarchiver/archive"
)

type tarArchive struct {
	file     *os.File
	gzip     *gzip.Reader
	tar      *tar.Reader
	numFiles int
}

func (a *tarArchive) Basename() string {
	return archive.Basename(a.file.Name())
}

func (a *tarArchive) NumFiles() int {
	return a.numFiles
}

func (a *tarArchive) Next() (archive.File, error) {
	hdr, err := a.tar.Next()
	if err != nil {
		return nil, err
	}

	return &file{
		tar: a.tar,
		hdr: hdr,
	}, nil
}

func (a *tarArchive) Reset() error {
	a.file.Seek(0, 0)
	r, err := gzip.NewReader(a.file)
	if err != nil {
		return err
	}

	a.tar = tar.NewReader(r)

	return nil
}

type file struct {
	tar *tar.Reader
	hdr *tar.Header
}

func (f *file) Path() string {
	return f.hdr.Name
}

func (f *file) Mode() os.FileMode {
	return f.hdr.FileInfo().Mode()
}

func (f *file) Read(p []byte) (n int, err error) {
	return f.tar.Read(p)
}

func (f *file) Close() error {
	// TODO: check if anything needs to be closed here
	return nil
}

func matchesTarGz(filename string, r io.Reader) bool {
	// Check zip suffix
	if strings.HasSuffix(strings.ToLower(filename), ".tar.gz") {
		//|| // Only .tar.gz for now...
		//strings.HasSuffix(strings.ToLower(filename), ".tar.bz2") {
		return true
	}

	r, err := gzip.NewReader(r)
	if err != nil {
		return false
	}

	tr := tar.NewReader(r)

	// TODO: Check file signature (many different versions :-/)
	// For now, just try reading and see if it fails

	_, err = tr.Next()
	if err != nil {
		return false
	}

	return true
}

func decodeTarGz(file *os.File) (archive.Archive, error) {
	r, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	tr := tar.NewReader(r)

	return &tarArchive{
		file: file,
		tar:  tr,
	}, nil
}

func init() {
	archive.RegisterFormat("tar.gz", matchesTarGz, decodeTarGz)
}
