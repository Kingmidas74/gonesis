import {Either} from "../monads/index.js";

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
     * @type {JSON} Json provider
     */
    #jsonProvider

    #isInitialized = false

    /**
     * Creates an instance of Engine.
     * @param {string} wasmFile Name of wasm file
     * @param {Window} windowProvider Object implements window interface
     */
    constructor(wasmFile, windowProvider) {
        this.#wasmFile = new URL(wasmFile, new URL(import.meta.url)).href;
        this.#go = new Go();
        this.#windowProvider = windowProvider;
        this.#jsonProvider = windowProvider.JSON
    }

    /**
     * Initialize wasm module
     * @returns {Promise<void>}
     */
    async init() {
        if (this.#isInitialized) {
            return new Promise(resolve => resolve())
        }
        const result = await this.#windowProvider.WebAssembly.instantiateStreaming(fetch(this.#wasmFile), this.#go.importObject)
        this.#go.run(result.instance).then(_ => this.#isInitialized = true).catch(err => console.error(err));
    }

    /**
     * Initialize world
     * @param {number} width Width of the world
     * @param {number} height Height of the world
     * @param {Configuration} configuration Configuration of the world
     * @returns {Either<World, Error>} World instance
     */
    initWorld(width, height, configuration) {
        const response = this.#windowProvider.initWorld(width, height, this.#jsonProvider.stringify(configuration))
        return this.#parseResponse(response)
    }

    /**
     * Step of the game
     * @returns {Either<World, Error>} World instance
     */
    step() {
        const response = this.#windowProvider.step()
        return this.#parseResponse(response)
    }

    #parseResponse(response) {
        const parsedResponse = this.#jsonProvider.parse(response)
        if (parsedResponse.code !== 0) {
            return Either.exception(new Error(parsedResponse.message))
        }

        return Either.value(this.#jsonProvider.parse(parsedResponse.message))
    }
}

export { Engine }