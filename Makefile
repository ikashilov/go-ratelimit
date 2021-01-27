.PHONY: build
build:
	go build -o bin/cmd cmd/*

.PHONY: run
run:
	./bin/cmd 

.PHONY: test
test:
	go build -o bin/cmd cmd/* && ./bin/cmd 
