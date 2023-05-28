import {Colors} from "../render/colors.js";
import {MazeGenerators} from "../enums/maze-generator-types.js";

class AgentConfiguration {
    constructor({
                    MaxEnergy = 100,
                    InitialCount = 0,
                    Color = Colors.GREEN,
                } = {}) {
        /**
         * The maximum energy an agent can have.
         * @type {number}
         */
        this.MaxEnergy = MaxEnergy;

        /**
         * The initial number of agents.
         * @type {number}
         */
        this.InitialCount = InitialCount;

        /**
         * The color of the agents.
         * @type {string}
         */
        this.Color = Color;
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
                    plantConfiguration = new AgentConfiguration({
                        InitialCount: 10,
                        Color: Colors.GREEN,
                    }),
                    herbivoreConfiguration = new AgentConfiguration({
                        InitialCount: 10,
                        Color: Colors.BLUE,
                    }),
                    carnivoreConfiguration = new AgentConfiguration({
                        InitialCount: 10,
                        Color: Colors.RED,
                    }),
                    decomposerConfiguration = new AgentConfiguration({
                        InitialCount: 0,
                        Color: Colors.BROWN,
                    }),
                    omnivoreConfiguration = new AgentConfiguration({
                        InitialCount: 10,
                        Color: Colors.PURPLE,
                    }),
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
         * The configuration for the agents of type 'plant'
         * @see {@link AgentType}.
         * @type {AgentConfiguration}
         */
        this.PlantConfiguration = plantConfiguration;

        /**
         * The configuration for the agents of type 'herbivore'
         * @type {AgentConfiguration}
         * @see {@link AgentType}.
         */
        this.HerbivoreConfiguration = herbivoreConfiguration;

        /**
         * The configuration for the agents of type 'carnivore'
         * @type {AgentConfiguration}
         * @see {@link AgentType}.
         */
        this.CarnivoreConfiguration = carnivoreConfiguration;

        /**
         * The configuration for the agents of type 'decomposer'
         * @type {AgentConfiguration}
         * @see {@link AgentType}.
         */
        this.DecomposerConfiguration = decomposerConfiguration;

        /**
         * The configuration for the agents of type 'omnivore'
         * @type {AgentConfiguration}
         * @see {@link AgentType}.
         */
        this.OmnivoreConfiguration = omnivoreConfiguration;
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

        if(newConfig.PlantConfiguration) {
            Object.assign(ConfigurationProvider.#instance.PlantConfiguration, newConfig.PlantConfiguration);
            delete newConfig.PlantConfiguration;
        }

        if(newConfig.HerbivoreConfiguration) {
            Object.assign(ConfigurationProvider.#instance.HerbivoreConfiguration, newConfig.HerbivoreConfiguration);
            delete newConfig.HerbivoreConfiguration;
        }

        if(newConfig.CarnivoreConfiguration) {
            Object.assign(ConfigurationProvider.#instance.CarnivoreConfiguration, newConfig.CarnivoreConfiguration);
            delete newConfig.CarnivoreConfiguration;
        }

        if(newConfig.DecomposerConfiguration) {
            Object.assign(ConfigurationProvider.#instance.DecomposerConfiguration, newConfig.DecomposerConfiguration);
            delete newConfig.DecomposerConfiguration;
        }

        if(newConfig.OmnivoreConfiguration) {
            Object.assign(ConfigurationProvider.#instance.OmnivoreConfiguration, newConfig.OmnivoreConfiguration);
            delete newConfig.OmnivoreConfiguration;
        }

        Object.assign(ConfigurationProvider.#instance, newConfig);
    }
}

export {Configuration, ConfigurationProvider};