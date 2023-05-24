import {Agent, Cell, World} from "./domain.js";


class JsonProvider {
    parse(json) {
        const object = JSON.parse(json);
        return object
        const cells = object.cells.map(cellData => new Cell(cellData.cellType, cellData.energy));

        const agents = object.agents.map(agentData =>
            new Agent(agentData.commands, agentData.x, agentData.y, agentData.energy, agentData.agentType));

        return new World(object.width, object.height, cells, agents);
    }

    stringify(object) {
        return JSON.stringify(object)
    }
}

export { JsonProvider }
