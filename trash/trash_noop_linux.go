// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux dummy/noop implementation of the trash api allow build on linux based CI container.

package trash

// MoveToTrash should move the given file to the OS provided trash.
// THIS IS A NOOP IMPLEMENTATION TO ALLOW COMPILATION ON LINUX. DOES NOT ACTUALLY WORK!
func MoveToTrash(path string) error {
	return nil
}
