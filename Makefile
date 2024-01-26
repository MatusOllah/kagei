# settings
IS_RELEASE = false

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# tools
GO = go

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE = $(BINARY)/kagei$(shell go env GOEXE)

# flags
GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true

.PHONY: all
all: build

.PHONY: build
build: clean
	mkdir -p $(BINARY)

	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE)

.PHONY: clean
clean:
	rm -rf $(BINARY)
