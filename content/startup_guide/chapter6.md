---
title: "Functions"
weight: 6
---

## Using Functions

A **function call** causes the function to produce a result and/or perform an action based on
its arguments.  As we have seen with the `println` function, we can call a function by using
its name followed by a pair of parentheses in which we place a set of a comma delimited arguments.

    fun(2, "as")
    do_something() # no arguments
    get_factors(35)

Many functions will return a value that can be used in a computation or stored in variable.
This is the value the function will evaluate to when called.

    let sum = add(2, 3)

    factorial(5) * zeta(3) + 45

Some functions (such as `println`) will not return anything.  Unlike in many other languages,
it is possible to extract a value from such functions: this value is simply the `nothing` or
unit type.  This type has no behavior and will be discussed in more detail in a later chapter.

{{< alert theme="info" >}}The unit type is primarily used for things like generics where
a type may choose whether or not it actually stores a value.{{< /alert >}}

## Defining a Function

We have already seen what function definitions look like when we defined the `main` function.
They being with the `func` keyword followed by the name of the function.

However, the main function is especially simple in that it (by default) takes no arguments and
returns no value.  However, the majority of functions you will likely write will do one or
both of those things.  

The **arguments** of a function are enclosed in parentheses and are placed immediately after
the name.  

{{< alert theme="info" >}}When a function takes no arguments, the parentheses are empty.{{< /alert >}}

These arguments are a sequence of names followed by type extensions.  Each name defines one
argument that may be accessed by that name in the body of the function.

    func fn(x: int, y: float) do
        ...

The function `fn` defines two arguments: `x` and `y` where `x` is an integer and `y` is a
float. 

If a function takes two arguments of the same type, those two arguments may share a type label.

    func fn2(a, b: string) do
        ...

`fn2` takes two arguments `a` and `b` that are both strings.  

The next piece of the puzzle is the **return type**.  This is the type of value the function
gives back (or returns) when called.  This type is placed after the arguments and before the
`do`. 

{{< alert theme="info" >}}When a function returns no value, it specifies no return type.{{< /alert >}}

If we specify that our function returns a value, then we must explicitly return such a value
in the body of the function *on all codepaths* (branchs our code can take).  We return values
using the `return` statement.

    func times_two(n: int) int do
        return n * 2

Functions that only require a single expression to perform their computation can use the
**expression return** syntax which is essentially a shorter form of the function above.

    func times_two(n: int) int -> n * 2

{{< alert theme="info" >}}The expression can be on a newline and indented if it is too long
to be reasonable by on the same line as the function signature.{{< /alert >}}.

## Optional Arguments and Named Initializers

An **optional argument** is an argument that doesn't need to passed to the function every time
it is called.  Optional arguments are given a predetermined value as a part of the function's
signature.  This is done using an initializer.

    func newtons_sqrt(n: double, k: int = 10) double do
        let x = n

        for _ in 1..k do
            x = x - (x * x - n) / (2 * x)

        return x

The argument `k` is used to determine how many iterations the algorithm should be ran for and
defaults to `10`.  Notice that `k` still needs a type label -- all optional argument initializers
require type labels.  Optional arguments must *always* be placed after required arguments.

A **named initializer** is a way of explicitly specifying the values of the arguments of a
function by name during the function call.  This done by using an initializer in the function call.

    my_func(2, 3, arg4=2, arg3=0)

The arguments to `my_func`, `arg4` and `arg3` are specified with named initializers.  These
allow us to specify the arguments out of order (presumably `arg3` comes before `arg4` in
the definition) and allows us to choose which (if any) of the optional arguments to a function
we'd like to override.

{{< alert theme="info" >}}Named initializers must always come after the unnamed arguments.{{< /alert >}}

For example, consider a function defined like so:

    func fn(arg1: string, arg2, arg3: int = 0, arg4: bool = false) do
        ...

Say you only wanted to override `arg3` and leave the others untouched.  Without named initializers,
it would be impossible to do without knowing the value of `arg2`.  Luckily, we have named initializers,
so we can do this easily.

    fn("Hello", arg3=2)

## Recursion

**Recursion** occurs when a function calls itself from within its own body.  This is a commonly
supported feature in most programming languages today and is, of course, something Whirlwind supports.

If you are unfamiliar with the concept, it can be a bit difficult to understand and is not really
something that this guide is going to explain in much detail since it is already understood well
enough by most programmers and not a feature that is unique to Whirlwind.  However, for
the purposes of demonstration, we will include a sample use case of recursion: the factorial function.

The factorial of a number is simply the product of the sequence of a whole numbers between itself and
one.  For example, the factorial of `4` is `4 * 3 * 2 * 1`.  We can use recursion to implement this
function cleanly.

    func factorial(n: int) int do
        if n < 3 do
            return n

        return n * factorial(n-1)

For all values less than 3, the factorial is simply an identity (`2 * 1` = 2).  For all other values,
the factorial can simply be considered to be `n` times the factorial of `n - 1` which is what is leveraged
in the recursive definition above.  

An iterative definition of factorial might look like the following:

    func factorial(n: int) int do
        let result = 1

        for i in 1..n do
            result *= i

        return result

It's a lot more cluttered than the recursive method, but it accomplishes the same result.

