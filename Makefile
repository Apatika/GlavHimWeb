include .env

ifeq ($(OS), Windows_NT)
	EXECUTABLE=${APP_NAME}.exe
else
	EXECUTABLE=${APP_NAME}
endif

all: build

build:
	@cp ./.env ${BUILD_PATH}/config/.env
	@cp ${CONFIG_PATH} ${BUILD_PATH}/${CONFIG_PATH}
	@cd cmd; go build -o .${BUILD_PATH}/${EXECUTABLE}

run: build
	@cd ${BUILD_PATH}; ./${APP_NAME}