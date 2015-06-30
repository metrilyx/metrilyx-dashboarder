SHELL=/bin/bash


NAME = metrilyx-dashboarder
BUILD_DIR = ./build/${NAME}

INSTALL_DIR = ${BUILD_DIR}/opt/metrilyx
BIN_DIR = ${INSTALL_DIR}/bin
WEBROOT = webroot

clean:
	rm -rf ./build
	go clean -i ./...

test: clean
	go test -v -cover ./...

.build_web:
	[ -d ${INSTALL_DIR} ] || mkdir -p ${INSTALL_DIR}
	git submodule init
	cp -a metrilyx-web ${INSTALL_DIR}/${WEBROOT}
	rm -f ${INSTALL_DIR}/${WEBROOT}/.git

build: clean .build_web
	[ -d ./build ] || mkdir -p ${BUILD_DIR}
	go get -d -v ./...
	
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${NAME} ${NAME}.go 

	cp -a ./etc ${INSTALL_DIR}
