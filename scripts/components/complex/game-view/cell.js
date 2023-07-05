import {AgentType} from "../../../domain/enum.js";

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
        throw new Error("Abstract method!");
    }
}

/**
 * Strategy for drawing an empty cell
 */
class EmptyDrawStrategy extends CellDrawStrategy {

    /**
     * A canvas wrapper.
     * @type {CanvasWrapper}
     * @protected
     */
    canvasWrapper;

    /**
     * The configuration provider.
     * @type {ConfigurationProvider}
     * @protected
     */
    configurationProvider;

    /**
     * The cell.
     * @private
     */
    #cell

    /**
     * EmptyDrawStrategy constructor.
     * @param {CanvasWrapper} canvasWrapper
     * @param {ConfigurationProvider} configurationProvider
     * @param {Cell | null} cell
     */
    constructor(canvasWrapper, configurationProvider, cell = null) {
        super();

        this.canvasWrapper = canvasWrapper;
        this.configurationProvider = configurationProvider;

        this.#cell = cell;
    }


    draw(x, y) {
        this.canvasWrapper.drawRect(
            (x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize),
            (y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize),
            this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
            this.configurationProvider.getInstance().WorldConfiguration.EmptyColor,
            );



        if(this.#cell?.nortWall) {
            this.canvasWrapper.drawLine(
                x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (x + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                this.configurationProvider.getInstance().WorldConfiguration.MazeColor
            );
        }

        if(this.#cell?.eastWall) {
            this.canvasWrapper.drawLine(
                (x + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (x + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (y + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                this.configurationProvider.getInstance().WorldConfiguration.MazeColor
            );
        }

        if(this.#cell?.southWall) {
            this.canvasWrapper.drawLine(
                x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (y + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (x + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (y + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                this.configurationProvider.getInstance().WorldConfiguration.MazeColor
            );
        }

        if(this.#cell?.westWall) {
            this.canvasWrapper.drawLine(
                x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                (y + 1) * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize,
                this.configurationProvider.getInstance().WorldConfiguration.MazeColor
            );
        }
    }

    withCell(cell) {
        return new EmptyDrawStrategy(this.canvasWrapper, this.configurationProvider, cell);
    }
}

/**
 * Strategy for drawing an agent cell
 */
class AgentDrawStrategy extends EmptyDrawStrategy {
    /**
     * The agent.
     * @private
     * @nullable
     */
    #agent;

    /**
     * AgentDrawStrategy constructor.
     * @param {CanvasWrapper} canvasWrapper
     * @param {ConfigurationProvider} configurationProvider
     * @param {Cell | null} cell
     */
    constructor(canvasWrapper, configurationProvider, cell = null) {
        super(canvasWrapper, configurationProvider, cell);

        this.#agent = cell?.agent;
    }


    draw(x, y) {
        super.draw(x, y);
        this.canvasWrapper.drawCircle(
            (x * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize)+1,
            (y * this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize)+1,
            (this.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize / 2) - 1,
            this.#getAgentColor());
    }

    withCell(cell) {
        return new AgentDrawStrategy(this.canvasWrapper, this.configurationProvider, cell);
    }

    #getAgentColor() {
        let agentColor = null;
        switch (this.#agent.agentType) {
            case AgentType.CARNIVORE:
                agentColor = this.configurationProvider.getInstance().CarnivoreConfiguration.Color;
                break;
            case AgentType.HERBIVORE:
                agentColor = this.configurationProvider.getInstance().HerbivoreConfiguration.Color;
                break;
            case AgentType.DECOMPOSER:
                agentColor = this.configurationProvider.getInstance().DecomposerConfiguration.Color;
                break;
            case AgentType.PLANT:
                agentColor = this.configurationProvider.getInstance().PlantConfiguration.Color;
                break;
            case AgentType.OMNIVORE:
                agentColor = this.configurationProvider.getInstance().OmnivoreConfiguration.Color;
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

export {Cell, CellDrawStrategy, EmptyDrawStrategy, AgentDrawStrategy};