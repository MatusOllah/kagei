GO=go

BINARY=./bin

FLAGS=-v

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

.PHONY: install
install:
	$(GO) get
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) install $(FLAGS)

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)