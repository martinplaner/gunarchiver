# gunarchiver

gunarchiver (Go Unarchiver) is a minimalistic archive extraction tool with a few simple goals:

- Single static binary (no external dependencies) with simple user interaction (just double click archive file).
- Extract all contained file to the same folder as the archive.
- Create a new folder, if there is more than on top-level file in the archive.
- Move the archive to the recycle bin after successful extraction.

# Future Work:

(not necessarily implemented soon or ever)

- Handle symlinks and hardlinks properly (currently not supported)
- UI and trash implementation for other platforms (Linux, macOS). But low priority since other platform already have similar (better) tools that inspired me to develop this one.
- Preferences pane for customization (+ persistent config).
- i18n (currently hardcoded to english)
- Add (nicer) application icon
- Fine(r) grained progress updates (currently updates only after every file in the archive)

# License

Copyright 2017 Martin Planer. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
