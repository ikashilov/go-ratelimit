.PHONY: build
build:
	go build -o bin/example cmd/*

.PHONY: run
run:
	./bin/example 

.PHONY: test
test:
	go build -o bin/example cmd/* && ./bin/example 
