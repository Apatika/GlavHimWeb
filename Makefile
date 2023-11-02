build:
	cp ./.env ./bin/.env
	cp ./config.yaml ./bin/config.yaml
	cd cmd; go build -o ../bin/glavhim

run: build
	./bin/glavhim