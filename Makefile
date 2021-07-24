export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

ifdef GOBIN
$(info go bin $(GOBIN))
endif

all:
	@echo ok

.PHONY: test
test:
	go test -race ./...

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@go fmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "go fmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: golint
golint:
	@cd tools && go install golang.org/x/lint/golint
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html
