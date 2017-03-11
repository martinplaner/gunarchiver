# gunarchiver

gunarchiver (Go Unarchiver) is a minimalistic archive extraction tool with a few simple goals:

- Single static binary (no external dependencies) with simple user interaction (just double click archive file).
- Extract all contained file to the same folder as the archive.
- Create a new folder, if there is more than on top-level file in the archive.
- Move the archive to the recycle bin after successful extraction.

This was primarily a project born out of the frustration of having to navigate through several layers of context menus of 7zip
and manually moving files, just to have them structured the way I like it. This tool does it all in one double-click.

![screenshot](https://raw.githubusercontent.com/martinplaner/gunarchiver/master/doc/screenshot.png)

NOTE: The program is currently windows only due to better alternatives already available for other platforms.

## Supported formats

Currently supported formats are: .zip, .tar.gz and .rar

No support for multi-volume archives or password protected files (yet).

## Installation

Either `go get github.com/martinplaner/gunarchiver` and then `go build -ldflags="-H windowsgui"` in the project directory OR download the pre-compiled binary from the GitHub releases page.

Put the binary anywhere you want and associate the proper file extensions with the program.
You're done, no third step.

## Possible future work and known problems:

(not necessarily implemented soon or ever -- works for me in the current state ;)

- Handle symlinks and hardlinks properly (currently not supported)
- UI and trash implementation for other platforms (Linux, macOS). But low priority since other platform already have similar (better) tools that inspired me to develop this one.
- Preferences pane for customization (+ persistent config).
- i18n (currently hardcoded to english)
- Add (nicer) application icon
- Fine(r) grained progress updates (currently updates only after every file in the archive)
- Maybe multi-volume RAR support

## License

Copyright 2017 Martin Planer. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
