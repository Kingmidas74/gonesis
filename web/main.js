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

const handleTabClick = (e) => {
    const clickedTab = e.target.closest('[data-target]');
    if (!clickedTab) return;

    const container = clickedTab.parentNode;
    const sidebar = container.closest('.settings');

    Array.from(container.children).forEach(tab => {
        tab.classList.remove('active');
    });

    clickedTab.classList.add('active');

    const siblingFieldsets = Array.from(container.nextElementSibling.children);
    siblingFieldsets.forEach(fieldset => fieldset.classList.remove('active'));

    const parentFieldset = container.closest('.form__fieldset');
    if (parentFieldset) {
        const allNestedFieldsets = parentFieldset.querySelectorAll('.form__fieldset');
        allNestedFieldsets.forEach(fieldset => fieldset.classList.remove('active'));
    }

    const targetFieldset = sidebar.querySelector(`#${clickedTab.getAttribute('data-target')}`);
    if (targetFieldset) targetFieldset.classList.add('active');

};

sideBar.addEventListener('click', handleTabClick);

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