import {Configuration, Topologies} from "../../../application/configuration/configuration.js";

export class GAME_SETTINGS extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #controls = {}

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });
        this.#shadow.addEventListener('click', this.#onClickHandler);

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = GAME_SETTINGS.documentProvider.createElement("template");
                template.innerHTML = GAME_SETTINGS.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#controls = {
                    plantSettings: this.#shadow.getElementById("plantSettings"),
                    herbivoreSettings: this.#shadow.getElementById("herbivoreSettings"),
                    carnivoreSettings: this.#shadow.getElementById("carnivoreSettings"),
                    omnivoreSettings: this.#shadow.getElementById("omnivoreSettings"),
                    saveBtn: this.#shadow.getElementById("saveBtn"),
                }
            })
            .catch((err) => {
                GAME_SETTINGS.logger.error(err);
            });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            GAME_SETTINGS.windowProvider.fetch(
                new URL(GAME_SETTINGS.stylePath, new URL(import.meta.url)).href
            ),
            GAME_SETTINGS.windowProvider.fetch(
                new URL(GAME_SETTINGS.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = GAME_SETTINGS.documentProvider.createElement("style");
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
            this.#shadow.getElementById('terrain-settings').value = config.WorldConfiguration.MazeType;
            this.#shadow.getElementById(`topology-types`).data = Object.entries(Topologies).map(([name, value]) => ({ name, value }))
            this.#shadow.getElementById(`topology-types`).value = config.WorldConfiguration.Topology;
            this.#shadow.getElementById('cellSize').value = config.WorldConfiguration.CellSize;

            this.#controls.plantSettings.config = config.PlantConfiguration
            this.#controls.herbivoreSettings.config = config.HerbivoreConfiguration
            this.#controls.carnivoreSettings.config = config.CarnivoreConfiguration
            this.#controls.omnivoreSettings.config = config.OmnivoreConfiguration

            this.#setFirstTabsActive();
        })
    }

    /**
     * Get config data
     * @returns {Configuration}
     */
    get config() {
        return new Configuration({
            worldConfiguration: {
                CellSize: this.#shadow.getElementById('cellSize').value,
                MazeType: this.#shadow.querySelector('#terrain-settings').value,
                Topology: this.#shadow.querySelector('#topology-types').value,
            },
            plantConfiguration: this.#controls.plantSettings.config,
            herbivoreConfiguration: this.#controls.herbivoreSettings.config,
            carnivoreConfiguration: this.#controls.carnivoreSettings.config,
            omnivoreConfiguration: this.#controls.omnivoreSettings.config,
        })
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    #setFirstTabsActive() {
        const topLevelTabsContainer = this.#shadow.querySelector('.settings-header .tab-container');

        const firstTopLevelTab = topLevelTabsContainer?.querySelector('.tab');
        this.#activateTab(firstTopLevelTab);

        const nestedTabGroups = this.#shadow.querySelectorAll('.form__fieldset[data-tab="true"]');
        nestedTabGroups.forEach(tabGroup => {
            const firstNestedTab = tabGroup.querySelector('.tab');
            this.#activateTab(firstNestedTab);
        });
    }

    #activateTab(tab) {
        if (!tab) {
            return;
        }

        const targetFieldset = this.#shadow.querySelector(`#${tab.getAttribute('data-target')}`);
        if (!targetFieldset) {
            return;
        }

        tab.classList.add('active');
        targetFieldset.classList.add('active');
    }

    #onClickHandler = (e) => {
        const updateBtn = e.target.closest('#saveBtn');
        if (updateBtn) {
            e.preventDefault();
            e.stopPropagation();

            this.dispatchEvent(new GAME_SETTINGS.windowProvider.CustomEvent('update', {
                detail: { value: this.config }
            }))
        }

        const clickedTab = e.target.closest('[data-target]');
        if (!clickedTab) return;

        const container = clickedTab.parentNode;

        Array.from(container.children).forEach(tab => {
            tab.classList.remove('active');
        });

        clickedTab.classList.add('active');

        const parentFieldset = container.closest('.form__fieldset');
        if (parentFieldset) {
            const allNestedFieldsets = parentFieldset.querySelectorAll('.form__fieldset');
            allNestedFieldsets.forEach(fieldset => fieldset.classList.remove('active'));
        }

        const form = this.#shadow.querySelector('.settings--form--content');
        const allSiblings = Array.from(form.children).filter(child => child !== parentFieldset);
        allSiblings.forEach(sibling => sibling.classList.remove('active'));

        const targetFieldset = form.querySelector(`#${clickedTab.getAttribute('data-target')}`);
        if (targetFieldset) targetFieldset.classList.add('active');
    }

}
