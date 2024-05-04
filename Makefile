include .env

ifeq ($(OS), Windows_NT)
	EXECUTABLE=${APP_NAME}.exe
else
	EXECUTABLE=${APP_NAME}
endif

all: build

build:
	@rm ${BUILD_PATH}/static/assets/*
	@cp ./.env ${BUILD_PATH}/config/.env
	@cp ${CONFIG_PATH} ${BUILD_PATH}/${CONFIG_PATH}
	@cd ${FRONTEND_PATH}; npm run build
	@cp ${STATIC_SOURCE_PATH}/index.html ${BUILD_PATH}/static/index.html
	@cp ${STATIC_SOURCE_PATH}/assets/* ${BUILD_PATH}/static/assets
	@cd cmd; go build -o .${BUILD_PATH}/${EXECUTABLE}

run: build
	@cd ${BUILD_PATH}; ./${EXECUTABLE}