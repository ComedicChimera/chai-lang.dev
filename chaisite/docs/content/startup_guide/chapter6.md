# Functions 

Functions are the backbone of the Chai programming language: they are its most
important powerful construct and one with which many programmers are already
familiar.  In essence, a **function** is a reusable construct that takes in some
number of values, performs some action, and returns a new value.  In this
section, we will see how to use and define them in Chai.

## Defining Functions

Function definitions in Chai always start with the `def` keyword followed by the
name of the function.  You have already seen this syntax with the `main`
function.

    def main

Next, you will open a pair of parentheses and describe the **arguments** to the
function: these are the values you pass into the function.  You label them much
like variables: a name or series of named separated by commas followed by a type
label.  A function can take many arguments that accept many different types.

Finally, you will include a **return type**.  This is the type of value the
function returns.  Some functions return no value in which case the return type
may be elided.

Putting these two pieces together, we can write the signature for a function
that takes in two numbers and multiplies them.

    def multiply(a, b: int) int

Now, we need to specify the body of the function.  The body is an expression
that is run when the function is called and its yielded value is returned from
the function.  The body is placed at the end of the function prefixed by an `=`.


    def multiply(a, b: int) int = a * b

For longer expressions, it is common to place a newline after the `=` and
indent the expression.  For example, if we wanted to implement the
[absolute value function](https://en.wikipedia.org/wiki/Absolute_value),
we would likely indent the block of code used for the logic.

    def abs(v: f64) f64 =
        if v >= 0 -> v
        else -> -v

As you have already seen, it is very common for functions to be defined as do
blocks.  For these functions, we typically place the `do` immediately after the
`=` and then indent the body of the do block.

    def run_actions(num: int, str: string) = do
        println("Number Action:", num)
        println("String Action:", str)

## Calling Functions

Functions can be called by placing `()` after their name and by placing a set
of argument values in order inside of the parentheses.  For example, a call to
our multiply function from before might look like the following:

    multiply(2, 3)

However, in addition to passing arguments by position, it is often useful to be
able to pass arguments by name.  We can do this by using the `=` syntax inside
the function call -- assigning the argument name to its value.

    multiply(a=2, b=3)

One benefit from passing arguments by name is that we can specify them in any
order we want.

    run_actions(str="test", num=4)

You can combine named and positional arguments to give your code extra
readability.

    # NOT A BUILTIN FUNCTION :)
    google_search("kittens", allow_ads=false, num_results=20)

When calling functions, particularly those take many, complex arguments, it is
often useful to split calls over multiple lines.  Chai will allow you to do
this with no issue.

    google_search(
        "kittens",
        num_results=20,
        allow_ads=false
    )

> Whenever you are inside grouping symbols, Chai will completely ignore
> whitespaces such as newlines or indentations meaning you can structure calls,
> tuples, and other accumulate types however you want from a whitespacing
> perspective.  Note that any whitespace sensitive blocks such as match blocks
> or do blocks will "reenable" whitespace awareness in their bodies.  This is
> called indentation framing.

Importantly, named arguments must always come after positioned arguments.

    fn(x=32, "test")  # ERROR

## The Return Statement

The **return statement** allows you to return from a function immediately from
wherever you are in its body.  The statement begins with keyword `return`
followed by any value(s) the function returns.  

    def sqrt(v: f64) (f64, bool) = do
        if v < 0 do
            return -1, false

        # -- SNIP: calculate sqrt; store in `result` --

        return result, true

Notice that the return statement will let you return a tuple of values without
needing to wrap those values in parentheses.

> The style of error handling shown above is valid in Chai but not really the
> best solution.  In a later chapter, we will cover a better solution to error
> handling.

Notably, if the function returns no value at all, you can simply place `return`
to exit the function early.

    def do_something() = do
        if cond do
            return

        # other function logic

Note that the return statement will exit the function from any depth within it.

## Optional and Variadic Arguments

Functions in Chai also support **optional arguments** -- arguments that don't
need to be supplied with every function call.  Optional arguments can be created
by simply initializing the argument in the function signature.

    def google_search(search_string: string, num_results: int, allow_ads: bool = false) = do
        ...

Importantly, when defining optional arguments, the argument data type is still required.

Now, we can call `google_search` with or without the `allow_ads` argument.  When
we choose to elide the argument: the default value provided is substituted in.

    google_search("kittens", 20)
    google_search("super cute kittens", 15, true)

Optional arguments must be defined after required arguments, but functions can
accept multiple optional arguments.  This is another great use case for named
arguments: only specifying the options you care about.

    def call_api(endpoint: string, timeout: int = 200, token: string = "") = do
        ...

    def main() = do
        call_api("/get", timeout=500)
        call_api("/admin/analytics", token="DO NOT USE THIS AS A TOKEN")

TODO: variadic arguments

## Recursion

TODO: recursion

