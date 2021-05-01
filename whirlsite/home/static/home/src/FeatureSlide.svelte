<script>
    import CodeView from 'common/CodeView.svelte'

    import { fade } from 'svelte/transition'

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

        .feature-content {
            padding-right: 2rem;
            width: 30%;

            .feature-title {
                display: flex;
                flex-direction: column;
                justify-content: space-evenly;
                align-items: flex-start;

                margin-bottom: 1rem;

                .feature-title-text {
                    font-family: 'Open Sans', sans-serif;
                    color: #212121;
                    font-size: 2rem;
                    margin-bottom: 0.2rem;
                }            
            }

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
        <div class="feature-title">
            <span class="feature-title-text">{title}</span>
            <svg width="80" height="5">
                <rect width="60" height="5" fill="#00a8ec"></rect>
                <rect x="60" width="20" height="5" fill="#032f55"></rect>
            </svg>
        </div>
        <div class="feature-description">
            {#if title == "Type System"}
                Whirlwind's versatile, strong, static type system allows you to write
                clean, expressive code with little hassle and complete compile-time
                type safety.  Moreover, Whirlwind's powerful type inferencer ensures
                that type labels won't clutter up your code.
            {:else}
                Temp
            {/if}
        </div>
    </div>
    <div id="feature-sample" class="feature-sample">
        {#if title == "Type System"}
            <CodeView language="whirlwind">
type Expr
    | Add(Expr, Expr)
    | Div(Expr, Expr)
    | Val(int)
    
func evaluate(e: Expr) Option&lt;int&gt;
    -&gt; match e to
        Val(x) -&gt; Some(x)
        Add(e1, e2) -&gt; evaluate(e1)? + evaluate(e2)?
        Div(e1, e2) -&gt; match evaluate(e2) to
            Some(0) -&gt; None
            Some(x) -&gt; evaluate(e1)? // x
            None -&gt; None
            </CodeView>
        {:else if title == "Data Processing"}
            <CodeView language="whirlwind">
func radix_sort(list: [int]) [int] do
    let max = list.max()

    while let it = 0; 10 ** it &lt; max; it++ do
        let buckets = [make [int] for _ in 0..10]

        for item in list do
            buckets[item % (10 ** it) // 10].push(item)

        list = buckets.flatten().to_list()

    return list
            </CodeView>
        {:else}
            Temp
        {/if}
    </div>
</div>

