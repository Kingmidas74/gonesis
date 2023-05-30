import {Application, Configuration} from "./scripts/application/application.js";
import {GameController, UIController } from './scripts/controllers/index.js';

const config = new Configuration({
    isPlayable: false,
});

const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")

const gameController = new GameController(application.game, application.configurationProvider);
const uiController = new UIController(window, document, gameController);
uiController.OnWindowResizeListener = async () => {
    uiController.togglePlayPause(false);
    await gameController.generateGame();
}

uiController.OnSettingsUpdateListener = async (config) => {
    await gameController.pauseGame()
    uiController.togglePlayPause(false);
    await application.configurationProvider.updateConfiguration(config);
    await gameController.generateGame();
}

application.game.addOnGameOverListener((e) => {
    uiController.togglePlayPause(false)
    e.map((_) => {
        console.log("Game over");
    }).orElse((err) => {
        console.error(err);
    });
})


uiController.init()
await gameController.generateGame();