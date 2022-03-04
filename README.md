## Console Wordle
A simple implementation of Wordle for use on the console, primarily targeting Wasm32-wasi.

* Compile with [TinyGo](https://tinygo.org/): `tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go`
* Test with [Wasmtime](https://wasmtime.dev/): `wasmtime main.wasm`
* Run with [Enarx](https://enarx.dev/): `enarx run main.wasm`
