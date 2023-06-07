import {Application} from "../../../application/application.js";
import {Either} from "../../../application/monads/index.js";

export class GAME_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #isGenerated = false;

    /**
     * @type {Application}
     */
    #application;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
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
    }

    async initializeTemplateParser() {
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

    #setup = async () => {
        this.#application = new Application();
        await this.#application.configure(window, document, this.#shadow.querySelector("canvas"), "engine.wasm")
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
            this.generateGame().catch((error) => {
                this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('over', {
                    detail: { value: Either.exception(error) }
                }))
            })
        }
        this.#application.configurationProvider.getInstance().Playable = true;
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
                const stepResult = await this.#application.game.step();
                stepResult.map(shouldContinue => {
                    if(!shouldContinue) {
                        this.#application.configurationProvider.getInstance().Playable = false;
                        this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('over', {
                            detail: { value: stepResult }
                        }))
                    }

                }).orElse(err => {
                    this.dispatchEvent(new GAME_VIEW.windowProvider.CustomEvent('over', {
                        detail: { value: stepResult }
                    }))
                });
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
     * @returns {Promise<Either<null, Error>>}
     */
    generateGame = async () => {
        this.pauseGame();
        return (await this.#application.game.init())
            .map(() => {
                this.#isGenerated = true;
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


}
