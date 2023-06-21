import {Engine} from "../../../application/engine/engine.js";
import {AgentConfiguration, WorldConfiguration, Configuration} from "../../../application/configuration/configuration.js";
import {Application} from "../../../application/application.js";

export class BRAIN_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #application;

    #isGenerated = false;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = BRAIN_VIEW.documentProvider.createElement("template");
                template.innerHTML = BRAIN_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .then(this.#setup)
            .then(this.generateGame)
            .then(this.#showBrain)
            .catch((err) => {
                BRAIN_VIEW.logger.error(err);
            });
    }

    #setActiveCommand = (el) => {
        Array.from(el.parentElement.children).forEach((c) => {
            c.classList.remove('active');
        })
        el.classList.add('active');
    }


    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            BRAIN_VIEW.windowProvider.fetch(
                new URL(BRAIN_VIEW.stylePath, new URL(import.meta.url)).href
            ),
            BRAIN_VIEW.windowProvider.fetch(
                new URL(BRAIN_VIEW.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = BRAIN_VIEW.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
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
                return this.#application.game.agents()[0]
            })
    }

    #setup = () => {
        return new Promise((resolve, reject) => {
            BRAIN_VIEW.windowProvider.setTimeout(async () => {
                this.#application = new Application();
                await this.#application.configure(window, document, this.#shadow.querySelector("canvas"), "agent.wasm")
                this.#application.configurationProvider.updateConfiguration(new Configuration({
                    worldConfiguration: {
                        CellSize: 20,
                    },

                    plantConfiguration: {
                        InitialCount:0
                    },
                    herbivoreConfiguration: {
                        InitialCount:0
                    },
                    carnivoreConfiguration: {
                        InitialCount: 0
                    },
                    omnivoreConfiguration: {
                        InitialCount:1
                    }
                }))
                resolve();
            }, 10)
        })
    }

    #showBrain = (eitherAgent) => {
        eitherAgent.map(agent => {
            let matrixContainer = this.#shadow.getElementById('matrix');
            matrixContainer.style.setProperty('--N', Math.sqrt(this.#application.configurationProvider.getInstance().OmnivoreConfiguration.BrainVolume));
            matrixContainer.addEventListener('click', (e) => {
                if (e.target.classList.contains('square')) {
                    this.#setActiveCommand(e.target)
                }
            });
            agent.commands.forEach(item => {
                let newDiv = BRAIN_VIEW.documentProvider.createElement('li');
                newDiv.className = 'square';
                newDiv.setAttribute('data-command', item)
                newDiv.textContent = item.toString();
                matrixContainer.appendChild(newDiv);
            });
            this.#setActiveCommand(this.#shadow.querySelector('.square:nth-of-type(1)'))
        });
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

}
