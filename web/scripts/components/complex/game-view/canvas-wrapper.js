/**
 * @interface
 */
class CanvasWrapper {
    constructor() {
    }

    /**
     * Draws a circle.
     * @param x
     * @param y
     * @param radius
     * @param color
     */
    drawCircle(x, y, radius, color) {
        throw new Error("Not implemented");
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
        throw new Error("Not implemented");
    }

    drawLine(xStart, yStart, xEnd, yEnd, color) {
        throw new Error("Not implemented");
    }

    /**
     * Gets the width of the canvas element.
     * @returns {number} The width of the canvas.
     */
    get width() {
        throw new Error("Not implemented");
    }

    /**
     * Gets the height of the canvas element.
     * @returns {number} The height of the canvas.
     */
    get height() {
        throw new Error("Not implemented");
    }

    /**
     * Renders the buffer canvas to the main canvas.
     */
    render() {
        throw new Error("Not implemented");
    }

    /**
     * Initialize canvas
     * @param {number} width
     * @param {number} height
     */
    init = (width, height) => {
        throw new Error("Not implemented");
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
     * @type {HTMLCanvasElement} canvas for buffer
     * @private
     */
    #bufferCanvas;

    /**
     * @type {CanvasRenderingContext2D} ctx for buffer canvas
     * @private
     */
    #bufferCtx;

    /**
     * @type {Map}
     */
    #previousFrame

    /**
     *
     * @param {HTMLCanvasElement} canvasElement
     * @param {Document} documentProvider
     */
    constructor(canvasElement, documentProvider) {
        super();
        this.#canvas = canvasElement;
        this.#previousFrame = new Map();

        if(this.#canvas.getContext)
        {
            this.#bufferCanvas = documentProvider.createElement("canvas");
            this.#bufferCtx = this.#bufferCanvas.getContext("2d");
            this.ctx = this.#canvas.getContext("2d");
            this.#canvas.style.transform = 'translate3d(0, 0, 0)';
        }
    }

    /**
     * Initialize canvas
     * @param {number} width
     * @param {number} height
     */
    init = (width, height) => {
        this.#canvas.width = width;
        this.#canvas.height = height;
        this.#bufferCanvas.width = this.#canvas.width;
        this.#bufferCanvas.height = this.#canvas.height;
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
     * Renders the buffer canvas to the main canvas.
     */
    render() {
        //this.ctx.clearRect(0, 0, this.#bufferCanvas.width, this.#bufferCanvas.height);
        this.ctx.drawImage(this.#bufferCanvas, 0, 0);
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
        let key = `${x},${y}`;

        //if(!this.#previousFrame.has(key) || this.#previousFrame.get(key) !== color) {
        this.#bufferCtx.fillStyle = color;
        this.#bufferCtx.fillRect(x, y, width, height);
        this.#previousFrame.set(key, color);
        //}
    }

    drawLine(xStart, yStart, xEnd, yEnd, color) {
        this.#bufferCtx.beginPath();
        this.#bufferCtx.strokeStyle = color;
        this.#bufferCtx.moveTo(xStart, yStart);
        this.#bufferCtx.lineTo(xEnd, yEnd);
        this.#bufferCtx.stroke();
    }

    /**
     * Draws a circle.
     * @param x
     * @param y
     * @param radius
     * @param color
     */
    drawCircle(x, y, radius, color) {
        let key = `${x},${y}`;
        //if(!this.#previousFrame.has(key) || this.#previousFrame.get(key) !== color) {
        this.#bufferCtx.beginPath();
        this.#bufferCtx.arc(x + radius, y + radius, radius, 0, 2 * Math.PI, false);
        this.#bufferCtx.fillStyle = color;
        this.#bufferCtx.fill();
        this.#previousFrame.set(key, color);
        //}
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
        this.#canvas = canvas;

        const gl = this.#canvas.getContext("webgl");
        if (!gl) {
            throw new Error("WebGL not supported");
        }

        this.#context = gl;
    }

    init = (width, height) => {
        this.#canvas.width = width;
        this.#canvas.height = height;

        this.#context.viewport(0, 0, this.#context.canvas.width, this.#context.canvas.height);

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
     * Clear the canvas.
     */
    clear() {
        this.#context.clearColor(0.0, 0.0, 0.0, 0.0);
        this.#context.clear(this.#context.COLOR_BUFFER_BIT);
    }

    /**
     * Draw a rectangle.
     * @param {number} x - The x position.
     * @param {number} y - The y position.
     * @param {number} width - The width of the rectangle.
     * @param {number} height - The height of the rectangle.
     * @param {string} color - The color of the rectangle in RGBA format.
     * @return {never}
     */
    drawRect(x, y, width, height, color) {
        width-=2;
        height-=2;
        const gl = this.#context;
        // Convert the color from RGB to 0.0-1.0
        const {
            r,g,b,a
        } = this.#hslaToHex(color);
        const colors = [r, g, b, a];

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

    #previousFrame = new Map();

    /**
     * Draws a circle.
     * @param xStart
     * @param yStart
     * @param xEnd
     * @param yEnd
     * @param color
     */
    drawLine(xStart, yStart, xEnd, yEnd, color) {
        const gl = this.#context;
        // Convert the color from RGB to 0.0-1.0
        const { r, g, b, a } = this.#hslaToHex(color);
        const colors = [r, g, b, a];

        const positions = [
            xStart, yStart,
            xEnd, yEnd,
        ];

        // Convert pixel positions to clip space
        for (let i = 0; i < positions.length; i += 2) {
            positions[i] = 2 * positions[i] / this.#canvas.width - 1;
            positions[i + 1] = 1 - 2 * positions[i + 1] / this.#canvas.height;
        }

        const positionBuffer = gl.createBuffer();
        gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);
        gl.bufferData(gl.ARRAY_BUFFER, new Float32Array(positions), gl.STATIC_DRAW);

        gl.useProgram(this.#shaderProgram);

        gl.enableVertexAttribArray(this.#positionAttributeLocation);
        gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);
        gl.vertexAttribPointer(this.#positionAttributeLocation, 2, gl.FLOAT, false, 0, 0);

        gl.uniform4fv(this.#colorUniformLocation, colors);

        gl.drawArrays(gl.LINES, 0, 2);
    }

    drawCircle(x, y, radius, color) {
        radius-=2;
        x-=4;
        y-=4;
        const segments = 100; // increase for a more precise circle
        const positions = new Array(segments * 2);
        const angleIncrement = Math.PI * 2 / segments;

        for (let i = 0; i < segments; ++i) {
            positions[2 * i] = x + radius * Math.cos(i * angleIncrement);
            positions[2 * i + 1] = y + radius * Math.sin(i * angleIncrement);
        }

        // Convert pixel positions to clip space
        for (let i = 0; i < positions.length; i += 2) {
            positions[i] = 2 * (positions[i] - radius) / this.#canvas.width - 1;
            positions[i + 1] = 1 - 2 * (positions[i + 1] - radius) / this.#canvas.height;
        }

        const { r, g, b, a } = this.#hslaToHex(color);
        const colors = [r, g, b, a];

        const positionBuffer = this.#context.createBuffer();
        this.#context.bindBuffer(this.#context.ARRAY_BUFFER, positionBuffer);
        this.#context.bufferData(this.#context.ARRAY_BUFFER, new Float32Array(positions), this.#context.STATIC_DRAW);

        this.#context.useProgram(this.#shaderProgram);

        this.#context.enableVertexAttribArray(this.#positionAttributeLocation);
        this.#context.bindBuffer(this.#context.ARRAY_BUFFER, positionBuffer);
        this.#context.vertexAttribPointer(this.#positionAttributeLocation, 2, this.#context.FLOAT, false, 0, 0);

        this.#context.uniform4fv(this.#colorUniformLocation, colors);

        this.#context.drawArrays(this.#context.TRIANGLE_FAN, 0, segments);
    }


    /**
     * Renders the buffer canvas to the main canvas.
     */
    render() {

    }

    #hslaToHex = (hsla) => {
        const hslaInArray = hsla.substring(5, hsla.length-1).replace(/ /g, '').split(',');

        let h = parseInt(hslaInArray[0]) / 360; // we need to convert it to be between 0 to 1
        let s = parseInt(hslaInArray[1]) / 100; // we need to convert it to be between 0 to 1
        let l = parseInt(hslaInArray[2]) / 100; // we need to convert it to be between 0 to 1
        let a = parseFloat(hslaInArray[3]); // we need to convert it to be between 0 to 1

        return this.#hslaToRgba(h, s, l, a);
    }

    #hslaToRgba = (h, s, l, a) => {
        let r, g, b;

        if(s === 0){
            r = g = b = l; // achromatic
        } else {
            const hue2rgb = (p, q, t) => {
                if(t < 0) t += 1;
                if(t > 1) t -= 1;
                if(t < 1/6) return p + (q - p) * 6 * t;
                if(t < 1/2) return q;
                if(t < 2/3) return p + (q - p) * (2/3 - t) * 6;
                return p;
            };
            let q = l < 0.5 ? l * (1 + s) : l + s - l * s;
            let p = 2 * l - q;
            r = hue2rgb(p, q, h + 1/3);
            g = hue2rgb(p, q, h);
            b = hue2rgb(p, q, h - 1/3);
        }
        return {r,g,b,a}

    }

}

class WebGLWrapper extends CanvasWrapper{
    /**
     * @type {HTMLCanvasElement}
     * @private
     */
    #canvas;

    /**
     * @type {WebGLRenderingContext}
     * @private
     */
    #gl;

    /**
     *
     * @param {HTMLCanvasElement} canvasElement
     */
    constructor(canvasElement) {
        super();
        this.#canvas = canvasElement;

        if (this.#canvas.getContext) {
            this.#gl = this.#canvas.getContext("webgl");
            if (!this.#gl) {
                console.error("Unable to initialize WebGL. Your browser or machine may not support it.");
                return;
            }
        }
    }

    /**
     * Initialize WebGL
     * @param {number} width
     * @param {number} height
     */
    init = (width, height) => {
        this.#canvas.width = width;
        this.#canvas.height = height;
        this.#gl.viewport(0, 0, this.#gl.canvas.width, this.#gl.canvas.height);
    }

    /**
     * Gets the width of the WebGL canvas element.
     * @returns {number} The width of the WebGL canvas.
     */
    get width() {
        return this.#canvas.width;
    }

    /**
     * Gets the height of the WebGL canvas element.
     * @returns {number} The height of the WebGL canvas.
     */
    get height() {
        return this.#canvas.height;
    }

    /**
     * Clears the WebGL screen.
     */
    clearScreen() {
        this.#gl.clearColor(0.0, 0.0, 0.0, 1.0);  // Clear to black, fully opaque
        this.#gl.clearDepth(1.0);                 // Clear everything
        this.#gl.enable(this.#gl.DEPTH_TEST);           // Enable depth testing
        this.#gl.depthFunc(this.#gl.LEQUAL);            // Near things obscure far things
        this.#gl.clear(this.#gl.COLOR_BUFFER_BIT | this.#gl.DEPTH_BUFFER_BIT);
    }

    /**
     * Create shader program
     * @param {string} vsSource - vertex shader source
     * @param {string} fsSource - fragment shader source
     * @returns {WebGLProgram} program
     */
    initShaderProgram(vsSource, fsSource) {
        const vertexShader = this.loadShader(this.#gl.VERTEX_SHADER, vsSource);
        const fragmentShader = this.loadShader(this.#gl.FRAGMENT_SHADER, fsSource);
        const shaderProgram = this.#gl.createProgram();
        this.#gl.attachShader(shaderProgram, vertexShader);
        this.#gl.attachShader(shaderProgram, fragmentShader);
        this.#gl.linkProgram(shaderProgram);
        if (!this.#gl.getProgramParameter(shaderProgram, this.#gl.LINK_STATUS)) {
            alert('Unable to initialize the shader program: ' + this.#gl.getProgramInfoLog(shaderProgram));
            return null;
        }
        return shaderProgram;
    }

    /**
     * creates a shader of the given type, uploads the source and compiles it.
     * @param {number} type
     * @param {string} source
     * @returns {WebGLShader}
     */
    loadShader(type, source) {
        const shader = this.#gl.createShader(type);
        this.#gl.shaderSource(shader, source);
        this.#gl.compileShader(shader);
        if (!this.#gl.getShaderParameter(shader, this.#gl.COMPILE_STATUS)) {
            alert('An error occurred compiling the shaders: ' + this.#gl.getShaderInfoLog(shader));
            this.#gl.deleteShader(shader);
            return null;
        }
        return shader;
    }
}

export {CanvasWrapper2D, CanvasWrapperWebGL, WebGLWrapper};