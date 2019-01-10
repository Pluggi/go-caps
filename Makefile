GO ?= go

all: mycaps test

test:
	$(GO) test

mycaps:
	$(GO) build -o mycaps
