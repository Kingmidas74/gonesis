export class BRAIN_VIEW extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = BRAIN_VIEW.documentProvider.createElement("template");
                template.innerHTML = BRAIN_VIEW.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .catch((err) => {
                BRAIN_VIEW.logger.error(err);
            });
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
