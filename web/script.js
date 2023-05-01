const go = new Go();
WebAssembly.instantiateStreaming(fetch("engine.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});