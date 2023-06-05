const CACHE_NAME = 'gonesis-cache';
const urlsToCache = [
    //paths to files
    './',
    'index.html',
    'main.css',
    'main.js',
    'assets/manifest.json',
    'assets/sprite.svg',
    'assets/icons/icon-128x128.png',
    'assets/icons/icon-144x144.png',
    'assets/icons/icon-152x152.png',
    'assets/icons/icon-192x192.png',
    'assets/icons/icon-384x384.png',
    'assets/icons/icon-512x512.png',
    'assets/icons/icon-72x72.png',
    'assets/icons/icon-96x96.png',
    'scripts/application/application.js',
    'scripts/application/configuration/configuration.js',
    'scripts/application/engine/engine.js',
    'scripts/application/engine/engine.wasm',
    'scripts/application/engine/wasm_exec.js',
    'scripts/application/game/game.js',
    'scripts/application/game/index.js',
    'scripts/application/game/world-manager.js',
];

self.addEventListener('install', event => {
    event.waitUntil((async () => {
        const cache = await caches.open(CACHE_NAME);
        await cache.addAll(urlsToCache);
    })());
});

self.addEventListener('fetch', event => {
    // only cache GET requests
    if (event.request.method !== 'GET') return;

    event.respondWith((async () => {
        try {
            // Try to fetch from network first
            const freshResponse = await fetch(event.request);
            // Then cache the response
            const cache = await caches.open(CACHE_NAME);
            cache.put(event.request, freshResponse.clone());
            return freshResponse;
        } catch (error) {
            // If network fails, try to fetch from cache
            const cacheResponse = await caches.match(event.request);
            if (cacheResponse) return cacheResponse;
            // Handle scenario when both network and cache fail
            throw Error('Both network and cache have failed.');
        }
    })());
});

self.addEventListener('activate', event => {
    event.waitUntil((async () => {
        const cacheNames = await caches.keys();
        await Promise.all(
            cacheNames.map(cacheName => {
                if (cacheName !== CACHE_NAME) {
                    return caches.delete(cacheName);
                }
            })
        );
    })());
});
