<script>
    import {onMount} from 'svelte'

    import prism from 'prismjs'

    if (prism.languages.whirlwind == undefined) {
        prism.languages.whirlwind = window.language_whirlwind
    }

    export let language
    let el

    onMount(() => {
        prism.highlightElement(el, false)
    })
</script>

<style lang='scss'>
    pre {
        margin-top: 2rem;
        margin-bottom: 2rem;
    }
</style>

<pre>
    <code class="language-{language}" bind:this={el}>
        <slot></slot>
    </code>
</pre>