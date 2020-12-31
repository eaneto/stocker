all:
	rm -rf bin/*
	mkdir -p bin
	go build -o bin/

run:
	./bin/stocker

test:
	go test -v ./...

codecov:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic
