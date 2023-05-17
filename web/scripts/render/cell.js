/**
 * Represents a cell of the world
 */
export default class Cell {

    /**
     *
     * @param {number} x
     * @param {number} y
     * @param {number} size
     * @param {string} color
     */
    constructor(x, y, size, color) {
        this.x = x;
        this.y = y;
        this.size = size;
        this.color = color;
    }

    /**
     * Draws the cell on the canvas.
     * @param {CanvasWrapper} canvasWrapper - The canvas wrapper.
     */
    draw(canvasWrapper) {
        canvasWrapper.context.fillStyle = this.color;
        canvasWrapper.context.fillRect(this.x, this.y, this.size, this.size);
    }
}