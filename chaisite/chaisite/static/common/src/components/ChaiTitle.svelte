<script>
    // used for overriding empty slot
    export let titleString = ""
</script>

<style lang='scss'>
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";

    .chai-title {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: flex-start;

        margin-bottom: 1rem;

        .chai-title-text {
            font-family: 'Open Sans', sans-serif;
            color: $text-color;
            font-size: 2rem;
            margin-bottom: 0.2rem;
            font-weight: normal;
        }            
    }
</style>

<div class="chai-title">
    <span class="chai-title-text">
        {#if titleString == ""}
            <slot></slot>
        {:else}
            {titleString}
        {/if}
    </span>
    <svg width="80" height="5">
        <rect width="60" height="5" fill="#be6a15"></rect>
        <rect x="60" width="20" height="5" fill="#ac3f21"></rect>
    </svg>
</div>