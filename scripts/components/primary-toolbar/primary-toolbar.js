export class PRIMARY_TOOLBAR extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #controls = {}

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });
        this.#shadow.addEventListener('click', this.#clickEventHandler);

        this.#template = this.initializeTemplateParser()
            .then((templateContent) => {
                const template = PRIMARY_TOOLBAR.documentProvider.createElement("template");
                template.innerHTML = PRIMARY_TOOLBAR.templateParser?.parse(templateContent);
                this.#shadow.appendChild(template.content.cloneNode(true));

                this.#controls = {
                    nextStepBtn: this.#shadow.getElementById("nextStepBtn"),
                    playBtn: this.#shadow.getElementById("playBtn"),
                    pauseBtn: this.#shadow.getElementById("pauseBtn"),
                    generateBtn: this.#shadow.getElementById("generateBtn"),
                    settingsBtn: this.#shadow.getElementById("settingsBtn")
                }
            })
            .catch((err) => {
                PRIMARY_TOOLBAR.logger.error(err);
            });
    }

    async initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            PRIMARY_TOOLBAR.windowProvider.fetch(
                new URL(PRIMARY_TOOLBAR.stylePath, new URL(import.meta.url)).href
            ),
            PRIMARY_TOOLBAR.windowProvider.fetch(
                new URL(PRIMARY_TOOLBAR.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = PRIMARY_TOOLBAR.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    #clickEventHandler = (e) => {
        const clickedBtn = e.target.closest('button');
        if (!clickedBtn) return;
        e.preventDefault();
        e.stopPropagation();
        let customEvent;
        switch (clickedBtn.id) {
            case this.#controls?.nextStepBtn?.id:
                customEvent = new PRIMARY_TOOLBAR.windowProvider.CustomEvent('nextStep');
                break;
            case this.#controls?.playBtn?.id:
                customEvent = new PRIMARY_TOOLBAR.windowProvider.CustomEvent('play');
                break;
            case this.#controls?.pauseBtn?.id:
                customEvent = new PRIMARY_TOOLBAR.windowProvider.CustomEvent('pause');
                break;
            case this.#controls?.generateBtn?.id:
                customEvent = new PRIMARY_TOOLBAR.windowProvider.CustomEvent('generate');
                break;
            case this.#controls?.settingsBtn?.id:
                customEvent = new PRIMARY_TOOLBAR.windowProvider.CustomEvent('settings');
                break;
            default:
                return;
        }

        this.dispatchEvent(customEvent);
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

    togglePlayPause(isPlaying) {
        this.#controls.settingsBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#controls.generateBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#controls.nextStepBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#controls.playBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#controls.pauseBtn.parentElement.classList.toggle("hidden", !isPlaying);
    }

    toggleActions(hide) {
        this.#controls.generateBtn.disabled = hide;
        this.#controls.nextStepBtn.disabled = hide;
        this.#controls.playBtn.disabled = hide;
        this.#controls.pauseBtn.disabled = hide;
    }

}
