build:
	@cp ./.env ./bin/config/.env
	@cp ./config.yaml ./bin/config/config.yaml
	@cd cmd; go build -o ../bin/glavhim

run: build
	@cd bin; ./glavhim