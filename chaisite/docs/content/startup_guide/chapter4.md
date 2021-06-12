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

The first basic control flow contruct in Chai is the **if statement** which runs
its body if the condition in its header is true.

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

## If Statements as Expressions


    
