import {Engine} from "../../../application/engine/engine.js";

export class BRAIN_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
            .then(async (templateContent) => {
                const template = BRAIN_VIEW.documentProvider.createElement("template");
                template.innerHTML = BRAIN_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));

             /*   const engine = new Engine("engine.wasm", BRAIN_VIEW.windowProvider);
                await engine.init()

                let matrixContainer = this.#shadow.getElementById('matrix');
                let N = 4;  // Count of divs per row and count of rows
                matrixContainer.style.setProperty('--N', N);
                matrixContainer.addEventListener('click', (e) => {
                    if(e.target.classList.contains('square')) {
                        this.#setActiveCommand(e.target)
                    }
                });
                for(let i=0; i < N*N; i++) {
                    let newDiv = BRAIN_VIEW.documentProvider.createElement('li');
                    newDiv.className = 'square';
                    newDiv.textContent = Math.floor(BRAIN_VIEW.windowProvider.Math.random()*100).toString();
                    matrixContainer.appendChild(newDiv);
                }

              */
            })
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

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

}
