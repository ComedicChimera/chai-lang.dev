<script>
    import ChaiTitle from "common/ChaiTitle.svelte";

    export let title;
</script>

<style lang="scss">
    .doc-row {
        width: 60%;

        .row-items {
            padding-left: 1.5rem;
            padding-right: 1.5rem;

            display: flex;
            flex-direction: row;
            justify-content: space-evenly;
            align-items: stretch;
        }
    }
</style>

<div class="doc-row">
    <ChaiTitle>{title}</ChaiTitle>
    <div class="row-items">
        <slot></slot>
    </div>
</div>