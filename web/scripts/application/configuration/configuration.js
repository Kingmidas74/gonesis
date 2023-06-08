const MazeGenerators = Object.freeze({
    AldousBroder:   "aldous_broder",
    Binary:  "binary",
    Grid: "grid",
    SideWinder: "side_winder",
    Border: "border",
    Empty: "empty",
})

const Topologies = Object.freeze({
    Moore:   "moore",
    Neumann: "neumann",
})

const ReproductionTypes = Object.freeze({
    Budding:   "budding",
})

const Colors = Object.freeze({
    RED:    "hsla(0, 60%, 60%, 1.0)",
    BLUE:   "hsla(240, 60%, 60%, 1.0)",
    GREEN:  "hsla(120, 60%, 60%, 1.0)",
    BROWN:  "hsla(30, 60%, 60%, 1.0)",
    PURPLE: "hsla(279, 60%, 60%, 1.0)",
    DARK:   "hsla(0, 0%, 0%, 1.0)",
    YELLOW: "hsla(60, 0%, 100%, 1.0)"
});

class WorldConfiguration {
    constructor({
                    MazeType = MazeGenerators.Empty,
                    Topology = Topologies.Moore,
                    CellSize = 10,
                    MazeColor = Colors.DARK,
                } = {}) {
        /**
         * The type of maze to generate.
         * @type {string}
         * @enum {MazeGenerators}
         * @see {@link MazeGenerators}
         */
        this.MazeType = MazeType;

        /**
         * The topology of the world.
         * @type {string}
         * @enum {Topologies}
         * @see {@link Topologies}
         */
        this.Topology = Topology;


        /**
         * The size of each cell in the maze.
         * @type {number}
         */
        this.CellSize = CellSize;

        /**
         * The color of the maze.
         * @type {string}
         */
        this.MazeColor = MazeColor;
    }
}

class AgentConfiguration {
    constructor({
                    MaxEnergy = 100,
                    InitialCount = 0,
                    Color = Colors.GREEN,
                    ReproductionType = ReproductionTypes.Budding,
                    InitialEnergy = 20,
                    ReproductionEnergyCost = 10,
                    ReproductionChance = 0.5,
                    MutationChance = 0.1,
                    BrainVolume = 64,
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

        /**
         * The type of reproduction.
         * @type {string}
         * @enum {ReproductionType}
         * @see {@link ReproductionTypes}
         */
        this.ReproductionType = ReproductionType;

        /**
         * The energy cost of reproduction.
         * @type {number}
         */
        this.ReproductionEnergyCost = ReproductionEnergyCost

        /**
         * The probability of reproduction.
         * @type {number}
         */
        this.ReproductionChance = ReproductionChance

        /**
         * The probability of mutation.
         * @type {number}
         */
        this.MutationChance = MutationChance

        /**
         * The initial energy of the agents.
         */
        this.InitialEnergy = InitialEnergy

        /**
         * The volume of the brain.
         */
        this.BrainVolume = BrainVolume
    }
}

/**
 * Represents game settings.
 */
class Configuration {

    /**
     * Creates an instance of Settings.
     */
    constructor({
                    isPlayable = true,
                    drawRequired = true,
                    worldConfiguration = new WorldConfiguration(),
                    plantConfiguration = new AgentConfiguration({
                        InitialCount: 10,
                        Color: Colors.GREEN,
                    }),
                    herbivoreConfiguration = new AgentConfiguration({
                        InitialCount: 20,
                        Color: Colors.BLUE,
                    }),
                    carnivoreConfiguration = new AgentConfiguration({
                        InitialCount: 40,
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
         * Indicates if the game is playable.
         * @type {boolean}
         */
        this.Playable = isPlayable;

        /**
         * Indicates if the game should be drawn.
         * @type {boolean}
         */
        this.DrawRequired = drawRequired;

        /**
         * The configuration for the world.
         * @type {WorldConfiguration}
         * @see {@link WorldConfiguration}
         */
        this.WorldConfiguration = worldConfiguration;

        /**
         * The configuration for the agents of type 'plant'
         * @type {AgentConfiguration}
         * @see {@link AgentConfiguration}.
         * @see {@link AgentType}
         */
        this.PlantConfiguration = plantConfiguration;

        /**
         * The configuration for the agents of type 'herbivore'
         * @type {AgentConfiguration}
         * @see {@link AgentConfiguration}.
         * @see {@link AgentType}
         */
        this.HerbivoreConfiguration = herbivoreConfiguration;

        /**
         * The configuration for the agents of type 'carnivore'
         * @type {AgentConfiguration}
         * @see {@link AgentConfiguration}.
         * @see {@link AgentType}
         */
        this.CarnivoreConfiguration = carnivoreConfiguration;

        /**
         * The configuration for the agents of type 'decomposer'
         * @type {AgentConfiguration}
         * @see {@link AgentConfiguration}.
         * @see {@link AgentType}
         */
        this.DecomposerConfiguration = decomposerConfiguration;

        /**
         * The configuration for the agents of type 'omnivore'
         * @type {AgentConfiguration}
         * @see {@link AgentConfiguration}.
         * @see {@link AgentType}
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
        if(newConfig.WorldConfiguration) {
            Object.assign(ConfigurationProvider.#instance.WorldConfiguration, newConfig.WorldConfiguration);
            delete newConfig.WorldConfiguration;
        }

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

export {Configuration, ConfigurationProvider, AgentConfiguration, MazeGenerators, Topologies, Colors, ReproductionTypes};