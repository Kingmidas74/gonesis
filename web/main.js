import {Application, Configuration} from "./scripts/application/application.js";
import {GameController, UIController } from './scripts/controllers/index.js';

const config = new Configuration({
    cellSize: 30,
    isPlayable: false,
    herbivoreConfiguration: {
        InitialCount: 30,
    }
});

const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")
const game = await application.run(config);

const gameController = new GameController(game, application.configurationProvider);

const uiController = new UIController(window, document, gameController);
uiController.OnWindowResizeListener = async () => {
    await gameController.generateGame();
    uiController.togglePlayPause(false);
}

uiController.OnSettingsUpdateListener = async (config) => {
    await gameController.pauseGame()
    uiController.togglePlayPause(false);
    await application.configurationProvider.updateConfiguration(config);
    await gameController.generateGame();
}

uiController.init();