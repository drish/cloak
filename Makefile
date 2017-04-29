all: test

test: 
	go test -v ./crypt/... 
