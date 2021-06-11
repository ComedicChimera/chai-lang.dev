<script>
    import GuideAsideUnit from './GuideAsideUnit.svelte'

    const guideLayout = [
        {
            title: "Unit I - Variables and Expressions",
            chapters: [
                "Hello World",
                "Basic Types"
            ],
            startIndex: 1
        }
    ]

    function redirectHome() {
        window.location.href = "/docs/guide"
    }
</script>

<style lang="scss">
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";

    .guide-aside {
        padding: 1rem;
        padding-left: 2rem;
        margin: 1rem;
        background-color: $light-bg;
        border-radius: 0.3rem;

        .guide-title {
            color: $primary-color;
            font-size: 2rem;
            font-family: 'Roboto', sans-serif;
            cursor: pointer;
        }
    }
</style>

<aside class="guide-aside">
    <div class="guide-title" on:click={redirectHome}>Getting Started Guide</div>

    {#each guideLayout as unit}
        <GuideAsideUnit unit={unit} selectedChapter={window.chapterNumber}></GuideAsideUnit>
    {/each}
</aside>