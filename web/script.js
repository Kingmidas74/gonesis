const Colors = Object.freeze({
    RED:   Symbol("red"),
    BLUE:  Symbol("blue"),
    GREEN: Symbol("green")
});

class Settings {
    constructor() {
        this.CellSize = 20;
        this.MazeColor = Colors.BLUE;

        this.Playable = false;
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
        }
    }

    clear() {
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }
}

class Wall {
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

    generateSideWinderMaze(width, height) {
        return generateSideWinderMaze(width, height)
    }
}

class Game {
    constructor(canvas, engine, config, math) {
        this.canvas = canvas;
        this.engine = engine;
        this.config = config;
        this.math = math;

        this.cellSize = this.config.getInstance().CellSize;
        this.walls = [];
    }

    async init() {
        await this.engine.init()
    }

    draw() {
        this.canvas.clear();
        for (let i = 0; i < this.walls.length; i++) {
            this.walls[i].draw(this.canvas.ctx);
        }
    }

    update() {
        let mazeWidth = this.math.floor(this.canvas.canvas.width / this.cellSize);
        let mazeHeight = this.math.floor(this.canvas.canvas.height / this.cellSize);

        this.walls.splice(0,this.walls.length)

        this.maze = JSON.parse(this.engine.generateSideWinderMaze(mazeWidth,mazeHeight));

        for (let row = 0; row < mazeHeight; row++) {
            for (let col = 0; col < mazeWidth; col++) {
                if (this.maze.Content[row*mazeHeight+col] === false) {
                    const wall = new Wall(
                        col * this.cellSize,
                        row * this.cellSize,
                        this.cellSize,
                        this.config.getInstance().MazeColor.description,
                    );
                    this.walls.push(wall);
                }
            }
        }
    }

    async run() {
        await this.init()
        const config = this.config.getInstance();

        if(config.Playable) {
            let lastFrameTime = performance.now();

            const loop = () => {
                this.update();
                this.draw();

                requestAnimationFrame(loop);
            }

            loop();
            return;
        }

        this.update();
        this.draw();
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