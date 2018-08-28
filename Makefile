VERSION=`git rev-parse --short HEAD`

deps:
	dep ensure

test: deps
	go test ./...

build-darwin: deps
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags="-X main.Version=$(VERSION)" -o remap

build-linux: deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X main.Version=$(VERSION)" -o remap

build-windows: deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X main.Version=$(VERSION)" -o remap.exe

dist: deps
	rm -rf build
	rm -rf dist

	mkdir -p build/linux/amd64 && GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/linux/amd64/remap 
	mkdir -p build/linux/i386 && GOOS=linux GOARCH=386 go build -a -installsuffix cgo -o build/linux/i386/remap 
	mkdir -p build/linux/armel && GOOS=linux GOARCH=arm GOARM=5 go build -a -installsuffix cgo -o build/linux/armel/remap 
	mkdir -p build/linux/armhf && GOOS=linux GOARCH=arm GOARM=6 go build -a -installsuffix cgo -o build/linux/armhf/remap 
	mkdir -p build/linux/arm-7 && GOOS=linux GOARCH=arm GOARM=7 go build -o build/linux/arm-7/remap 
	mkdir -p build/linux/arm64 && GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o build/linux/arm64/remap 
	mkdir -p build/darwin/amd64 && GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o build/darwin/amd64/remap 
	mkdir -p build/darwin/i386 && GOOS=darwin GOARCH=386 go build -a -installsuffix cgo -o build/darwin/i386/remap 
	mkdir -p build/windows/i386 && GOOS=windows GOARCH=386 go build -a -installsuffix cgo -o build/windows/i386/remap.exe 
	mkdir -p build/windows/amd64 && GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o build/windows/amd64/remap.exe 

	mkdir -p dist/

	tar -cvzf dist/remap-linux-amd64-$(VERSION).tar.gz -C build/linux/amd64 remap
	tar -cvzf dist/remap-linux-i386-$(VERSION).tar.gz -C build/linux/i386 remap
	tar -cvzf dist/remap-linux-armel-$(VERSION).tar.gz -C build/linux/armel remap
	tar -cvzf dist/remap-linux-armhf-$(VERSION).tar.gz -C build/linux/armhf remap
	tar -cvzf dist/remap-linux-arm-7-$(VERSION).tar.gz -C build/linux/arm-7 remap
	tar -cvzf dist/remap-linux-arm64-$(VERSION).tar.gz -C build/linux/arm64 remap
	tar -cvzf dist/remap-darwin-amd64-$(VERSION).tar.gz -C build/darwin/amd64 remap
	tar -cvzf dist/remap-darwin-i386-$(VERSION).tar.gz -C build/darwin/i386 remap
	zip dist/remap-windows-i386-$(VERSION).zip build/windows/i386/remap.exe
	zip dist/remap-windows-amd64-$(VERSION).zip build/windows/amd64/remap.exe
	rm -rf build