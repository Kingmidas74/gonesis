import {WallDrawStrategy, Cell, AgentDrawStrategy, EmptyDrawStrategy} from "../render/index.js";
import {AgentType, CellType} from "../domain/enum.js";

class CellFactory {

    /**
     * @type {ConfigurationProvider} The configuration of the game.
     */
    #configuration;

    #drawingStrategies = new Map();

    /**
     * Constructs a new instance of CellFactory.
     * @param {ConfigurationProvider} configurationProvider
     * @param {CanvasWrapper} canvasWrapper
     */
    constructor(configurationProvider, canvasWrapper) {
        this.#configuration = configurationProvider;
        this.#drawingStrategies.set(CellType.WALL, new WallDrawStrategy(canvasWrapper, this.#configuration));
        this.#drawingStrategies.set(CellType.EMPTY, new EmptyDrawStrategy(canvasWrapper, this.#configuration));

        this.#drawingStrategies.set(AgentType.PLANT, new AgentDrawStrategy(canvasWrapper, this.#configuration));
        this.#drawingStrategies.set(AgentType.HERBIVORE, new AgentDrawStrategy(canvasWrapper, this.#configuration));
        this.#drawingStrategies.set(AgentType.CARNIVORE, new AgentDrawStrategy(canvasWrapper, this.#configuration));
        this.#drawingStrategies.set(AgentType.OMNIVORE, new AgentDrawStrategy(canvasWrapper, this.#configuration));
        this.#drawingStrategies.set(AgentType.DECOMPOSER, new AgentDrawStrategy(canvasWrapper, this.#configuration));
    }

    /**
     * Creates a wall cell.
     * @param {number} x - The x coordinate of the cell.
     * @param {number} y - The y coordinate of the cell.
     * @returns {import('../render/cell.js').Cell}
     */
    createWall(x, y) {
        return new Cell(x, y, this.#drawingStrategies.get(CellType.WALL));
    }

    /**
     * Creates an empty cell.
     * @param {number} x - The x coordinate of the cell.
     * @param {number} y - The y coordinate of the cell.
     * @returns {import('../render/cell.js').Cell}
     */
    createEmpty(x, y) {
        return new Cell(x, y, this.#drawingStrategies.get(CellType.EMPTY));
    }

    /**
     * Creates an agent.
     * @param {Agent} agent
     * @returns {import('../render/cell.js').Cell}
     */
    createAgent(agent) {
        let drawingStrategies = this.#drawingStrategies.get(agent.agentType).withAgent(agent);
        return new Cell(agent.x, agent.y, drawingStrategies);
    }
}

export { CellFactory };