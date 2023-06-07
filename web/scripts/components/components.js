import {Component} from "./decorators.js";

import { RADIO_GROUP_TOGGLE } from "./basic/radio-group-toggle/radio-group-toggle.js";
import { COLOR_PICKER } from "./basic/color-picker/color-picker.js";
import { RANGE_SLIDER } from "./basic/range-slider/range-slider.js";

import { TERRAIN_SETTINGS } from "./complex/terrain-settings/terrain-settings.js";
import { AGENT_SETTINGS } from "./complex/agent-settings/agent-settings.js";
import { PRIMARY_TOOLBAR } from "./complex/primary-toolbar/primary-toolbar.js";
import { GAME_SETTINGS } from "./complex/game-settings/game-settings.js";
import { TOAST } from "./complex/toast/toast.js";
import { TAB_LAYOUT } from "./basic/tab-layout/tab-layout.js";

import { GAME_VIEW } from "./complex/game-view/game-view.js";
import { BRAIN_VIEW} from "./complex/brain-view/brain-view.js";
import { CHART_VIEW } from "./complex/chart-view/chart-view.js";


import { APPLICATION } from "./application/application.js";

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
        },
        {
            name: "app-game-view",
            component: GAME_VIEW,
            templatePath: "./game-view.html",
            stylePath: "./game-view.css",
        },
        {
            name: "app-brain-view",
            component: BRAIN_VIEW,
            templatePath: "./brain-view.html",
            stylePath: "./brain-view.css",
        },
        {
            name: "app-chart-view",
            component: CHART_VIEW,
            templatePath: "./chart-view.html",
            stylePath: "./chart-view.css",
        },
        {
            name: "app-tab-layout",
            component: TAB_LAYOUT,
            templatePath: "./tab-layout.html",
            stylePath: "./tab-layout.css",
        },
        {
            name: "app-genesis",
            component: APPLICATION,
            templatePath: "./application.html",
            stylePath: "./application.css",
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
