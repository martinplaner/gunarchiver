// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trash

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTrash(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "gunarchiver")
	if err != nil {
		t.Fatalf("could not create temp file; %v", err)
	}
	file.WriteString("foobarbaz")
	if err := file.Close(); err != nil {
		t.Errorf("could not close temp file: %v", err)
	}
	defer func() {
		if err := os.Remove(file.Name()); err != nil && !os.IsNotExist(err) {
			t.Errorf("could not remove temp file: %v", err)
		}
	}()

	if err := MoveToTrash(file.Name()); err != nil {
		t.Errorf("could not execute syscall: %v", err)
	}

	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Errorf("temp file should be gone, but is still there")
	}
}
