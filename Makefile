build:
	cd cmd; go build -o ../bin/glavhim

run: build
	go run ./cmd/main.go