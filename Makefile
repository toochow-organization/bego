
# Default Go binary.
ifndef GOROOT
  GOROOT = /usr/local/go
endif

# Determine the OS to build.
ifeq ($(OS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS = darwin
  else
    GOOS = linux
  endif
else
  GOOS = $(OS)
endif

# Allow optional arguments for `proto-gen`
ifeq (proto-gen, $(firstword $(MAKECMDGOALS)))
  runargs := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
  $(eval $(runargs):;@true)
endif

GOCMD = GOOS=$(GOOS) go
GOTEST = $(GOCMD) test -race
GO_PKGS?=$$(go list ./... | grep -v /vendor/)

# See golangci.yml in extra folder for linters setup
linter:			## Run linter
		@golangci-lint run -c golangci.yml ./...;
		@cd sdlc/codegen && golangci-lint run --path-prefix="sdlc/codegen" -c ../../golangci.yml ./...;

## Generate Go code (pb, grpc and grpc-gateway) of our api
#  Supports optional name of directory for which to run the command. E.g.
# make proto-gen b2c
proto-gen:
	@if [ -z $(filter-out $@,$(MAKECMDGOALS)) ]; then \
        cd protocol && buf generate; \
    else \
        cd protocol && buf generate --path $(runargs); \
    fi

test:			## Run unit test (without integration tests)
		$(GOTEST) -v $(GO_PKGS)

bench:			## Run go benchmarks
		$(GOCMD) test -tags integration -bench=. ./... -benchmem