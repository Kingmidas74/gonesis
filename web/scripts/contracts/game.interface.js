/**
 * @interface
 */
class IGame {

    /**
     *
     * @param {ConfigurationProvider} configurationProvider
     */
    async initWorld(configurationProvider) {
        throw new Error("Not implemented");
    }

    async step() {
        throw new Error("Not implemented");
    }

    calculateGeneration(worldInstance) {
        throw new Error("Not implemented");
    }

    agents(worldInstance) {
        throw new Error("Not implemented");
    }

    livingAgentsCount(worldInstance) {
        throw new Error("Not implemented");
    }

    isOnlyOneAgentTypeAlive(worldInstance) {
        throw new Error("Not implemented");
    }

    cell(x, y, worldInstance) {
        throw new Error("Not implemented");
    }
}

export {IGame}