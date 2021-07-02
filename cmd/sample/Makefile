CURDIR := $(shell pwd)

GO        := go
# GOBUILD   := GOPATH=$(GOPATH) CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)
GOBUILD   := GOPATH=$(GOPATH) $(GO) build $(BUILD_FLAG)
GOTEST    := GOPATH=$(GOPATH) CGO_ENABLED=1 $(GO) test -p 3

LINUX:=env GOOS="linux" GOHOSTARCH="amd64" 

LDFLAGS += -X "github.com/linnv/simdog.Version=$(shell git describe --tags --dirty)"
LDFLAGS += -X "github.com/linnv/simdog.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "github.com/linnv/simdog.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "github.com/linnv/simdog.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "github.com/linnv/simdog.Owner=jialin"
LDFLAGS += -X "github.com/linnv/simdog.BuildHost=$(shell hostname)"
LDFLAGS += -X "github.com/linnv/simdog.BuildUser=$(shell whoami)"

all: build

BUILDDIR=$(CURDIR)
OUTPUTDIR=$(CURDIR)
TARGETBIN=$(OUTPUTDIR)/simdog
build: 
	@mkdir -p $(OUTPUTDIR)
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o $(TARGETBIN) $(BUILDDIR)/main.go

linux: 
	@mkdir -p $(OUTPUTDIR)
	$(LINUX) $(GOBUILD) -ldflags '$(LDFLAGS)' -o $(TARGETBIN)_linux $(BUILDDIR)/main.go
clean: 
	@rm -rf $(TARGETBIN)*

