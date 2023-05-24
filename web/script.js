import Application from "./scripts/application.js";
import { Configuration } from "./scripts/configuration/configuration.js";

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


document.getElementById("nextStepBtn").addEventListener("click", async (e) => {
    await game.step()
});

document.getElementById("playBtn").addEventListener("click", async (e) => {
    application.configurationProvider.getInstance().Playable = true
    await game.run()
});


document.getElementById("generateBtn").addEventListener("click", async (e) => {
    await game.init()
});

document.getElementById("settingsBtn").addEventListener("click", async (e) => {
    //application.configurationProvider.getInstance().Playable = false
    document.getElementById("settings").classList.toggle("active")
});
