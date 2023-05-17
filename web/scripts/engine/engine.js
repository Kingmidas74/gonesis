/** WebAssembly Go instance */
class Engine {
    /**
     * @type {Go} Go instance
     */
    #go
    /**
     * @type {string} Path to wasm file
     */
    #wasmFile
    /**
     * @type {Window}
     * @private
     */
    #windowProvider

    /**
     * Creates an instance of Engine.
     * @param {string} wasmFile Name of wasm file
     * @param {Window} windowProvider Object implements window interface
     */
    constructor(wasmFile, windowProvider) {
        this.#wasmFile = new URL(wasmFile, new URL(import.meta.url)).href;
        this.#go = new Go();
        this.#windowProvider = windowProvider;
    }

    /**
     * Initialize wasm module
     * @returns {Promise<void>}
     */
    async init() {
        const result = await this.#windowProvider.WebAssembly.instantiateStreaming(fetch(this.#wasmFile), this.#go.importObject)
        this.#go.run(result.instance).catch(err => console.error(err));
    }

    /**
     * Initialize world
     * @param {number} width
     * @param {number} height
     * @param {number} agentsCount The value should be less than width * height
     * @returns {*}
     */
    initWorld(width, height, agentsCount) {
        return this.#windowProvider.initWorld(width, height, agentsCount)
    }

    step() {
        return this.#windowProvider.step()
    }
}

export { Engine }