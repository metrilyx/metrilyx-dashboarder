SHELL=/bin/bash
NGOOS = $(go env GOOS)

NAME = metrilyx-dashboarder
VERSION = $(shell cat VERSION)
BUILD_DIR_BASE = ./build
BUILD_DIR = ${BUILD_DIR_BASE}/${NAME}

INSTALL_DIR = ${BUILD_DIR}/opt/metrilyx
BIN_DIR = ${INSTALL_DIR}/bin
WEBROOT = webroot

.clean:
	rm -rf ${BUILD_DIR_BASE}
	go clean -i ./...

.deps:
	go get -d -v ./...

.test: .clean .deps
	go test -v -cover ./...

.build_webroot:
	[ -d ${INSTALL_DIR} ] || mkdir -p ${INSTALL_DIR}
	git submodule init
	cp -a metrilyx-web ${INSTALL_DIR}/${WEBROOT}
	rm -f ${INSTALL_DIR}/${WEBROOT}/.git

.build_osx:
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}

.build_linux:
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}

.build_native: 
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}

.rpm:
	[ -d ${BUILD_DIR_BASE}/el ] || mkdir -p ${BUILD_DIR_BASE}/el
	cd ${BUILD_DIR_BASE} &&  fpm -s dir -t rpm -n ${NAME} --version ${VERSION} ${NAME}
	mv *.rpm ${BUILD_DIR_BASE}/el/

.deb:
	[ -d ${BUILD_DIR_BASE}/ubuntu ] || mkdir -p ${BUILD_DIR_BASE}/ubuntu
	cd ${BUILD_DIR_BASE} &&  fpm -s dir -t deb -n ${NAME} --version ${VERSION} ${NAME}
	mv *.deb ${BUILD_DIR_BASE}/ubuntu/


all: .clean .deps .build_native .build_webroot
