ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
endif

test:
	go test -race ./internal/... ./cmd/...

build:
	go build -o bin/coagulation1d ./cmd/coagulation1d

run: build
	./bin/coagulation1d

