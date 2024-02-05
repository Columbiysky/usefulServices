export default defineNuxtConfig({
    css: [
        'primevue/resources/themes/lara-dark-purple/theme.css',
        'primevue/resources/primevue.css',
        'primeicons/primeicons.css',
        'primeflex/primeflex.css'
    ],
    build: {
        transpile: ['primevue']
    },
    modules: [
        'nuxt-primevue'
    ],
    primevue: {
        usePrimeVue: true,
        options: {
            ripple: true,
        },
        importPT: undefined,
        cssLayerOrder: 'tailwind-base, primevue, tailwind-utilities',
    },
    ssr: false,
})
