all: test

test: 
	go test -v ./crypt/... 

build:
	go build -v .