SHELL=/bin/bash


NAME = metrilyx-dashboarder
BUILD_DIR = ./build/metrilyx

INSTALL_DIR = ${BUILD_DIR}/opt/metrilyx
BIN_DIR = ${INSTALL_DIR}/bin

clean:
	rm -rf ./build
	go clean -i ./...

test: clean
	go test -v -cover ./...

build: clean
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	go get -d -v ./...
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}
