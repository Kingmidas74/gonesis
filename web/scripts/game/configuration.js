import {Colors} from "../render/colors.js";
import {MazeGenerators} from "../enums/maze-generator-types.js";


/**
 * Represents game settings.
 */
class Configuration {
    /**
     * Creates an instance of Settings.
     */
    constructor({
                    cellSize = 20,
                    mazeColor = Colors.DARK,
                    mazeGenerator = MazeGenerators.SideWinder,
                    initialAgentsCount = 100,
                    agentsColor = Colors.BLUE,
                    isPlayable = true,
                    foodColor = Colors.GREEN,
                    poisonColor = Colors.RED,
                } = {}) {
        /**
         * The size of each cell in the maze.
         * @type {number}
         */
        this.CellSize = cellSize;

        /**
         * The color of the maze.
         * @type {string}
         * @see Colors
         */
        this.MazeColor = mazeColor;

        /**
         * The maze generator algorithm to use.
         * @type {string}
         * @see MazeGenerators
         */
        this.MazeGenerator = mazeGenerator;

        /**
         * The initial number of agents in the game.
         * @type {number}
         */
        this.InitialAgentsCount = initialAgentsCount;

        /**
         * The color of the agents.
         * @type {string}
         * @see Colors
         */
        this.AgentsColor = agentsColor;

        /**
         * Indicates if the game is playable.
         * @type {boolean}
         */
        this.Playable = isPlayable;

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

    /**
     * Updates the game settings.
     * @param {Configuration} newConfig
     */
    updateConfiguration(newConfig) {
        Object.assign(ConfigurationProvider.#instance, newConfig);
    }
}

export {Configuration, ConfigurationProvider};