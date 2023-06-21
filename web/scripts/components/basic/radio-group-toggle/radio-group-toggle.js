export class RadioGroupToggleItemData {
    /**
     *
     * @param {String} name The item name.
     * @param {String} value The item value.
     */
    constructor(name, value) {
        this.name = name;
        this.value = value;
    }
}

export class RADIO_GROUP_TOGGLE extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #data = [];

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });
        this.#shadow.addEventListener('change', this.#changeEventHandler);

        this.#template = this.initializeTemplateParser().catch((err) => {
            RADIO_GROUP_TOGGLE.logger.error(err);
        });
    }

    static get observedAttributes() { return ['title']; }
    attributeChangedCallback(attr, oldVal, newVal) {
        if (oldVal === newVal) return; // nothing to do
        switch (attr) {
            case 'title':
                break;
        }
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            RADIO_GROUP_TOGGLE.windowProvider.fetch(
                new URL(RADIO_GROUP_TOGGLE.stylePath, new URL(import.meta.url)).href
            ),
            RADIO_GROUP_TOGGLE.windowProvider.fetch(
                new URL(RADIO_GROUP_TOGGLE.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = RADIO_GROUP_TOGGLE.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    /**
     * @param {Array<RadioGroupToggleItemData>} data
     */
    set data(data) {
        this.#data = data;
        if (!this.isConnected) {
            console.log('not connected', data)
            this.#pendingData = data;
            return;
        }

        this.#template
            .then((templateContent) => {
                const template = RADIO_GROUP_TOGGLE.documentProvider.createElement("template");
                template.innerHTML = RADIO_GROUP_TOGGLE.templateParser?.parse(templateContent, {
                    title: this.getAttribute('data-title'),
                    items: data,
                });
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .catch((err) => {
                RADIO_GROUP_TOGGLE.logger.error(err);
            });
    }

    #changeEventHandler = (e) => {
        if (e.target.type !== 'radio') {
            return;
        }

        this.dispatchEvent(new RADIO_GROUP_TOGGLE.windowProvider.CustomEvent('change', {
            detail: { value: e.target.value }
        }))
    }

    set value(value) {
        this.#template.then(_ => {
            const target = this.#shadow.querySelector(`input[value="${value}"]`);
            if (!target) {
                return;
            }
            target.checked = true
        });
    }

    get value() {
        return this.#shadow.querySelector(`input:checked`)?.value;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }
}
