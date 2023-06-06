import {Component} from "./decorators.js";

import { RADIO_GROUP_TOGGLE } from "./radio-group-toggle/radio-group-toggle.js";
import { COLOR_PICKER } from "./color-picker/color-picker.js";
import { RANGE_SLIDER } from "./range-slider/range-slider.js";

import { TERRAIN_SETTINGS } from "./terrain-settings/terrain-settings.js";
import { AGENT_SETTINGS } from "./agent-settings/agent-settings.js";
import { PRIMARY_TOOLBAR } from "./primary-toolbar/primary-toolbar.js";
import { GAME_SETTINGS } from "./game-settings/game-settings.js";
import { TOAST } from "./toast/toast.js";

/**
 * Initializes custom components by defining them in the customElements registry.
 * @param {Window} window - The window object of the page.
 * @param {Document} document - The document object of the page.
 * @param {Console} [logger=console] - The logger object used for logging. Defaults to the console object.
 */
const initializeComponents = (window, document, logger = console) => {
    const componentOptions = {
        templatePath: null,
        stylePath: null,
        windowProvider: window,
        documentProvider: document,
        logger,
    };

    const components = [
        {
            name: "app-radio-group-toggle",
            component: RADIO_GROUP_TOGGLE,
            templatePath: "./radio-group-toggle.html",
            stylePath: "./radio-group-toggle.css",
        },
        {
            name: "app-color-picker",
            component: COLOR_PICKER,
            templatePath: "./color-picker.html",
            stylePath: "./color-picker.css",
        },
        {
            name: "app-range-slider",
            component: RANGE_SLIDER,
            templatePath: "./range-slider.html",
            stylePath: "./range-slider.css",
        },
        {
            name: "app-terrain-settings",
            component: TERRAIN_SETTINGS,
            templatePath: "./terrain-settings.html",
            stylePath: "./terrain-settings.css",
        },
        {
            name: "app-agent-settings",
            component: AGENT_SETTINGS,
            templatePath: "./agent-settings.html",
            stylePath: "./agent-settings.css",
        },
        {
            name: "app-primary-toolbar",
            component: PRIMARY_TOOLBAR,
            templatePath: "./primary-toolbar.html",
            stylePath: "./primary-toolbar.css",
        },
        {
            name: "app-game-settings",
            component: GAME_SETTINGS,
            templatePath: "./game-settings.html",
            stylePath: "./game-settings.css",
        },
        {
            name: "app-toast",
            component: TOAST,
            templatePath: "./toast.html",
            stylePath: "./toast.css",
        }
    ];

    components.forEach(({ name, component, templatePath, stylePath }) => {
        if (templatePath && stylePath) {
            componentOptions.templatePath = templatePath;
            componentOptions.stylePath = stylePath;
            window.customElements.define(
                name,
                Component(componentOptions)(component)
            );
        } else {
            window.customElements.define(name, component);
        }
    });
};

export { initializeComponents };
