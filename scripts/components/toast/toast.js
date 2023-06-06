export class TOAST extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });


        this.#shadow.addEventListener('click', this.#onClickHandler);

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = TOAST.documentProvider.createElement("template");
                template.innerHTML = TOAST.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));

            })
            .catch((err) => {
            TOAST.logger.error(err);
        });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            TOAST.windowProvider.fetch(
                new URL(TOAST.stylePath, new URL(import.meta.url)).href
            ),
            TOAST.windowProvider.fetch(
                new URL(TOAST.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = TOAST.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    #onClickHandler = (event) => {
        if (event.target === this.#shadow.querySelector("button")) {
            this.#shadow.querySelector(".toast").classList.toggle("show")
        }
    }

    /**
     * @param {any} value - color in hsla format
     */
    set value(value) {
        if (!this.isConnected) {
            this.#pendingData = value;
            return;
        }

        this.#shadow.querySelector(".toast--message").textContent = value;

        this.#shadow.querySelector(".toast").classList.toggle("show")

        TOAST.windowProvider.setTimeout(() => {
            this.#shadow.querySelector(".toast").classList.toggle("show")
        }, 3000);
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
