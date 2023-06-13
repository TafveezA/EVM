build:
	go build -o ./bin/evm
run:
	./bin/evm

test:
	go test -v ./...