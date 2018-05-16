CURDIR := $(shell pwd)

GO        := go
GOBUILD   := GOPATH=$(GOPATH) CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)
GOTEST    := GOPATH=$(GOPATH) CGO_ENABLED=1 $(GO) test -p 3


LDFLAGS += -X "github.com/linnv/manhelp.Version=$(shell git describe --tags --dirty)"
LDFLAGS += -X "github.com/linnv/manhelp.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "github.com/linnv/manhelp.Branch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "github.com/linnv/manhelp.GitHash=$(shell git rev-parse HEAD)"

all: build

BUILDDIR=$(CURDIR)/bin
TARGETBIN="simdog"
build: 
	@mkdir -p $(BUILDDIR)
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o $(BUILDDIR)/$(TARGETBIN) $(BUILDDIR)/main.go

clean: 
	@rm $(BUILDDIR)/$(TARGETBIN)

