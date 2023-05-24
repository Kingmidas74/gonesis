import { MathProvider } from "./providers/math.provider.js";

import { Engine } from "./engine/engine.js";

import { CanvasWrapper, Renderer } from "./render/index.js";
import { WorldManager, ConfigurationProvider, CellFactory, Game } from "./game/index.js";

/** Startup class for the application. */
export default class Application {
    /**
     * @type {Game} The game instance
     */
    #game

    /**
     * @type {ConfigurationProvider}
     */
    configurationProvider


    /**
     * Configures the application.
     * @param {Window} window Implementation of window
     * @param {Document} document Implementation of document
     * @param {HTMLCanvasElement} canvasElement The canvas html element for drawing.
     * @param {string} wasmFile The path to the WebAssembly file
     * @returns {Promise<void>} A promise that resolves when the application is configured.
     */
    async configure(window, document, canvasElement, wasmFile) {
        const configurationProvider = new ConfigurationProvider();

        const mathProvider = new MathProvider();

        const canvas = new CanvasWrapper(canvasElement);
        const engine = new Engine(wasmFile, window);
        const worldManager = new WorldManager(engine, configurationProvider, mathProvider);
        const cellFactory = new CellFactory(configurationProvider);
        const renderer = new Renderer(canvas, configurationProvider);

        this.#game = new Game({
            canvas: canvas,
            worldManager: worldManager,
            configurationProvider: configurationProvider,
            windowProvider: window,
            renderer: renderer,
            cellFactory: cellFactory,
        });

        this.configurationProvider = configurationProvider
    }

    /**
     * Runs the game.
     * @param {Configuration} config The configuration of the game.
     * @returns {Promise<Game>} A promise that resolves when the game is running.
     */
    async run(config){
        const configurationProvider = new ConfigurationProvider();
        configurationProvider.updateConfiguration(config);
        await this.#game.init()
        console.log("Ready");
        await this.#game.run()
        return this.#game
    }
}