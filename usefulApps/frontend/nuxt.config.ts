export default defineNuxtConfig({
    css: [
        'primevue/resources/themes/saga-blue/theme.css',
        'primevue/resources/primevue.css',
        'primeicons/primeicons.css',
        'primeflex/primeflex.css'
    ],
    modules: ['@nuxt/http', '@nuxt/types'],
    build: {
        transpile: ['primevue']
    },
    ssr: false
})
