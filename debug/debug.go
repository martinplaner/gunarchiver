// +build debug

// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "time"

func IsDebug() bool {
	return true
}

func Wait(delay time.Duration) {
	<-time.After(delay)
}
