all: build-dev
	./buildy

build-dev:
	go build

build-release:
	go build -tags release

push:
	git add .
	#m=""
	git commit -m "🚀 $(m)"
	git push
