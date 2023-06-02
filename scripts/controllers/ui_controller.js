import {Configuration, MazeGenerators} from "../application/configuration/configuration.js";

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
        this.#elements.sideBar.querySelectorAll('.color-picker__color').forEach(picker => {
            picker.addEventListener('change', this.#handleSettingsUpdate);
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

        this.#document.querySelector('.terrain-types')?.addEventListener('change', (event) => {
            if (!event.target.matches('input[type="radio"]')) {
                return
            }

            this.#document.querySelector('.maze-types').classList.toggle('hidden', event.target.value !== 'maze');
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
        let mazeType = this.#document.querySelector('input[name="terrain"]:checked').value;
        if (mazeType === 'maze') {
                    mazeType = this.#document.querySelector('input[name="maze"]:checked').value;
        }
        return new Configuration({
            worldConfiguration: {
                CellSize: this.#document.getElementById('cellSize').value,
                MazeType: mazeType,
                Topology: this.#document.querySelector('input[name="topology"]:checked').value,
            },
            herbivoreConfiguration: {
                InitialCount: this.#window.parseInt(this.#document.getElementById('herbivoreInitialCount').value),
                Color: this.#document.getElementById('herbivoreColor').value,
            },
            carnivoreConfiguration: {
                InitialCount: this.#window.parseInt(this.#document.getElementById('carnivoreInitialCount').value),
                Color: this.#document.getElementById('carnivoreColor').value,
            },
            plantConfiguration: {
                InitialCount: this.#window.parseInt(this.#document.getElementById('plantInitialCount').value),
                Color: this.#document.getElementById('plantColor').value,
            },
            omnivoreConfiguration: {
                InitialCount: this.#window.parseInt(this.#document.getElementById('omnivoreInitialCount').value),
                Color: this.#document.getElementById('omnivoreColor').value,
            }
        })
    }

    #handleSettingsUpdate = (e) => {
        if(e.target.classList.contains('range-slider__range')
        || e.target.classList.contains('color-picker__color')) {
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
        if([MazeGenerators.AldousBroder, MazeGenerators.Binary, MazeGenerators.SideWinder].includes(settings.WorldConfiguration.MazeType)) {
            this.#document.querySelector(`input[name="terrain"][value="maze"]`).checked = true;
            this.#document.querySelector(`input[name="maze"][value="${settings.WorldConfiguration.MazeType}"]`).checked = true;
            this.#document.querySelector(`.maze-types`).classList.toggle('hidden', false);
        } else {
            this.#document.querySelector(`input[name="terrain"][value="${settings.WorldConfiguration.MazeType}"]`).checked = true;
            this.#document.querySelector(`.maze-types`).classList.toggle('hidden', true);
        }

        this.#document.querySelector(`input[name="topology"][value="${settings.WorldConfiguration.Topology}"]`).checked = true;
        this.#document.getElementById('cellSize').value = settings.WorldConfiguration.CellSize;

        this.#document.getElementById('herbivoreInitialCount').value = settings.HerbivoreConfiguration.InitialCount;
        this.#document.getElementById('herbivoreColor').value = this.#hslaToHex(settings.HerbivoreConfiguration.Color);
        this.#document.getElementById('herbivoreColor').parentNode.querySelector('.range-slider__value').innerHTML = this.#document.getElementById('herbivoreColor').value;
        this.#document.getElementById('carnivoreInitialCount').value = settings.CarnivoreConfiguration.InitialCount;
        this.#document.getElementById('carnivoreColor').value = this.#hslaToHex(settings.CarnivoreConfiguration.Color);
        this.#document.getElementById('carnivoreColor').parentNode.querySelector('.range-slider__value').innerHTML = this.#document.getElementById('carnivoreColor').value;
        this.#document.getElementById('plantInitialCount').value = settings.PlantConfiguration.InitialCount;
        this.#document.getElementById('plantColor').value = this.#hslaToHex(settings.PlantConfiguration.Color);
        this.#document.getElementById('plantColor').parentNode.querySelector('.range-slider__value').innerHTML = this.#document.getElementById('plantColor').value;
        this.#document.getElementById('omnivoreInitialCount').value = settings.OmnivoreConfiguration.InitialCount;
        this.#document.getElementById('omnivoreColor').value = this.#hslaToHex(settings.OmnivoreConfiguration.Color);
        this.#document.getElementById('omnivoreColor').parentNode.querySelector('.range-slider__value').innerHTML = this.#document.getElementById('omnivoreColor').value;
    }

    #hslaToHex = (hsla) => {
        const hslaInArray = hsla.substring(5, hsla.length-1).replace(/ /g, '').split(',');

        let h = parseInt(hslaInArray[0]) / 360; // we need to convert it to be between 0 to 1
        let s = parseInt(hslaInArray[1]) / 100; // we need to convert it to be between 0 to 1
        let l = parseInt(hslaInArray[2]) / 100; // we need to convert it to be between 0 to 1
        let a = parseFloat(hslaInArray[3]); // we need to convert it to be between 0 to 1

        let rgba = this.#hslaToRgba(h, s, l, a);
        let hex = this.#rgbaToHex(rgba);

        return hex;
    }

    #hslaToRgba = (h, s, l, a) => {
        let r, g, b;

        if(s === 0){
            r = g = b = l; // achromatic
        } else {
            const hue2rgb = (p, q, t) => {
                if(t < 0) t += 1;
                if(t > 1) t -= 1;
                if(t < 1/6) return p + (q - p) * 6 * t;
                if(t < 1/2) return q;
                if(t < 2/3) return p + (q - p) * (2/3 - t) * 6;
                return p;
            };
            let q = l < 0.5 ? l * (1 + s) : l + s - l * s;
            let p = 2 * l - q;
            r = hue2rgb(p, q, h + 1/3);
            g = hue2rgb(p, q, h);
            b = hue2rgb(p, q, h - 1/3);
        }
        return 'rgba('+Math.round(r * 255)+','+Math.round(g * 255)+','+Math.round(b * 255)+','+a+')';

    }

    #rgbaToHex = (rgba) => {
        const sep = rgba.indexOf(",") > -1 ? "," : " ";
        rgba = rgba.substr(5).split(")")[0].split(sep);

        // Strip the slash if using space-separated syntax
        if (rgba.indexOf("/") > -1)
            rgba.splice(3,1);

        for (let R in rgba) {
            let r = rgba[R];
            if (r.indexOf("%") > -1) {
                const p = r.substr(0,r.length - 1) / 100;

                if (R < 3) {
                    rgba[R] = Math.round(p * 255);
                } else {
                    rgba[R] = p;
                }
            }
        }

        let r = (+rgba[0]).toString(16),
            g = (+rgba[1]).toString(16),
            b = (+rgba[2]).toString(16);

        if (r.length === 1)
            r = "0" + r;
        if (g.length === 1)
            g = "0" + g;
        if (b.length === 1)
            b = "0" + b;

        return "#" + r + g + b;
    }
}
