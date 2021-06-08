import Hero from './index-src/Hero.svelte'
import Features from './index-src/Features.svelte'

new Hero({
    target: document.getElementById('hero')
})

new Features({
    target: document.getElementById('features')
})
