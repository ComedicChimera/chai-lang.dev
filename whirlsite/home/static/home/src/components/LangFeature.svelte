<script>
    export let title
</script>

<style lang="scss">
    .feature {
        width: 100%;

        .feature-title {
            font-family: 'Roboto', sans-serif;
            font-size: 2.0rem;
            color: #383838;

            margin-bottom: 0.5rem;
        }

        .feature-content {
            display: flex;
            flex-direction: row;
            justify-content: space-evenly;
            align-items: center;

            .feature-description {
                font-family: 'Open Sans', sans-serif;
                font-size: 1.2rem;
                color: #383838;
            }
        }
    }
</style>

<div class="feature">
    <div class="feature-title">{title}</div>
    <div class="feature-content">
        <div class="feature-description">
            <slot name="description"></slot>
        </div>
        <div class="feature-sample">
            <slot name="sample"></slot>
        </div>
    </div>
</div>