import {JsonProvider} from "./json.provider.js";

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
     * @type {JsonProvider} Json provider
     */
    #jsonProvider

    /**
     * Creates an instance of Engine.
     * @param {string} wasmFile Name of wasm file
     * @param {Window} windowProvider Object implements window interface
     */
    constructor(wasmFile, windowProvider) {
        this.#wasmFile = new URL(wasmFile, new URL(import.meta.url)).href;
        this.#go = new Go();
        this.#windowProvider = windowProvider;
        this.#jsonProvider = new JsonProvider();
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
     * @param {number} width Width of the world
     * @param {number} height Height of the world
     * @param {Configuration} configuration Configuration of the world
     * @returns {World} World instance
     */
    initWorld(width, height, configuration) {
        return this.#jsonProvider.parse(this.#windowProvider.initWorld(width, height, this.#jsonProvider.stringify(configuration)))
    }

    /**
     * Step of the game
     * @returns {World} World instance
     */
    step() {
        return this.#jsonProvider.parse(this.#windowProvider.step())
    }
}

export { Engine }