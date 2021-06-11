<script>
    import { onMount } from "svelte";

    let subsectionTitles = []

    onMount(() => {
        subsectionTitles = [...document.getElementsByTagName("article")[0]
            .getElementsByTagName("h2")]
            .map(item => item.innerHTML)
    })
</script>

<style lang="scss">
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";

    .guide-summary {
        padding: 2rem;
        padding-left: 1rem;

        .guide-summary-title {
            color: $primary-color;
            font-weight: bold;
            font-family: 'Open Sans', sans-serif;
            font-size: 1.3rem;

            margin-bottom: 0.5rem;

            overflow: none;
            white-space: nowrap;
        }

        .guide-subsections {
            .h2-links {
                display: flex;
                flex-direction: column;
                align-items: flex-start;
                justify-content: space-evenly;

                a {
                    font-family: 'Open Sans', sans-serif;
                    color: $text-color;
                    text-decoration: none;
                    font-size: 1rem;
                    font-weight: 300;

                    overflow: none;
                    white-space: nowrap;

                    margin-bottom: 0.3rem;
                }

                a:hover {
                    color: $primary-color;
                }
            }    
        }
    }
</style>

<div class="guide-summary">
    <div class="guide-summary-title">Page Summary</div>
    <div class="guide-subsections">
        {#each subsectionTitles as item}
            <div class="h2-links">
                <a href="#{item.replaceAll(" ", "-").toLowerCase()}">{item}</a>
            </div>
        {/each}
    </div>
</div>