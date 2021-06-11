<script>
    import FeatherIcon from 'common/FeatherIcon.svelte'

    function navLeft() {
        if (window.chapterNumber == 1) {
            window.location.href = "/docs/guide"
        } else {
            window.location.href = `/docs/guide/${window.chapterNumber-1}`
        }
    }

    function navRight() {
        window.location.href = `/docs/guide/${window.chapterNumber+1}`
    }
</script>

<style lang="scss">
    .guide-nav-buttons {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
    }
</style>

<div class="guide-nav-buttons">
    {#if window.chapterNumber != 0}
        <div class="guide-nav-button" on:click={navLeft}>
            <FeatherIcon iconName="chevron-left" color="black" size=80></FeatherIcon>
        </div>
    {:else}
        <div class="guide-nav-button">
            <FeatherIcon iconName="chevron-left" color="none" size=80></FeatherIcon>
        </div>
    {/if}
    {#if window.chapterNumber != 25}
        <div class="guide-nav-button" on:click={navRight}>
            <FeatherIcon iconName="chevron-right" color="black" size=80></FeatherIcon>
        </div>
    {:else}
        <div class="guide-nav-button">
            <FeatherIcon iconName="chevron-right" color="none" size=80></FeatherIcon>
        </div>
    {/if}
</div>