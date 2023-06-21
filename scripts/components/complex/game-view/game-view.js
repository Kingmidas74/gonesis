import {Application} from "../../../application/application.js";
import {Configuration, TerrainCellSizes} from "../../../application/configuration/configuration.js";
import {Either} from "../../../application/monads/either.js";

export class GAME_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #isGenerated = false;

    /**
     * @type {Application}
     */
    #application;

    #width = 0;
    #height = 0;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then((templateContent) => {
                const template = GAME_VIEW.documentProvider.createElement("template");
                template.innerHTML = GAME_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .then(this.#setup)
            .then(this.generateGame)
            .catch((err) => {
                GAME_VIEW.logger.error(err);
            });
        GAME_VIEW.windowProvider.addEventListener('resize', this.#windowResizeListener);
    }

    #windowResizeListener = async () => {
        const { width, height } = this.#adjustCanvasSize();
        this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Width = width;
        this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Height = height;
        await this.generateGame();
    }

    #adjustCanvasSize = () => {

        const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));

        const lcm = (a, b) => (a * b) / gcd(a, b);

        const lcmOfArray = (arr) => arr.reduce((a, b) => lcm(a, b));

        const lsmRatio = lcmOfArray(Object.values(TerrainCellSizes));

        const canvas = this.#shadow.querySelector('canvas');
        if (!canvas) return;
        const parentWidth = this.offsetWidth;
        const widthToSet = parentWidth - (parentWidth % lsmRatio);
        canvas.style.width = `${widthToSet}px`;

        const parentHeight = this.offsetHeight;
        const heightToSet = parentHeight - (parentHeight % lsmRatio);
        canvas.style.height = `${heightToSet}px`;

        this.#width = widthToSet;
        this.#height = heightToSet;

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

    #setup = () => {
        return new Promise((resolve, reject) => {
            GAME_VIEW.windowProvider.setTimeout(async () => {
                const {width, height} = this.#adjustCanvasSize();
                this.#application = new Application();
                await this.#application.configure(GAME_VIEW.windowProvider, GAME_VIEW.documentProvider, this.#shadow.querySelector("canvas"), "engine.wasm")
                this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Width = width;
                this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Height = height;
                resolve();
            }, 1)
        })
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    /**
     * Next step of the game.
     * @returns {Promise<void>}
     */
    async nextStep() {
        try {
            if(!this.#isGenerated) {
                (await this.generateGame())
            }
            await this.#application.game.step();
        } catch (error) {
            console.error(error);
        }
    }

    /**
     * Plays the game.
     * @returns {Promise<void>}
     */
    async playGame() {
        if(!this.#isGenerated) {
            (await this.generateGame())
                .orElse(error => {
                    this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('over', {
                        detail: { value: error }
                    }))
                })
        }

        this.#application.configurationProvider.getInstance().Playable = true;
        this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('start'))
        const desiredFPS = 10;
        const timeStep = 1000 / desiredFPS;
        let lastTime = GAME_VIEW.windowProvider.performance.now();

        const loop = async (currentTime) => {
            const deltaTime = currentTime - lastTime;

            if(deltaTime >= timeStep) {
                if(!this.#application.configurationProvider.getInstance().Playable) {
                    return;
                }
                lastTime = currentTime - (deltaTime % timeStep);
                (await this.#application.game.step())
                    .map(worldInstance => {
                        if(this.#application.game.livingAgentsCount(worldInstance) === 0 ||
                            (!this.#application.configurationProvider.getInstance().WorldConfiguration.OneAgentTypeMode &&
                                this.#application.game.isOnlyOneAgentTypeAlive(worldInstance))) {
                            this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('finish', {
                                detail: {value: worldInstance}
                            }))
                            this.pauseGame();
                        } else {
                            this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('update', {
                                detail: {value: worldInstance}
                            }))
                        }
                        return worldInstance;
                    }).orElse(_ => {
                        this.pauseGame();
                    })
            }
            GAME_VIEW.windowProvider.requestAnimationFrame(loop);
        }
        GAME_VIEW.windowProvider.requestAnimationFrame(loop);
    }

    /**
     * Pauses the game.
     * @returns {void}
     */
    pauseGame() {
        this.#application.configurationProvider.getInstance().Playable = false;
    }

    /**
     * Generates a new world.
     * @returns {Promise<Either<World, Error>>}
     */
    generateGame = async () => {
        this.pauseGame();
        const {width, height} = this.#adjustCanvasSize()
        this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Width = width / this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize;
        this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.Height = height / this.#application.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize;
        return (await this.#application.game.init())
            .map((world) => {
                this.#isGenerated = true;
                return world;
            })
    }

    /**
     * Updates the settings of the game.
     * @param {Configuration} settings
     */
    updateSettings(settings) {
        this.#application.configurationProvider.updateConfiguration(settings);
    }

    /**
     * @returns {Promise<Configuration>}
     */
    get config() {
        return this.#template.then(() => this.#application.configurationProvider.getInstance());
    }

    disconnectedCallback() {
        GAME_VIEW.windowProvider.removeEventListener('resize', this.#windowResizeListener);
    }
}
