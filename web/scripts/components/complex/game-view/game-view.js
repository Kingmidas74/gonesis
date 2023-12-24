import {TerrainCellSizes} from "../../../configuration/configuration.js";
import {Either} from "../../../monads/either.js";
import {CellFactory} from "./cell.factory.js";
import {CanvasWrapper2D, CanvasWrapperWebGL} from "./canvas-wrapper.js";
import {CellType} from "../../../domain/enum.js";

export class GAME_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #width = 0;
    #height = 0;

    #cellFactory;
    #canvasWrapper;

    #canvas;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then((templateContent) => {
                const template = GAME_VIEW.documentProvider.createElement("template");
                template.innerHTML = GAME_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
                return this.#setup();
            })
            .catch((err) => {
                GAME_VIEW.logger.error(err);
            });
    }

    #setup = () => {
        return new Promise((resolve) => {
            GAME_VIEW.windowProvider.setTimeout(async () => {
                this.#canvas = this.#shadow.querySelector("canvas");
                resolve();
            }, 1)
        })
    }

    #clickHandler = async (event) => {
        event.stopPropagation();
        event.preventDefault();

        const rect = this.#canvas.getBoundingClientRect();
        const x = event.clientX - rect.left;
        const y = event.clientY - rect.top;
        this.dispatchEvent(new CustomEvent("click", { detail: { x: x, y: y } }));
    }

    #windowResizeListener = GAME_VIEW.debounce(async () => {
        await this.generateGame();
        this.dispatchEvent(new CustomEvent("resize", { detail: { width: this.#width, height: this.#height } }));
    }, 250)

    #adjustCanvasSize = () => {
        if (!this.#canvas) return;

        const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));
        const lcm = (a, b) => (a * b) / gcd(a, b);
        const lcmOfArray = (arr) => arr.reduce((a, b) => lcm(a, b));

        const lsmRatio = lcmOfArray(Object.values(TerrainCellSizes));

        const parentWidth = this.offsetWidth;
        const widthToSet = parentWidth - (parentWidth % lsmRatio);
        this.#canvas.style.width = `${widthToSet}px`;

        const parentHeight = this.offsetHeight;
        const heightToSet = parentHeight - (parentHeight % lsmRatio);
        this.#canvas.style.height = `${heightToSet}px`;

        return { width: widthToSet, height: heightToSet};
    }

    async #initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            GAME_VIEW.windowProvider.fetch(
                new URL(GAME_VIEW.stylePath, new URL(import.meta.url)).href
            ),
            GAME_VIEW.windowProvider.fetch(
                new URL(GAME_VIEW.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = GAME_VIEW.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    /**
     * Generates a new world.
     * @returns {Promise<Either<World, Error>>}
     */
    generateGame = async () => {
        const {width, height} = this.#adjustCanvasSize()
        this.#canvasWrapper.init(width, height);
        this.#width = width;
        this.#height = height;
    }

    get width() {
        return this.#width;
    }

    get height() {
        return this.#height;
    }

    /**
     * Configures the view with the given configuration provider.
     * @param {ConfigurationProvider} provider
     */
    set configurationProvider(provider) {
        return this.#template.then(async () => {
            const {width, height} = this.#adjustCanvasSize();

            this.#width = width;
            this.#height = height;

            this.#canvasWrapper = new CanvasWrapper2D(this.#canvas, GAME_VIEW.documentProvider);
            this.#cellFactory = new CellFactory(provider, this.#canvasWrapper)

            await this.generateGame()

            GAME_VIEW.windowProvider.addEventListener('resize', this.#windowResizeListener);

            //this.#canvas.removeEventListener('click', this.#clickHandler);
            this.#canvas.addEventListener('click', this.#clickHandler);
        })
    }

    disconnectedCallback() {
        GAME_VIEW.windowProvider.removeEventListener('resize', this.#windowResizeListener);
    }

    /**
     * Update the world with the given world data.
     * @param {World} worldInstance - The world data in object format.
     * @private
     * @returns {void}
     */
    update(worldInstance) {
        const width = worldInstance.width;
        const height = worldInstance.height;

        for (let row = 0; row < height; row++) {
            for (let col = 0; col < width; col++) {
                if (worldInstance.cells[row*width+col].agent) {
                    this.#cellFactory.createAgent(worldInstance.cells[row*width+col]).draw();
                    continue
                }
                this.#cellFactory.createEmpty(worldInstance.cells[row*width+col]).draw();
            }
        }

        this.#canvasWrapper.render();
    }
}
