<script>
    import LangFeature from './components/LangFeature.svelte'
    import SectionTitle from './components/SectionTitle.svelte'
</script>

<style lang="scss">
    .feature-container {
        background-color: #eeeeee;
        width: 100%;

        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;

        .title-container {
            width: 60%;
        }

        .features {
            width: 60%;

            display: flex;
            flex-direction: column;
            justify-content: space-evenly;
            align-items: center;
        }
    }
</style>

<div class="feature-container">
    <div class="title-container">
        <SectionTitle>Features</SectionTitle>
    </div>
    <div class="features">
        <LangFeature title="Type Inference">
            <span slot="description">Lorem Ipsum</span>
            <div slot="sample">Lorem Ipsum</div>
        </LangFeature>
    </div>

    
</div>