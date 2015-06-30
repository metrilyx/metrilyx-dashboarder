SHELL=/bin/bash

NAME = metrilyx-dashboarder
VERSION = $(shell cat VERSION)
BUILD_DIR_BASE = ./build
BUILD_DIR = ${BUILD_DIR_BASE}/${NAME}

INSTALL_DIR = ${BUILD_DIR}/opt/metrilyx
BIN_DIR = ${INSTALL_DIR}/bin
WEBROOT = webroot

.clean:
	rm -rf ${BUILD_DIR}
	go clean -i ./...

.deps:
	go get -d -v ./...

test: .clean .deps
	go test -v -cover ./...

.build_web:
	[ -d ${INSTALL_DIR} ] || mkdir -p ${INSTALL_DIR}
	git submodule init
	cp -a metrilyx-web ${INSTALL_DIR}/${WEBROOT}
	rm -f ${INSTALL_DIR}/${WEBROOT}/.git

.build_osx: .clean .build_web .deps
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}

	cd ${BUILD_DIR_BASE} && tar -czf ${NAME}.darwin.x86_64.tgz ${NAME}

.build_linux: .clean .build_web .deps
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	go get -d -v ./...
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}

	cd ${BUILD_DIR_BASE} && tar -czf ${NAME}.linux.x86_64.tgz ${NAME}

