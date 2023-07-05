export class SEED_GENERATOR extends HTMLElement {

    #shadow;

    #template;
    #templateContent = "";

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#shadow.addEventListener("click", this.#clickHandler);

        this.#template = this.#initializeTemplateParser()
            .then((_) => {
                this.#render();
            })
        .catch((err) => {
            SEED_GENERATOR.logger.error(err);
        });
    }

    async #initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            SEED_GENERATOR.windowProvider.fetch(
                new URL(SEED_GENERATOR.stylePath, new URL(import.meta.url)).href
            ),
            SEED_GENERATOR.windowProvider.fetch(
                new URL(SEED_GENERATOR.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = SEED_GENERATOR.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        this.#templateContent = templateContent;
        return templateContent;
    }

    /**
     * @param {any} value - color in hsla format
     */
    set value(value) {
        if (!this.isConnected) {
            this.#pendingData = value;
            return;
        }

        this.#template
            .then((templateContent) => {
                this.#render(value);
            })
            .catch((err) => {
                SEED_GENERATOR.logger.error(err);
            });
    }

    get value() {
        if(this.#shadow.querySelector('input[type="checkbox"]').checked) {
            return this.#shadow.querySelector('input[type="text"]')?.value;
        }
        return (Date.now() * 1000000).toString();
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    #clickHandler = (e) => {
        if(e?.target?.type === 'checkbox') {
            this.#shadow.querySelector('input[type="text"]').disabled = !e.target.checked;
            this.#shadow.querySelector('button').disabled = !e.target.checked;
            return
        }
        if(e?.target?.closest('button')) {
            this.#shadow.querySelector('input[type="text"]').value = (SEED_GENERATOR.windowProvider.Date.now() * 1000000).toString();
        }
    }

    #render = (data = {}) => {
        const template = SEED_GENERATOR.documentProvider.createElement("template");
        template.innerHTML = SEED_GENERATOR.templateParser?.parse(this.#templateContent, {
            value: data?.value || SEED_GENERATOR.windowProvider.Date.now() * 1000000,
        });
        this.#shadow.appendChild(template.content.cloneNode(true));
    }
}
