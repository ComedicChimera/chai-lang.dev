# Control Flow

Control flow is an essential part of any program: it determines both its
structure and its behavior.  For this reason, Whirlwind offers a versatile set
of control flow structures, the building blocks of which will be covered in this
section.

## If, Elif, and Else

The first and most basic control flow construct is the `if` statement.  If
statements run the code in the block if the condition in their header is true.
They are written like so:

```whirlwind
    if some_cond do
        ...
```

Notice that our block begins with the keyword `do` and continues in the next
indentation level much like the main function examined in the previous section.

> The `...` is actual part of Whirlwind's syntax denoting that a given block is
> incomplete.  The compiler will simply insert the most logical code possible in
> such a block to not cause an error.  It is akin to the *pass* statement in
> Python.

If statements can be combined with `elif` and `else` clauses to create a tower
of conditional blocks to be executed if the previous block fails.  

For example, the following code uses and if/else tower to make decisions about
some user input (read in the `readln` function).

```whirlwind
    import println, readln from io::std

    func main do
        println("Enter your favorite color:")
        let user_color = readln()

        if user_color == "green" do
            println("Hey, that's my favorite color too!")
        elif user_color == "blue" do
            println("Like the ocean!")
        elif user_color == "red" do
            println("Purple is a better light saber color.")
        else
            println("Not RGB? PCMR would not be happy.")
```

In this program, the `elif` branches only run if the previous branch doesn't.
Of course the `else` branch runs if all of the other conditions fail.

> It is also, of course, possible to have just an `if` and an `else` or an `if`
> with some number of `elif` blocks and no `else.

Notice also the `else` clause does not require a `do` after it (although once
can be placed there if you so choose).  You will see a number of single keyword
block declarators do not require a `do` to be placed after them.

## Conditional Loops

A loop is simply a section of code that gets executing repeatedly.  Whirlwind
offers several kinds of loops, the simplest being the conditional loop.  A
conditional loop is begun with the `loop` keyword and repeats the code inside
its block so long as the condition in its header is true.

```whirlwind
    let x = 0
    loop x < 10 do
        x++
```

The code above simply repeated increments `x` until it is equal to `10`.  

Loops also allow the use of the `break` and `continue` keywords to control the
flow of the loop by exiting the loop immediately and continuing to the next
iteration (respectively).  Here is an example use of the `break` keyword:

```whirlwind
    import println, readln from io::std

    const password = "password123"

    func main do
        let resp = ""

        loop resp != "QUIT" do
            println("Enter the password:")
            resp = readln()

            if resp == password:
                println("You guessed the password!")
                break
```

All loops in Whirlwind have the ability to utilize a `nobreak` clause that is
run only if the loop exits normally (ie. if the loop is not broken).  This
clause is not only cleaner but more efficient than using a flag variable.  Here
is a common use case of the `nobreak` keyword.

```whirlwind
    import println from io::std

    const magic_number = 12

    func main do
        let list = get_some_list()

        loop list.len() > 0 do
            if list[0] == magic_number do
                println("Found the magic number!")
                break

            list = list[1:]
        nobreak
            // notice the omission of the `do`
            println("Did not find the magic number.")
```

> Generally when searching through a list, one would want to use a `for` loop
> which we will cover later.  Moreover, although in this example, one could
> return from the function for the same effect as `nobreak`, in a more complex
> control flow structure or situation, such a return may not be feasible which
> is where `nobreak` comes in.  Recall that the code above is simply an example.

As a final note about conditional loops, it is possible to omit the condition to
achieve the effect of an infinite loop (instead of writing `loop true do`) like
so:

```whirlwind
    import print from io::std

    func main do
        print("Forever ")

        loop
            print("and ever ")
