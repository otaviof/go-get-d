APP = go-get-d

GOFLAGS ?= -a -v
GOFLAGS_TEST ?= -race -cover

ARGS ?=

.EXPORT_ALL_VARIABLES:

.PHONY: $(APP)
$(APP):
	go build .

build: $(BIN)

default: build

install:
	go install .

run:
	go run . $(ARGS)

test: test-unit

.PHONY: test-unit
test-unit:
	go test $(GOFLAGS_TEST) . $(ARGS)
