# git diff-index --quiet HEAD --        // returns 1 if UNclean
# you can also use "go generate" to trigger a script that fills in a "version.go" file ???
# check for set SYSTEMROOT envvar --> windows

build:
	go build -ldflags="-H windowsgui"

resources:
	rsrc -manifest gunarchiver.manifest -ico gunarchiver.ico -o rsrc.syso
	
clean:
	rm gunarchiver.exe
	rm rsrc.syso

dist:
	zip gunarchiver.zip gunarchiver.exe README.md LICENSE
