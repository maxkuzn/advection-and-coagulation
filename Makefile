.PHONY: build
build:
	cd ./build; pwd; cmake ..; make;

.PHONY: run
run:
	./build/bin/main

