import {Colors, AgentConfiguration, Topologies} from "../../../configuration/configuration.js";

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

                this.#shadow.getElementById("initialCount").value = {min:0, max:500, value: config.InitialCount};
                this.#shadow.getElementById("color").value = config?.Color || Colors.DARK;
                this.#shadow.getElementById("initialEnergy").value = {min:0, max:config.MaxEnergy, value: config.InitialEnergy};
                this.#shadow.getElementById("maxEnergy").value = {min:config.InitialEnergy, max:1000, value: config.MaxEnergy};
                this.#shadow.getElementById("reproductionEnergyCost").value = {min:0, max:500, value: config.ReproductionEnergyCost};

                this.#shadow.getElementById(`brain-volume`).data = [{name: 'small', value: 16}, {name: 'medium', value: 64}, {name: 'large', value: 256}];
                this.#shadow.getElementById(`brain-volume`).value = 16;

                /*this.#shadow.getElementById("initialCount").addEventListener("change", (event) => {
                    event.stopPropagation();
                    this.dispatchEvent(new AGENT_SETTINGS.windowProvider.CustomEvent('change', {
                        detail: { value: this.config }
                    }))
                });*/
/*
                this.#shadow.getElementById("initialEnergy").addEventListener("change", (event) => {
                    event.stopPropagation();
                    this.#shadow.getElementById("maxEnergy").value = {min:event.detail.value, max:1000, value: config.MaxEnergy};
                });

                this.#shadow.getElementById("maxEnergy").addEventListener("change", (event) => {
                    event.stopPropagation();
                    this.#shadow.getElementById("initialEnergy").value = {min:0, max:event.detail.value, value: config.InitialEnergy};
                });*/


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
                BrainVolume: parseInt(this.#shadow.getElementById(`brain-volume`).value),
            })
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
