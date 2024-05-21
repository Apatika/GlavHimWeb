-include .env

CONFIG_PATH?=./config/config.yaml
BUILD_PATH?=./bin
APP_NAME?=glavhim
STATIC_SOURCE_PATH?=./frontend/svelte/dist
FRONTEND_PATH?=./frontend/svelte

ifeq ($(OS), Windows_NT)
	EXECUTABLE=${APP_NAME}.exe
else
	EXECUTABLE=${APP_NAME}
endif

all: build

build:
	@mkdir ${BUILD_PATH} || true
	@mkdir ${BUILD_PATH}/static || true
	@mkdir ${BUILD_PATH}/config || true
	@rm ${BUILD_PATH}/static/assets/* || true
	@cd ${FRONTEND_PATH}; npm run build
	@cp -a ${STATIC_SOURCE_PATH}/. ${BUILD_PATH}/static/
	@cd cmd; go build -o .${BUILD_PATH}/${EXECUTABLE}

run: build
	@cd ${BUILD_PATH}; ./${EXECUTABLE}