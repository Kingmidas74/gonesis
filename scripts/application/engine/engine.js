import {Either} from "../monads/either.js";

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
     * @private
     */
    #jsonProvider

    /**
     * The flag that indicates whether the engine is initialized
     * @type {boolean}
     * @private
     */
    #isInitialized = false

    /**
     * Creates an instance of Engine.
     * @param {string} wasmFile Name of wasm file
     * @param {Window} windowProvider Object implements window interface
     * @param {JSON} jsonProvider Object implements JSON interface
     */
    constructor(wasmFile, windowProvider, jsonProvider) {
        this.#wasmFile = new URL(wasmFile, new URL(import.meta.url)).href;
        this.#go = new Go();
        this.#windowProvider = windowProvider;
        this.#jsonProvider = jsonProvider
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
        this.#go.run(result.instance).then(_ => this.#isInitialized = true);
    }

    /**
     * Initialize world
     * @param {Configuration} configuration Configuration of the world
     * @returns {Either<World, Error>} World instance
     */
    initWorld(configuration) {
        const response = this.#windowProvider.initWorld(configuration.WorldConfiguration.Ratio.Width, configuration.WorldConfiguration.Ratio.Height, this.#jsonProvider.stringify(configuration), "1687351067538000000")
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

    /**
     * Parse response
     * @param response
     * @returns {Either<World, Error>}
     * @private
     */
    #parseResponse(response) {
        try {

            const parsedResponse = this.#jsonProvider.parse(response)
            if (parsedResponse.code !== 0) {
                return Either.exception(new Error(parsedResponse.message))
            }

            return Either.value(this.#jsonProvider.parse(parsedResponse.message))
        } catch (e) {
            console.log(response)
            return Either.exception(e)
        }

    }
}

export { Engine }