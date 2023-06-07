import {Colors, AgentConfiguration } from "../../../application/configuration/configuration.js";

export class AGENT_SETTINGS extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser().catch((err) => {
            AGENT_SETTINGS.logger.error(err);
        });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            AGENT_SETTINGS.windowProvider.fetch(
                new URL(AGENT_SETTINGS.stylePath, new URL(import.meta.url)).href
            ),
            AGENT_SETTINGS.windowProvider.fetch(
                new URL(AGENT_SETTINGS.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = AGENT_SETTINGS.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    /**
     * @param {AgentConfiguration} config - color in hsla format
     */
    set config(config) {
        if (!this.isConnected) {
            this.#pendingData = config;
            return;
        }

        this.#template
            .then((templateContent) => {
                const template = AGENT_SETTINGS.documentProvider.createElement("template");
                template.innerHTML = AGENT_SETTINGS.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#shadow.getElementById("initialCount").value = {min:0, max:500, value: config.InitialCount, title:"Count"};
                this.#shadow.getElementById("color").value = config?.Color || Colors.DARK;
                this.#shadow.getElementById("initialEnergy").value = {min:0, max:500, value: config.InitialEnergy, title:"Energy"};
                this.#shadow.getElementById("maxEnergy").value = {min:0, max:1000, value: config.MaxEnergy, title:"Max Energy"};
                this.#shadow.getElementById("reproductionEnergyCost").value = {min:0, max:500, value: config.ReproductionEnergyCost, title:"Reproduction cost"};


            })
            .catch((err) => {
                AGENT_SETTINGS.logger.error(err);
            });
    }

    /**
     * @returns {AgentConfiguration}
     */
    get config() {
        return new AgentConfiguration({
                InitialCount: this.#shadow.getElementById("initialCount").value,
                Color: this.#shadow.getElementById("color").value,
                InitialEnergy: this.#shadow.getElementById("initialEnergy").value,
                MaxEnergy: this.#shadow.getElementById("maxEnergy").value,
                ReproductionEnergyCost: this.#shadow.getElementById("reproductionEnergyCost").value,
            })
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
