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

    static of(a) {
        return Either.value(a);
    }

    static fromNullable(val) {
        return val != null ? Either.value(val) : Either.exception(val);
    }

    static attempt(f) {
        try {
            return Either.value(f());
        } catch(e) {
            return Either.exception(e);
        }
    }

    map(fn) {
        return this;
    }

    orElse(fn) {
        return this;
    }

    chain(fn) {
        return this;
    }
}

class Exception extends Either {
    map(fn) {
        return this;  // ignore the map function and return the current instance
    }

    get value() {
        throw new Error('Cannot extract the value of a Left(a).');

    }

    getOrElse(other) {
        return other;
    }

    orElse(fn) {
        return fn(this._value);
    }

    chain(fn) {
        return this;
    }

    getOrElseThrow(a) {
        throw new Error(a);
    }

    filter(fn) {
        return this;
    }
}

class Value extends Either {
    map(fn) {
        return Either.of(fn(this._value));
    }

    get value() {
        return this._value;
    }

    getOrElse(other) {
        return this._value;
    }

    orElse() {
        return this;
    }

    chain(fn) {
        return fn(this._value);
    }

    getOrElseThrow() {
        return this._value;
    }

    filter(fn) {
        return Either.fromNullable(fn(this._value) ? this._value : null);
    }
}

export { Either };
