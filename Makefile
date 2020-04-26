programname = zerostick

.PHONY: clean build run test

default: zerostick

zerostick: certs
	go build -a -o $(programname) *.go

certs:
	if [ ! -d zerostick_web/certs ]; then ./scripts/generate_certs.sh; fi

install:
	go install .

generate:
	mkdir -p build/bin
	go generate

image: build_arm6 rclone dms
	./scripts/build_image.sh

dist: image
	mv build/pi-gen/deploy/*.img .
	zip -9 ZeroStick.zip *.img

rclone:
	./scripts/build_rclone.sh

dms:
	./scripts/build_dms.sh

build_darwin: generate certs
	GOOS=darwin GOARCH=amd64 go build -tags=deploy_build -a -o ./build/bin/$(programname) *.go
	#zip ./build/$(programname)_darwin64.zip ./build/$(programname)

build_linux: generate certs
	GOOS=linux GOARCH=amd64 go build -tags=deploy_build -a -o ./build/bin/$(programname) *.go
	#zip ./build/$(programname)_linux64.zip ./build/$(programname)

# The Raspberry Pi Zero
build_arm6: generate certs
	GOOS=linux GOARM=6 GOARCH=arm go build -tags=deploy_build -a -o ./build/bin/$(programname) *.go
	#zip ./build/$(programname)_linux_arm6.zip ./build/$(programname)

build_arm7: generate certs
	GOOS=linux GOARM=7 GOARCH=arm go build -tags=deploy_build -a -o ./build/bin/$(programname) *.go
	#zip ./build/$(programname)_linux_arm7.zip ./build/$(programname)

build_win64: generate certs
	GOOS=windows GOARCH=amd64 go build -tags=deploy_build -a -o ./build/bin/$(programname).exe *.go
	#zip ./build/$(programname)_win64.zip ./build/$(programname).exe

build_win32: generate certs
	GOOS=windows GOARCH=386 go build -tags=deploy_build -a -o ./build/bin/$(programname).exe *.go
	#zip ./build/$(programname)_win32.zip ./build/$(programname).exe

for_snap: generate certs
	GOOS=linux GOARM=7 GOARCH=arm go build -tags=deploy_build -a -o ./$(programname) *.go


# all: build_darwin build_linux build_arm6 build_arm7 build_win64 build_win32
# 	rm ./build/$(programname)
# 	rm ./build/$(programname).exe

ui:
	cd ./zerostick_web/ui/; \
	if [ ! -d node_modules ]; then yarn install; fi ;\
	rm -rf build ;\
	yarn build

ui_dev:
	cd ./zerostick_web/ui; \
	yarn start

test:
	PATH="`pwd`/test/mock:$(PATH)" go test -v test/*.go

run: zerostick
	./$(programname) -d serve

runmock: zerostick
	PATH="`pwd`/test/mock:$(PATH)" ./$(programname) -d serve

# Development target; Build, push to zerostick.local and restart service
device: certs
	GOOS=linux GOARM=6 GOARCH=arm go build -tags=deploy_build -a -o ./build/bin/$(programname) *.go
	scp build/bin/zerostick pi@zerostick.local:
	scp -r zerostick_web pi@zerostick.local:
	ssh pi@zerostick.local "sudo mv zerostick /opt/zerostick/ && sudo rm -rf /opt/zerostick/zerostick_web && sudo mv zerostick_web /opt/zerostick/ && sudo systemctl restart zerostick.service"

clean:
	- rm -rf build
	- rm -f zerostick
	- docker rm -v pigen_work || true

real_clean: clean
	- rm -rf cache

