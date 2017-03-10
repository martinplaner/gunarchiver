package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/zip"

	"github.com/martinplaner/gunarchiver/progress"
	"github.com/martinplaner/gunarchiver/ui"
)

var userInterface ui.UserInterface

func main() {
	var extractErr error
	var uiErr error

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <archive>\n", os.Args[0])
		return
	}

	archivePath := os.Args[1]
	progressChan := make(chan progress.Progress)
	progressWindow := userInterface.NewProgressWindow()

	// Kick off extraction
	go func() {
		extractErr = extractArchiveAndDelete(archivePath, progressChan, progressWindow.RequestedCancel)
		close(progressChan)
	}()

	// Synchronize extraction and UI progress
	go progress.Sync{
		UpdateCloser: progressWindow,
		Progress:     progressChan,
	}.Run()

	uiErr = progressWindow.Show()

	if extractErr != nil {
		errorWindow := userInterface.NewErrorWindow(extractErr.Error())
		errorWindow.Show()
		log.Fatalln("could not extract archive:", extractErr)
	}

	if uiErr != nil {
		log.Fatalln("could not show user interface:", uiErr)
	}
}

func extractArchiveAndDelete(path string, progressChan chan progress.Progress, shouldCancel func() bool) error {
	if err := extractArchive(path, progressChan, shouldCancel); err != nil {
		return err
	}

	// TODO: uncomment [DEBUG]
	//if err := trash.MoveToTrash(path); err != nil {
	//	return err
	//}

	return nil
}

func extractArchive(path string, progressChan chan progress.Progress, shouldCancel func() bool) error {
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

	if err := archive.Extract(a, baseDir, progressChan, shouldCancel); err != nil {
		return fmt.Errorf("could not extract archive: %v", err)
	}

	return nil
}
