package archive

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Archive interface {
	Basename() string

	// Next returns the next file in the archive. err == io.EOF on end of archive.
	Next() (File, error)
	Reset() error
}

type File interface {
	Path() string // Path returns the file path in the archive. Only forward slashes are allowed as separators.
	Mode() os.FileMode
	io.ReadCloser
}

func Extract(a Archive, path string) error {
	for {
		_ /*file*/, err := a.Next()
		if err == io.EOF {
			// end of archive; finished
			return nil
		}
		if err != nil {
			return err
		}

	}
}

// Basename returns the base name for a given file path, without a file extension or the path to the containing directory.
// e.g. Basename("/home/user/test.txt") returns "test" and Basename("user/test.tar.gz") returns "test"
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
	return true
}
