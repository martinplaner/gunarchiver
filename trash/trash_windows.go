// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows specific trash implementation using WINAPI calls.

package trash

import (
	"fmt"
	"syscall"
	"unsafe"
)

var ()

//type (
//	HWND         uintptr
//	UINT         uint32
//	PCZZTSTR     *uint16
//	FILEOP_FLAGS uint16
//)

type SHFILEOPSTRUCT struct {
	hwnd                  uintptr // HWND         hwnd;
	wFunc                 uint32  // UINT         wFunc; -- 32bit; win32 api after all...
	pFrom                 *uint16 // PCZZTSTR     pFrom; -- double zero terminated string; 16bit char pointer??? (let's try this...)
	pTo                   *uint16 // PCZZTSTR     pTo;
	fFlags                uint16  // FILEOP_FLAGS fFlags; -- typedef WORD FILEOP_FLAGS
	fAnyOperationsAborted bool    // BOOL         fAnyOperationsAborted;
	hNameMappings         uintptr // LPVOID       hNameMappings;
	lpszProgressTitle     *uint16 // PCTSTR       lpszProgressTitle; -- also 16bit char pointer???
}

const (
	_ = iota
	FO_MOVE
	FO_COPY
	FO_DELETE
	FO_RENAME
)

const (
	FOF_ALLOWUNDO = 0x0040
)

func MoveToTrash(path string) error {
	// TODO: Investigate potential DLL preloading attacks
	shell32, err := syscall.LoadDLL("Shell32.dll")
	if err != nil {
		return err
	}
	defer syscall.FreeLibrary(shell32.Handle)
	shFileOperationW, err := syscall.GetProcAddress(shell32.Handle, "SHFileOperationW")
	if err != nil {
		return err
	}

	pFrom, err := syscall.UTF16PtrFromString(path)

	if err != nil {
		return fmt.Errorf("could not convert path to WINAPI compatible pointer: %v", err)
	}

	fileOp := &SHFILEOPSTRUCT{
		//hwnd : nil?,
		wFunc: FO_DELETE,
		pFrom: pFrom, // how to convert Go string to *uint16 ?
		//pTo : nil,                 // should be set to NULL according to doc. == nil?
		fFlags: FOF_ALLOWUNDO,
		// fAnyOperationsAborted:  // out-Param? result
	}

	_, _, errno := syscall.Syscall(shFileOperationW, 1, uintptr(unsafe.Pointer(fileOp)), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}
