---
title: "Algebraic Types"
weight: 9
---

## Case Study: Expression Type

We find that the best way to introduce the concept of an algebraic type to those who
are unfamiliar with the concept is use a case study to demonstrate their usage.  In
this section, we will write a simple expression type that we can use to represent
and evaluate mathematical expressions.

Put simply, an **algebraic type** is a composite type made up of multiple distinct,
named **instances** that may store one or more values.  The type can assume any one of
its different instances at any time.  This can be useful for variety of things; however,
the case of evaluating mathematical expressions demonstrates their power perfectly.

A mathematical expression can be composed of many different parts.  For example,
the expression `(2 * 3) + 5` is first composed of a product of two values followed
by a sum of that product and another value.  Notice that this structure is inherently
recursive and that it seems to be made up of multiple forms: a sum form, a product
form, and a value form.  This is exactly the situation in which one should employ an
algebraic type.  

We are going to define the `Expr` type to represent our expression.  Algebraic types,
like struct types, begin with the `type` keyword followed by a name.  However, instead
of curly braces, we indent and enumerate out the various possible instances of the type
beginning with `|` and separated by newlines.

    type Expr
        | Value(double)
        | Add(Expr, Expr)
        | Mul(Expr, Expr)

Our `Expr` type has three instances: `Value`, `Add`, and `Mul`.  All of them store values
and represent the corresponding operations. Much like our `Option<T>` type, we can refer
to these instance directly, as values, in our program.   For example, if we wanted to
represent the expression `(2 * 3) + 5` using this type, we could write the following:

    let expr = Add(Mul(Value(2), Value(3)), Value(5))

As you might guess, the variable `expr` has a type of `Expr`.  Note that all of the
instances of `Expr` are considered to be of the same type -- we can assign any one of the
three instances to `expr` without causing any errors.

Now, let's put our `Expr` type to good use and write a simple function that will evaluate
any expression.  To do this, we need to use pattern matching to extract specific instances
and evaluate them.  Here is what such an `evaluate` function would look like:

    func evaluate(expr: Expr) double
        => match expr to
            Value(x) => x
            Add(a, b) => evaluate(a) + evaluate(b)
            Mul(a, b) => evaluate(a) * evaluate(b)

Notice that we reference the instance we are matching against as part of the pattern and that
we need to call `evaluate` recursively in our `Add` and `Mul` cases -- our data structure is
inately recursive, and we need to handle that recursion.  

{{< alert theme="info" >}}Our match is exhaustive since we covered every possible instance and
combination of values in those instances.{{< /alert >}}

If we run the `evaluate` function on `expr`, the result is `11` which is exactly what we
expect.  However, `evaluate` will work on any expression built up of sums and products.

    // (7 + 1.2) * (5.1 * -2.3) => -96.186
    evaluate(Mul(Add(Value(7), Value(1.2)), Mul(Value(5.1), Value(-2.3))))

There are two key takeaways from this "Case Study".  The first is that algebraic types are
useful in situations where you a have a finite set of different states or values that you need
to represent.  The second is that we can use pattern matching to unpack and manipulate algebraic
types.  

## Closed Algebraic Types

All the types we have looked at so far have been **open** meaning their instances are readily
accessible in the global namespace.  However, this is not always desireable.  Consider the
example of an **enum**, a special case of algebraic types where none of the instances take values.
It is rare that we want such an enum to expose its instances globally.  Thus, it would be
advantageous for us to use a **closed** algebraic type -- whose instances can only be accessed
through the parent explicitly instead.

To demonstrate this point, let's consider the simple enum `Color` below.

    type Color
        | Red
        | Blue
        | Green

By default, `Color` is defined as an open type, meaning all of the instances: `Red`, `Blue`, and
`Green` will be defined globally.  However, this is not really what we are looking for so instead
we can mark `Color` has closed using the `closed` keyword.

    closed type Color
        | Red
        | Blue
        | Green

Now, we need to explicitly access the values of `Color` using the `::` syntax in order to use them.

    let c = Color::Red

    c = Color::Blue

    c = Green // ERROR

{{< alert theme="info">}}We can also use this syntax with open algebraic types although it is rarely
necessary.{{< /alert >}}

## The Test-Match Expression

The **test-match expression** is a general pattern matching construct used to test if a type matches
a given pattern.  This expression is a most commonly used with algebraic types which why we waited
to demonstrate it.  It is constructed using the `match` keyword separating two values -- the value
and the pattern.  The expression will evaluate to `true` if the match succeeds (and all variables will
be extracted as necessary) and false if it doesn't.  This is used to match against a single case as
opposed to multiple.

Let us consider a custom Option type just for `int` called `OptionInt`.

{{< alert theme="info" >}}We introduce this new type for demonstration purposes -- the regular `Option<T>`
can support any type of value.{{< /alert >}}

closed type OptionInt 
    | Some(int)
    | None

Notice that we mark this type as closed so as not to conflict with `Option`.  Now, let's say we
only cared about whether or not the value was `None`.  We could check this single case using the test-match
expression.

    if val match OptionInt::None do
        ...

If `val` (which is assumed to be type `OptionInt`) is `None`, then the pattern will match, and we will
proceed. 

Conversely, we can do a similar test-match against the `Some` instance as well.

    if val match OptionInt::Some(v) do
        // `v` is visible here
        ...
