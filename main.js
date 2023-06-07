import {initializeComponents} from "./scripts/components/components.js";

/*
if ('serviceWorker' in navigator) {
    navigator.serviceWorker
        .register('./sw.js')
        .catch((err) => { console.log('Service Worker Failed to Register', err); });
}

 */

initializeComponents(window, document);