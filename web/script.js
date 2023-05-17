import Application from "./scripts/application.js";

const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")
await application.run();