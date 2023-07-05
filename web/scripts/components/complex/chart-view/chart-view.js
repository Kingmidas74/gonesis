import {AgentType} from "../../../domain/enum.js";

export class CHART_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;


    #myChart;
    #datasets = []

    #config

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then((templateContent) => {
                const template = CHART_VIEW.documentProvider.createElement("template");
                template.innerHTML = CHART_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .then(this.#setup)
            .catch((err) => {
                CHART_VIEW.logger.error(err);
            });
    }

    async #initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            CHART_VIEW.windowProvider.fetch(
                new URL(CHART_VIEW.stylePath, new URL(import.meta.url)).href
            ),
            CHART_VIEW.windowProvider.fetch(
                new URL(CHART_VIEW.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = CHART_VIEW.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    #setup = (templateContent) => {
        return new Promise((resolve, reject) => {
            CHART_VIEW.windowProvider.setTimeout(() => {
                this.#datasets = Object.entries(AgentType)
                    .map(([name, value]) => ({
                        label: value,
                        data: [],
                        fill: false,
                        tension: 0.1
                    }));

                let ctx = this.shadowRoot.querySelector('canvas').getContext('2d');
                if(!CHART_VIEW.windowProvider.Chart) {
                    return
                }
                this.#myChart = new Chart(ctx, {
                    type: 'line',
                    data: {
                        labels: [],
                        datasets: this.#datasets
                    },
                    options: {

                        responsive:true,
                        interaction: {
                            intersect: false,
                        },
                        plugins: {
                            zoom: {
                                zoom: {
                                    wheel: {
                                        enabled: false,
                                    },
                                    pinch: {
                                        enabled: true
                                    },
                                    mode: 'x',
                                }
                            }
                        }
                    }
                });
                resolve();
            }, 1)
        });
    }

    /**
     * @param {World} worldInstance
     */
    updateData(worldInstance) {
        if(!CHART_VIEW.windowProvider.Chart) {
            return
        }

        if(worldInstance.currentDay % 1 !== 0) {
            return;
        }
        let newCounts = worldInstance.cells.filter(c => c.agent).map(c => c.agent).reduce((counts, agent) => {
            if (!counts[agent.agentType]) {
                counts[agent.agentType] = 0;
            }
            counts[agent.agentType]++;
            return counts;
        }, {});

        // Update the agentCounts and datasets
        for (let agentType in newCounts) {
            const currentSet = this.#datasets.filter(d => d.label === agentType)[0];
            const currentColor = this.#getAgentTypeColor(agentType);
            currentSet.data.push(newCounts[agentType]);
            currentSet.borderColor = currentColor;
            currentSet.backgroundColor = currentColor;
        }

        this.#myChart.data.labels.push(worldInstance.currentDay);
        this.#myChart?.update();
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData
            this.#pendingData = null;
        }
    }

    /**
    * Set config data
    * @param {Configuration} config
    */
    set config(config)
    {
        if (!this.isConnected) {
            this.#pendingData = config;
            return
        }

        this.#config = config;
    }

    clean() {
        for(let i = 0; i < this.#datasets.length; i++) {
            this.#datasets[i].data = [];
        }
        this.#myChart?.update();
    }

    #getAgentTypeColor = (agentType) => {
        switch (agentType) {
            case AgentType.PLANT:
                return this.#config.PlantConfiguration.Color;
            case AgentType.HERBIVORE:
                return this.#config.HerbivoreConfiguration.Color;
            case AgentType.CARNIVORE:
                return this.#config.CarnivoreConfiguration.Color;
            case AgentType.OMNIVORE:
                return this.#config.OmnivoreConfiguration.Color;
            default:
                return "rgb(75, 192, 192)";
        }
    }

}
