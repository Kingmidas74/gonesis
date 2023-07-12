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
     */
    async initWorld(configurationProvider) {
        throw new Error("Not implemented");
    }

    /**
     * @param {ConfigurationProvider} configurationProvider
     * @return {Promise<void>}
     */
    async step(configurationProvider) {
        throw new Error("Not implemented");
    }

}

export { IDataClient }