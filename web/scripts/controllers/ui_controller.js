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
            nextStepBtn: document.getElementById("nextStepBtn"),
            playBtn: document.getElementById("playBtn"),
            pauseBtn: document.getElementById("pauseBtn"),
            generateBtn: document.getElementById("generateBtn"),
            settingsBtn: document.getElementById("settingsBtn"),
            sideBar: document.getElementById("settings"),
            saveBtn: document.getElementById("saveBtn"),
            toast: document.getElementById('toast'),
            toastMessage: document.querySelector('.toast--message'),
            toastClose: document.querySelector('.toast--close'),
        };
    }

    /**
     * Game configuration.
     * @param {Configuration} configuration
     */
    init(configuration) {
        this.#elements.nextStepBtn.addEventListener("click", this.#raiseOnNextStepButtonClickEvent);
        this.#elements.playBtn.addEventListener("click", this.#raiseOnPlayButtonClickEvent);
        this.#elements.pauseBtn.addEventListener("click", this.#raiseOnPauseButtonClickEvent);
        this.#elements.generateBtn.addEventListener("click", this.#raiseOnGenerateButtonClickEvent);

        this.#elements.settingsBtn.addEventListener("click", () => {
            this.toggleSideBar(!this.#elements.sideBar.classList.contains("active"))
        });

        this.#elements.saveBtn.addEventListener("click", this.#handleSettingsUpdate);
        this.#elements.sideBar.addEventListener('click', this.#handleTabClick);
        this.#window.addEventListener('resize', this.#handleResize);
        this.setFirstTabsActive();

        this.#elements.sideBar.querySelectorAll('.range-slider__range').forEach(range => {
            range.addEventListener('input', this.#handleSettingsUpdate);
        })

        this.#document.addEventListener('click', (event) => {
            const isClickInside = this.#elements.sideBar.contains(event.target);
            const isSettingsBtn = this.#elements.settingsBtn.contains(event.target);

            if (!isClickInside && !isSettingsBtn) {
                this.#elements.sideBar.classList.remove('active');
                this.toggleSideBar(false)
            }
        });

        this.#elements.toastClose.addEventListener('click', () => {
            this.#elements.toast.classList.toggle('show');
        });

        this.#setupSettings(configuration);
    }

    togglePlayPause(isPlaying) {
        this.#elements.settingsBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#elements.generateBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#elements.nextStepBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#elements.playBtn.parentElement.classList.toggle("hidden", isPlaying);
        this.#elements.pauseBtn.parentElement.classList.toggle("hidden", !isPlaying);

        if(isPlaying && this.#elements.sideBar.classList.contains("active")) {
            this.#elements.sideBar.classList.remove("active");
        }
    }

    /**
     * Get all settings from the UI.
     * @returns {Configuration} The settings.
     */
    collectAllSettings() {
        return new Configuration({
            worldConfiguration: {
                CellSize: this.#document.getElementById('cellSize').value,
                MazeType: this.#document.querySelector('#terrain-settings').value,
                Topology: this.#document.querySelector('#topology-types').value,
            },
            plantConfiguration: this.#document.getElementById('plantSettings').config,
            carnivoreConfiguration: this.#document.getElementById('carnivoreSettings').config,
            herbivoreConfiguration: this.#document.getElementById('herbivoreSettings').config,
            omnivoreConfiguration: this.#document.getElementById('omnivoreSettings').config,
        })
    }

    #handleSettingsUpdate = (e) => {
        if(e.target.classList.contains('range-slider__range')) {
            e.target.parentNode.querySelector('.range-slider__value').innerHTML = e.target.value;
            return
        }

        this.#window.clearTimeout(this.#settingsUpdateEventTimeout);
        this.#settingsUpdateEventTimeout = this.#window.setTimeout(() => {
            const settings = this.collectAllSettings();
            this.#raiseOnSettingsUpdateEvent(settings);
        }, 250);
    }

    #handleResize = () => {
        this.#window.clearTimeout(this.#windowResizeEventTimeout);
        this.#windowResizeEventTimeout = this.#window.setTimeout(() => {
            this.#raiseOnWindowResizeEvent();
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

    /**
     * Shows a toast message and hides it automatically after 3 seconds.
     * @param {string} message The message to show in the toast.
     */
    showToast(message) {
        // Set the toast message
        this.#elements.toastMessage.textContent = message;

        this.#elements.toast.classList.toggle('show');

        this.#window.setTimeout(() => {
            this.#elements.toast.classList.toggle('show');
        }, 3000);
    }

    toggleSideBar = (isOpen) => {
        this.#elements.sideBar.classList.toggle('active', isOpen);
        this.#elements.generateBtn.disabled = isOpen;
        this.#elements.nextStepBtn.disabled = isOpen;
        this.#elements.playBtn.disabled = isOpen;
        this.#elements.pauseBtn.disabled = isOpen;
    }

    /**
     * Sets up the UI with the given settings.
     * @param {Configuration} settings
     */
    #setupSettings = (settings)=> {
        this.#document.getElementById('terrain-settings').value = settings.WorldConfiguration.MazeType;
        this.#document.getElementById(`topology-types`).data = Object.entries(Topologies).map(([name, value]) => ({ name, value }))
        this.#document.getElementById(`topology-types`).value = settings.WorldConfiguration.Topology;
        this.#document.getElementById('cellSize').value = settings.WorldConfiguration.CellSize;

        this.#document.getElementById('plantSettings').config = settings.PlantConfiguration
        this.#document.getElementById('herbivoreSettings').config = settings.HerbivoreConfiguration
        this.#document.getElementById('carnivoreSettings').config = settings.CarnivoreConfiguration
        this.#document.getElementById('omnivoreSettings').config = settings.OmnivoreConfiguration
    }
}
