import Application from "./scripts/application/application.js";
import {Configuration} from "./scripts/application/configuration/configuration.js";

const newConfig = new Configuration({
    cellSize: 10,
    isPlayable: false,
    agentConfiguration: {
        InitialCount: 200,
    }
});


const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")
const game = await application.run(newConfig);

const nextStepBtn = document.getElementById("nextStepBtn")
const playBtn = document.getElementById("playBtn")
const pauseBtn = document.getElementById("pauseBtn")
const generateBtn = document.getElementById("generateBtn")
const settingsBtn = document.getElementById("settingsBtn")

const sideBar = document.getElementById("settings")

let tabs = document.querySelector('.settings--tabs');
let tabFieldsets = document.querySelectorAll('.settings--form > *[data-tab="true"]');

tabs.addEventListener('click', (e) => {
    if (!e.target.matches('.tab')) return;

    for (let tab of tabs.children) {
        tab.classList.remove('active');
    }

    e.target.classList.add('active');

    let target = e.target.getAttribute('data-target');
    tabFieldsets.forEach(fieldset => {
        if (fieldset.getAttribute('id') === target) {
            fieldset.classList.add('active');
        } else {
            fieldset.classList.remove('active');
        }
    });
});


let agentTabs = document.querySelector('.agent--tabs');
let agentTabFieldsets = document.querySelectorAll('#agents > *[data-tab="true"]');

agentTabs.addEventListener('click', (e) => {
    if (!e.target.matches('.tab')) return;

    for (let tab of agentTabs.children) {
        tab.classList.remove('active');
    }

    e.target.classList.add('active');

    let target = e.target.getAttribute('data-target');
    agentTabFieldsets.forEach(fieldset => {
        if (fieldset.getAttribute('id') === target) {
            fieldset.classList.add('active');
        } else {
            fieldset.classList.remove('active');
        }
    });
});

nextStepBtn.addEventListener("click", async (e) => {
    await game.step()
});

playBtn.addEventListener("click", async (e) => {
    nextStepBtn.disabled = true;
    generateBtn.disabled = true;
    application.configurationProvider.getInstance().Playable = true
    playBtn.parentElement.classList.toggle("hidden")
    pauseBtn.parentElement.classList.toggle("hidden")
    await game.run()
});

pauseBtn.addEventListener("click", async (e) => {
        nextStepBtn.disabled = false;
        generateBtn.disabled = false;
        application.configurationProvider.getInstance().Playable = false;
        playBtn.parentElement.classList.toggle("hidden")
        pauseBtn.parentElement.classList.toggle("hidden")
});


generateBtn.addEventListener("click", async (e) => {
    console.log("Generate")
    await game.init()
});

settingsBtn.addEventListener("click", async (e) => {
    sideBar.classList.toggle("active")
});