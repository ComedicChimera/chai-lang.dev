<script>
    
</script>

<style lang='scss'>
    .whirl-title {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: flex-start;

        margin-bottom: 1rem;

        .whirl-title-text {
            font-family: 'Open Sans', sans-serif;
            color: #212121;
            font-size: 2rem;
            margin-bottom: 0.2rem;
        }            
    }
</style>

<div class="whirl-title">
    <span class="whirl-title-text"><slot></slot></span>
    <svg width="80" height="5">
        <rect width="60" height="5" fill="#00a8ec"></rect>
        <rect x="60" width="20" height="5" fill="#032f55"></rect>
    </svg>
</div>