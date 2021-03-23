
.PHONY: shared 
shared:
	mkdir -p dist;
	# OSX
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o dist/lib-osx.so -buildmode=c-shared c-main/main.go;
	# Linux (fails on my mac)
	# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o dist/lib-linux.so -buildmode=c-shared c-main/main.go;
	# Raspberry (fails on my mac)
	# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o dist/lib-rpi.so -buildmode=c-shared c-main/main.go;