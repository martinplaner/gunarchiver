# git diff-index --quiet HEAD --        // returns 1 if UNclean
# you can also use "go generate" to trigger a script that fills in a "version.go" file ???
# check for set SYSTEMROOT envvar --> windows

LD_FLAGS = -s -w -H windowsgui

build:
	go build -ldflags="$(LD_FLAGS)"

# Needs: go get github.com/akavel/rsrc
resources:
	rsrc -manifest gunarchiver.manifest -ico gunarchiver.ico -o rsrc.syso
	
clean:
	rm gunarchiver.exe
	rm rsrc.syso

dist:
	zip gunarchiver.zip gunarchiver.exe README.md LICENSE