```

The output of this program is *forever* followed by an infinite sequence of *and
ever*.  Once again, notice the omission of the `do` after the single keyword.

## For Loops

The next kind of loop in Whirlwind is the for loop.  These loops are used to perform
some kind of iteration and come in two variants: the c-style for loop and the iterable
for loop.  

### The C-Style For Loop

This is the for loop that most programmers will be most familiar with coming
into Whirlwind. It consists of three parts: the iterator variable declarations,
the condition, and update statement declared in that order.  A sample for loop
to print out the numbers zero through nine written C-Style would look like so:

```whirlwind
    for i := 0; i < 10; i++ do
        println(i)
```

> The `:=` operator is used in a few places to denote an implicit declaration.
> It cannot be used in a statement on its own but can be used when information
> about the named value can be reliably inferred (eg. iterator variables will
> always be mutable).

The first statement declared our iterator variable `i` (bound inside the scope
of the loop). The second expression is the condition that determines whether or
not the loop should execute. The final statement is executed at the end of the
loop.

The for loop above is mostly just syntactic sugar for the following code:

```whirlwind
    let i = 0
    loop i < 10 do
        println(i)
        i++
```

The only notable differences are that in a for loop, `i` is bound within the
scope of the loop and can't be used outside it and that the update (`i++`) is
executed at the end of all iterations (unless the loop is broken) whereas in the
loop above, if a continue were to be encountered, the update would be skipped.

It is also worth mentioning the various pieces of the for loop can be elided
if they are not necessary or desirable.  For example, if you want access to the
iterator variable outside of the loop scope, you can use a predeclared variable
like so:

```whirlwind
    let i = 0
    for ; i < 10; i++ do
        println(i)
```

### The Iterable For Loop

The iterable for loop is by far the more powerful and common counterpart to the
C-Style for loop.  It is used to iterate through the elements of an iterable
(such as a list or dictionary).  They are written like so:

```whirlwind
    for item in iterable do
        ...
```

Using a *for-iter loop* (shorthand for iterable for loop), we could rewrite our
search code from earlier like so:

```whirlwind
    import println from io::std

    const magic_number = 12

    func main do
        let list = get_some_list()

        for item in list do
            if item == magic_number do
                println("Found the magic number!")
                break
        nobreak
            println("Did not find the magic number.")
```

Not only is this code cleaner, it also doesn't modify the underlying list and is
considerably faster.

Another salient feature of the for-iter loop is that they support pattern
matching for the iterator variables.  For example, if we wanted to iterate
through the index and value of a list at the same time, we would the use the
iterable method `enumerate` along with for-iter pattern matching to cleanly and
efficiently do this:

```whirlwind
    let fib_numbers = [1, 1, 2, 3, 5, 8, 13, 21]

    for i, num in fib_numbers.enumerate() do
        println(i, num)
```

### When to Use Each Kind of Loop

Given that Whirlwind supports two different kinds of for-loop, it is important
to mention when to use one over the other.  A good rule is that whenever you
want to iterate over something, be it a list, dictionary, or other iterable, you
should always use the for-iter loop.  If you want to generate a sequence of
values, you should also see if you can use an iterable.  However, if you simply
want to set, check, and update (ie. to get a range of numbers), the c-style for
loop is likely the better (and faster) choice.

All in all, if you can't decide, it is generally considered more idiomatic
(and clean) to use a for-iter loop.

## Match Statements

The final basic control flow structure in Whirlwind is the match statement. The
match statement performs a case analysis based on a given value.  

For example, the if/else tree matching different color names could be easily
rewritten to utilize a match statement instead.

```whirlwind
    import println, readln from io::std

    func main do
        println("Enter your favorite color:")
        let user_color = readln()

        match user_color to
            case "green" do
                println("Hey, that's my favorite color too!")
            case "blue" do
                println("Like the ocean!")
            case "red" do
                println("Purple is a better light saber color.")
            case _ do
                println("Not RGB? PCMR would not be happy.")
