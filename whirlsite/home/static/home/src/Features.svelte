<script>
    // import LangFeature from './components/LangFeature.svelte'
    // import SectionTitle from './components/SectionTitle.svelte'

    let features = [0, 1, 2, 3, 4]
    let currentFeature = 0

    function switchToFeature(id) {

    }

    function nextFeature() {
        if (currentFeature < features.length-1) {
            switchToFeature(currentFeature+1)
        } else {
            switchToFeature(0)
        }
    }

    function prevFeature() {
        if (currentFeature > 0) {
            switchToFeature(currentFeature-1)
        } else {
            switchToFeature(features.length-1)
        }
    }
</script>

<style lang="scss">
    .feature-container {
        background-color: #ececec;
        width: 100%;

        margin-top: 2rem;
        margin-bottom: 2rem;

        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;

        .feature-selector {
            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;

            .feature-select-button {
                padding: 0.5rem;
                border: none;

                svg {
                    fill: #8E8E8E;
                    transition-duration: 0.5s;
                }

                svg:hover {
                    transition-duration: 0.5s;
                    fill: #00a8ec;
                }
            }
        }

        .feature-deck {
            .feature-slide-button {
                border: none;
            }
        }
    }
</style>

<div class="feature-container">
    <div class="feature-deck">
        <button id="feature-left" class="feature-slide-button" on:click={prevFeature}>
            <i class="icon-grey-hover" data-feather="chevron-left" width="100" height="100" stroke-width="0.7" fill="none"></i>
        </button>
        <button id="feature-right" class="feature-slide-button" on:click={nextFeature}>
            <i class="icon-grey-hover" data-feather="chevron-right" width="100" height="100" stroke-width="0.7" fill="none"></i>
        </button>
    </div>
    <div class="feature-selector">
        {#each features as feature}
            <button class="feature-select-button" on:click={() => switchToFeature(feature)}>
                <svg width="16" height="16"><circle cx="8" cy="8" r="8"></circle></svg>
            </button>
        {/each}
    </div>
</div>