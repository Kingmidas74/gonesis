/**
 * An interface for access to golang functions
 *
 * @interface
 */
class IDataClient {

    /**
     *
     * @param { ConfigurationProvider } configurationProvider
     * @see Configuration
     * @return {Promise<Either<World, Error>>} World if world is initialized successfully, error otherwise
     */
    async initWorld(configurationProvider) {
        throw new Error("Not implemented");
    }

    /**
     * @param {ConfigurationProvider} configurationProvider
     * @return {Promise<Either<World, Error>>} World if world is initialized successfully, error otherwise
     */
    async step(configurationProvider) {
        throw new Error("Not implemented");
    }

}

export { IDataClient }