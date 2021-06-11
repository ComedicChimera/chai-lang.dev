<script>
    import { fade } from 'svelte/transition'
    import outdent from 'outdent'

    import CodeView from 'common/CodeView.svelte'
    import ChaiTitle from 'common/ChaiTitle.svelte'

    import {features} from './feature-list'

    export let title
</script>

<style lang='scss'>
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";
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
                font-family: $text-font;
                color: $text-color;
                font-size: 1rem;
            }
        }
    }
</style>

<!--this component is reinstantiated any time the feature changes-->
<div class="feature" in:fade>
    <div class="feature-content">
        <ChaiTitle>{title}</ChaiTitle>
        <div class="feature-description">
            {features[title].desc}
        </div>
    </div>
    <div id="feature-sample" class="feature-sample">
        {#if features[title].sample.type == "code"}
            <CodeView language="chai">
                {outdent.string(features[title].sample.data)}
            </CodeView>
        {/if}
    </div>
</div>

