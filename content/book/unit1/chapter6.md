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

## The Return Statement

## Recursion

## Globals


    