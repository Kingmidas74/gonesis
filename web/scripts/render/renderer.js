/** Render new state on canvas */
export default class Renderer {
    /**
     * @type {CanvasWrapper}
     * @private
     */
    #canvas

    /**
     * @param {CanvasWrapper} canvas   - The canvas wrapper.
     */
    constructor(canvas) {
        this.#canvas = canvas;
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
                cellSet[i].draw(this.#canvas);
            }
        }
    }
}