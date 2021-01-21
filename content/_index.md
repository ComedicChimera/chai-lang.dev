---
title: Home
---

# {{< color "#00a8ec" >}}Whirl{{< /color >}}{{< color "#032f55" >}}wind{{< /color >}}

{{< heading-card >}}

Whirlwind is a compiled, modern, and multipurpose language designed with
intentionality. It is strongly-typed, versatile, expressive, concurrent, and
relatively easy to learn. It boasts numerous new and old features and is
designed to represent the needs of any software developer.  

## Goals

Whirlwind has several main goals all of which can be enscapsulated in the idea of speed.

- Speed of writing.
- Speed of thinking.
- Speed of compiling.
- Speed of running.

Whirlwind recognizes that performance while important is not the only feature that
matters to modern programmer.  Being able to develop swiftly and fluidly is critical
to designing robust and elegant applications.  Whirlwind strives to allow you to
translate your ideas into code as quickly and effortlessly as possible without incurring
a significant performance cost.  Moreover, by removing the clutter of the
standard development workflow, we allow more time for optimization, testing, and polishing.

## Platform Features

- Simplicity (it just works)
- Performance
- Concurrency

## Language Features

#### Versatile Type System

Whirlwind's type system is designed to adapt to a wide range of applications and situations
fluidly and logically.  It features both structured types as well as algebraic types coupled
with a clever interface model.  Moreover, the language supports and was designed with generics
in mind for writing widely reusable code.

    type Option<T>
        | Some(T)
        | None

    interf<T> for Option<T> of
        func chain<R>(f: func(T)(Option<R>)) Option<R> =>
            match this to
                Some(v) => f(v)
                None => None

    func main do
        let r = option_sqrt(2)

        r.chain(|x| => println(x))

#### Rich Functional Programming

In Whirlwind, all functions are first-class types and support partial function calling (implicit
abstraction).  Whirlwind also features anonymous functions (lambdas), closures, and explicit
composition by default.  

    func fibonacci func()(int) do
        let a = 0, b = 1

        return || do
            yield a
            a, b = b, a + b

    func print_number(prefix: string, n: int) do
        println(prefix + ":", n)

    func main do
        let fib_print = print_number("The next fibonacci number is", _)

        let fib = fibonacci()
        for _ in 1..10 do
            fib_print(fib())

*This code is made more complex for the sake of example and showing off more features.*
#### Elegant Error Handling

While exceptions are a powerful tool in the right hands, they can often lead to sloppy code and cause
unexpected changes in control flow which can cause resources not to be cleaned up or important state
not to be updated.  Whirlwind prefers to indicate failure by return value, thereby making the ability to
fail part of the function's definition and requiring the caller to appropriately handle the failure state.
However, it also tries to make this process as smooth as possible by leveraging its type system and
some syntactic sugar.

    func main do
        with
            resp <- call_api_endpoint("/get-stuff")
            raw_data <- extract_data(resp)
            data <- validate_raw(raw_data)
        do
            println(data["field"])
        else match Err(e) do
            println("Error:", e)

- Pervasive Pattern Matching
- Baked-In Vectors (SIMD)
- Powerful Package System


