// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package archive_test

// Tests for archive extraction.
// Currently file system dependent, since the zip reader needs os.File.
// This may change in the future if I find a way to abstract this away.

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"errors"

	"github.com/martinplaner/gunarchiver/progress"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/rar"
	_ "github.com/martinplaner/gunarchiver/archive/tar"
	_ "github.com/martinplaner/gunarchiver/archive/zip"
)

var testDataDir = "testdata"
var basenames = []string{"single", "multiple", "deep", "subfolder"}
var formats = map[string][]string{
	"zip":    []string{"zip"},
	"tar.gz": []string{"tar.gz"},
	"rar":    []string{"rar"},
}

func TestArchives(t *testing.T) {
	for _, exts := range formats {
		for _, ext := range exts {
			for _, basename := range basenames {
				t.Run(basename+"."+ext, func(t *testing.T) {
					if err := testArchive(basename, ext); err != nil {
						t.Error(err)
					}
				})
			}
		}
	}
}

func testArchive(basename, ext string) error {
	filename := basename + "." + ext
	path := filepath.Join(testDataDir, filename)
	comparePath := filepath.Join(testDataDir, basename)

	tempDir, err := ioutil.TempDir(os.TempDir(), "gunarchiver")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	a, _, err := archive.Decode(file)
	if err != nil {
		return err
	}

	nopCancel := func() bool { return false }
	pChan := make(chan progress.Progress, 10000) // Leave enough room that this does not block during testing
	if err := archive.Extract(a, tempDir, pChan, nopCancel); err != nil {
		return err
	}

	if !compareDirs(tempDir, comparePath) {
		return errors.New("dirs to not match!")
	}

	return nil
}

func compareDirs(a, b string) bool {
	aFiles, aErr := getFileList(a)
	bFiles, bErr := getFileList(b)

	if aErr != nil || bErr != nil {
		return false
	}

	if len(aFiles) != len(bFiles) {
		return false
	}

	for i, _ := range aFiles {
		aRel, aErr := filepath.Rel(a, aFiles[i])
		bRel, bErr := filepath.Rel(b, bFiles[i])

		if aErr != nil || bErr != nil {
			return false
		}

		if aRel != bRel {
			return false
		}
	}

	return true
}

func getFileList(path string) ([]string, error) {
	var files []string
	add := func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}
	if err := filepath.Walk(path, add); err != nil {
		return nil, err
	}

	// Return all but root dir
	return files[1:], nil
}
