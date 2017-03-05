package main

import (
	"fmt"
	"os"
	"path/filepath"

	"log"

	"context"

	"sync"

	"time"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/zip"
	"github.com/martinplaner/gunarchiver/progress"
	"github.com/martinplaner/gunarchiver/ui"
)

func main() {
	var extractErr error
	var uiErr error
	var wg sync.WaitGroup

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <archive>\n", os.Args[0])
		return
	}

	archivePath := os.Args[1]
	pChan := make(chan progress.Progress)
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), progress.ProgressChan, pChan))

	wg.Add(1)
	go func() {
		extractErr = extractArchiveAndDelete(ctx, archivePath, true)
		pChan <- progress.Progress{Done: true}
		close(pChan)
		wg.Done()
	}()

	progressWindow := ui.Default.NewProgressWindow(cancel)

	wg.Add(1)
	go func() {
		var currentProgress progress.Progress
		var progressInterval = 500 * time.Millisecond
	LOOP:
		for {
			select {
			case p, ok := <-pChan:
				if !ok {
					break LOOP
				}
				currentProgress = p
			case <-time.After(progressInterval):
				progressWindow.Update(currentProgress)
			}
		}
		progressWindow.Close()
		wg.Done()
	}()

	uiErr = progressWindow.Show()

	wg.Wait()

	if extractErr != nil {
		errorWindow := ui.Default.NewErrorWindow()
		errorWindow.Show(extractErr.Error())
		log.Fatalln(extractErr)
	}

	if uiErr != nil {
		errorWindow := ui.Default.NewErrorWindow()
		errorWindow.Show(uiErr.Error())
		log.Fatalln(uiErr)
	}
}

func extractArchiveAndDelete(ctx context.Context, path string, deleteAfter bool) error {
	if err := extractArchive(ctx, path); err != nil {
		return err
	}

	if deleteAfter {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}

func extractArchive(ctx context.Context, path string) error {
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

	if err := archive.Extract(ctx, a, baseDir); err != nil {
		return fmt.Errorf("could not extract archive: %v", err)
	}

	return nil
}
