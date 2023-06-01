import {Either} from "../application/monads/index.js";

export default class GameController {

    /**
     * @type {Game} The game.
     * @private
     */
    #game

    /**
     * @type {ConfigurationProvider} The configuration provider.
     * @private
     */
    #configurationProvider

    /**
     * @type {Window} The provider for window operations.
     * @private
     */
    #windowProvider

    #isGenerated = false;


    #onGameOverEventHandlers = [];

    addOnGameOverEventListener = (listener) => {
        this.#onGameOverEventHandlers.push(listener);
    }

    removeOnGameOverEventListener = (listener) => {
        this.#onGameOverEventHandlers = this.#onGameOverEventHandlers.filter(l => l !== listener);
    }

    #raiseOnGameOverEvent = (...args) => {
        this.#onGameOverEventHandlers.forEach(listener => listener(...args));
    }


    /**
     * @param {Game} game - The game.
     * @param {ConfigurationProvider} configurationProvider - The configuration provider.
     * @param {Window} windowProvider - The provider for window operations.
     */
    constructor(game, configurationProvider, windowProvider) {
        this.#game = game;
        this.#configurationProvider = configurationProvider;
        this.#windowProvider = windowProvider;
    }

    /**
     * Next step of the game.
     * @returns {Promise<void>}
     */
    async nextStep() {
        try {
            if(!this.#isGenerated) {
                (await this.generateGame())
            }
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
        if(!this.#isGenerated) {
            this.generateGame().catch((error) => {
                this.#raiseOnGameOverEvent(Either.exception(error))
            })
        }
        this.#configurationProvider.getInstance().Playable = true;
        const desiredFPS = 10;
        const timeStep = 1000 / desiredFPS;
        let lastTime = this.#windowProvider.performance.now();

        const loop = async (currentTime) => {
            const deltaTime = currentTime - lastTime;

            if(deltaTime >= timeStep) {
                if(!this.#configurationProvider.getInstance().Playable) {
                    return;
                }
                lastTime = currentTime - (deltaTime % timeStep);
                const stepResult = await this.#game.step();
                stepResult.map(shouldContinue => {
                    if(!shouldContinue) {
                        this.#configurationProvider.getInstance().Playable = false;
                        this.#raiseOnGameOverEvent(stepResult)
                    }

                }).orElse(err => {
                    this.#raiseOnGameOverEvent(stepResult)
                });
            }

            this.#windowProvider.requestAnimationFrame(loop);
        }

        this.#windowProvider.requestAnimationFrame(loop);
    }

    /**
     * Pauses the game.
     * @returns {void}
     */
    pauseGame() {
        this.#configurationProvider.getInstance().Playable = false;
    }

    /**
     * Generates a new world.
     * @returns {Promise<Either<null, Error>>}
     */
    async generateGame() {
        this.pauseGame();
        return (await this.#game.init())
            .map(() => {
              this.#isGenerated = true;
            })
    }

    /**
     * Updates the settings of the game.
     * @param {Configuration} settings
     */
    updateSettings(settings) {
        this.#configurationProvider.updateConfiguration(settings);
    }
}
