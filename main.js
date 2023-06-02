import {Application, Configuration} from "./scripts/application/application.js";
import {GameController, UIController } from './scripts/controllers/index.js';

const config = new Configuration({
    isPlayable: false,
});

const application = new Application();
await application.configure(window, document, document.getElementById("canvas"), "engine.wasm")

const uiController = new UIController(window, document);
const gameController = new GameController(application.game, application.configurationProvider, window);

uiController.addOnGenerateButtonClickEventListener(async () => {
    (await gameController.generateGame()).map((_) => {
        console.log("Game generated");
    }).orElse((err) => {
        uiController.showToast(err.message)
    });
})

uiController.addOnPauseButtonClickEventListener( () => {
    gameController.pauseGame();
    uiController.togglePlayPause(false);
})

uiController.addOnPlayButtonClickEventListener( () => {
  gameController.playGame()
      .then(_=> uiController.togglePlayPause(true))
      .catch((err) => {
          uiController.togglePlayPause(false)
          uiController.showToast(err.message)
      })
})

uiController.addOnNextStepButtonClickEventListener( async () => {
    await gameController.nextStep()
})

uiController.addOnSettingsUpdateListener(async (config) => {
    gameController.pauseGame()
    await application.updateConfiguration(config);
    (await gameController.generateGame()).orElse((err) => {
        uiController.showToast(err.message)
    });
})

gameController.addOnGameOverEventListener((state) => {
    uiController.togglePlayPause(false)
    state.map((_) => {
        console.log("Game over");

    }).orElse((err) => {
        uiController.showToast(err.message)
    })
});

uiController.addOnWindowResizeListener(async () => {
   uiController.togglePlayPause(false);
   gameController.pauseGame();
    (await gameController.generateGame())
        .orElse((err) => {
            uiController.showToast(err.message)
        })
});

uiController.init(application.configurationProvider.getInstance());
(await gameController.generateGame())
    .orElse((err) => {
        uiController.showToast(err.message)
    })