<script>
    import GuideAsideUnit from './GuideAsideUnit.svelte'

    const guideLayout = [
        {
            title: "Variables and Expressions",
            chapters: [
                "Hello World",
                "Basic Types",
                "Variables",
                "Control Flow",
                "Pattern Matching"
            ],
            startIndex: 0
        },
        {
            title: "Functions and Collections",
            chapters: [
                "Functions",
                "Lists and Dictionaries",
                "References",
                "Loops and Comprehensions",
                "Vectors and Matrices",
                "Lambdas and Closures",
                "Sequence Functions"
            ],
            startIndex: 5
        },
        {
            title: "Defined Types",
            chapters: [
                "Structure Types",
                "Algebraic Types",
                "Method Spaces",
                "Type Classes",
                "Generics",
                "Monads and Error Handling"
            ],
            startIndex: 12
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