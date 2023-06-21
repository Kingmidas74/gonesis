import { CellType, AgentType } from "../domain/enum.js";

class Cell {

    /**
     * Cell constructor.
     * @param {CellType} cellType - The type of the cell.
     * @param {number} energy - The energy of the cell.
     */
    constructor(cellType, energy) {
        this.cellType = cellType;
        this.energy = energy;
    }
}

class Agent {

    /**
     * Agent constructor.
     * @param {Array<number>} commands - The commands of the agent.
     * @param {number} x - The x coordinate of the agent.
     * @param {number} y - The y coordinate of the agent.
     * @param {number} energy - The energy of the agent.
     * @param {AgentType} agentType - The type of the agent.
     * @param {number} generation - The generation of the agent.
     */
    constructor(commands, x, y, energy, agentType, generation) {
        this.commands = commands;
        this.x = x;
        this.y = y;
        this.energy = energy;
        this.agentType = agentType;
        this.generation = generation;
    }
}

class World {

    /**
     * World constructor.
     * @param {number} width - The width of the world.
     * @param {number} height - The height of the world.
     * @param {Array<Cell>} cells - The cells of the world.
     * @param {Array<Agent>} agents - The agents of the world.
     * @param {number} currentDay - The current day of the world.
     */
    constructor(width, height, cells, agents, currentDay) {
        this.width = width;
        this.height = height;
        this.cells = cells;
        this.agents = agents;
        this.currentDay = currentDay;
    }
}



export { Cell, Agent, World };