.PHONY: all build clean

all: build

build:
	go build -o reconsuite-ai main.go

clean:
	rm -f reconsuite-ai