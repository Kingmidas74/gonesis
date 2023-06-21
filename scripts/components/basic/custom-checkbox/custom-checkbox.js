export class CheckboxData {
    /**
     *
     * @param {String} name The item name.
     * @param {Boolean} value The item value.
     */
    constructor(name, value) {
        this.name = name;
        this.value = value;
    }
}

export class CUSTOM_CHECKBOX extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });
        this.#shadow.addEventListener('change', this.#changeEventHandler);

        this.#template = this.initializeTemplateParser().catch((err) => {
            CUSTOM_CHECKBOX.logger.error(err);
        });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            CUSTOM_CHECKBOX.windowProvider.fetch(
                new URL(CUSTOM_CHECKBOX.stylePath, new URL(import.meta.url)).href
            ),
            CUSTOM_CHECKBOX.windowProvider.fetch(
                new URL(CUSTOM_CHECKBOX.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = CUSTOM_CHECKBOX.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    #changeEventHandler = (e) => {
        if (e.target.type !== 'checkbox') {
            return;
        }

        this.dispatchEvent(new CUSTOM_CHECKBOX.windowProvider.CustomEvent('change', {
            detail: { value: e.target.checked }
        }))
    }

    set value(value) {
        this.#template
            .then((templateContent) => {
                const template = CUSTOM_CHECKBOX.documentProvider.createElement("template");
                template.innerHTML = CUSTOM_CHECKBOX.templateParser?.parse(templateContent, {
                    title: this.getAttribute('data-title'),
                    value: value,
                });
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .catch((err) => {
                CUSTOM_CHECKBOX.logger.error(err);
            });
    }

    get value() {
        return this.#shadow.querySelector(`input`)?.checked;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
