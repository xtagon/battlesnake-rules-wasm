.PHONY: clean

main: main.go
	GOOS=js GOARCH=wasm go build -o dist/main.wasm main.go && cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" dist/wasm_exec.js

clean:
	rm -fv dist/main.wasm dist/wasm_exec.js
