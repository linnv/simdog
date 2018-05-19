CURDIR := $(shell pwd)

GO        := go
# GOBUILD   := GOPATH=$(GOPATH) CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)
GOBUILD   := GOPATH=$(GOPATH) $(GO) build $(BUILD_FLAG)
GOTEST    := GOPATH=$(GOPATH) CGO_ENABLED=1 $(GO) test -p 3

LINUX:=env GOOS="linux" GOHOSTARCH="amd64" 

LDFLAGS += -X "github.com/linnv/manhelp.Version=$(shell git describe --tags --dirty)"
LDFLAGS += -X "github.com/linnv/manhelp.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "github.com/linnv/manhelp.Branch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "github.com/linnv/manhelp.GitHash=$(shell git rev-parse HEAD)"

all: build

BUILDDIR=$(CURDIR)
OUTPUTDIR=$(CURDIR)/"bin"
TARGETBIN=$(OUTPUTDIR)/"simdog"
build: 
	@mkdir -p $(OUTPUTDIR)
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o $(TARGETBIN) $(BUILDDIR)/main.go

linux: 
	@mkdir -p $(OUTPUTDIR)
	$(LINUX) $(GOBUILD) -ldflags '$(LDFLAGS)' -o $(TARGETBIN)_linux $(BUILDDIR)/main.go
clean: 
	@rm $(TARGETBIN)*

