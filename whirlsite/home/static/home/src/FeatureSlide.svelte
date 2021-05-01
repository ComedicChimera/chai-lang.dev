<script>
    import { fade } from 'svelte/transition'

    export let title
</script>

<style lang='scss'>
    .feature {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        align-items: center;

        padding-left: 2rem;
        padding-right: 2rem;

        .feature-content {
            padding-right: 2rem;

            .feature-title {
                display: flex;
                flex-direction: column;
                justify-content: space-evenly;
                align-items: flex-start;

                margin-bottom: 1rem;

                .feature-title-text {
                    font-family: 'Open Sans', sans-serif;
                    color: #212121;
                    font-size: 2rem;
                    margin-bottom: 0.2rem;
                }            
            }

            .feature-description {
                font-family: 'Open Sans', sans-serif;
                color: #212121;
                font-size: 1rem;
                font-weight: 300;
            }
        }
        
    }
</style>

<!--this component is reinstantiated any time the feature changes-->
<div class="feature" in:fade>
    <div class="feature-content">
        <div class="feature-title">
            <span class="feature-title-text">{title}</span>
            <svg width="80" height="5">
                <rect width="60" height="5" fill="#00a8ec"></rect>
                <rect x="60" width="20" height="5" fill="#032f55"></rect>
            </svg>
        </div>
        <div class="feature-description">
            Temp
        </div>
    </div>
    <div class="feature-sample">
        Temp
    </div>
</div>

