<script>
    import {slide} from 'svelte/transition'

    import FeatherIcon from 'common/FeatherIcon.svelte'

    export let unit
    export let selectedChapter

    let open = selectedChapter > unit.startIndex && selectedChapter <= unit.startIndex + unit.chapters.length

    function changeAccordion() {
        open = !open

        feather.replace({
            // initialize these as empty properties so CSS can override them at will
            stroke: "",
            fill: ""
        })
    }
</script>

<style lang="scss">
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";

    .guide-aside-unit {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: space-evenly;

        .unit-title {
            display: flex;
            flex-direction: row;
            align-items: center;
            width: 100%;
            justify-content: space-between;
            margin-bottom: -1rem;

            .unit-title-text {
                color: $text-color;
                font-family: 'Open Sans', sans-serif;
                font-size: 1.3rem;
                overflow: hidden;
                white-space: nowrap;
            }

            .unit-accordion-icon {
                padding: 1rem;
            }
        }

        .unit-chapters {
            margin-left: 0.5rem;

            .unit-chapter a {
                font-family: 'Open Sans', sans-serif;
                font-size: 1.3rem;
                font-weight: 300;
                color: $text-color;
                text-decoration: none;
                transition-duration: 0.5s;
            }

            .unit-chapter .selected {
                color: $primary-color;
            }

            .unit-chapter a:hover {
                color: $primary-color;
                transition-duration: 0.5s;
            }
        }
    }
</style>

<div class="guide-aside-unit">
    <div class="unit-title">
        <span class="unit-title-text">{unit.title}</span>
        <span class="unit-accordion-icon" on:click={changeAccordion}>
            {#if open}
                <FeatherIcon color="black" iconName="chevron-down" size="40"></FeatherIcon>
            {:else}
                <FeatherIcon color="black" iconName="chevron-up" size="40"></FeatherIcon>
            {/if}     
        </span>
    </div>
    {#if open}
        <div class="unit-chapters" transition:slide>
            {#each unit.chapters as chapter, index}
                <div class="unit-chapter">
                    {#if index == unit.startIndex + selectedChapter - 1}
                        <a class="chapter-title selected" href="/docs/guide/{index + unit.startIndex + 1}">{chapter}</a>
                    {:else}
                        <a class="chapter-title" href="/docs/guide/{index + unit.startIndex + 1}">{chapter}</a>
                    {/if}
                </div>
            {/each}
        </div> 
    {/if}
</div>