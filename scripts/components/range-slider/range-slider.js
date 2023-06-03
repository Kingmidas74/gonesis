export class RANGE_SLIDER extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser().catch((err) => {
            RANGE_SLIDER.logger.error(err);
        });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            RANGE_SLIDER.windowProvider.fetch(
                new URL(RANGE_SLIDER.stylePath, new URL(import.meta.url)).href
            ),
            RANGE_SLIDER.windowProvider.fetch(
                new URL(RANGE_SLIDER.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = RANGE_SLIDER.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
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
                const template = RANGE_SLIDER.documentProvider.createElement("template");
                template.innerHTML = RANGE_SLIDER.templateParser?.parse(templateContent, {
                    min: value?.min || 0,
                    max: value?.max || 1000,
                    value: value?.value || 0,
                    title: value?.title || 'Value'
                });
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#shadow.querySelector("input").value = value?.value || 0;
                this.#shadow.querySelector(".value").textContent = value?.value || 0;

                this.#shadow.addEventListener('change', (e) => {
                    if (e.target.type !== 'range') {
                        return;
                    }

                    this.#shadow.querySelector(".value").textContent = e.target.value;

                    this.dispatchEvent(new RANGE_SLIDER.windowProvider.CustomEvent('change', {
                        detail: { value: e.target.value }
                    }))
                });

            })
            .catch((err) => {
                RANGE_SLIDER.logger.error(err);
            });
    }

    get value() {
        return this.#shadow.querySelector("input")?.value
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
