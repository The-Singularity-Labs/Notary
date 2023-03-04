
GO_ROOT := $(shell go env GOROOT)

all: run

dev:
	cd ui && yarn run dev --host 127.0.0.1 --port 41119

build: build-wasm build-linux build-windows build-osx

build-wasm:
	mkdir -p ui/src/assets/wasm
	cp $(GO_ROOT)/misc/wasm/wasm_exec.js ui/src/wasm_exec.js
	cd lib/wasm && GOOS=js GOARCH=wasm go build -o ../../ui/src/assets/wasm/golib.wasm ./main.go

app:
	guark run

build-linux:
	guark build  --target linux --rm

build-windows:
	guark build  --target windows --rm

build-osx:
	guark build  --target darwin --rm