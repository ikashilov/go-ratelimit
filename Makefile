.PHONY: build
build:
	go build -o bin/cmd cmd/*

.PHONY: run
run:
	./bin/cmd 