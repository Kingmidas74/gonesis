import {IJSONProvider} from "../contracts/json.provider.interface.js";

/**
 * A class implementing a IJSONProvider interface
 * @see {IJSONProvider}
 */
class JSONProvider extends IJSONProvider {
    /**
     * Parses a JSON string, constructing the JavaScript value or object described by the string.
     * @param {string} text A valid JSON string.
     * @returns {any} The JavaScript value or object described by the string.
     */
    parse(text) {
        return JSON.parse(text);
    }

    /**
     * Converts a JavaScript value to a JSON string.
     * @param {any} value A JavaScript value, usually an object or array, to be converted.
     * @returns {string} A JSON string representing the given value.
     */
    stringify(value) {
        return JSON.stringify(value);
    }
}

export { JSONProvider }