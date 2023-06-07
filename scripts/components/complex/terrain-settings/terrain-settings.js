import {MazeGenerators} from "../../../application/configuration/configuration.js";

export class TerrainSettingsData {
    /**
     *
     * @param {String} name The item name.
     * @param {String} value The item value.
     */
    constructor(name, value) {
        this.name = name;
        this.value = value;
    }
}

export class TERRAIN_SETTINGS extends HTMLElement {

    #shadow;

    #template;

    #terrainTypesToggle
    #mazeTypesToggle

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = TERRAIN_SETTINGS.documentProvider.createElement("template");
                template.innerHTML = TERRAIN_SETTINGS.templateParser?.parse(templateContent, {});
                this.#shadow.appendChild(template.content.cloneNode(true));
                this.#terrainTypesToggle = this.#shadow.querySelector("#terrain-types")
                this.#mazeTypesToggle = this.#shadow.querySelector("#maze-types")
                const terrainTypes = Object.entries(MazeGenerators).map(([name, value]) => ({ name, value })).filter(({ value }) => [MazeGenerators.Border, MazeGenerators.Empty, MazeGenerators.Grid].includes(value));
                terrainTypes.push({ name: "Maze", value: "maze" })
                const mazeTypes = Object.entries(MazeGenerators).map(([name, value]) => ({ name, value })).filter(({ value }) => [MazeGenerators.AldousBroder, MazeGenerators.Binary, MazeGenerators.SideWinder].includes(value));
                this.#terrainTypesToggle.addEventListener("change", (event) => {
                    this.#mazeTypesToggle.classList.toggle("hidden", event?.detail?.value !== "maze");
                });
                this.#terrainTypesToggle.data = terrainTypes;
                this.#mazeTypesToggle.data = mazeTypes;
                this.#mazeTypesToggle.value = MazeGenerators.AldousBroder;

            })
            .catch((err) => {
                TERRAIN_SETTINGS.logger.error(err);
            });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            TERRAIN_SETTINGS.windowProvider.fetch(
                new URL(TERRAIN_SETTINGS.stylePath, new URL(import.meta.url)).href
            ),
            TERRAIN_SETTINGS.windowProvider.fetch(
                new URL(TERRAIN_SETTINGS.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = TERRAIN_SETTINGS.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    set value(value) {
        this.#template.then(_ => {
            this.#mazeTypesToggle.classList.toggle("hidden", ![MazeGenerators.SideWinder, MazeGenerators.Binary, MazeGenerators.AldousBroder].includes(value));
            this.#mazeTypesToggle.value = value;
            this.#terrainTypesToggle.value = value;
        });
    }

    get value() {
        let terrainType = this.#terrainTypesToggle.value;
        if (terrainType === "maze") {
           terrainType = this.#mazeTypesToggle.value;
        }
        return terrainType;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
