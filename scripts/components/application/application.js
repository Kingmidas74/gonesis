import {ApplicationGameService} from "./game.service.js";
import {Engine} from "../../services/data-access/wasm/engine.js";

export class APPLICATION extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #controls = {};

    #applicationGameService;

    constructor() {
        super();

        this.#applicationGameService = new ApplicationGameService(APPLICATION.windowProvider, APPLICATION.windowProvider.performance);

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then(this.#render)
            .then(this.#setup)
            .then(this.#generateGame)
            .catch(APPLICATION.logger.error);
    }

    async #initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            APPLICATION.windowProvider.fetch(
                new URL(APPLICATION.stylePath, new URL(import.meta.url)).href
            ),
            APPLICATION.windowProvider.fetch(
                new URL(APPLICATION.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = APPLICATION.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    #render = (templateContent) => {
        const template = APPLICATION.documentProvider.createElement("template");
        template.innerHTML = APPLICATION.templateParser?.parse(templateContent, {
            title: this.getAttribute('data-title'),
            brainSlotAvailable: false,
            chartSlotAvailable: !!APPLICATION.windowProvider.Chart,
            gameSlotAvailable: true,
        });
        this.#shadow.appendChild(template.content.cloneNode(true));
    }

    #setup = async () => {
        this.#controls = {
            primaryToolbar: this.#shadow.querySelector('app-primary-toolbar'),
            gameView: this.#shadow.querySelector('app-game-view'),
            brainView: this.#shadow.querySelector('app-brain-view'),
            chartView: this.#shadow.querySelector('app-chart-view'),
            gameSettings: this.#shadow.querySelector('app-game-settings'),
            tabLayout: this.#shadow.querySelector('app-tab-layout'),
            aside: this.#shadow.querySelector('aside'),
            toast: this.#shadow.querySelector('app-toast'),
        };

        this.#controls.primaryToolbar.addEventListener("play", async () => {
            (await this.#applicationGameService.playGame())
                .map(_ => {
                    this.#togglePlayPause(true);
                })
                .orElse(err=> {
                    this.#showToast(err.message);
                })
        });

        this.#controls.primaryToolbar.addEventListener("pause", (e) => {
            this.#applicationGameService.pauseGame();
            this.#togglePlayPause(false);
        });

        this.#controls.primaryToolbar.addEventListener('generate', this.#generateGame);

        this.#controls.primaryToolbar.addEventListener("nextStep", async () => {
            (await this.#applicationGameService.nextStep())
                .map(worldInstance => {
                    this.#controls.gameView.update(worldInstance);
                })
                .orElse(err=> {
                    this.#showToast(err.message);
                })
        });

        this.#controls.primaryToolbar.addEventListener("settings", () => {
            this.#toggleSideBar(!this.#controls.aside.classList.contains("active"));
        });

        this.#controls.tabLayout.addEventListener("change", (e) => {

        })

        this.#controls.gameSettings.addEventListener("update", this.#handleSettingsUpdate);

        this.#shadow.addEventListener('click', this.#onClickHandler);

        this.#applicationGameService.addEventListener('update', (e) => {
            this.#controls.gameView.update(e.detail.value);
            this.#controls.chartView.update(e.detail.value);
        });

        this.#applicationGameService.addEventListener('finish', (e) => {
            this.#togglePlayPause(false);
            this.#controls.gameView.update(e.detail.value);
        });

        this.#applicationGameService.addEventListener('error', (e) => {
            this.#togglePlayPause(false);
            this.#showToast(e.detail.value.message);
        })

        this.#controls.gameView.addEventListener('click', (e) => {
            const cellX = APPLICATION.windowProvider.Math.floor(e.detail.x / this.#applicationGameService.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize);
            const cellY = APPLICATION.windowProvider.Math.floor(e.detail.y / this.#applicationGameService.configurationProvider.getInstance().WorldConfiguration.Ratio.CellSize);


            const clickedCell = this.#applicationGameService.cell(cellX,cellY);
            if(clickedCell) {
                console.log(clickedCell);
            } else {
                console.log('Clicked outside of cells area');
            }
        })

        const dataClient = new Engine("engine.wasm", APPLICATION.windowProvider, APPLICATION.windowProvider.JSON);
        await dataClient.init()

        this.#applicationGameService.configure(dataClient);

        this.#controls.gameView.configurationProvider = this.#applicationGameService.configurationProvider;
        this.#controls.chartView.configurationProvider = this.#applicationGameService.configurationProvider;
        this.#controls.gameSettings.configuration = this.#applicationGameService.configurationProvider.getInstance();
    }

    #generateGame = async () => {
        return new Promise((resolve, reject) => {
            APPLICATION.windowProvider.setTimeout(async () => {
                const config = this.#controls.gameSettings.configuration;
                if(!config) {
                    return;
                }

                this.#applicationGameService.configurationProvider.updateConfiguration(config);

                const width = this.#controls.gameView.width;
                const height = this.#controls.gameView.height;
                (await this.#applicationGameService.generateGame(width, height))
                    .map(worldInstance => {
                        this.#controls.gameView.update(worldInstance);
                    })
                    .orElse(err=> {
                        this.#showToast(err.message);
                    })
                resolve();
            }, 500);
        });
    }

    #onClickHandler = (e) => {
        const isClickOutsideAside = !this.#controls.aside.contains(e.target);
        if (isClickOutsideAside) {
            this.#toggleSideBar(false);
        }
    }

    #togglePlayPause = (isPlaying) => {
        this.#controls.primaryToolbar.togglePlayPause(isPlaying);
    }

    #handleSettingsUpdate = async (e) => {
        this.#toggleSideBar(false);
        await this.#generateGame();
    }

    /**
     * Shows a toast message and hides it automatically after 3 seconds.
     * @param {string} message The message to show in the toast.
     */
    #showToast(message) {
        this.#controls.toast.value = message;
    }

    #toggleSideBar = (isOpen) => {
        this.#controls.aside.classList.toggle('active', isOpen);
        this.#controls.primaryToolbar.toggleActions(isOpen);
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    disconnectedCallback() {
        this.#shadow.removeEventListener('click', this.#onClickHandler);
    }

}
