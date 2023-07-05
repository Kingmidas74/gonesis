import {MazeGenerators, TerrainRatio, TerrainCellSizes} from "../../../configuration/configuration.js";

export class TERRAIN_SETTINGS extends HTMLElement {

    #shadow;

    #template;

    #elements = {}


    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = TERRAIN_SETTINGS.documentProvider.createElement("template");
                template.innerHTML = TERRAIN_SETTINGS.templateParser?.parse(templateContent, {});
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#elements = {
                    terrainTypesToggle: this.#shadow.querySelector("#terrain-types"),
                    mazeTypesToggle: this.#shadow.querySelector('#maze-types'),
                    terrainRatioToggle: this.#shadow.querySelector('#terrain-ratio'),
                }

                const terrainTypes = Object.entries(MazeGenerators).map(([name, value]) => ({ name, value })).filter(({ value }) => [MazeGenerators.Border, MazeGenerators.Empty, MazeGenerators.Grid].includes(value));
                terrainTypes.push({ name: "Maze", value: "maze" })
                const mazeTypes = Object.entries(MazeGenerators).map(([name, value]) => ({ name, value })).filter(({ value }) => [MazeGenerators.AldousBroder, MazeGenerators.Binary, MazeGenerators.SideWinder].includes(value));
                this.#elements.terrainTypesToggle.addEventListener("change", (event) => {
                    this.#elements.mazeTypesToggle.classList.toggle("hidden", event?.detail?.value !== "maze");
                });
                this.#elements.terrainTypesToggle.data = terrainTypes;
                this.#elements.mazeTypesToggle.data = mazeTypes;
                this.#elements.mazeTypesToggle.value = MazeGenerators.AldousBroder;

                this.#elements.terrainRatioToggle.data = Object.entries(TerrainCellSizes).map(([name, value]) => ({
                        name: name,
                        value: value
                }));

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

    set mazeType(value) {
        this.#template.then(_ => {
            const isMaze = [MazeGenerators.SideWinder, MazeGenerators.Binary, MazeGenerators.AldousBroder].includes(value);
            this.#elements.mazeTypesToggle.classList.toggle("hidden", !isMaze);
            this.#elements.mazeTypesToggle.value = value;
            this.#elements.terrainTypesToggle.value = isMaze ? "maze" : value;
        });
    }

    get mazeType() {
        let terrainType = this.#elements.terrainTypesToggle.value;
        if (terrainType === "maze") {
            terrainType = this.#elements.mazeTypesToggle.value;
        }
        return terrainType;
    }

    set size(value) {
        this.#template.then(_ => {
            this.#elements.terrainRatioToggle.value = value;
        });
    }

    get size() {
        return this.#elements.terrainRatioToggle.value;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
