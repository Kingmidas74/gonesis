import {CellType} from "../domain/enum.js";
import {Either} from "../monads/either.js";

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
     * @type {WorldManager} Manager of the world
     * @private
     */
    #worldManager

    /**
     * @type {CellFactory} Factory for creating cells
     * @private
     */
    #cellFactory

    /**
     * @type {World} The world data in object format.
     * @private
     */
    #worldInstance

    /**
     * Constructs a new instance of Game.
     * @param {CanvasWrapper} canvas - The canvas for drawing.
     * @param {WorldManager} worldManager - The manager of the world.
     * @param {ConfigurationProvider} configurationProvider - The configuration of the game.
     * @param {Window} windowProvider - The provider for window operations.
     * @param {CellFactory} cellFactory - The factory for creating cells.
     */
    constructor({canvas, worldManager, configurationProvider, windowProvider, cellFactory}) {
        this.#windowProvider = windowProvider;

        this.#canvas = canvas;
        this.#configuration = configurationProvider;

        this.#worldManager = worldManager;
        this.#cellFactory = cellFactory;
        this.#cells = [];
    }

    /**
     * Initialize game's world
     * @returns {Promise<Either<World, Error>>} void if world is initialized successfully, error otherwise
     */
    async init() {
        return (await this.#worldManager.initWorld(this.#canvas))
            .bind((world) => {
                this.#worldInstance = world;
                this.#cells = Array(world.cells.length);

                return this.#fillCells(world);
            })
            .map(this.#drawCells)
            .map(_ => this.#worldInstance);
    }

    /**
     * Step game
     * @returns {Promise<Either<World, Error>>}
     */
    async step() {
        return this.#worldManager
            .updateWorld()
            .bind(this.#fillCells)
            .map(world => {
                this.#worldInstance = world;
                return this.#drawCells(world);
            })
            .map(_ => this.#worldInstance);
    }

    /**
     * @param {World} world - The world data in object format.
     * @returns {number}
     */
    calculateGeneration(world) {
        let maxGeneration = 0;
        for (const a of this.#cells) {
            if (a?.generation > maxGeneration) {
                maxGeneration = a?.generation;
            }
        }
        return maxGeneration;
    }

    /**
     * @param {World} world - The world data in object format.
     * @returns {Array<Agent>}
     */
    agents = (world) => {
        return this.#worldInstance?.agents ?? []
    }

    /**
     * Fill cells array based on world data.
     * @param {World} worldInstance - The world data in object format.
     * @private
     * @returns {Either<World, Error>} world if world is filled successfully, error otherwise
     */
    #fillCells = (worldInstance) => {
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
                    this.#cells[row*width+col] = this.#cellFactory.createEmpty(col, row);
                }
            }
        }

        for (const agent of worldInstance.agents) {
            this.#cells[agent.y*width+agent.x] = this.#cellFactory.createAgent(agent);
        }

        return Either.value(worldInstance)
    }

    #drawCells = () => {
        for (const cell of this.#cells) {
            cell.draw();
        }
        this.#canvas.render()
    }

    /**
     * @param {World} world - The world data in object format.
     * @return {number} count of living agents
     */
    livingAgentsCount(world) {
        return world.agents.filter(a => a.energy > 0).length;
    }

    isOnlyOneAgentTypeAlive(world) {
        const firstLivingAgentType = world.agents.find(a => a.energy > 0)?.agentType;
        const everyAgentIsSameType = world.agents.every(a => a.agentType === firstLivingAgentType);
        return everyAgentIsSameType && firstLivingAgentType !== undefined;
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
}

export {Game}