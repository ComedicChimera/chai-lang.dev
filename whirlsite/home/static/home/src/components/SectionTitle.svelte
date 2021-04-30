<script>

</script>

<style lang='scss'>
    .section-title {
        margin-bottom: 1rem;
        margin-top: 1rem;
        
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: flex-start;

        .title-text {
            font-family: 'Roboto', sans-serif;
            font-size: 2.5rem;
            color: #383838;
        }
    }
</style>

<div class="section-title">
    <span class="title-text">
        <slot></slot>
    </span>
    <svg width="100" height="10">
        <rect width="70" height="10" style="fill:#00a8ec;"/>
        <rect x="70" width="30" height="10" style="fill:#032f55;"/>
    </svg>
</div>