```

> As the name would imply, the match statement fully supports pattern matching
> on all of its cases, and we are, in fact, using it above to convey the idea of
> a default case.  We will see more about pattern matching a bit later.

Whirlwind match statements do not require (or allow) a break to occur at the end
of a given case.  Instead, each case simply matches and exits.  This affords us
several benefits including concision and the ability to properly use break
statements to exit loops *inside* the match statement.

Because our match statement does not fall through by default (like in C or
Java), we have two features to supplement the behavior lost by the removal of
that default fall through behavior.

Firstly, we can match multiple expressions in a single case by simply including
multiple expressions in the case block header separated by commas.

```whirlwind
    match user_color to
        case "red", "green", "blue" do
            println("An RGB color")
        case "yellow", "cyan", "magenta", "black" do
            println("A CMYK color")
```

Secondly, we can cause a match statement to fall through to the next case using
the `fallthrough` keyword at the end of a case block.

```whirlwind
    match 2 to
        case 1 do
            println("One")
            fallthrough
        case 2 do
            println("Two")
            fallthrough
        case 3 do
            println("Three")
        case _ do
            println("Another number.")
```

The output of this code looks like this:

```terminal
    Two
    Three
```

## Pattern Matching

We have already seen a little bit of pattern matching thus far in variables,
iterators, and in the default case of the match statement.  Now we will take a
bit of a deeper look into pattern matching and how it works inside of the match
statement.

The first thing to know is that we can use pattern matching to declare case
specific variables like so:

```whirlwind
    match (2, 3, 4) to
        case (x, _, 7) do
            println(x)
        case (2, y, 4) do
            println(y)
        case (_, 3, z) do
            println(z)
```

This code will match to the third case and declare the variable `y` within that
case to be `3` because of the value at its position in the original tuple.
Whenever a pattern variable occurs, the slot it occupies can be any value (just
like with `_`).  The only different is that now we are extracting that value as
we watch so we can do something with it.

> It is worth noting that pattern variables will not shadow any variables with
> the same name declared in any enclosing scope.  For example, if the variable
> `y` were to be declared visibly somewhere else in the above code, the compiler
> would take the existing value of `y` over creating a pattern variable.

There are some limitations on the usage of pattern matching in a match
statement: specifically of pattern variables.  Pattern variables can only be
used in single expression cases that cannot be fallen through to from a previous
case.  This is because Whirlwind might not known which variables to populate
with values or in the case of the `fallthrough`, how to populate them.

## Control Expressions

Control expressions are simpler forms of some of the conditional control flow
structures intended for use within expressions.  There are two such control
expressions: the match expression and the conditional expression.

The match expression works similarly to its match statement counterpart with two
key differences. The first is that the branches themselves must be expressions,
and the second is that they must be *exhaustive* meaning they have a branch for
every possible value of the expression they are matching over.

Here is an example of a simple match expression:

```whirlwind
    let vp = match (-1, 1) to
        (x, 0) => x
        (0, y) => y
        (x, y) => x * y
```

Each `=>` designates the expression that branch returns (instead of the `case`
syntax).  

Because the match expression only supports expressions on its branches, it is
not possible to use `fallthrough` in a match expression.  However, we can match
against multiple expressions on a single branch.

```whirlwind
    let v2 = match vp to
        1, -1 => 0
        0 => 1
        2 => -1
        _ => 2  // our match must be exhaustive!
```

It is possible to use match expressions inside more complex expressions like so:

```whirlwind
    let y = [
        match (vp, v2) to
            (a, 0) => a
            (-1, b) => b
            _ => 1
        , 4, 5, 6]  // notice the proper end to the indentation suite
```

> You should generally avoid using the match expression like that if you can,
> but it can sometimes prove unavoidable (within reason) and so it is considered
> possible even if not beautiful.

The second and last control expression is much simpler: the conditional
expression chooses its left branch if its condition is true and its right branch
if its condition is false (like an if statement).  They are written like so:

```whirlwind
    -v2 if v2 < 0 else v2
```

As you can see, they are fairly simple but occasionally useful.  The same idea
can expression in a slightly more verbose form by a match expression.

```whirlwind
    match v2 < 0 to
        true => -v2
        _ => v2
```
