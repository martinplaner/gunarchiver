// +build !debug

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "time"

// IsDebug returns true if the program is executed in debug mode.
func IsDebug() bool {
	return false
}

// Wait pauses the execution for the given duration. The effect only really applies when running in debug mode.
func Wait(delay time.Duration) {}
