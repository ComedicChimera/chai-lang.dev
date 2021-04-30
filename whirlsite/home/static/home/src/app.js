import Hero from './Hero.svelte'
import Features from './Features.svelte'

new Hero({
    target: document.getElementById('hero')
})

new Features({
    target: document.getElementById('features')
})
