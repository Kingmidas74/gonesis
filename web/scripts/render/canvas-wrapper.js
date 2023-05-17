/** Class representing a canvas wrapper. */
export default class CanvasWrapper {
    /**
     * @type {HTMLCanvasElement}
     * @private
     */
    #canvas

    /**
     *
     * @param {HTMLCanvasElement} canvasElement
     */
    constructor(canvasElement) {
        this.#canvas = canvasElement;
        if(this.#canvas.getContext)
        {
            this.ctx = this.#canvas.getContext("2d");

            this.#canvas.width = document.body.clientWidth;
            this.#canvas.height = document.body.clientHeight;

            this.#canvas.style.transform = 'translate3d(0, 0, 0)';
        }
    }

    /**
     * Gets the width of the canvas element.
     * @returns {number} The width of the canvas.
     */
    get width() {
        return this.#canvas.width;
    }

    /**
     * Gets the height of the canvas element.
     * @returns {number} The height of the canvas.
     */
    get height() {
        return this.#canvas.height;
    }

    /**
     * Gets the canvas context.
     * @returns {CanvasRenderingContext2D} The canvas context.
     */
    get context() {
        return this.ctx;
    }

    /**
     * Clears the entire canvas by erasing the contents.
     */
    clear() {
        this.ctx.clearRect(0, 0, this.#canvas.width, this.#canvas.height);
    }
}