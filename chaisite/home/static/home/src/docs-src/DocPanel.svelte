<script>
    export let title
    export let href
</script>

<style lang="scss">
    @import "chaisite/chaisite/static/common/src/scss/globals.scss";

    .doc-panel {
        padding: 1rem;
        margin: 1rem;
        border-radius: 0.3rem;
        background-color: $light-bg;

        display: flex;
        flex-direction: column;
        align-items: center;

        width: 30%;

        flex: 1;

        .doc-panel-title {
            margin-bottom: 0.7rem;

            .doc-panel-title-text {
                color: $primary-color;
                font-size: 1.5rem;
                text-decoration: none;
                font-family: 'Open Sans', sans-serif;
                text-align: center;
            }
        }

        .doc-panel-content {
            font-family: 'Open Sans', sans-serif;
            font-size: 1rem;
            color: $text-color;
            text-align: left;
            font-weight: 300;
        }
    }
</style>

<div class="doc-panel">
    <div class="doc-panel-title">
        <a class="doc-panel-title-text" href="{href}">{title}</a>
    </div>
    <div class="doc-panel-content">
        <slot></slot>
    </div>
</div>