/** Render new state on canvas */
export default class Renderer {
    /**
     * @type {CanvasWrapper}
     * @private
     */
    #canvas

    /**
     * @type {ConfigurationProvider} The configuration of the game.
     */
    #configurationProvider

    /**
     * @param {CanvasWrapper} canvas   - The canvas wrapper.
     * @param {ConfigurationProvider} configurationProvider - The configuration of the game.
     */
    constructor(canvas, configurationProvider) {
        this.#canvas = canvas;
        this.#configurationProvider = configurationProvider;
    }

    /** Clear canvas */
    clear() {
        this.#canvas.clear();
    }

    /**
     * Draw cell sets on canvas
     * @param {Array<Cell>} cellSets - The cell sets.
     */
    draw(...cellSets) {
        this.clear()
        for (let cellSet of cellSets) {
            for (let i = 0; i < cellSet.length; i++) {
                cellSet[i].draw(this.#canvas, this.#configurationProvider.getInstance().WorldConfiguration.CellSize);
            }
        }
    }
}