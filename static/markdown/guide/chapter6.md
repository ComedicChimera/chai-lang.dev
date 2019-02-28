# Functions

A function is a callable block of code that can accept an external state in the form of arguments.
In Whirlwind, a function is composed of three primary parts: the name, arguments, and return
types.  

## Basic Functions

All functions begin with the `func` keyword and take the following form.

    func add(x, y: int) int {
        return x + y;
    }

`add` is the name of the function and `x` and `y` are the arguments, both are integers.
The return type of the function is an integer.

As you can see from its body, this function take in two arguments, adds them, and returns
their sum.  The `return` keyword is used to return a value or values.  

Since this function only has one statement in its body which is a return statement,
you can rewrite the using an expression return.

    func add(x, y: int) int => x + y;

This is just shorthand for the first function definition, as it is common for functions to
return single, complex expressions.

Functions can also take no or as many arguments as you choose of multiple types.  For example,
if you have a function that returns the time, its signature might look like this.

    func getTimeInHours() int {
        // get time in hours
    }

Similarly, a function that sums a list, might look like the following.

    func sum(lst: list[int]) int {
        let s = 0;

        for (n <- lst)
            s += n;

        return s;
    }

Functions can take as many or as few arguments as necessary.  Arguments
can also be marked as constant by putting the const keyword before the
name.

    func sum(const lst: list[int]) int {
        let s = 0;

        for (n <- lst)
            s += n;

        return s;
    }

Constant arguments are only constant within the function body and do not
necessarily have to be constant outside the function.  However, constant
values must be passed or respected as constant within a function.

## Function Calling

Calling functions is fairly simple.  You simply have to
access to function by name and follow it with parentheses.

    function();

If a function takes argument, each argument should be provided
to the function in order, separated by commas.

    function2(3, "a");

Finally, arguments can be specified by name for additional clarity
or do accomodate more complex types of arguments.

    function2(a=3, b="a");

As you can see, the traditional assignment operator is used here.
The arguments must be correctly named and any non-existent argument
names will cause in an error.

> Certain types of complex arguments cannot be specifically assigned
> using this method.

Besides the three forms shown, there are no other ways to call functions.
This will be the consistent function call form for all functions regardless
of what type of function they are.

## Optional Arguments

Optional arguments are another method by which functions can be
made more versatile.  An optional argument is an argument with a default value.  
These arguments do not have to be specified and can be left empty if
the creator so chooses.

These arguments are declared by specifying a default value in
functions signature.

    func someFunc(x: int, y: double = 1.0) {
        // -- snip --
    }

As you can see, these signatures look similar to the way arguments
as explicitly specified in a function call.  The only
difference being the optional type annotation.  An argument declaration without
the type annotation would look like so.

    func someFunc(x: int, y = 1.0) {

    }

It is important to note that type inference is used here, so `y` is still typed.
It has a type of `float`

## Indefinite and Variadic Arguments

**Indefinite arguments** are arguments that can take in multiple values.  They are
designated by placing `...` before the argument name.

    func makeList(... args: int) list[int]
        => args;

The function would the be called like so.

    makeList(4, 5, 3, 1); // [4, 5, 3, 1]

Indefinite argument values are accepted without any form of enclosing collection
and as many can be provided as possible.  However, you may notice that
the indefinite arguments are converted to a list before they are passed to the function.

Additionally, indefinite arguments must be the last argument in the function.
Due to that fact, a function can only have one indefinite argument.  Finally,
indefinite arguments cannot be specified by a argument specifier.  So doing the following
is invalid.

    makeList(args=4); // ERROR

The next kind of argument is similar to an indefinite argument with one big difference:
they do not have a type.  These arguments are called **variadic arguments**.

    func varFunc(...args) {
        // -- snip --
    }

These arguments can only be accessed and managed by using the va_args builtin
function.  You can than use the `get()` method to operate upon your
variadic argument list.

    func varFunc(...args) {
        let vArgs = va_args(args);

        let first = vArgs.get<int>(0);
        let second = vArgs.get<str>(1);
    }

The type in angle brackets is the type of the original argument passed in
and the value passed into the `get()` function is the position of the argument
in the argument list (bounds checked).  

> The `get()` method on variadic argument lists does not allow negative indices.

## Recursion

Recursion is the process by which a function can call itself within its
own body.

    func recur() {
        // infinite loop
        recur();
    }

This can be used to add additional power and flexibility to any given function.
For example, the factorial can be calculated using recursion.

    func fac(n: int) int {
        if (n > 1)
            return n * fac(n - 1);

        return n;
    }

Just be careful when using recursion as it is easy to create infinite loops
and stack overflow exceptions.