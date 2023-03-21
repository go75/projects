GOOS=js GOARCH=wasm go build -o wasm/main.wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" wasm