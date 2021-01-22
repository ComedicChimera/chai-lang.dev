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

In line with these goals, Whirlwind is beautiful by design -- it is designed with minimal
syntactic clutter while remaining clear but concise.

*insert some sample code*

## Platform Features

#### Simplicity

Whirlwind attempts to abide by the "it just works" philosophy.  The full toolchain is
super light-weight and easy to use.  There are no unnecessary or over-complicated build
tools to negotiate with and the dependencies are easy to manage even without any external
tools.

Most of the tools you will need to build applications are part of the standard library and
obtaining and installing new packages to supplement your needs is a trivial endeavor.

For larger, more sophisticated projects, Whirlwind ships with a project management
tool (think *much* smaller NPM) called Blend *insert link*.  This will not only manage
local dependencies but also handle static and dynamic linking for you should you need
it for your project.

Finally, Whirlwind supports C binding natively, and Blend is even capable of automatically
generating bindings for you based on a preexisting C library. 

#### Performance

Whirlwind doesn't just look good: it also runs blazing-fast.  Here is how it performs
against some of the top-dogs in its "performance bracket."

*insert graph of speed and of memory usage*

Whirlwind's compiler back-end is built using the [LLVM IR compiler infrastructure](https://llvm.org/docs/index.html).
This means that the IR Whirlwind generates will be transformed into efficient assembly using one of
the most performant and powerful tools in the compilation world, and, with optimization enabled,
your code is put through LLVM's battle-hardened optimizers to produce a ridiculously well-optimized
output.

#### Concurrency

Whirlwind is an innately concurrent language and thus sports a robust concurrency
model.

Whirlwind programs are executed in Strands *insert link* which are like light-weight
threads that entirely managed by the Whirlwind runtime as opposed to by the operating
system.  They are designed to act cooperatively: preferring to voluntarily exchange
data and yield execution as opposed to be interrupted.

*insert some concurrent sample code*

Whirlwind's model is also designed with fault tolerance in mind.  Information is
communicated through managed message queues and groups of Strands can be managed
by a Supervisor *insert link* to ensure that tasks are properly distributed between
them and act appropriately should would Strand panic or fail at any point.

*insert some supervisor sample code*


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

{{< alert theme="info">}}This code is made more complex for the sake of example and showing off more features.{{< /alert >}}

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

#### Baked-In Vectorization

Whirlwind supports vectorization out of the box and comes with a powerful vector data type that enables
efficient numeric computation and fully leverages any system that supports SIMD.  Whirlwind also enables
auto-vectorization where possible.

    import vec_sum from vecutil
    
    func main do
        let matrix = {
            <{1, 2, 3}>,
            <{4, 5, 6}>,
            <{7, 8, 9}>
        }

        let x = <{10, 11, 12}>

        let b = <int : 3>
        for i in 0..2 do
            b[i] = vec_sum(matrix[i] * x)

{{< alert theme="warning" >}}Not all systems support vectorization.  However, having a built-in vector type
that can allow for efficiently applying operations to multiple numeric values is still useful.  Vectors are
not required to use Whirlwind.{{< /alert >}}
        
#### Powerful Package System

In Whirlwind, each directory consistutes a package that contains multiple individual files.  All internal
definitions are shared between files.  Whirlwind can resolve most import cycles automatically and even
allows individual symbols in separate files to mutually reference each other provided such referencing
wouldn't lead to recursive data structure.  This allows you to organize large projects in a very logical
manner and sub-divide the elements of your application as you see fit without any compiler-negotiating.


