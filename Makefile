default: build

all: fidel
dev: fidel debug build install clean distcleann
debug: fidel
build: fidel
install: fidel
clean: fidel


.PHONY: all dev debug build install clean distclean

###Build ###

BUILD_THREADS=4
BUILD_IPFS_PATH=./ipfs
BUILD_IPFS_BIN=./ipfs
BUILD_IPFS_BIN_PATH=./ipfs
BUILD_IPFS_BIN_NAME=ipfs
BUILD_IPFS_BIN_VERSION=v0.4.24
BUILD_IPFS_BIN_URL=https://dist.ipfs.io/ipfs/${BUILD_IPFS_BIN_VERSION}/go-ipfs-${BUILD_IPFS_BIN_VERSION}.linux-amd64
FIDEL_PATH=./fidel
FIDEL_BIN=./fidel
FIDEL_BIN_PATH=./fidel
FIDEL_BIN_NAME=fidel
FIDEL_BIN_VERSION=v0.4.24
FIDEL_EINSTEINDB_PATH=./fidel/einsteindb
FIDEL_EINSTEINDB_BIN=./fidel/einsteindb
FIDEL_EINSTEINDB_BIN_PATH=./fidel/einsteindb
FIDEL_EINSTEINDB_BIN_NAME=einsteindb


###Build ###


###Install ###

install:
	mkdir -p ${BUILD_IPFS_PATH}
	mkdir -p ${BUILD_IPFS_BIN_PATH}
	mkdir -p ${FIDEL_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_BIN_PATH}

	wget -O ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_URL}
	wget -O ${FIDEL_BIN} ${FIDEL_BIN_URL}

	cp ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_PATH}/${BUILD_IPFS_BIN_NAME}
	cp ${FIDEL_BIN} ${FIDEL_BIN_PATH}/${FIDEL_BIN_NAME}

	cp ${BUILD_IPFS_BIN} ${FIDEL_BIN} ${FIDEL_EINSTEINDB_BIN} ${FIDEL_EINSTEINDB_BIN_PATH}
	cp ${BUILD_IPFS_BIN} ${FIDEL_BIN} ${FIDEL_EINSTEINDB_BIN} ${FIDEL_EINSTEINDB_BIN_PATH}


###Install ###

###Clean ###

clean:
	rm -rf ${BUILD_IPFS_PATH}
	rm -rf ${BUILD_IPFS_BIN_PATH}
	rm -rf ${FIDEL_PATH}
	rm -rf ${FIDEL_EINSTEINDB_PATH}
	rm -rf ${FIDEL_EINSTEINDB_BIN_PATH}

###Clean ###

###DistClean ###

distclean:
	rm -rf ${BUILD_IPFS_PATH}
	rm -rf ${BUILD_IPFS_BIN_PATH}
	rm -rf ${FIDEL_PATH}
	rm -rf ${FIDEL_EINSTEINDB_PATH}
	rm -rf ${FIDEL_EINSTEINDB_BIN_PATH}
	rm -rf ${BUILD_IPFS_BIN}
	rm -rf ${FIDEL_BIN}
	rm -rf ${FIDEL_EINSTEINDB_BIN}

###DistClean ###

###Build ###

		build:
	mkdir -p ${BUILD_IPFS_PATH}
	mkdir -p ${BUILD_IPFS_BIN_PATH}
	mkdir -p ${FIDEL_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_BIN_PATH}
	wget -O ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_URL}
	wget -O ${FIDEL_BIN} ${FIDEL_BIN_URL}
	cp ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_PATH}/${BUILD_IPFS_BIN_NAME}
	cp ${FIDEL_BIN} ${FIDEL_BIN_PATH}/${FIDEL_BIN_NAME}
	cp ${BUILD_IPFS_BIN} ${FIDEL_BIN} ${FIDEL_EINSTEINDB_BIN} ${FIDEL_EINSTEINDB_BIN_PATH}


###Build ###


###Debug ###

debug:
	mkdir -p ${BUILD_IPFS_PATH}
	mkdir -p ${BUILD_IPFS_BIN_PATH}
	mkdir -p ${FIDEL_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_BIN_PATH}
	wget -O ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_URL}
	wget -O ${FIDEL_BIN} ${FIDEL_BIN_URL}
	cp ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_PATH}/${BUILD_IPFS_BIN_NAME}
	cp ${FIDEL_BIN} ${FIDEL_BIN_PATH}/${FIDEL_BIN_NAME}
	cp ${BUILD_IPFS_BIN} ${FIDEL_BIN} ${FIDEL_EINSTEINDB_BIN} ${FIDEL_EINSTEINDB_BIN_PATH}
	${BUILD_IPFS_BIN} --debug
	${FIDEL_BIN} --debug
	${FIDEL_EINSTEINDB_BIN} --debug

###Debug ###

###Build ###

		build:
	mkdir -p ${BUILD_IPFS_PATH}
	mkdir -p ${BUILD_IPFS_BIN_PATH}
	mkdir -p ${FIDEL_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_PATH}
	mkdir -p ${FIDEL_EINSTEINDB_BIN_PATH}
	wget -O ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_URL}
	wget -O ${FIDEL_BIN} ${FIDEL_BIN_URL}
	cp ${BUILD_IPFS_BIN} ${BUILD_IPFS_BIN_PATH}/${BUILD_IPFS_BIN_NAME}