import {MathProvider} from "../providers/math.provider.js";
import {Engine} from "../engine/engine.js";

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
     * @type {ConfigurationProvider}
     * @private
     */
    #config

    /**
     * Constructs a new instance of WorldManager.
     * @param {Engine} engine - The engine for work with WebAssembly.
     * @param {ConfigurationProvider} configurationProvider - The configuration of the game.
     * @param {MathProvider} mathProvider - The provider for math operations.
     */
    constructor(engine, configurationProvider, mathProvider) {
        this.#engine = engine;
        this.#config = configurationProvider;
        this.#mathProvider = mathProvider;
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
     * @returns {Promise<World>} The world data in object format.
     */
    async initWorld(canvas) {
        await this.#initEngine();
        canvas.init();
        const width = this.#mathProvider.floor(canvas.width / this.#config.getInstance().CellSize);
        const height = this.#mathProvider.floor(canvas.height / this.#config.getInstance().CellSize);
        return  this.#engine.initWorld(width, height, this.#config.getInstance());
    }

    /**
     * Update world
     * @returns {World} The world data in object format.
     */
    updateWorld() {
        return this.#engine.step();
    }
}