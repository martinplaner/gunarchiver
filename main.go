package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/zip"
)

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <archive>\n", os.Args[0])
		return
	}

	archivePath := os.Args[1]

	err = extractArchive(archivePath)
	if err != nil {
		log.Fatalln(err)
		//showError(...)
		//os.Exit(1)
	}

	if err := os.Remove(archivePath); err != nil {
		log.Fatalln(err)
	}
}

func extractArchive(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	a, _, err := archive.Decode(f)
	if err != nil {
		return fmt.Errorf("could not decode archive: %v", err)
	}

	baseDir := filepath.Dir(path)
	singleRoot := archive.HasSingleRoot(a)

	if !singleRoot {
		baseDir = filepath.Join(baseDir, a.Basename())
		archive.CreateDir(baseDir)
	}

	if err := archive.Extract(a, baseDir); err != nil {
		return fmt.Errorf("could not extract archive: %v", err)
	}

	return nil
}

//var progressWindow ui.ProgressWindow = windows.NewProgressWindow()
//p := ui.Progress{}
//
//go func() {
//	//<-time.After(2 * time.Second)
//	for p.Percentage < 100 {
//		<-time.After(500 * time.Millisecond)
//		p.Percentage += 10
//		progressWindow.Update(p)
//
//		if progressWindow.RequestedCancel() {
//			break
//		}
//	}
//	progressWindow.Close()
//}()
//
//progressWindow.Show(p)
