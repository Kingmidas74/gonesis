import {Configuration} from "../application/application.js";

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
     * @type {GameController} The game controller.
     * @private
     */
    #gameController

    /**
     * @type {Object} The elements of the UI.
     * @private
     */
    #elements

    OnWindowResizeListener = null;
    #resizeTimeout = null;

    OnSettingsUpdateListener = null;
    #settingsUpdateTimeout = null;

    /**
     * Creates a new instance of this class.
     * @param {Window} window - The window.
     * @param {Document} document - The document.*
     * @param {GameController} gameController - The game controller.
     */
    constructor(window, document, gameController) {
        this.#window = window;
        this.#document = document;
        this.#gameController = gameController;

        this.#elements = {
            nextStepBtn: document.getElementById("nextStepBtn"),
            playBtn: document.getElementById("playBtn"),
            pauseBtn: document.getElementById("pauseBtn"),
            generateBtn: document.getElementById("generateBtn"),
            settingsBtn: document.getElementById("settingsBtn"),
            sideBar: document.getElementById("settings"),
            saveBtn: document.getElementById("saveBtn"),
        };
    }

    init() {
        this.#elements.nextStepBtn.addEventListener("click", () => this.#gameController.nextStep());

        this.#elements.playBtn.addEventListener("click", () => {
            const isPlaying = this.#gameController.playGame();
            this.togglePlayPause(isPlaying);
        });

        this.#elements.pauseBtn.addEventListener("click", () => {
            const isPlaying = this.#gameController.pauseGame();
            this.togglePlayPause(isPlaying);
        });

        this.#elements.generateBtn.addEventListener("click", () => this.#gameController.generateGame());

        this.#elements.saveBtn.addEventListener("click", this.#handleSettingsUpdate);

        this.#elements.settingsBtn.addEventListener("click", () => {
            this.#elements.sideBar.classList.toggle("active");
        });

        this.#elements.sideBar.addEventListener('click', this.#handleTabClick);
        this.#window.addEventListener('resize', this.#handleResize);
        this.setFirstTabsActive();
    }

    togglePlayPause(isPlaying) {
        this.#elements.nextStepBtn.disabled = isPlaying;
        this.#elements.generateBtn.disabled = isPlaying;
        this.#elements.playBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#elements.pauseBtn.parentElement.classList.toggle("hidden", !isPlaying);
    }

    /**
     * Get all settings from the UI.
     * @returns {Configuration} The settings.
     */
    collectAllSettings() {
        return new Configuration({
            cellSize: this.#document.getElementById('cellSize').value,
        })
    }

    #handleSettingsUpdate = () => {
        this.#window.clearTimeout(this.#settingsUpdateTimeout);
        this.#settingsUpdateTimeout = this.#window.setTimeout(() => {
            const settings = this.collectAllSettings();
            this.OnSettingsUpdateListener?.(settings);
        }, 250);
    }

    #handleResize = () => {
        this.#window.clearTimeout(this.#resizeTimeout);
        this.#resizeTimeout = this.#window.setTimeout(() => {
            this.OnWindowResizeListener?.();
        }, 250);
    }

    #handleTabClick = (e) => {
        const clickedTab = e.target.closest('[data-target]');
        if (!clickedTab) return;

        const container = clickedTab.parentNode;
        const sidebar = container.closest('.settings');

        Array.from(container.children).forEach(tab => {
            tab.classList.remove('active');
        });

        clickedTab.classList.add('active');

        const parentFieldset = container.closest('.form__fieldset');
        if (parentFieldset) {
            const allNestedFieldsets = parentFieldset.querySelectorAll('.form__fieldset');
            allNestedFieldsets.forEach(fieldset => fieldset.classList.remove('active'));
        }

        const form = sidebar.querySelector('.settings--form--content');
        const allSiblings = Array.from(form.children).filter(child => child !== parentFieldset);
        allSiblings.forEach(sibling => sibling.classList.remove('active'));

        const targetFieldset = form.querySelector(`#${clickedTab.getAttribute('data-target')}`);
        if (targetFieldset) targetFieldset.classList.add('active');
    }

    setFirstTabsActive() {
        const topLevelTabsContainer = this.#document.querySelector('.settings-header .tab-container');

        const firstTopLevelTab = topLevelTabsContainer?.querySelector('.tab');
        this.#activateTab(firstTopLevelTab);

        const nestedTabGroups = this.#elements.sideBar.querySelectorAll('.form__fieldset[data-tab="true"]');
        nestedTabGroups.forEach(tabGroup => {
            const firstNestedTab = tabGroup.querySelector('.tab');
            this.#activateTab(firstNestedTab);
        });
    }

    #activateTab(tab) {
        if (!tab) {
            return;
        }

        const targetFieldset = this.#elements.sideBar.querySelector(`#${tab.getAttribute('data-target')}`);
        if (!targetFieldset) {
            return;
        }

        tab.classList.add('active');
        targetFieldset.classList.add('active');
    }
}
