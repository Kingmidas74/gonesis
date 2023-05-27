import {Colors} from "../render/colors.js";
import {MazeGenerators} from "../enums/maze-generator-types.js";

class AgentConfiguration {
    constructor({
                    maxEnergy = 100,
                    initialCount = 100,
                    carnivoreColor = Colors.BLUE,
                    herbivoreColor = Colors.GREEN,
                    plantColor = Colors.GREEN,
                    decomposerColor = Colors.BROWN,
                    omnivoreColor=Colors.PURPLE,
                } = {}) {

        /**
         * The maximum energy an agent can have.
         * @type {number}
         */
        this.MaxEnergy = maxEnergy;

        /**
         * The initial number of agents.
         * @type {number}
         */
        this.InitialCount = initialCount;

        /**
         * The color of the carnivore agents.
         * @type {string}
         */
        this.CarnivoreColor = carnivoreColor;

        /**
         * The color of the herbivore agents.
         * @type {string}
         */
        this.HerbivoreColor = herbivoreColor;

        /**
         * The color of the plants.
         * @type {string}
         */
        this.PlantColor = plantColor;

        /**
         * The color of the decomposers.
         * @type {string}
         */
        this.DecomposerColor = decomposerColor;

        /**
         * The color of the omnivore
         * @type {string}
         */
        this.OmnivoreColor = omnivoreColor;
    }
}

/**
 * Represents game settings.
 */
class Configuration {

    /**
     * The width of the world
     * @type {number}
     */
    Width
    /**
     * The height of the world
     * @type {number}
     */
    Height

    /**
     * Creates an instance of Settings.
     */
    constructor({
                    cellSize = 20,
                    mazeColor = Colors.DARK,
                    mazeGenerator = MazeGenerators.SideWinder,
                    isPlayable = true,
                    agentConfiguration = new AgentConfiguration(
                        {
                            maxEnergy: 100,
                            initialCount: 2,
                            carnivoreColor: Colors.RED,
                            herbivoreColor: Colors.BLUE,
                            plantColor: Colors.GREEN,
                            decomposerColor: Colors.BROWN,
                        }
                    ),
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
         * Indicates if the game is playable.
         * @type {boolean}
         */
        this.Playable = isPlayable;

        /**
         * The agent configuration.
         * @type {AgentConfiguration}
         * @see AgentConfiguration
         */
        this.AgentConfiguration = agentConfiguration;
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
        if (newConfig.AgentConfiguration) {
            Object.assign(ConfigurationProvider.#instance.AgentConfiguration, newConfig.AgentConfiguration);
            delete newConfig.AgentConfiguration;
        }

        Object.assign(ConfigurationProvider.#instance, newConfig);
    }
}

export {Configuration, ConfigurationProvider};