import {Either} from "../../monads/either.js";
import {Game} from "../../services/game/game.js";
import {ConfigurationProvider} from "../../configuration/configuration.js";

class ApplicationGameService {


    #performanceProvider;
    #windowProvider;

    /**
     * @type {IGame}
     */
    #gameService;

    #isGenerated = false;

    #configurationProvider;

    #events = {};

    #lastWorldInstance;

    /**
     * Creates an instance of ApplicationGameService.
     * @param { Window } windowProvider
     * @param { Performance } performanceProvider
     */
    constructor(windowProvider, performanceProvider) {
        this.#windowProvider = windowProvider;
        this.#performanceProvider = performanceProvider;

        this.#configurationProvider = new ConfigurationProvider();
    }

    /**
     * Configures the application.
     * @param { IDataClient } dataClient
     * @returns {Promise<void>}
     */
    configure(dataClient){
        this.#gameService = new Game(dataClient);
    }

    /**
     * Gets configuration provider.
     * @returns {ConfigurationProvider}
     */
    get configurationProvider() {
        return this.#configurationProvider;
    }

    /**
     * Sets configuration
     * @param { ConfigurationProvider } configurationProvider
     */
    set configurationProvider(configurationProvider) {
        console.trace('gh')
        this.#configurationProvider = configurationProvider;
    }

    addEventListener(eventName, listener) {
        if (!this.#events[eventName]) {
            this.#events[eventName] = [];
        }
        this.#events[eventName].push(listener);
    }

    emit(eventName, data) {
        if (this.#events[eventName]) {
            this.#events[eventName].forEach(listener => listener({detail:{value: data}}));
        }
    }

    /**
     * Pauses the game.
     * @returns {void}
     */
    pauseGame() {
        this.configurationProvider.getInstance().Playable = false;
    }

    /**
     * Generates a new world.
     * @param {Number} width
     * @param {Number} height
     * @returns {Promise<Either<World, Error>>}
     */
    generateGame = async (width, height) => {
        this.pauseGame();

        const cellSize = this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize;
        this.#configurationProvider.getInstance().WorldConfiguration.Ratio.Width = width / cellSize;
        this.#configurationProvider.getInstance().WorldConfiguration.Ratio.Height = height / cellSize;
        return (await this.#gameService.initWorld(this.#configurationProvider))
            .map(worldInstance => {
                this.#lastWorldInstance = worldInstance;
                this.#isGenerated = true;
                return worldInstance;
            });
    }

    /**
     * Next step of the game.
     * @returns {Either<World, Error>}
     */
    async nextStep() {
        if(!this.#isGenerated) {
            return Either.exception(new Error('Game is not generated yet'));
        }

        return (await this.#gameService.step(this.#configurationProvider))
            .map(worldInstance => {
                this.#lastWorldInstance = worldInstance;
                return worldInstance;
            })
    }

    /**
     * Plays the game.
     * @returns {Either<void, Error>}
     */
    async playGame() {
        if(this.#configurationProvider.getInstance().Playable) {
            return Either.value()
        }

        if(!this.#isGenerated) {
            return Either.exception(new Error('Game is not generated yet'));
        }

        this.#configurationProvider.getInstance().Playable = true;
        const desiredFPS = 10;
        const timeStep = 1000 / desiredFPS;
        let lastTime = this.#performanceProvider?.now();

        const loop = async (currentTime) => {
            if(!this.#configurationProvider.getInstance().Playable) {
                return;
            }

            const deltaTime = currentTime - lastTime;

            if(deltaTime >= timeStep) {
                lastTime = currentTime - (deltaTime % timeStep);
                (await this.nextStep())
                    .map(worldInstance => {
                        if(this.#gameService.livingAgentsCount(worldInstance) === 0 ||
                            (!this.#configurationProvider.getInstance().WorldConfiguration.OneAgentTypeMode &&
                                this.#gameService.isOnlyOneAgentTypeAlive(worldInstance))) {
                            this.pauseGame();
                            this.emit('finish', worldInstance);
                        } else {
                            this.emit('update', worldInstance);
                        }

                    }).orElse(err => {
                        this.pauseGame();
                        this.#raiseError(err);
                    })
            }
            this.#windowProvider.requestAnimationFrame(loop);
        }
        this.#windowProvider.requestAnimationFrame(loop);
        return Either.value();
    }

    cell(x, y) {
        return this.#gameService.cell(x, y,this.#lastWorldInstance);
    }

    #raiseError = (error) => {
        console.error(error);
    }
}

export { ApplicationGameService }