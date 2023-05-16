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

    generateBinaryMaze(width, height) {
        return generateBinaryMaze(width, height)
    }

    generateBorder(width, height) {
        return generateBorder(width, height)
    }

    generateGridMaze(width, height) {
        return generateGridMaze(width, height)
    }

    generateAldousBroderMaze(width, height) {
        return generateAldousBroderMaze(width, height)
    }

    runGame(width, height, agentsCount) {
        return runGame(width, height, agentsCount)
    }

    update(state) {
        return updateState(state)
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
        this.walls = [];

        this.#daysQueue = [];
        this.#pastDays =[];
    }

    lighter = (color, percent) => {
        // Remove the "#" symbol from the color
        color = color.substring(1);

        // Parse the color values
        const r = parseInt(color.substring(0, 2), 16);
        const g = parseInt(color.substring(2, 4), 16);
        const b = parseInt(color.substring(4, 6), 16);

        // Calculate the lighter values
        const lighterR = Math.round(r + (255 - r) * (percent / 100));
        const lighterG = Math.round(g + (255 - g) * (percent / 100));
        const lighterB = Math.round(b + (255 - b) * (percent / 100));

        // Convert the lighter values back to hexadecimal
        return `#${(lighterR.toString(16)).padStart(2, '0')}${(lighterG.toString(16)).padStart(2, '0')}${(lighterB.toString(16)).padStart(2, '0')}`;
    }

    async init() {
        await this.engine.init()

        let mazeWidth = this.math.floor(this.canvas.canvas.width / this.cellSize);
        let mazeHeight = this.math.floor(this.canvas.canvas.height / this.cellSize);


        this.world = JSON.parse(this.engine.runGame(mazeWidth, mazeHeight, this.config.getInstance().InitialAgensCount))
        for (let row = 0; row < mazeHeight; row++) {
            for (let col = 0; col < mazeWidth; col++) {
                if (this.world.cells[row*mazeWidth+col].cellType === 3) {
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

        for (const agent of this.world.agents) {
            const wall = new Wall(
                agent.x * this.cellSize,
                agent.y * this.cellSize,
                this.cellSize,
                this.config.getInstance().AgentsColor.description,
            );
            this.walls.push(wall);
        }
        //TODO: call this.engine.update(state).
        /*this.maze = JSON.parse(((alg, width, height) => {
            switch (alg) {
                case MazeGenerators.AldousBroder: {
                    return this.engine.generateAldousBroderMaze(width, height)
                }
                case MazeGenerators.SideWinder: {
                    return this.engine.generateSideWinderMaze(width, height)
                }
                case MazeGenerators.Binary: {
                    return this.engine.generateBinaryMaze(width, height)
                }
                case MazeGenerators.Border: {
                    return this.engine.generateBorder(width, height)
                }
                default: {
                    return this.engine.generateGridMaze(width, height)
                }
            }
        })(this.config.getInstance().MazeGenerator, mazeWidth, mazeHeight))
        ;

        for (let row = 0; row < mazeHeight; row++) {
            for (let col = 0; col < mazeWidth; col++) {
                if (this.maze.content[row*mazeWidth+col] === false) {
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
         */
    }

    draw() {
        this.canvas.clear();
        for (let i = 0; i < this.walls.length; i++) {
            this.walls[i].draw(this.canvas.ctx);
        }
    }

    update() {
        if(this.#daysQueue.length === 0) {
            return false;
        }
        const newState = this.#daysQueue.shift();
        if (this.livingAgentsCount(newState) === 0) {
            return false;
        }
        this.walls = [];
        for (let row = 0; row < newState.height; row++) {
            for (let col = 0; col < newState.width; col++) {
                if (newState.cells[row*newState.width+col].cellType === 3) {
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

        for (const agent of newState.agents) {
            const wall = new Wall(
                agent.x * this.cellSize,
                agent.y * this.cellSize,
                this.cellSize,
                this.config.getInstance().AgentsColor.description,
            );
            this.walls.push(wall);
        }
        this.#pastDays.push(newState);
        return true;
    }

    saveState(state) {
        this.#daysQueue.push(state)
    }

    /**
     * @param {World} world
     * @return {number} count of living agents
     */
    livingAgentsCount(world) {
        let count = 0;
        for (const a in world.agents) {
            if (a.energy > 0) {
                count++;
            }
        }
    }

    async run() {
        await this.init()
        const config = this.config.getInstance();

        if(config.Playable) {
            let lastFrameTime = performance.now();

            const loop = () => {
                const updateResult = this.update();
                this.draw();

                if (updateResult) {
                    requestAnimationFrame(loop);
                }
                else {
                    console.log("Game over");
                    console.log(this.#pastDays)
                }


            }

            loop();
            return;
        }

        this.update();
        this.draw();
    }
}

/**
 * @type {Game}
 */
let game;

function fromGo(day, data) {
    game.saveState(JSON.parse(data))
}

(async (canvasID, wasmFile) => {
    const settingsProvider = new SettingsProvider();

    const mathProvider = new MathProvider();

    const canvas = new Canvas(canvasID);
    const engine = new Engine(wasmFile);

    game = new Game(canvas, engine, settingsProvider, mathProvider);
    await game.run()
})("canvas", "engine.wasm").then(() => console.log("Ready"))