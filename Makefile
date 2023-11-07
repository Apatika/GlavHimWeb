include .env

APP_NAME=glavhim

build:
	@cp ./.env ${BUILD_PATH}/config/.env
	@cp ${CONFIG_PATH} ${BUILD_PATH}/${CONFIG_PATH}
	@cd cmd; go build -o ../bin/${APP_NAME}

run: build
	@cd ${BUILD_PATH}; ./${APP_NAME}