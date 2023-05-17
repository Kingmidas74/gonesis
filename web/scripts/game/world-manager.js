import {MathProvider} from "../providers/math.provider.js";
import {Engine} from "../engine/engine.js";
import {Configuration} from "./configuration.js";


/**
 * Manages the world in the game.
 */
export default class WorldManager {
    /**
     * @type {MathProvider}
     * @private
     */
    #mathProvider
    /**
     * @type {JSON}
     * @private
     */
    #jsonProvider

    /**
     * @type {Engine}
     * @private
     */
    #engine
    /**
     * @type {Configuration}
     * @private
     */
    #config

    /**
     * Constructs a new instance of WorldManager.
     * @param {Engine} engine - The engine for work with WebAssembly.
     * @param {Configuration} configuration - The configuration of the game.
     * @param {MathProvider} mathProvider - The provider for math operations.
     * @param {JSON} jsonProvider - The provider for JSON operations.
     */
    constructor(engine, configuration, mathProvider, jsonProvider) {
        this.#engine = engine;
        this.#config = configuration;
        this.#mathProvider = mathProvider;
        this.#jsonProvider = jsonProvider;
    }

    /**
     * Initialize engine
     * @returns {Promise<void>}
     */
    async #initEngine() {
        await this.#engine.init();
    }

    /**
     * Initialize world
     * @param {CanvasWrapper} canvas - The canvas for drawing.
     * @returns {Promise<any>} The world data in object format.
     */
    async initWorld(canvas) {
        await this.#initEngine();
        const width = this.#mathProvider.floor(canvas.width / this.#config.CellSize);
        const height = this.#mathProvider.floor(canvas.height / this.#config.CellSize);
        const worldData = this.#engine.initWorld(width, height, this.#config.InitialAgentsCount);
        return this.#parseWorldData(worldData);
    }

    /**
     * Parse world data from JSON to object.
     * @param {any} worldData - The world data in JSON format.
     * @returns {any} The world data in object format.
     */
    #parseWorldData(worldData) {
        return this.#jsonProvider.parse(worldData);
    }

    /**
     * Update world
     * @returns {any} The world data in object format.
     */
    updateWorld() {
        return this.#parseWorldData(this.#engine.step());
    }
}