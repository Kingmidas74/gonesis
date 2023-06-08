export class TAB_LAYOUT extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #elements = {}

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then((templateContent) => {
                const template = TAB_LAYOUT.documentProvider.createElement("template");
                template.innerHTML = TAB_LAYOUT.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .then(this.#setup)
            .catch((err) => {
                TAB_LAYOUT.logger.error(err);
            });
    }

    #setup = async () => {
        this.#shadow.querySelector('footer').addEventListener('click', (event) => {
            const clickedBtn = event.target.closest('[data-target]');
            if (!clickedBtn) return;

            const menu = clickedBtn.closest('menu');

            Array.from(menu.children).forEach(tab => {
                tab.classList.remove('active');
            });

            clickedBtn.parentElement.classList.add('active');

            Array.from(this.#shadow.querySelectorAll("main > section")).forEach(section => section.classList.remove('active'))
            this.#shadow.querySelector(`#${clickedBtn.getAttribute('data-target')}`)?.classList.add('active');

            this.dispatchEvent(new TAB_LAYOUT.windowProvider.CustomEvent('change', {
                detail: { value: clickedBtn.getAttribute('data-target') }
            }))
        })
    }

    #initializeTemplateParser = async ()=> {
        const [cssResponse, htmlResponse] = await Promise.all([
            TAB_LAYOUT.windowProvider.fetch(
                new URL(TAB_LAYOUT.stylePath, new URL(import.meta.url)).href
            ),
            TAB_LAYOUT.windowProvider.fetch(
                new URL(TAB_LAYOUT.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = TAB_LAYOUT.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

}
