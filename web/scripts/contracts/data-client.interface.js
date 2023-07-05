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

    async step() {
        throw new Error("Not implemented");
    }

}

export { IDataClient }