import {AgentType} from "../domain/enum.js";

/**
 * Strategy for drawing a cell
 * @interface
 */
class CellDrawStrategy {
    /**
     * Draw the cell on the canvas.
     * @param {number} x
     * @param {number} y
     * @abstract
     * @returns {void}
     */
    draw(x, y) {
        throw new Error("Not implemented");
    }
}

/**
 * Strategy for drawing a wall cell
 */
class WallDrawStrategy extends CellDrawStrategy {

    /**
     * A canvas wrapper.
     * @type {CanvasWrapper}
     * @private
     */
    #canvasWrapper;

    /**
     * The configuration provider.
     * @type {ConfigurationProvider}
     * @private
     */
    #configurationProvider;

    /**
     * WallDrawStrategy constructor.
     * @param {CanvasWrapper} canvasWrapper
     * @param {ConfigurationProvider} configurationProvider
     */
    constructor(canvasWrapper, configurationProvider) {
        super();

        this.#canvasWrapper = canvasWrapper;
        this.#configurationProvider = configurationProvider;
    }


    draw(x, y) {
        this.#canvasWrapper.drawRect(
            x * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            y * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.MazeColor);
    }
}

/**
 * Strategy for drawing an empty cell
 */
class EmptyDrawStrategy extends CellDrawStrategy {

    /**
     * A canvas wrapper.
     * @type {CanvasWrapper}
     * @private
     */
    #canvasWrapper;

    /**
     * The configuration provider.
     * @type {ConfigurationProvider}
     * @private
     */
    #configurationProvider;

    /**
     * WallDrawStrategy constructor.
     * @param {CanvasWrapper} canvasWrapper
     * @param {ConfigurationProvider} configurationProvider
     */
    constructor(canvasWrapper, configurationProvider) {
        super();

        this.#canvasWrapper = canvasWrapper;
        this.#configurationProvider = configurationProvider;
    }


    draw(x, y) {
        this.#canvasWrapper.drawRect(
            x * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            y * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            "hsla(0, 100%, 100%, 1.0)"
            );
    }
}

/**
 * Strategy for drawing an agent cell
 */
class AgentDrawStrategy extends CellDrawStrategy {

    /**
     * A canvas wrapper.
     * @type {CanvasWrapper}
     * @private
     */
    #canvasWrapper;

    /**
     * The configuration provider.
     * @type {ConfigurationProvider}
     * @private
     */
    #configurationProvider;

    /**
     * The agent.
     * @type {Agent}
     * @private
     * @nullable
     */
    #agent;

    /**
     * WallDrawStrategy constructor.
     * @param {CanvasWrapper} canvasWrapper
     * @param {ConfigurationProvider} configurationProvider
     * @param {Agent | null} agent
     */
    constructor(canvasWrapper, configurationProvider, agent = null) {
        super();

        this.#canvasWrapper = canvasWrapper;
        this.#configurationProvider = configurationProvider;

        this.#agent = agent;
    }


    draw(x, y) {
        this.#canvasWrapper.drawCircle(
            x * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            y * this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.#configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize / 2,
            this.#getAgentColor());
    }

    withAgent(agent) {
        return new AgentDrawStrategy(this.#canvasWrapper, this.#configurationProvider, agent);
    }

    #getAgentColor() {
        let agentColor = null;
        switch (this.#agent.agentType) {
            case AgentType.CARNIVORE:
                agentColor = this.#configurationProvider.getInstance().CarnivoreConfiguration.Color;
                break;
            case AgentType.HERBIVORE:
                agentColor = this.#configurationProvider.getInstance().HerbivoreConfiguration.Color;
                break;
            case AgentType.DECOMPOSER:
                agentColor = this.#configurationProvider.getInstance().DecomposerConfiguration.Color;
                break;
            case AgentType.PLANT:
                agentColor = this.#configurationProvider.getInstance().PlantConfiguration.Color;
                break;
            case AgentType.OMNIVORE:
                agentColor = this.#configurationProvider.getInstance().OmnivoreConfiguration.Color;
                break;
            default:
                throw "Unknown agent type: " + this.#agent.agentType;
        }
        return agentColor;
    }
}

/**
 * Represents a cell of the world
 */
class Cell {

    /**
     * The x coordinate of the cell.
     * @type {number}
     * @private
     */
    #x;

    /**
     * The y coordinate of the cell.
     * @type {number}
     * @private
     */
    #y;

    /**
     * The draw strategy of the cell.
     * @type {CellDrawStrategy}
     * @private
     */
    #drawStrategy;


    /**
     * Cell constructor.
     * @param {number} x
     * @param {number} y
     * @param {CellDrawStrategy} drawStrategy
     */
    constructor(x, y, drawStrategy) {
        this.#x = x;
        this.#y = y;
        this.#drawStrategy = drawStrategy;
    }

    /**
     * Draw the cell on the canvas.
     * @returns {void}
     */
    draw() {
        this.#drawStrategy.draw(this.#x, this.#y);
    }
}

export {Cell, CellDrawStrategy, WallDrawStrategy, EmptyDrawStrategy, AgentDrawStrategy};