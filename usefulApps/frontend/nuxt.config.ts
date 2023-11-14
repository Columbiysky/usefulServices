export default defineNuxtConfig({
    css: [
        'primevue/resources/themes/saga-blue/theme.css',
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
        options: {},
        importPT: undefined,
        cssLayerOrder: 'tailwind-base, primevue, tailwind-utilities',
        components: {
            prefix: '',
            name: undefined,
            include: undefined,
            exclude: undefined
        },
        directives: {
            prefix: '',
            name: undefined,
            include: undefined,
            exclude: undefined
        },
        composables: {
            prefix: '',
            name: undefined,
            include: undefined,
            exclude: undefined
        }
    },
    ssr: false,
})
