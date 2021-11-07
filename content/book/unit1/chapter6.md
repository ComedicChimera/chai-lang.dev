# Functions

A **function** is a reusable block of code that can be **called** with some
**arguments** to produce a **return value**.  They are one of the only two
places programmers actually define the logic of their code in Chai: all the code
we have covered thus far has been occurring inside functions.

Most programmers are already familiar with the basic idea of a function:
however, as you will see in later chapters, Chai pushes the idea of a function
to its absolute extreme.  In this chapter, however, we will stick the more
mundane and basic aspects of functions in Chai.

## Calling Functions

As you have likely inferred from previous chapters, functions can be called by
placing parentheses after the function name.

    my_func()

To pass arguments to a function, we simply place them inside the parentheses
separated by commas.

    add(2, 3)

All functions in Chai will return a value.  The result of the call expression is
always that value.

    let sum = add(2, 3)  # sum = 5

However, some functions such as the main function may return the special value
`()` or **nothing** which really means that they don't return anything. The
nothing value can still be used and treated as a value, but it really has no
real use: you can't operate on its and it doesn't really have any meaning
besides being nothing.  It's type label is simply `nothing` -- you will rarely
find this type used explicitly.

> The main use of the nothing value is to allow code that deals with generic
> functions (and other generic constructs) which can return a variety of types
> to not have to treat nothing as a special case.  We will learn more about
> generics in later chapters.

## Defining Functions

A function definition is made of four key parts:

1. The Name
2. The Arguments
3. The Return Type
4. The Function Body

All function definitions begin with the `def` keyword followed by the name of
the function.  Function definitions can only occur at the top level of the
program: ie. you cannot define functions inside of functions.

The most simple function definition would look like so:

    def my_func()
        # ...
    end

You have already seen such a definition: the `main` function.  This function
takes no arguments and returns nothing.  This, of course, means that it has the
inferred return type of `nothing`: when a return type is not specified, the
function is assumed to return the nothing value.

### Function Bodies

The function body is what comes after the return type.  The function above has a
**block body** which means that its body is a block.  These kinds of bodies are
denoted by placing a newline immediately after the return type.

Functions can also have an **expression body** which begins with the `=`
operator followed by the expression to return.  We saw this kind of body in our
hello world program where the expression returned was the result of the call to
`println`.

    def print_hello() = println("Hello!")

Expression bodies can also be written on a separate line from the function
definition by simply placing an optional newline after the `=`.

    def print_hello() =
        println("Hello!")

Note that block bodies are semantically identical to returning a do block via an
expression body:

    # *identical* to the above `my_func`
    def my_func() = do
        # ...
    end

Block-bodied functions are used so often that the special syntax of placing a
newline immediately after the return type was created to avoid having to type `=
do` over and over again.

### Return Types

The **return type** of the function is placed immediately after the `()` enclosing
the arguments.  This type may be elided to denote that the function returns nothing.

The type of the body is required to match the return type of the function.

    def get_pi() f64 = 3.14159265  # correct

    def get_h() f64 = "Plank's constant"  # TYPE ERROR

This is true even for block-bodied functions: recall that blocks yield the type
of the value of the last statement or expression in them.

    def get_e() f64
        println("User requested Euler's number")
        2.7182  # correct
    end

However, functions that return a type of `nothing` do *not* have to have a
matching return value in their body. 

    def print_help_message()
        println("Help message!")
        5  # ok
    end

The logic behind this is two fold: firstly, there is no variation among nothing
values: one nothing is identical to another; and secondly, this is often
convenient: it means that you don't have tack on an extra `()` at the end of
such functions.

Furthermore, functions that return nothing have the expressions that make up
their bodies treated as statements (including the last statement/expression of
their block). This means that if expressions and match expressions yielded from
functions that return return nothing may be non-exhaustive and type
inconsistent.

    def my_func() = if some_cond -> println("some_cond is true") end  # ok

    def my_func2()
        println("Something...")

        if some_cond
            println("Something else...")
        end  # also ok
    end

    def my_func3() i32 = if some_cond -> 2 end  # ERROR: not exhaustive

    def my_func4() f32
        if some_cond
            println("message")
        else
            43.3
        end  # ERROR: type inconsistent
    end

Effectively, functions that return nothing treat their expressions as pure
statements: allowing your code to be written more procedurally.

### Arguments

The arguments that a function takes are specified inside the parentheses of the
function definition.  Each argument begins with a name followed by type
extension.  

    def double(x: i32) i32 = x * 2

Functions can take multiple arguments separated by commas, and arguments can
share a type extension like in variable definitions.

    def sum_and_scale(a, b: i64, s: f64) f64 = (a + b as f64) * s

All function arguments must have a distinct name from other arguments to the
same function.  However, function arguments can be shadowed by other variables
of the same name within their bodies.

    def exclaim_message(msg: string) string
        let msg = msg + "!"

        println(msg)

        msg
    end

## The Return Statement

The **return statement** is a special control flow statement that causes the
enclosing function to immediately return when encountered. 

