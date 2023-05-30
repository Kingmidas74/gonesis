import {CellType} from "../engine/domain.js";
import {Either} from "../monads/index.js";

class Game {

    /**
     * @type {Window} implementation of window
     * @private
     */
    #windowProvider
    /**
     * @type {ConfigurationProvider} The configuration of the game.
     * @private
     */
    #configuration
    /**
     * @type {CanvasWrapper} Object for work with canvas
     * @private
     */
    #canvas
    /**
     * @type {Array<Cell>} Array of cells
     * @private
     */
    #cells

    /**
     * @type {Renderer} Size of cell
     * @private
     */
    #renderer
    /**
     * @type {WorldManager} Manager of the world
     * @private
     */
    #worldManager

    /**
     * @type {CellFactory} Factory for creating cells
     * @private
     */
    #cellFactory

    onGameOverListeners = [];

    addOnGameOverListener(listener) {
        this.onGameOverListeners.push(listener);
    }

    removeOnGameOverListener(listener) {
        this.onGameOverListeners = this.onGameOverListeners.filter(l => l !== listener);
    }

    #raiseOnGameOverEvent(...args) {
        this.onGameOverListeners.forEach(listener => listener(...args));
    }

    /**
     * Constructs a new instance of Game.
     * @param {CanvasWrapper} canvas - The canvas for drawing.
     * @param {WorldManager} worldManager - The manager of the world.
     * @param {ConfigurationProvider} configurationProvider - The configuration of the game.
     * @param {Window} windowProvider - The provider for window operations.
     * @param {Renderer} renderer - The renderer for drawing.
     * @param {CellFactory} cellFactory - The factory for creating cells.
     */
    constructor({canvas, worldManager, configurationProvider, windowProvider, renderer, cellFactory}) {
        this.#windowProvider = windowProvider;

        this.#renderer = renderer;

        this.#canvas = canvas;
        this.#configuration = configurationProvider;

        this.#worldManager = worldManager;
        this.#cellFactory = cellFactory;
        this.#cells = [];
    }

    /**
     * Fill cells array based on world data.
     * @param {World} worldInstance - The world data in object format.
     * @private
     * @returns {Either<null, Error>} null if world is filled successfully, error otherwise
     */
    #fillCells(worldInstance) {
        const width = worldInstance.width;
        const height = worldInstance.height;

        if(worldInstance.cells.length < worldInstance.agents.length) {
            return Either.exception(new Error("World is corrupted"));
        }

        for (let row = 0; row < height; row++) {
            for (let col = 0; col < width; col++) {
                if (worldInstance.cells[row*width+col].cellType === CellType.WALL) {
                    this.#cells[row*width+col] = this.#cellFactory.createWall(col, row);
                }
                if (worldInstance.cells[row*width+col].cellType === CellType.EMPTY) {
                    const energyPercent = (worldInstance.cells[row*width+col].energy / (worldInstance.height - 1));
                    this.#cells[row*width+col] = this.#cellFactory.createEmpty(col, row, energyPercent);
                }
            }
        }

        for (const agent of worldInstance.agents) {
            this.#cells[agent.y*width+agent.x] = this.#cellFactory.createAgent(agent.x, agent.y, agent.energy, agent.agentType);
        }

        return Either.value()
    }

    /**
     * Initialize game's world
     * @returns {Promise<Either<null, Error>>} null if world is initialized successfully, error otherwise
     */
    async init() {
        const eitherWorld = await this.#worldManager.initWorld(this.#canvas);
        return eitherWorld
            .map((world) => {
                this.#cells = Array(world.cells.length);
                return this.#fillCells(world).map(() => {
                    this.#renderer.draw(this.#cells);
                    return  world
                })
            })
    }

    /**
     * Update game's world
     * @returns {Either<boolean, Error>} true if game is not over, false otherwise
     * @private
     */
    async #update() {
        return this.#worldManager.updateWorld().map(world => {
            this.#fillCells(world);
            return this.#livingAgentsCount(world) > 0
        });
    }

    /**
     * @param {World} world - The world data in object format.
     * @return {number} count of living agents
     * @private
     */
    #livingAgentsCount(world) {
        let count = 0;
        for (const a of world.agents) {
            if (a.energy > 0) {
                count++;
            }
        }
        return count;
    }

    /**
     * @param {World} world - The world data in object format.
     * @return {number} count of empty cells
     * @private
     */
    #emptyCellsCount(world) {
        let count = 0;
        for (const a of world.cells) {
            if (a.cellType === CellType.EMPTY) {
                count++;
            }
        }
        return count;
    }

    /**
     * Step game
     * @returns {Either<boolean, Error>}
     */
    async step() {
        return (await this.#update()).map((shouldContinue) => {
            this.#renderer.draw(this.#cells);
            return shouldContinue;
        });
    }

    /**
     * Run game
     * @returns {Promise<void>}
     */
    async run() {
        const desiredFPS = 10;
        const timeStep = 1000 / desiredFPS;
        let lastTime = this.#windowProvider.performance.now();

        const loop = async (currentTime) => {
            const deltaTime = currentTime - lastTime;

            if(deltaTime >= timeStep) {
                if(!this.#configuration.getInstance().Playable) {
                    return;
                }
                lastTime = currentTime - (deltaTime % timeStep);
                const stepResult = await this.step();
                stepResult.map(shouldContinue => {
                    if(!shouldContinue) {
                        this.#renderer.draw(this.#cells);
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
}

export {Game}