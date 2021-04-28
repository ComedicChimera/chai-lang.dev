// NOTE: This file is completely unnecessary from the standpoint of actually
// building the application.  It is mainly here so VSCode won't complain about
// me using SCSS inside svelte components.
const sveltePreprocess = require('svelte-preprocess')

module.exports = {
    preprocess: sveltePreprocess()
}