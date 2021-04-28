<script>
    function handleButtonPress() {
        console.log("Redirecting...")
    }
</script>

<style lang="scss">
    .hero {
        width: 70%;

        .hero-side-by-side {
            display: flex;
            flex-direction: row;
            justify-content: space-evenly;
            align-items: center;

            padding: 10px;            
        }
    }
</style>

<svelte:options tag="app-hero"/>

<div class="hero">
    <div class="hero-side-by-side">
        <img src="/static/common/images/logo.png" height="250em" width="250em" alt="Whirlwind logo"/>
        <div class="hero-right-side">
            <div class="hero-tag-line">
                An elegant language for <span class="light-blue">speed</span>
                without <span class="dark-blue">sacrifice</span>.
            </div>
            <div class="hero-button-container">
                <button class="hero-button" on:click={handleButtonPress}>Get Started</button>
            </div>
        </div>
    </div>
    <div class="hero-main-title">
        <span class="light-blue">Whirl</span><span class="dark-blue">wind</span>
    </div>
</div>