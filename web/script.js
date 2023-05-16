const Colors = Object.freeze({
    RED:   Symbol("#ff0000"),
    BLUE:  Symbol("#999999"),
    GREEN: Symbol("#00ff00"),
});

const MazeGenerators = Object.freeze({
    AldousBroder:   Symbol("AldousBroder"),
    Binary:  Symbol("Binary"),
    Grid: Symbol("Grid"),
    SideWinder: Symbol("SideWinder"),
    Border: Symbol("Border")
});

class Settings {
    constructor() {
        this.CellSize = 20;
        this.MazeColor = Colors.BLUE;
        this.MazeGenerator = MazeGenerators.SideWinder;

        this.InitialAgensCount = 1;
        this.AgentsColor = Colors.RED;

        this.Playable = true;
    }

    static getInstance() {
        if (!this.instance) {
            this.instance = new Settings();
        }

        return this.instance;
    }
}

class SettingsProvider {
    getInstance() {
        return Settings.getInstance()
    }
}

class MathProvider {
    floor(x) {
        return Math.floor(x)
    }
}


class Canvas {
    constructor(elementId) {
        this.canvas = document.getElementById(elementId);
        if(this.canvas.getContext)
        {
            this.ctx = this.canvas.getContext("2d");

            this.canvas.width = document.body.clientWidth;
            this.canvas.height = document.body.clientHeight;

            this.canvas.style.transform = 'translate3d(0, 0, 0)';
        }
    }

    clear() {
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }
}

class Cell {
    constructor(x, y, size, color) {
        this.x = x;
        this.y = y;
        this.size = size;
        this.color = color;
    }

    draw(ctx) {
        ctx.fillStyle = this.color;
        ctx.fillRect(this.x, this.y, this.size, this.size);
    }
}

class Engine {

    #go
    #wasmFile

    constructor(wasmFile) {
        this.#wasmFile = wasmFile;
        this.#go = new Go();
    }

    async init() {
        const result = await WebAssembly.instantiateStreaming(fetch(this.#wasmFile), this.#go.importObject)
        this.#go.run(result.instance);
    }

    /**
     *
     * @param {number} width
     * @param {number} height
     * @param {number} agentsCount The value should be less than width * height
     * @returns {*}
     */
    initWorld(width, height, agentsCount) {
        return initWorld(width, height, agentsCount)
    }

    step() {
        return step()
    }
}

class Game {

    #daysQueue
    #pastDays

    constructor(canvas, engine, config, math) {
        this.canvas = canvas;
        this.engine = engine;
        this.config = config;
        this.math = math;

        this.cellSize = this.config.getInstance().CellSize;
        this.cells = [];
    }

    fillCells(worldInstance) {
        const width = worldInstance.width;
        const height = worldInstance.height;

        for (let row = 0; row < height; row++) {
            for (let col = 0; col < width; col++) {
                if (worldInstance.cells[row*width+col].cellType === 3) {
                    const cell = new Cell(
                        col * this.cellSize,
                        row * this.cellSize,
                        this.cellSize,
                        this.config.getInstance().MazeColor.description,
                    );
                    this.cells.push(cell);
                }
            }
        }

        for (const agent of worldInstance.agents) {
            const cell = new Cell(
                agent.x * this.cellSize,
                agent.y * this.cellSize,
                this.cellSize,
                this.config.getInstance().AgentsColor.description,
            );
            this.cells.push(cell);
        }
    }

    async init() {
        await this.engine.init()

        let mazeWidth = this.math.floor(this.canvas.canvas.width / this.cellSize);
        let mazeHeight = this.math.floor(this.canvas.canvas.height / this.cellSize);

        const world = JSON.parse(this.engine.initWorld(mazeWidth, mazeHeight, this.config.getInstance().InitialAgensCount))
        this.fillCells(world)

    }

    draw() {
        this.canvas.clear();
        for (let i = 0; i < this.cells.length; i++) {
            this.cells[i].draw(this.canvas.ctx);
        }
    }

    update() {
        const world = JSON.parse(this.engine.step())
        console.log("livingAgentsCount", this.livingAgentsCount(world))
        if(this.livingAgentsCount(world) === 0) {
            console.log(world);
            return false;
        }
        this.cells = [];
        this.fillCells(world)
        return true;
    }

    /**
     * @param {*} world
     * @return {number} count of living agents
     */
    livingAgentsCount(world) {
        let count = 0;
        for (const a of world.agents) {
            if (a.energy > 0) {
                count++;
            }
        }
        return count;
    }

    async run() {
        const config = this.config.getInstance();

        await this.init();
        this.draw();

        if(!config.Playable) {
            console.log("Game is not playable");
            return;
        }

        let lastFrameTime = performance.now();

        const loop = () => {
            const updateResult = this.update();
            this.draw();
            console.log("updateResult", updateResult);
            if (updateResult) {
                requestAnimationFrame(loop);
            }
            else {
                console.log("Game over");
            }


        }

        loop();
    }
}

(async (canvasID, wasmFile) => {
    const settingsProvider = new SettingsProvider();

    const mathProvider = new MathProvider();

    const canvas = new Canvas(canvasID);
    const engine = new Engine(wasmFile);

    const game = new Game(canvas, engine, settingsProvider, mathProvider);
    await game.run()
})("canvas", "engine.wasm").then(() => console.log("Ready"))