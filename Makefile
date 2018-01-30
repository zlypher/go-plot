PACKAGES := \
	github.com/zlypher/go-plot

all: build silent-test

build:
	go build -v .

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)
