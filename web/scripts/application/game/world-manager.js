import {MathProvider} from "../providers/math.provider.js";
import {Engine} from "../engine/engine.js";
import {Either} from "../monads/either.js";

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
        return this.#engine.init();
    }

    /**
     * Initialize world
     * @param {CanvasWrapper} canvas - The canvas for drawing.
     * @returns {Either<World, Error>} The world data in object format.
     */
    async initWorld(canvas) {
        try {
            await this.#initEngine();
            canvas.init(this.#config.getInstance().WorldConfiguration.Ratio.Width * this.#config.getInstance().WorldConfiguration.Ratio.CellSize, this.#config.getInstance().WorldConfiguration.Ratio.Height * this.#config.getInstance().WorldConfiguration.Ratio.CellSize);
            return this.#engine.initWorld(this.#config.getInstance());
        } catch (e) {
            return Either.exception(e)
        }
    }

    /**
     * Update world
     * @returns {Either<World, Error>} The world data in object format.
     */
    updateWorld() {
        return this.#engine.step();
    }
}