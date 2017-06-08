// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package archive implements functions for decoding and extracting archives such as ZIP, tar.gz or RAR.
package archive

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/martinplaner/gunarchiver/debug"
	"github.com/martinplaner/gunarchiver/progress"
)

// ErrCanceled is the error returned by Extract, when the process was canceled by the user
var ErrCanceled = errors.New("the process was canceled by user request")

// IsCanceled returns true if err represents that process was canceled by the user
func IsCanceled(err error) bool {
	return err == ErrCanceled
}

// Archive represent an archive and provides methods for iterating through its files.
type Archive interface {
	// Basename returns the base name of the archive file, without file extension or path.
	Basename() string
	// NumFiles returns the total number of files in the archive
	NumFiles() int
	// Next returns the next file in the archive. err == io.EOF on end of archive.
	Next() (File, error)
	// Reset resets the current position to the start of the archive.
	Reset() error
}

// FileHeader contains basic information about a file in the archive, such as Name, Path, and FileMode.
type FileHeader interface {
	Path() string // Path returns the file path in the archive. Only forward slashes are allowed as separators.
	Mode() os.FileMode
}

// File is equivalent to FileHeader but also provides a io.ReadCloser for accessing the file's content.
type File interface {
	FileHeader
	io.ReadCloser
}

// Extract extracts the Archive a to the specified path and reports progress to the supplied progressChan.
// If the supplied shouldCancel func returns true, the extraction will get canceled as soon as possible
// and return ErrCanceled. Callers should check using IsCanceled.
func Extract(a Archive, path string, progressChan chan progress.Progress, shouldCancel func() bool) error {
	progressChan <- progress.Progress{CurrentFile: "Starting extraction..."}

	currentFileNum := 0
	percentage := 0
	for {
		if shouldCancel() {
			// Aborting extraction as per user's request
			progressChan <- progress.Progress{CurrentFile: "Canceled extraction!", Percentage: 100}
			return ErrCanceled
		}
		file, err := a.Next()
		if err == io.EOF {
			// end of archive; finished
			progressChan <- progress.Progress{CurrentFile: "Finished extraction!", Percentage: 100}
			return nil
		}
		if err != nil {
			return err
		}

		currentFileNum++
		progressChan <- progress.Progress{CurrentFile: file.Path(), Percentage: percentage}
		debug.Wait(1 * time.Second)

		fullPath := filepath.Join(path, file.Path())

		switch {
		case file.Mode().IsDir():
			if err := CreateDir(fullPath); err != nil {
				return err
			}

		case file.Mode().IsRegular():
			if err := extractFile(file, fullPath); err != nil {
				return err
			}

			// TODO: handle hard-/symlinks and special files?
		}

		percentage = calcPercentage(currentFileNum, a.NumFiles())
		progressChan <- progress.Progress{CurrentFile: file.Path(), Percentage: percentage}
		file.Close()
	}
}

func calcPercentage(current, total int) int {
	if total == 0 {
		return 0
	}
	return int(float64(current) / float64(total) * 100)
}

func extractFile(file File, path string) error {
	if err := CreateDir(filepath.Dir(path)); err != nil {
		return err
	}

	out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, file.Mode())
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		return fmt.Errorf("could not write file: %v", err)
	}

	return nil
}

// Directory permissions for new folders; before umask
const dirPerm = 0777

// CreateDir creates the specified directory along with any necessary parent directories with 0777 permissions (before umask).
func CreateDir(path string) error {
	return os.MkdirAll(path, dirPerm)
}

// Basename returns the base name for a given file path, without a file extension or the path to the containing directory.
// e.g. Basename("/home/user/test.txt") returns "test" and Basename("user/test.tar.gz") returns "test".
// In general only single file extensions are supported, but this implementation contains a special case to remove the optional ".tar.*" suffix.
func Basename(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]

	// Remove additional ".tar" suffix, if present (.tar.gz, .tar.bz2, ...)
	if strings.HasSuffix(strings.ToLower(base), ".tar") {
		base = base[:len(base)-4]
	}

	return base
}

// HasSingleRoot returns true, if the archive contains only a single file/directory at the root of the archive.
// e.g. [/README, /a.bin] -> false, [/a/README, /a/a.bin] -> true.
// WARNING: This iterates the archive from the beginning and resets the current position!
func HasSingleRoot(a Archive) bool {
	a.Reset()
	defer a.Reset()
	rootElems := make(map[string]bool)
	numElems := 0
	for {
		file, err := a.Next()
		if err == io.EOF {
			return true
		}
		if err != nil {
			return false
		}
		numElems++
		first := strings.Split(file.Path(), "/")[0]
		if numElems > 1 && !rootElems[first] {
			return false
		}
		rootElems[first] = true
	}
}
