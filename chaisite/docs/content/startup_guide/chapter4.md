# Control Flow

This chapter is only going to cover the first and most basic control flow
construct in Chai because Chai handles control flow slightly differently then
you may be used.  Ergo, we will need to take time to explore each control flow
construct thoroughly.

## Block Expressions

Unlike most languages, Chai considers blocks to be expressions -- ie. things
that can yield a value.  In fact, the do blocks you were introduced to last
chapter are themselves expressions.

    let x = do
        ...

> As you might guess, the main function we defined earlier also accepts an
> expression after its `=`, but since do blocks are expressions, we had no
> issues using it. 

By default, do blocks will yield the value of the last statement contained
within them.  

    let x = do
        let y = 2 + 3
        y++  # increment yields the new value added

Many of the statements we have already covered yield values as well.  

> However, some statements and functions return `()` which is a special value of
> type `nothing`.  This value effectively signifies that no value is yielded but
> provides Chai with a type to process so your code still makes sense.

We can also specify the value the do block yields with the **yield statement**.

    x = do
        let y = 6 * 7
        yield y + 2  # 44 will be returned
        y++  # the yielded value is not effected by this update

As you can see, `yield` does not change the flow of the program -- it just
changes what value the do block will return.

## If, Elif, and Else

The first basic control flow contruct in Chai is the **if expression** which
runs its body if the condition in its header is true.

    # NOTE: `random_number` is NOT a real function in Chai
    let x = random_number()  

    if x < 5 do
        println("x is less than 5")

As we can see, its body is defined by a do block containing a single print
statement which will only run if x is less than 5.

We can also add an `else` condition that will execute if the condition is false.

    if x % 2 == 0 do
        println("`x` is even.")
    else
        println("`x` is odd.")

Notice that we do not need to put a `do` at the end of the else to begin the
block. This is common with many blocks that are begun by single keywords.  

> You can include a `do` if you want -- Chai doesn't care.

Finally, you can use the `elif` keyword to introduce a series of secondary
conditions. These conditions will be checked in order after the primary if, and
the block of the first one that is true will be executed, breaking the chain.

    # NOTE: `random_color` is also not a real function
    let color = random_color()

    if color == "red" do
        println("Blood for the blood god!")
    elif color == "green" do
        println("A fine choice.")
    elif color == "blue" do
        println("Ocean lover, I see...")
    else
        println("Not RGB? Quirky.")

You do not need to include an `else` at the end of an `elif` chain.  However,
the statement must be constructed beginning with an `if`, followed by any `elif`
blocks, and (optionally) ending with an `else`.

## If Expressions Yielding Values

If expressions can yield values, specifically they yield the type that is on
their branches.  Moreover, they can yield expressions instead of do blocks using
the `->` symbol.

    let x = random_number()

    let v = if x < 50 -> "small" else -> "big"

Note that the branches must all yield the same type (or types that can all be
unified to the same type -- ie. all coercible to the same type).

    let y = (
        if x < 0 -> "less than zero"
        elif x > 0 -> "greater than zero"
        else -> 1  # TYPE ERROR  
    )

> When you want to use block statements like if blocks as expressions but still
> want to separate the blocks over multiple lines, you can use parentheses to
> wrap the block statement.

However, if the block is used within a do block but not at its end or as the body
of a function or block that returns nothing, we don't need to obey this rule.

    def main() = do
        let x = random_number()

        if x < 0 ->
            "less than zero"
        elif x > 0 ->
            "greater than zero"
        else -> 1

Even though the if block is at the end of the do block and its branches yield
different typed values, the `main` function returns no value, and therefore
we are allowed to use it here.

> Notice that we indent the expression blocks inside if expressions -- this is
> not just for code appearance but actually necessary in order for Chai to
> properly process your code when you use `->`.  If you were to elide the
> indentations, you would get a syntax error on the `elif`.  In general, if you
> are not using do blocks as the bodies for your if expressions, you must either
> wrap the block in parentheses (as shown above) or use indentation to denote
> that you want the if block to continue (ie. have `elif` and `else` clauses).

All the if expressions we have looked at so far have been **exhaustive** meaning
they have a branch for every possible input value.  However, if branches do not
have to be exhaustive in order to be used as expressions.

    let x = if random_number() % 2 == 0 -> 5

However, if you were to examine the type of `x`, you find that it is not a type
of `int` but rather a type of `Option[int]`.  This is a special, builtin type
that is used to represent the idea of that a function or operation may or may
not return a value.  We will talk much more about how to use these types in
later chapters, but we wanted to mention this behavior so it doesn't catch you
off guard when you are writing your code.  

## Shadowing

**Shadowing** is an important but simple behavior of variables.  It occurs
within any scope hierarchy and allows a variable named identically to that in a
higher scope to "shadow" or stand in place of that higher variable.

> For those unfamiliar, the term scope refers to the area over which a variable
> exists or is defined.  For example, a variable defined in an if block is
> contained within the scope of that if block and cannot be used outside it.

Here is some example code that demonstrates this behavior.

    let x = 0

    if some_cond do
        let x = 2 # No error occurs here
        if some_other_cond do
            println(x) # Prints `2` not `0`

Although variables cannot be defined multiple times in the same scope, as you
can see above it is possible to override their definitions in lower scopes.  The
variable `x` is shadowed by a variable of the same name defined in a lower
scope.

You can also define variables in the head of the if expression itself.  Simple
write a variable declaration followed by a semicolon and the conditional
expression for the if expression inside the if header.

    if let x = random_number(); x % 2 == 0 do
        println(x, "is even")

Note that if variables will shadow variables in our scopes and are only visible
within the block or branch expression of the if expression.
