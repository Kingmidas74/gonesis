import {CellType} from "../../domain/enum.js";
import {Either} from "../../monads/either.js";
import {IGame} from "../../contracts/game.interface.js"

class Game extends IGame {

    /**
     * @type {IDataClient}
     * @private
     */
    #dataClient;

    /**
     *
     * @param {IDataClient} dataClient
     */
    constructor(dataClient) {
        super();

        this.#dataClient = dataClient;
    }

    /**
     * Initialize game's world
     * @param {ConfigurationProvider} configurationProvider
     * @returns {Promise<Either<World, Error>>} void if world is initialized successfully, error otherwise
     */
    async initWorld(configurationProvider) {
        return await this.#dataClient.initWorld(configurationProvider);
    }

    /**
     * Step game
     * @param {ConfigurationProvider} configurationProvider
     * @returns {Promise<Either<World, Error>>}
     */
    async step(configurationProvider) {
        return await this.#dataClient.step(configurationProvider);
    }

    /**
     * @param {World} worldInstance - The world data in object format.
     * @returns {number}
     */
    calculateGeneration(worldInstance) {
        let maxGeneration = 0;
        for (const c of worldInstance.cells) {
            //console.log(c?.agent?.generation, maxGeneration,"maxGeneration",c)
            if (c?.agent?.generation > maxGeneration) {
                maxGeneration = c?.agent?.generation;
            }
        }
        return maxGeneration;
    }

    /**
     * @param {World} worldInstance - The world data in object format.
     * @returns {Array<Agent>}
     */
    agents = (worldInstance) => {
        return worldInstance?.cells(c => c.cellType === CellType.AGENT && c.agent?.energy > 0).map(c=>c.a) ?? []
    }

    /**
     * @param {World} worldInstance - The world data in object format.
     * @return {number} count of living agents
     */
    livingAgentsCount(worldInstance) {
        return worldInstance.cells.filter(c => c.agent?.energy > 0).length;
    }

    isOnlyOneAgentTypeAlive(worldInstance) {
        const firstLivingAgentType = worldInstance.cells.find(c => c.agent?.energy > 0)?.agent?.agentType;
        const everyAgentIsSameType = worldInstance.cells.every(c => c.agent?.agentType === firstLivingAgentType);
        return everyAgentIsSameType && firstLivingAgentType !== undefined;
    }

    cell(x, y, worldInstance) {
        return worldInstance.cells.find(c => c.x === x && c.y === y);
    }
}

export {Game}