all: clean test
	gb build all

clean:
	@@rm -rf pkg/
	@@rm -rf bin/

test:
	gb test

