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
     * @param {any} worldInstance - The world data in object format.
     * @private
     */
    #fillCells(worldInstance) {
        const width = worldInstance.width;
        const height = worldInstance.height;

        for (let row = 0; row < height; row++) {
            for (let col = 0; col < width; col++) {
                if (worldInstance.cells[row*width+col].cellType === 3) {
                    this.#cells[row*width+col] = this.#cellFactory.createWall(col, row);
                }
                if (worldInstance.cells[row*width+col].cellType === 0) {
                    this.#cells[row*width+col] = this.#cellFactory.createEmpty(col, row);
                }
            }
        }

        for (const agent of worldInstance.agents) {
            this.#cells[agent.y*width+agent.x] = this.#cellFactory.createAgent(agent.x, agent.y, agent.energy);
        }
    }

    /**
     * Initialize game's world
     * @returns {Promise<void>}
     * @private
     */
    async #init() {
        const world = await this.#worldManager.initWorld(this.#canvas);
        this.#cells = Array(world.width * world.height);
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

        if(!this.#configuration.getInstance().Playable) {
            console.log("Game is not playable");
            return;
        }
        const desiredFPS = 40;
        const timeStep = 1000 / desiredFPS;
        let lastTime = this.#windowProvider.performance.now();

        const loop = (currentTime) => {
            const deltaTime = currentTime - lastTime;

            if(deltaTime >= timeStep) {
                lastTime = currentTime - (deltaTime % timeStep);
                const updateResult = this.#update();
                this.#renderer.draw(this.#cells);
                if (!updateResult) {
                    console.log("Game over");
                    return;
                }
            }

            this.#windowProvider.requestAnimationFrame(loop);
        }

        this.#windowProvider.requestAnimationFrame(loop);
    }
}

export {Game}