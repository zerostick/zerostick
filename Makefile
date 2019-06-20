programname = zerostick

.PHONY: clean build run

default: zerostick

zerostick: certs
	go build -a -o $(programname) *.go

certs:
	if [ ! -d zerostick_web/certs ]; then ./generate_certs.sh; fi

install:
	go install .

generate:
	mkdir -p build
	go generate

build_darwin: generate
	GOOS=darwin GOARCH=amd64 go build -tags=deploy_build -a -o ./build/$(programname) *.go
	zip ./build/$(programname)_darwin64.zip ./build/$(programname)

build_linux:
	GOOS=linux GOARCH=amd64 go build -tags=deploy_build -a -o ./build/$(programname) *.go
	zip ./build/$(programname)_linux64.zip ./build/$(programname)

# The Raspberry Pi Zero
build_arm5:
	GOOS=linux GOARM=5 GOARCH=arm go build -tags=deploy_build -a -o ./build/$(programname) *.go
	zip ./build/$(programname)_linux_arm5.zip ./build/$(programname)

build_arm7:
	GOOS=linux GOARM=7 GOARCH=arm go build -tags=deploy_build -a -o ./build/$(programname) *.go
	zip ./build/$(programname)_linux_arm7.zip ./build/$(programname)

build_win64:
	GOOS=windows GOARCH=amd64 go build -tags=deploy_build -a -o ./build/$(programname).exe *.go
	zip ./build/$(programname)_win64.zip ./build/$(programname).exe

build_win32:
	GOOS=windows GOARCH=386 go build -tags=deploy_build -a -o ./build/$(programname).exe *.go
	zip ./build/$(programname)_win32.zip ./build/$(programname).exe

all: build_darwin build_linux build_arm5 build_arm7 build_win64 build_win32
	rm ./build/$(programname)
	rm ./build/$(programname).exe

run: zerostick
	./$(programname)

clean:
	- rm -rf build
	- rm -f zerostick
