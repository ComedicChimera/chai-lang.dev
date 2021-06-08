// This file contains all the "source" for the home page features
// of the feature slider.
module.exports = {
    features: {
        "Algebraic Types": {
            desc: `
                Chai supports algebraic types: powerful data types
                that are made up of multiple variants representing
                different possible states or values.  These types
                dramatically reduce boiler-plate and allow you to
                express complex ideas concisely.
            `,
            sample: {
                type: "code",
                data: `
                    type Expr =
                        | Add(Expr, Expr)
                        | Div(Expr, Expr)
                        | Val(int)
                    
                    def evaluate(e: Expr) Option[int] =
                        match e to
                            Val(x) -> Some(x)
                            Add(e1, e2) -> evaluate(e1)? + evaluate(e2)?
                            Div(e1, e2) -> match evaluate(e2) to
                                Some(0) -> None
                                Some(x) -> evaluate(e1)? // x
                                None -> None
                `
            }
        },
        "Data Processing": {
            desc: `
                Collections are a fundmental tool of data processing
                and manipulation.  Chai makes taking and manipulating
                sets of data easy by offering a selection of versatile ways
                of generating, representing, and manipulating collections
                such as lists and dictionaries.
            `,
            sample: {
                type: "code",
                data: `
                    def radix_sort(list: List[int]) List[int] = do
                        let max = list.max()
                    
                        while let it = 0; 10 ** it < max; it++ do
                            let buckets = {List.[int].new() for _ in 0..10}
                    
                            for item in list do
                                buckets[item % (10 ** it) // 10].push(item)
                    
                            list = buckets.flatten().to_list()
                    
                        return list
                `
            }
        },
        "Concurrency": {
            desc: `
                Chai's concurrency model is facilitated by
                lightweight threads of execution called strands
                with can run in parallel but switch and share
                resources cooperatively.  They can be easily
                coordinated using Chai's simple yet powerful
                concurrency infrastructure to build fault-tolerant,
                data-race-free systems.
            `,
            sample: {
                type: "code",
                data: `
                    import pool_results from asyncutil
                    import get from net.http
                    import println from io.std

                    def main() = do
                        let results = pool_results(
                            | get "my-api.com/get-stuff" ),
                            | get "my-api.com/get-more-stuff" ),
                            | get "my-api.com/get-other-stuff" )
                        )

                        async for result in results do
                            if result match Ok(resp) do
                                println(resp.data)

                `
            }
        },
        "Performance": {
            desc: `
                Chai doesn't just look good; it runs blazingly
                fast too.  Built on the tried-and-true LLVM infrastructure
                and enhanced by its conditional garbage collection algorithm, 
                Chai holds its own both in both CPU and memory benchmarks.
            `,
            sample: {}
        },
        "Generics": {
            desc: `
                Generics are an essential mechanism for building reusable
                but expressive code.  Chai uses generics to their full
                effect allowing you to avoid duplication and build reusable
                and extensible APIs.
            `,
            sample: {}
        }
    }
}