/**
 * An interface implementing a JSON interface
 *
 * @interface
 * @see {JSON}
 */
class IJSONProvider {
    /**
     * Parses a JSON string, constructing the JavaScript value or object described by the string.
     * @param {string} text A valid JSON string.
     * @returns {any} The JavaScript value or object described by the string.
     */
    parse(text) {
        throw new Error("Not implemented");
    }

    /**
     * Converts a JavaScript value to a JSON string.
     * @param {any} value A JavaScript value, usually an object or array, to be converted.
     * @returns {string} A JSON string representing the given value.
     */
    stringify(value) {
        throw new Error("Not implemented");
    }
}

export { IJSONProvider }