import {AgentType, CellType} from "../../../domain/enum.js";
import {AgentDrawStrategy, Cell, EmptyDrawStrategy} from "./cell.js";

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
     * @param cell - The cell data.
     * @returns {import('../render/cell.js').Cell}
     */
    createEmpty(cell) {
        let drawingStrategy = this.#drawingStrategies.get(CellType.EMPTY).withCell(cell);
        return new Cell(cell.x, cell.y, drawingStrategy);
    }

    /**
     * Creates an agent.
     * @param cell - The cell data.
     * @returns {import('../render/cell.js').Cell}
     */
    createAgent(cell) {
        let drawingStrategy = this.#drawingStrategies.get(cell.agent.agentType).withCell(cell);
        return new Cell(cell.x, cell.y, drawingStrategy);
    }
}

export { CellFactory };