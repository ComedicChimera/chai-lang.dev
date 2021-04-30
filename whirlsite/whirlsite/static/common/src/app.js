import HeaderBar from './HeaderBar.svelte'

import feather from 'feather-icons'

new HeaderBar({
    target: document.getElementById('header'),
})

// This `app.js` is always loaded after all other bundles that use the
// `base.html` template so this call will update all feather icons used.
feather.replace({
    // initialize these as empty properties so CSS can override them at will
    stroke: "",
    fill: ""
})