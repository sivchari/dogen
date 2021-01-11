GOCMD=go
GOBUILD=$(GOCMD) build
GOOS=linux
GOARCH=amd64
BINARY_NAME=dogen

build:
	cd cmd \
		&& GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME) main.go

bench:
	cd dogen \
		&& $(GOCMD) test ./... -bench . -benchmem -trace gen.trace

trace:
	$(GOCMD) tool trace --http localhost:6060 gen.trace