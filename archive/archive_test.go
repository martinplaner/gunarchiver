// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package archive_test

// Tests for archive extraction.
// Currently file system dependent, since the zip reader needs os.File.
// This may change in the future if I find a way to abstract this away.

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/martinplaner/gunarchiver/progress"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/rar"
	_ "github.com/martinplaner/gunarchiver/archive/tar"
	_ "github.com/martinplaner/gunarchiver/archive/zip"
)

var testDataDir = "testdata"
var testFiles = []struct {
	basename string
	numFiles int
}{
	{
		"single",
		1,
	},
	{
		"multiple",
		3,
	},
	{
		"deep",
		7,
	},
	{
		"subfolder",
		7,
	},
}
var formats = map[string][]string{
	"zip":    {"zip"},
	"tar.gz": {"tar.gz"},
	"rar":    {"rar"},
}

func TestNumFiles(t *testing.T) {
	for _, exts := range formats {
		for _, ext := range exts {
			for _, testFile := range testFiles {
				filename := testFile.basename + "." + ext
				t.Run(filename, func(t *testing.T) {
					if err := testNumFiles(filename, testFile.numFiles); err != nil {
						t.Errorf("TestNumFiles - %s: %v", filename, err)
					}
				})
			}
		}
	}
}

func testNumFiles(filename string, expected int) error {
	path := filepath.Join(testDataDir, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	a, _, err := archive.Decode(file)
	if err != nil {
		return err
	}

	if got := a.NumFiles(); got != expected {
		return fmt.Errorf("numFiles mismatch: got %d, expected %d", got, expected)
	}

	return nil
}

func TestExtraction(t *testing.T) {
	for _, exts := range formats {
		for _, ext := range exts {
			for _, testFile := range testFiles {
				filename := testFile.basename + "." + ext
				t.Run(filename, func(t *testing.T) {
					if err := testExtraction(testFile.basename, ext); err != nil {
						t.Errorf("TestExtraction - %s: %v", filename, err)
					}
				})
			}
		}
	}
}

func testExtraction(basename, ext string) error {
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

	if err := compareDirs(tempDir, comparePath); err != nil {
		return fmt.Errorf("error comparing dir structure: %v", err)
	}

	return nil
}

func compareDirs(a, b string) error {
	aFiles, aErr := getFileList(a)
	bFiles, bErr := getFileList(b)

	if aErr != nil {
		return aErr
	}
	if bErr != nil {
		return bErr
	}

	aLen := len(aFiles)
	bLen := len(bFiles)
	if aLen != bLen {
		return fmt.Errorf("file list length does not match: %d != %d", aLen, bLen)
	}

	for i := range aFiles {
		aRel, aErr := filepath.Rel(a, aFiles[i])
		bRel, bErr := filepath.Rel(b, bFiles[i])

		if aErr != nil {
			return aErr
		}
		if bErr != nil {
			return bErr
		}

		if aRel != bRel {
			return fmt.Errorf("relative file paths do not match")
		}
	}

	return nil
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
