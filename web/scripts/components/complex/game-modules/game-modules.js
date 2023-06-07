import {Configuration, Topologies} from "../../../application/configuration/configuration.js";

export class GAME_MODULES extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #controls = {}

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });
       // this.#shadow.addEventListener('click', this.#onClickHandler);

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = GAME_MODULES.documentProvider.createElement("template");
                template.innerHTML = GAME_MODULES.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .catch((err) => {
                GAME_MODULES.logger.error(err);
            });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            GAME_MODULES.windowProvider.fetch(
                new URL(GAME_MODULES.stylePath, new URL(import.meta.url)).href
            ),
            GAME_MODULES.windowProvider.fetch(
                new URL(GAME_MODULES.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = GAME_MODULES.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    /**
     * Set config data
     * @param {Configuration} config
     */
    set config(config) {
        if (!this.isConnected) {
            this.#pendingData = config;
            return;
        }

        this.#template.then(_ => {

        })
    }

    /**
     * Get config data
     * @returns {Configuration}
     */
    get config() {
        return new Configuration()
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

}
