import {Configuration, Topologies} from "../application/configuration/configuration.js";

export default class UIController {

    /**
     * @type {Window} The window.
     * @private
     */
    #window

    /**
     * @type {Document} The document.
     * @private
     */
    #document

    /**
     * @type {Object} The elements of the UI.
     * @private
     */
    #elements

    #onWindowResizeListener = [];
    #windowResizeEventTimeout = null;

    addOnWindowResizeListener = (listener) => {
        this.#onWindowResizeListener.push(listener);
    }

    removeOnWindowResizeListener = (listener) => {
        this.#onWindowResizeListener = this.#onWindowResizeListener.filter(l => l !== listener);
    }

    #raiseOnWindowResizeEvent = (...args) => {
        this.#onWindowResizeListener.forEach(listener => listener(...args));
    }


    #onSettingsUpdateListener = [];
    #settingsUpdateEventTimeout = null;

    addOnSettingsUpdateListener = (listener) => {
        this.#onSettingsUpdateListener.push(listener);
    }

    removeOnSettingsUpdateListener = (listener) => {
        this.#onSettingsUpdateListener = this.#onSettingsUpdateListener.filter(l => l !== listener);
    }

    #raiseOnSettingsUpdateEvent = (...args) => {
        this.#onSettingsUpdateListener.forEach(listener => listener(...args));
    }

    #onPlayButtonClickEventHandlers = [];

    addOnPlayButtonClickEventListener = (listener) => {
        this.#onPlayButtonClickEventHandlers.push(listener);
    }

    removeOnPlayButtonClickEventListener = (listener) => {
        this.#onPlayButtonClickEventHandlers = this.#onPlayButtonClickEventHandlers.filter(l => l !== listener);
    }

    #raiseOnPlayButtonClickEvent = (...args) => {
        this.#onPlayButtonClickEventHandlers.forEach(listener => listener(...args));
    }

    #onPauseButtonClickEventHandlers = [];

    addOnPauseButtonClickEventListener = (listener) => {
        this.#onPauseButtonClickEventHandlers.push(listener);
    }

    removeOnPauseButtonClickEventListener = (listener) => {
        this.#onPauseButtonClickEventHandlers = this.#onPauseButtonClickEventHandlers.filter(l => l !== listener);
    }

    #raiseOnPauseButtonClickEvent = (...args) => {
        this.#onPauseButtonClickEventHandlers.forEach(listener => listener(...args));
    }

    #onGenerateButtonClickEventHandlers = [];

    addOnGenerateButtonClickEventListener = (listener) => {
        this.#onGenerateButtonClickEventHandlers.push(listener);
    }

    removeOnGenerateButtonClickEventListener = (listener) => {
        this.#onGenerateButtonClickEventHandlers = this.#onGenerateButtonClickEventHandlers.filter(l => l !== listener);
    }

    #raiseOnGenerateButtonClickEvent = (...args) => {
        this.#onGenerateButtonClickEventHandlers.forEach(listener => listener(...args));
    }

    #onNextStepButtonClickEventHandlers = [];

    addOnNextStepButtonClickEventListener = (listener) => {
        this.#onNextStepButtonClickEventHandlers.push(listener);
    }

    removeOnNextStepButtonClickEventListener = (listener) => {
        this.#onNextStepButtonClickEventHandlers = this.#onNextStepButtonClickEventHandlers.filter(l => l !== listener);
    }

    #raiseOnNextStepButtonClickEvent = (...args) => {
        this.#onNextStepButtonClickEventHandlers.forEach(listener => listener(...args));
    }

    /**
     * Creates a new instance of this class.
     * @param {Window} window - The window.
     * @param {Document} document - The document.*
     */
    constructor(window, document) {
        this.#window = window;
        this.#document = document;

        this.#elements = {
            primaryToolbar: document.getElementById("primaryToolbar"),
            gameSettings: document.getElementById("settings"),
            settings: document.getElementById("settings-container"),
            toast: document.getElementById('toast'),
        };
    }

    /**
     * Game configuration.
     * @param {Configuration} configuration
     */
    init(configuration) {
        this.#elements.primaryToolbar.addEventListener("play", this.#raiseOnPlayButtonClickEvent);
        this.#elements.primaryToolbar.addEventListener("pause", this.#raiseOnPauseButtonClickEvent);
        this.#elements.primaryToolbar.addEventListener("generate", this.#raiseOnGenerateButtonClickEvent);
        this.#elements.primaryToolbar.addEventListener("nextStep", this.#raiseOnNextStepButtonClickEvent);
        this.#elements.primaryToolbar.addEventListener("settings", () => {
            this.#toggleSideBar(!this.#elements.settings.classList.contains("active"));
        });
        this.#elements.gameSettings.addEventListener("update", this.#handleSettingsUpdate);

        this.#window.addEventListener('resize', this.#handleResize);

        this.#document.addEventListener('click', (event) => {
            const isClickOutside = !this.#elements.settings.contains(event.target);
            if (isClickOutside) {
                this.#elements.settings.classList.remove('active');
                this.#toggleSideBar(false)
            }
        });

        this.#elements.gameSettings.config = configuration;
    }

    togglePlayPause(isPlaying) {
        this.#elements.primaryToolbar.togglePlayPause(isPlaying);
    }

    #handleSettingsUpdate = (e) => {
        this.#window.clearTimeout(this.#settingsUpdateEventTimeout);
        this.#settingsUpdateEventTimeout = this.#window.setTimeout(() => {
            this.#raiseOnSettingsUpdateEvent(e.detail.value);
        }, 250);
    }

    #handleResize = () => {
        this.#window.clearTimeout(this.#windowResizeEventTimeout);
        this.#windowResizeEventTimeout = this.#window.setTimeout(() => {
            this.#raiseOnWindowResizeEvent();
        }, 250);
    }

    /**
     * Shows a toast message and hides it automatically after 3 seconds.
     * @param {string} message The message to show in the toast.
     */
    showToast(message) {
        this.#elements.toast.value = message;
    }

    #toggleSideBar = (isOpen) => {
        this.#elements.settings.classList.toggle('active', isOpen);
        this.#elements.primaryToolbar.toggleActions(isOpen);
    }
}
