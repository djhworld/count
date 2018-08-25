all: prepare test
	go install

prepare:
	go mod tidy
	gofmt -w ./

test:
	go test


