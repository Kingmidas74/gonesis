import Application from "./scripts/application.js";
import { Configuration } from "./scripts/game/configuration.js";

const newConfig = new Configuration({
    cellSize: 30,
    mazeColor: "#111111",
    agentsColor:"#000099",
});

const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")
await application.run(newConfig);