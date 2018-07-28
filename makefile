all: build-dev
	./buildy

build-dev:
	go build

build-release:
	go build -tags release
