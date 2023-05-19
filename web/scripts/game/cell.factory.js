import {Wall, Agent, Empty} from "../render/cell.js";

class CellFactory {
    /**
     * @type {ConfigurationProvider} The configuration of the game.
     */
    #configuration;

    #cellColors;

    /**
     * Constructs a new instance of CellFactory.
     * @param {ConfigurationProvider} configurationProvider
     */
    constructor(configurationProvider) {
        this.#configuration = configurationProvider;
    }

    createWall(x, y)   {
        return new Wall(x, y, this.#configuration.getInstance().MazeColor);
    }

    createEmpty(x, y)   {
        return new Empty(x, y, "#ffffff");
    }

    createAgent(x, y, energy) {
        return new Agent(x, y, this.#configuration.getInstance().AgentsColor, energy);
    }
}

export { CellFactory };