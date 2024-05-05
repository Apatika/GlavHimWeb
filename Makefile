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
	@rm ${BUILD_PATH}/static/assets/* || true
	@cd ${FRONTEND_PATH}; npm run build
	@cp ${STATIC_SOURCE_PATH}/index.html ${BUILD_PATH}/static/index.html
	@cp ${STATIC_SOURCE_PATH}/assets/* ${BUILD_PATH}/static/assets
	@cd cmd; go build -o .${BUILD_PATH}/${EXECUTABLE}

run: build
	@cd ${BUILD_PATH}; ./${EXECUTABLE}