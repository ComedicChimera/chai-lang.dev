<script>
    import { fade } from 'svelte/transition'
    import outdent from 'outdent'

    import CodeView from 'common/CodeView.svelte'
    import WhirlTitle from 'common/WhirlTitle.svelte'

    import {features} from './feature-list'

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

        width: 70%;
        height: 25rem;

        .feature-sample {
            width: 33rem;
        }

        .feature-content {
            padding-right: 2rem;
            width: 30%;

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
        <WhirlTitle>{title}</WhirlTitle>
        <div class="feature-description">
            {features[title].desc}
        </div>
    </div>
    <div id="feature-sample" class="feature-sample">
        {#if features[title].sample.type == "code"}
            <CodeView language="whirlwind">
                {outdent.string(features[title].sample.data)}
            </CodeView>
        {/if}
    </div>
</div>