It begins with the `return` keyword.

    def print_n_times(msg: string, n: i64)
        if n == 0
            return

        # -- snip --
    end

Return statements can also return values.  

    def exp(power: f64) f64
        if power == 0
            return 1

        # -- snip --
    end

Return statements must match the return type of the function.  If a function
returns nothing, then the return statement must either specify no value or `()`
as its return value.  Similarly, if a function does return a type, then the
expression of the return statement must be of that type.

    def some_func(s: string) i32
        return s  # TYPE ERROR
    end

The return statement can also take multiple values, separated by commas, to
return.  In this case, those values are automatically packaged into a tuple for
you.

    import sqrt from math
    
    def quadratic_formula(a, b, c: f64) (f64, f64)
        let discrim = sqrt(b ** 2 - 4 * a * c)

        return (-b + discrim) / (2 * a), (-b - discrim) / (2 * a)        
    end

Return statements also affect exhaustiveness checking: any branch that
definitively ends a return statement is considered exhaustive.  This includes
branches that occur inside expressions that are being immediately returned from
the enclosing function.

    def some_func(a, b: i32) i32
        let x = if a < b 
            return a
        else -> a + b end

        # -- snip --
    end

As you can observe, returns also circumvent type consistency requirements since
the branch that contains them never actually stores into the variable.
## Recursion

**Recursion** occurs when a function calls itself from within its own body.
Recursion is a widely supported feature in most programming languages, and, in
Chai, its behaves how you would expect.

A simple example to demonstrate this is to consider the 
[factorial function](https://en.wikipedia.org/wiki/Factorial).  For those
unfamiliar, the factorial is defined as follows:

```language-text
n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
```

As you can see, factorial is an inherently recursive function: `n!` = `n * (n - 1)!`

We can implement this function in Chai easily.  To begin, let's start with the
function signature.

    def fac(n: i32) i32 =
        # TODO

Then, we consider first the base case: `0! = 1`.  All other recursive cases will
reach this base eventually.

    def fac(n: i32) i32 =
        if n == 0 -> 1
        # TODO

Then, we consider the recursive case which is `fac(n) = n * fac(n-1)`.  Adding
this to our function, we get:

    def fac(n: i32) i32 =
        if n == 0 -> 1
        else -> n * fac(n-1)
    end 

> Putting the end of block statements inside expression function bodies an
> indentation level out is a fairly common stylistic practice in Chai: it
> denotes that both the expression and the function are ended.

We can check that this function works by expanding out an evaluation of it.
Consider `fac(4)` as an example:

```language-text
fac(4) =
    4 != 0 => 4 * fac(4-1) = 4 * fac(3)
fac(3) = 
    3 != 0 => 3 * fac(3-1) = 3 * fac(2)
fac(2) =
    2 != 0 => 2 * fac(2-1) = 2 * fac(1)
fac(1) =
    1 != 0 => 1 * fac(1-1) = 1 * fac(0)
fac(0) =
    0 == 0 => 1

Therefore:
fac(4) = 4 * fac(3)
       = 4 * 3 * fac(2)
       = 4 * 3 * 2 * fac(1)
       = 4 * 3 * 2 * 1 * fac(0)
       = 4 * 3 * 2 * 1 * 1
       = 24
```

As you can see, our factorial function matches up exactly with the mathematical
definition in execution.  Most programmers are always familiar with recursion so
we will leave the discussion here having demonstrated how recursion works in
Chai.  However, if you come from a language that does not support this paradigm,
then try playing around with this idea: recursion is extremely powerful when
used well and can allow to write very beautiful and intuitive code.

## Globals

**Globals** are variables and constants defined outside of the bodies of
functions. They use the exact same syntax as regular variables and constants but
have somewhat different semantics.  

    const pi: f64 = 3.14159264

    def circle_area(r: f64) f64 = pi * r ** 2

The `pi` constant is fully determined before any user-defined function,
including main, is ever run.  The same logic is true for mutable variables.

    let counter = 0

    def special_func()
        println("This function has been called", counter, "times.")
        counter++
    end

Notice that the value of `counter` is preserved between function calls as apart
of global state.

Global variables like regular variables can contain full block expressions in
their bodies.

    let tau = do
        println("initializing `tau`")

        pi * 2
    end

    def circumference(r: f64) f64 = tau * r
    
> Runtime initialization does occur before global variable initialization so you
> are free to use functions such as `println` in their bodies.

The reason we chose to discuss global variables now is that such variables do
have some distinctions from normal local variables that can only be appreciated
with an understanding of functions.  

Firstly, the return statement can not, for obvious reasons, be used inside a
global variable initializer.

    let my_global = do
        return  # ERROR
    end

Secondly, global variables will be shadowed by arguments and local variables
within functions.

    let n: i32

    def increment(n: i32) i32 = n + 1

> For the remainder of this book, you can assume that when we use variables,
> they are not global (ie. enclosed inside the main function) unless the context
> specifically requires them to be global.