class Either {
    constructor(value) {
        this._value = value;
    }

    static exception(a) {
        return new Exception(a);
    }

    static value(a) {
        return new Value(a);
    }

    /**
     * Transform value to Either
     * @param a
     * @returns {Either}
     */
    static of(a) {
        return Either.value(a);
    }

    bind(fn) {
        return this;
    }

    map(fn) {
        return this;
    }

    orElse(fn) {
        return this;
    }
}

class Value extends Either {
    /**
     * Get pure value
     * @returns {*}
     */
    get value() {
        return this._value;
    }

    /**
     * Execute function
     * @param fn
     * @returns {*}
     */
    bind(fn) {
        return fn(this._value);
    }

    /**
     * Transform result of the function to Either
     * @param fn
     * @returns {Either}
     */
    map(fn) {
        return Either.of(fn(this._value));
    }

    orElse() {
        return this;
    }
}

class Exception extends Either {
    /**
     * Get pure value
     * @returns {*}
     */
    get value() {
        throw new Error('Cannot extract the value of Exception');

    }

    /**
     * ignore the map function and return the current instance
     * @param fn
     * @returns {Exception}
     */
    map(fn) {
        return this;
    }

    /**
     * ignore the bind function and return the current instance
     * @param fn
     * @returns {Exception}
     */
    bind(fn) {
        return this;
    }

    orElse(fn) {
        return fn(this._value);
    }
}



export { Either };
