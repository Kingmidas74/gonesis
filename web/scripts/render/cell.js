/**
 * Represents a cell of the world
 */
class Cell {

    /**
     *
     * @param {number} x
     * @param {number} y
     * @param {string} color
     */
    constructor(x, y, color) {
        this.x = x;
        this.y = y;
        this.color = color;
    }

    /**
     * Draws the cell on the canvas.
     * @param {CanvasWrapper} canvasWrapper - The canvas wrapper.
     * @param {number} size - The size of the cell.
     */
    draw(canvasWrapper, size) {
        canvasWrapper.drawRect(this.x*size, this.y*size, size, size, this.color);
    }
}

// Wall cell
class Wall extends Cell {
    constructor(x, y, color) {
        super(x, y, color);
    }
    // Additional methods or properties for Wall cell if required
}

class Empty extends Cell {
    constructor(x, y, color) {
        super(x, y, color);
    }
    // Additional methods or properties for Wall cell if required
}

// Agent cell
class Agent extends Cell {
    constructor(x, y, color, energy) {
        super(x, y, color);
        this.energy = energy;
    }

    draw(canvasWrapper, size) {
        super.draw(canvasWrapper, size);
        // Additional drawing for energy if needed
    }
    // Additional methods or properties for Agent cell
}

export {Wall, Agent, Empty};