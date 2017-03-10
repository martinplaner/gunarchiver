// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trash provides the functionality to move files to the OS provided recycle bin (trash bin)
// instead of removing a file permanently through os.Remove.
// It is the user's responsibility to call Cleanup to unload the loaded System Libraries!
package trash
