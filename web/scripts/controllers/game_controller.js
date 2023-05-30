export default class GameController {

    /**
     * @type {Game} The game.
     * @private
     */
    #game

    /**
     * @type {ConfigurationProvider} The configuration provider.
     */
    #configurationProvider


    /**
     * @param {Game} game - The game.
     * @param {ConfigurationProvider} configurationProvider - The configuration provider.
     */
    constructor(game, configurationProvider) {
        this.#game = game;
        this.#game.addOnGameOverListener(this.#onGameOverHandler);
        this.#configurationProvider = configurationProvider;
    }

    #onGameOverHandler = (e) => {
        this.#configurationProvider.getInstance().Playable = false
    }

    /**
     * Next step of the game.
     * @returns {Promise<void>}
     */
    async nextStep() {
        try {
            await this.#game.step();
        } catch (error) {
            console.error(error);
        }
    }

    /**
     * Plays the game.
     * @returns {Promise<void>}
     */
    async playGame() {
        this.#configurationProvider.getInstance().Playable = true;
        try {
            console.log("Play", this.#game)
            await this.#game.run();
        } catch (error) {
            console.error(error);
        }
    }

    /**
     * Pauses the game.
     * @returns {Promise<void>}
     */
    pauseGame() {
        this.#configurationProvider.getInstance().Playable = false;
    }

    /**
     * Generates a new world.
     * @returns {Promise<void>}
     */
    async generateGame() {
        console.log("Generate");
        try {
            await this.pauseGame();
            (await this.#game.init())
                .map((_)=> {
                    console.log("Generated");
                })
                .orElse((err) => {
                    throw err
                });
        } catch (error) {
            console.error(error);
        }
    }

    /**
     * Updates the settings of the game.
     * @param {Configuration} settings
     */
    updateSettings(settings) {
        this.#configurationProvider.updateConfiguration(settings);
    }
}
