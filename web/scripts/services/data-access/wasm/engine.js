import {Either} from "../../../monads/either.js";
import {IDataClient} from "../../../contracts/data-client.interface.js";

/** WebAssembly Go instance */
class Engine extends IDataClient {

    /**
     * @type {Go} Go instance
     */
    #go

    /**
     * @type {string} Path to wasm file
     */
    #wasmFile

    /**
     * @type {JSON} Json provider
     * @private
     */
    #JSONProvider

    /**
     * @type {Window} Window provider
     */
    #windowProvider

    /**
     * The flag that indicates whether the engine is initialized
     * @type {boolean}
     * @private
     */
    #isInitialized = false

    /**
     * Creates an instance of Engine.
     * @param { string } wasmFile Name of wasm file
     * @param { Window } windowProvider Object implements window interface
     * @param { JSON } JSONProvider Object implements JSON interface
     */
    constructor(wasmFile, windowProvider, JSONProvider) {
        super();

        this.#wasmFile = new URL(wasmFile, new URL(import.meta.url)).href;
        this.#go = new windowProvider.Go();
        this.#windowProvider = windowProvider;
        this.#JSONProvider = JSONProvider
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
        this.#go.run(result.instance).then();
        this.#isInitialized=true;
    }

    /**
     * Initialize world
     * @param {ConfigurationProvider} configurationProvider Configuration (provider) of the world
     * @returns {Either<World, Error>} World instance
     */
    initWorld(configurationProvider) {
        const response = this.#windowProvider.initWorld(this.#JSONProvider.stringify(configurationProvider.getInstance()))
        return this.#parseResponse(response);
    }

    /**
     * Step of the game
     * @returns {Either<World, Error>} World instance
     */
    step() {
        const response = this.#windowProvider.updateWorld()
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

            const parsedResponse = this.#JSONProvider.parse(response)
            if (parsedResponse.code !== 0) {
                return Either.exception(new Error(parsedResponse.message))
            }

            return Either.value(this.#JSONProvider.parse(parsedResponse.message))
        } catch (e) {
            return Either.exception(e)
        }

    }
}

export { Engine }