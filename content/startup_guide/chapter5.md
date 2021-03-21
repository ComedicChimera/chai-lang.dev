---
title: 'Pattern Matching'
weight: 5
---

## Tuples

A **tuple** is a fixed-length, ordered grouping of data.  It stores multiple elements
of different types in specific positions.  

You can create a tuple by enclosing a sequence of values separated by commas in parentheses.

    let pair = (1, 2)

`pair` is a tuple of two integers.  Here are some other examples of tuples:

    ("abc", 3)
    (3.14, 2.718, 1.414)
    ('q', 0, 'r', 1)

The type label of a tuple is structured similarly to its literal value.
It is a sequence of type labels enclosed in parentheses.  For example, the type label for
pair would be `(int, int)`.  

You can access the elements of tuples using a `.` followed by the **index** of the element
in the tuple.  Tuples are zero indexed so to get the first element, you would use the index
`0`. 

    let first = pair.0

{{< alert >}}This style of accessing is unique to tuples -- a more standard syntax is used for
collections.  However, in Whirlwind, tuples are treated more like data structures than like
collections of values.{{< /alert >}}

Note that the indices must be constant integer values.  You cannot use an expression or variable
to access a tuple since the compiler would not know the type of the value being accessed.

Finally, it is impossible to change the fields of a tuple individually once it has been created.

    pair.0 = 4 # ERROR

## Tuple Deconstruction

**Tuple deconstruction** (also called tuple unpacking) is the first kind of **pattern matching**
that we will look at in this chapter.  Pattern matching is a mechanism for efficiently extracting
the values of more complex data structures such as tuples.  Tuple deconstruction is a kind of
pattern matching used to access the elements of the tuple without having to resort to the `.`
syntax every time.  

We can use tuple deconstruction to unpack the elements of the tuple into variables. Let's see what
this looks like in practice:

    let (x, y) = pair

The variables `x` and `y` will now hold the values in their corresponding positions `pair`.  Notice
that the pattern of variables of the left side of the initializer matches up with the pattern of
values on the right side.  This is where the term pattern matching comes from.

We can also use pattern matching in assignment.

    let
        triple1 = (1, 2, 0),
        triple2 = (0, 5, 6)

    let (x, y, z) = triple1

    # -- snip --

    x, y, z = triple2

If we don't want to extract all the values of a tuple and or don't care about all of them, we can
use the special variable `_` to ignore values during pattern matching.

    a, _, b = triple2

Both `a` and `b` will be populate with their corresponding tuple values.  However, the middle value
is simply ignored. 

{{< alert theme="warning" >}}You cannot use `_` as a variable nor can you access it as a value.{{< /alert >}}

## The Match Expression

The **match expression** is a control flow expression used to facilitate sophisticated pattern matching
in expressions.  It is begun with the `match` keyword, followed by an argument, the `of` keyword, and series
of cases to match against.

The cases are made up of a pattern and an expression separated by the `->` operator.  Let's see what this
expression looks like so we can break it down a bit more.

    let v = (1, 1)

    let is_std_basis = match v to
        (1, 0) -> true
        (0, 1) -> true
        _ -> false

`v` is our argument and match statement compares it to several different values to see if it matches.  This
introduces us to an aspect of pattern matching not present before: value matching.  The pattern matching
mechanism will also allows us to compare values directly. 

Notice that the match expression ends with `_ -> false`.  This is a special case called the **default case**
-- the variable `_` in pattern matching will match any value and ignore it (somewhat similar to what it does
in tuple deconstruction).  This case is provided to make the match expression **exhaustive**.  This means that
it will yield a value for every possible input.  All match expressions must be exhaustive so that the compiler
can ensure a value is returned.  

Additionally, notice that the branches of this expression all yield values of the same type.  This is another
part of the match statement: all the branches must yield values equal or coercible to the same type.

We can put multiple patterns in one branch -- so to shorten the sample code above, we could write:

    let is_std_basis = match v to
        (1, 0), (0, 1) -> true
        _ -> false

Notice that we separated the patterns with commas.  Now either pattern will yield the same branch.

We can elso extract values into variables while we are matching cases.

    let expr = ("add", 1, 2)

    let result = match expr to
        ("add", a, b) -> a + b
        ("mul", a, b) -> a * b
        (_, a, _) -> a

Each branch extracts the values it needs from the tuple to perform the calculate while simultanously
checking a condition.  Notice that we didn't use the default case at the end because the match expression
was already exhaustive: the last case always matches if none of the others do.

{{< alert theme="danger" >}}If you extract named values, you cannot use multiple patterns in one branch.{{< /alert >}}

## The Match Statement

The **match statement** is a statement-level extension of the match expression we already know.  It is
the other main conditional control flow operator in Whirlwind (like the if-statement).

Like the match expression, it begins with the `match` keyword followed by a value to match and the keyword `of`.
After which, we enter a new indent level and begin enumerating a series of cases.  Unlike the match expression,
these cases are formal blocks and using the keyword `case` to denote their start.

    let message = ("greet", "Carl")

    match message to
        case ("greet", name) do
            println("Hello there,", name)
        case ("goodbye", name) do
            println("Goodbye,", name)
        case ("birthday", name) do
            println("Happy Birthday,", name)
        case _ do
            println("Unable to display message.")

Notice that the patterns inside the cases have the same ability as the patterns inside the match expression.

{{< alert theme="info" >}}Whirlwind doesn't have a `default` keyword -- we use pattern matching with `_` to denote
the default case.{{< /alert >}}

Match statements contain blocks inside yielding values and so do not have to be exhaustive.  If we had wanted to
elide the default case above, we could have.  

If you are coming from another language, you may be used to `break` being required at the end of match statements
(often called switch statements in other languages).  In Whirlwind, this is not only not required but also not
permitted.  The `break` statement is reserved exclusively for exiting loops and will work to that end even if
placed at the end of a match statement.

To replace the fallthrough behavior of the conventional switch statement, Whirlwind introduces a `fallthrough`
statement that can be used anywhere inside a case to automatically jump to the next one.

    let n = 5

    match n to
        case 0 do
            println("Additive identity and")
            fallthrough
        case 2 do
            println("Even")
        case 1 do
            println("Multiplicative identity and")
            fallthrough
        case 3, 5 do
            println("Odd and prime")

When using `fallthrough`, the next case may not extract a named value in its pattern.  You may also notice that
our last cases matches multiple values: `3` and `5`.  This behavior works identically to the equivalent in
match expressions, and, once again, both patterns may not extract values.

## Variable Expressions

On occasion, it is necessary to match against other variables or expressions involving them.  Due to the
syntax used for pattern matching, we cannot use these variables directly (as otherwise they will match as
values and shadow their outer variables).  Luckily, we can simply wrap the variables in parentheses to
denote that they are to be interpreted as an expression not as a pattern variable.

    let x = 3
    
    let result = match (0, 1) to
        ((x), _) -> "Case 1"
        (_, 4) -> "Case 2"
        _ -> "Case 3"

The first case only matches a tuple that has the value `3` in first position.  

For expressions, this process is easier.  Even if the expression contains variables, you do not need to wrap
it in parentheses since it is too complex to be mistaken as a pattern.

    let a = 4, b = 3

    match (12, 10) to
        case (a * b, y) do
            println(y)
        // -- snip --

The above code matches the first value `12` to the product of `a` and `b`.