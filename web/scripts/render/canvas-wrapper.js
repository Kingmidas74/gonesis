

class CanvasWrapper {
    constructor() {
    }

    clear() {
    }

    /**
     * Draw a rectangle.
     * @param {number} x - The x position.
     * @param {number} y - The y position.
     * @param {number} width - The width of the rectangle.
     * @param {number} height - The height of the rectangle.
     * @param {string} color - The color of the rectangle in RGBA format.
     */
    drawRect(x, y, width, height, color) {
    }
}


/** Class representing a canvas wrapper. */
class CanvasWrapper2D extends CanvasWrapper{
    /**
     * @type {HTMLCanvasElement}
     * @private
     */
    #canvas

    /**
     *
     * @param {HTMLCanvasElement} canvasElement
     */
    constructor(canvasElement) {
        super();
        this.#canvas = canvasElement;
        if(this.#canvas.getContext)
        {
            this.ctx = this.#canvas.getContext("2d");

            this.#canvas.width = document.body.clientWidth;
            this.#canvas.height = document.body.clientHeight;

            this.#canvas.style.transform = 'translate3d(0, 0, 0)';
        }
    }

    /**
     * Gets the width of the canvas element.
     * @returns {number} The width of the canvas.
     */
    get width() {
        return this.#canvas.width;
    }

    /**
     * Gets the height of the canvas element.
     * @returns {number} The height of the canvas.
     */
    get height() {
        return this.#canvas.height;
    }

    /**
     * Gets the canvas context.
     * @returns {CanvasRenderingContext2D} The canvas context.
     */
    get context() {
        return this.ctx;
    }

    /**
     * Clears the entire canvas by erasing the contents.
     */
    clear() {
        this.ctx.clearRect(0, 0, this.#canvas.width, this.#canvas.height);
    }

    /**
     * Draw a rectangle.
     * @param {number} x - The x position.
     * @param {number} y - The y position.
     * @param {number} width - The width of the rectangle.
     * @param {number} height - The height of the rectangle.
     * @param {string} color - The color of the rectangle in RGBA format.
     */
    drawRect(x, y, width, height, color) {
        this.ctx.fillStyle = color;
        this.ctx.fillRect(x, y, width, height);
    }
}

class CanvasWrapperWebGL extends CanvasWrapper{
    /**
     * @type {HTMLCanvasElement}
     * @private
     */
    #canvas;

    /**
     * @type {WebGLRenderingContext}
     * @private
     */
    #context;

    /**
     * @type {WebGLProgram} The shader program.
     */
    #shaderProgram

    /**
     * @type {WebGLUniformLocation} The position attribute location.
     */
    #colorUniformLocation

    /**
     * @type {GLint} The position attribute location.
     */
    #positionAttributeLocation

    /**
     * @param {HTMLCanvasElement} canvas - The canvas.
     */
    constructor(canvas) {
        super();
        const actualCanvasWidth = canvas.offsetWidth;
        const actualCanvasHeight = canvas.offsetHeight;
        canvas.width = actualCanvasWidth;
        canvas.height = actualCanvasHeight;
        this.#canvas = canvas;

        const gl = this.#canvas.getContext("webgl");
        if (!gl) {
            throw new Error("WebGL not supported");
        }
        gl.viewport(0, 0, gl.canvas.width, gl.canvas.height);
        this.#context = gl;

        // Initialize a simple shader program
        const vsSource = `
            attribute vec4 a_position;
            void main() {
                gl_Position = a_position;
            }
        `;
        const fsSource = `
            precision mediump float;
            uniform vec4 u_color;
            void main() {
                gl_FragColor = u_color;
            }
        `;
        const vertexShader = this.#createShader(this.#context.VERTEX_SHADER, vsSource);
        const fragmentShader = this.#createShader(this.#context.FRAGMENT_SHADER, fsSource);

        this.#shaderProgram = this.#createProgram(vertexShader, fragmentShader);

        this.#positionAttributeLocation = this.#context.getAttribLocation(this.#shaderProgram, "a_position");
        this.#colorUniformLocation = this.#context.getUniformLocation(this.#shaderProgram, "u_color");
    }

    /**
     * Create a shader
     * @param {number} type - The type of the shader (VERTEX_SHADER | FRAGMENT_SHADER).
     * @param {string} source - The GLSL source code for the shader.
     * @returns {WebGLShader} The created shader.
     * @private
     */
    #createShader(type, source) {
        const shader = this.#context.createShader(type);
        this.#context.shaderSource(shader, source);
        this.#context.compileShader(shader);
        return shader;
    }

    /**
     * Create a shader program
     * @param {WebGLShader} vertexShader - The compiled vertex shader.
     * @param {WebGLShader} fragmentShader - The compiled fragment shader.
     * @returns {WebGLProgram} The created program.
     * @private
     */
    #createProgram(vertexShader, fragmentShader) {
        const program = this.#context.createProgram();
        this.#context.attachShader(program, vertexShader);
        this.#context.attachShader(program, fragmentShader);
        this.#context.linkProgram(program);
        return program;
    }

    /**
     * Gets the width of the canvas element.
     * @returns {number} The width of the canvas.
     */
    get width() {
        return this.#canvas.width;
    }

    /**
     * Gets the height of the canvas element.
     * @returns {number} The height of the canvas.
     */
    get height() {
        return this.#canvas.height;
    }

    /**
     * Gets the canvas context.
     * @returns {CanvasRenderingContext2D} The canvas context.
     */
    get context() {
        return this.ctx;
    }

    /**
     * Clear the canvas.
     */
    clear() {
        this.#context.clearColor(0.0, 0.0, 0.0, 0.0); // Set clear color to black, fully transparent
        this.#context.clear(this.#context.COLOR_BUFFER_BIT);
    }

    /**
     * Draw a rectangle.
     * @param {number} x - The x position.
     * @param {number} y - The y position.
     * @param {number} width - The width of the rectangle.
     * @param {number} height - The height of the rectangle.
     * @param {string} color - The color of the rectangle in RGBA format.
     */
    drawRect(x, y, width, height, color) {
        const gl = this.#context;

        // Convert the color from RGB to 0.0-1.0
        const r = parseInt(color.slice(1, 3), 16) / 255;
        const g = parseInt(color.slice(3, 5), 16) / 255;
        const b = parseInt(color.slice(5, 7), 16) / 255;
        const colors = [r, g, b, 1.0];

        const positions = [
            x, y,
            x + width, y,
            x, y + height,
            x, y + height,
            x + width, y,
            x + width, y + height,
        ];

        // Convert pixel positions to clip space


        for (let i = 0; i < positions.length; i += 2) {
            positions[i] = 2 * positions[i] / this.#canvas.width - 1;
            positions[i + 1] = 1 - 2 * positions[i + 1] / this.#canvas.height;
        }

        // Create a buffer for the positions
        const positionBuffer = gl.createBuffer();
        gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);
        gl.bufferData(gl.ARRAY_BUFFER, new Float32Array(positions), gl.STATIC_DRAW);

        gl.useProgram(this.#shaderProgram);

        // Enable the position attribute
        gl.enableVertexAttribArray(this.#positionAttributeLocation);
        gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);
        gl.vertexAttribPointer(this.#positionAttributeLocation, 2, gl.FLOAT, false, 0, 0);

        // Set the color uniform
        gl.uniform4fv(this.#colorUniformLocation, colors);

        gl.drawArrays(gl.TRIANGLES, 0, positions.length / 2);
    }
}
export {CanvasWrapper2D, CanvasWrapperWebGL};