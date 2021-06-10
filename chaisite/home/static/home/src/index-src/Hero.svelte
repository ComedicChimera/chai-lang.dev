<script>
    function handleButtonPress() {
        console.log("Redirecting...")
    }
</script>

<style lang="scss">
    @import "chaisite/chaisite/static/common/scss/globals.scss";

    .hero {
        margin-bottom: 5rem;
        
        .hero-side-by-side {
            display: flex;
            flex-direction: row;
            justify-content: space-evenly;
            align-items: center;

            padding: 2rem;
            margin-right: -8rem;
            
            .hero-right-side {
                display: flex;
                flex-direction: column;
                justify-content: space-evenly;
                align-items: center;

                .hero-tag-line {
                    font-family: 'Open Sans';
                    color: #212121;
                    font-size: 2rem;
                    text-align: center;

                    width: 50%;
                }

                .hero-button-container {
                    display: flex;
                    flex-direction: row;
                    justify-content: center;
                    align-items: center;

                    padding: 2rem;

                    button {
                        color: white;
                        background-color: $primary-color;

                        padding: 1rem;
                        border-radius: 0.5rem;
                        font-family: 'Open Sans';
                        font-weight: 300;
                        font-size: 2.5rem;
                        text-align: center;
                        border: none;

                        cursor: pointer;
                    }

                    button:active {
                        background-color: adjust-color($color: $primary-color, $lightness: 10%);

                    }
                }
            }
        }

        .hero-main-title {
            font-family: 'Roboto', sans-serif;
            font-size: 5rem;
            text-align: center;

            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;

            margin-top: 1rem;
            margin-left: -2rem;
        }
    }
</style>

<div class="hero">
    <div class="hero-side-by-side">
        <img src="/static/common/images/chai_logo.png" height="458em" width="400em" alt="Chai logo"/>
        <div class="hero-right-side">
            <div class="hero-tag-line">
                An elegant language for <span class="primary-color">speed</span>
                without <span class="secondary-color">sacrifice</span>.
            </div>
            <div class="hero-button-container">
                <button class="hero-button" on:click={handleButtonPress}>Get Started</button>
            </div>
        </div>
    </div>
    <div class="hero-main-title">
        <span class="primary-color">Cha</span><span class="secondary-color">i</span>
    </div>
</div>