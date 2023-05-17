import Cell from "../render/cell.js";

class Game {

    /**
     * @type {Window} implementation of window
     */
    #windowProvider
    /**
     * @type {Configuration}
     */
    #configuration
    /**
     * @type {CanvasWrapper} Object for work with canvas
     */
    #canvas
    /**
     * @type {Array<Cell>} Array of cells
     */
    #cells

    /**
     * @type {Renderer} Size of cell
     */
    #renderer
    /**
     * @type {WorldManager} Manager of the world
     */
    #worldManager

    /**
     * Constructs a new instance of Game.
     * @param {CanvasWrapper} canvas - The canvas for drawing.
     * @param {WorldManager} worldManager - The manager of the world.
     * @param {Configuration} configuration - The configuration of the game.
     * @param {Window} windowProvider - The provider for window operations.
     * @param {Renderer} renderer - The renderer for drawing.
     */
    constructor({canvas, worldManager, configuration, windowProvider, renderer}) {
        this.#windowProvider = windowProvider;

        this.#renderer = renderer;

        this.#canvas = canvas;
        this.#configuration = configuration;

        this.#worldManager = worldManager;
        this.#cells = [];
    }

    /**
     * Fill cells array based on world data.
     * @param {any} worldInstance - The world data in object format.
     * @private
     */
    #fillCells(worldInstance) {
        const width = worldInstance.width;
        const height = worldInstance.height;

        for (let row = 0; row < height; row++) {
            for (let col = 0; col < width; col++) {
                if (worldInstance.cells[row*width+col].cellType === 3) {
                    const cell = new Cell(
                        col * this.#configuration.CellSize,
                        row * this.#configuration.CellSize,
                        this.#configuration.CellSize,
                        this.#configuration.MazeColor,
                    );
                    this.#cells.push(cell);
                }
            }
        }

        for (const agent of worldInstance.agents) {
            const cell = new Cell(
                agent.x * this.#configuration.CellSize,
                agent.y * this.#configuration.CellSize,
                this.#configuration.CellSize,
                this.#configuration.AgentsColor,
            );
            this.#cells.push(cell);
        }
    }

    /**
     * Initialize game's world
     * @returns {Promise<void>}
     * @private
     */
    async #init() {
        const world = await this.#worldManager.initWorld(this.#canvas);
        this.#fillCells(world)
    }

    /**
     * Update game's world
     * @returns {boolean} true if game is not over, false otherwise
     * @private
     */
    #update() {
        const world = this.#worldManager.updateWorld();
        if(this.#livingAgentsCount(world) === 0) {
            return false;
        }
        this.#cells = [];
        this.#fillCells(world);
        return true;
    }

    /**
     * @param {any} world - The world data in object format.
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
     * Run game
     * @returns {Promise<void>}
     */
    async run() {
        await this.#init();
        this.#renderer.draw(this.#cells);

        if(!this.#configuration.Playable) {
            console.log("Game is not playable");
            return;
        }

        const loop = () => {
            const updateResult = this.#update();
            this.#renderer.draw(this.#cells);
            if (updateResult) {
                this.#windowProvider.requestAnimationFrame(loop);
            }
            else {
                console.log("Game over");
            }
        }

        loop();
    }
}

export {Game}