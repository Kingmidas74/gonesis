import {Colors} from "../render/colors.js";
import {MazeGenerators} from "../enums/maze-generator-types.js";


/**
 * Represents game settings.
 */
class Configuration {
    /**
     * Creates an instance of Settings.
     */
    constructor() {
        /**
         * The size of each cell in the maze.
         * @type {number}
         */
        this.CellSize = 20;

        /**
         * The color of the maze.
         * @type {string}
         * @see Colors
         */
        this.MazeColor = Colors.BLUE;

        /**
         * The maze generator algorithm to use.
         * @type {string}
         * @see MazeGenerators
         */
        this.MazeGenerator = MazeGenerators.SideWinder;

        /**
         * The initial number of agents in the game.
         * @type {number}
         */
        this.InitialAgentsCount = 100;

        /**
         * The color of the agents.
         * @type {string}
         * @see Colors
         */
        this.AgentsColor = Colors.RED;

        /**
         * Indicates if the game is playable.
         * @type {boolean}
         */
        this.Playable = true;

        this.RenderDebounceTime = 100;
    }
}

/**
 * Provides access to the game settings instance.
 */
class ConfigurationProvider {
    /**
     * The singleton instance of the Settings class.
     * @type {Configuration}
     * @private
     */
    static #instance;

    /**
     * Creates an instance of SettingsProvider. If the instance does not exist, it creates a new one.
     */
    constructor() {
        if (!ConfigurationProvider.#instance) {
            ConfigurationProvider.#instance = new Configuration();
        }
    }

    /**
     * Retrieves the singleton instance of the game settings.
     * @returns {Configuration} The singleton instance of the game settings.
     */
    getInstance() {
        return ConfigurationProvider.#instance;
    }
}

export {Configuration, ConfigurationProvider};