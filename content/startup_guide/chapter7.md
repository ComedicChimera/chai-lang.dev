---
title: "Error Handling"
weight: 7
---

## The Option Type

`Option<T>` is a builtin algebraic type that is used to manage boolean failure.
This kind of failure occurs when a value is either produced or not, and there is
no additional information to provide.

Algebraic types have their own chapter in this guide; however, this chapter should
serve a gentle introduction to them and their uses.  

The `Option<T>` type has two values: `Some` and `None`.  The `Some` value can also
store a value of type `T` whereas `None` stores no value.  `T` is a type parameter:
a type passed in that determines when generic form a type will assume -- in this case,
it is the type of the possible value returned.

{{< alert theme="tip" >}}All of these words will get more formal definitions later.{{< /alert >}}

    let op = Some(5)             // Type = Option<int>
    let op2: Option<bool> = None // Type = Option<bool>

As seen above, we initialize the value of `Some` by putting it in parentheses.  Notice
that when we use the `None` form, we need to give an explicit type since otherwise Whirlwind
can't infer the type of the type parameter.

It is possible to pattern match over algebraic types.  

    match op to
        case Some(x) do
            println(x)
        case None do
            println("No value.")

We can use this type to manage failure by return `Some` is we have a value to return and `None`
if we don't.

    func sqrt(n: double, k: int = 10) Option<double> do
        if n < 0 do
            return None

        let z = n
        for i in 1..k do
            z = z - (z * z - n) / (2 * z)

        return Some(z)

Now, if we pass in a negative number, we will simply get `None` back instead of causing an error.
Furthermore, this method of error handling forces us to deal with both the success and failure
states as we progress through our program.

    func quadratic_formula(a, b, c: double) Option<(double, double)>
        => match sqrt(b * b - 4 * a * c) to
            Some(r) => Some(((-b + r) / (2 * a), (-b - r) / (2 * a)))
            None => None

Notice that the "fallability" of the `sqrt` function pervades all of the functions that depend
on it (`quadratic formula` now also returns an `Option<T>` type).  It also worth noting that
we can use tuples to return multiple values even through an `Option` type.

## The Result Type

While `Option<T>` certainly covers a lot of cases when a function can fail, it also does not
have any way for us to return a descriptive error.  This is where the `Result<T>` type comes into
the picture.  It also has two values: `Ok` and `Err`.  The `Ok` value works just like `Some` --
it stores a value of type `T`.  However, `Err` also takes a value: a type of `Error` which is
a builtin type for representing an error.

    let r = Ok(4.31)
    let r2: Result<string> = Error(LogicError("Some error message"))

Notice that just like with `Option<T>` we need to provide a data type on the `Error` value since
Whirlwind once again can't infer the type parameter.  



