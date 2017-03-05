package main

import (
	"fmt"
	"os"

	"io"

	"github.com/martinplaner/gunarchiver/archive"
	_ "github.com/martinplaner/gunarchiver/archive/zip"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <archive>\n", os.Args[0])
		return
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("could not open file:", err)
		return
	}
	defer f.Close()

	a, format, err := archive.Decode(f)
	if err != nil {
		fmt.Println("could not decode archive:", err)
		return
	}
	fmt.Println("Opened file as archive format:", format)
	fmt.Println("Has single root?:", archive.HasSingleRoot(a))

	for {
		file, err := a.Next()
		if err == io.EOF {
			// end of archive
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(file.Path(), file.Mode())
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
}
