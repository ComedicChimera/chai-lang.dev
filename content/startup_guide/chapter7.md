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
    let r2: Result<string> = Err("Message")

Notice that just like with `Option<T>` we need to provide a data type on the `Error` value since
Whirlwind once again can't infer the type parameter.  

We can then return more informative error messages as we need them.

    func quadratic_formula(a, b, c: double) Result<(double, double)> do
        if a == 0 do
            return Err("`a` cannot be zero")

        match sqrt(b * b - 4 * a * c) to
            case Some(r) do
                return Ok(((-b + r) / (2 * a), (-b - r) / (2 * a)))
            case None do
                return Err("No real roots")

## Context Managers

A **context manager** is an advanced control flow construct that has a wide variety of uses.
Fundamentally, it does exactly what the name implies: it manages context.  The bigger question is:
what exactly does that mean?  Well, think of it this way, in certain parts of our program, we know
certain things to be true.  For example, inside an if-statement, we know the condition of that
if-statement is true.  Context managers help us to generalize this idea.

{{< alert theme="info" >}}Context managers can do a lot more than we are going to see in this
section -- they are also great for managing resources that require some amount of setup and
cleanup.{{< /alert >}}

One common use case for context managers is managing failure.  Let's consider a simple example.
Say we were writing a web applications, and we needed to make a series of calls to retrieve and
process some data from the web.  For this example, we will say that we need to use 3 functions to
do this: `fetch_data`, `validate_data`, and `extract_value`.  Each of these functions returns
a `Result` type.  Only using what we know now, our code would propably look something like this:

    func load_data(url: string) Result<string>
        => match fetch_data(url) to
            Ok(data) => match validate_data(data) to
                Ok(validated_data) => extract_value(validated_data)
                Err(e) => Err(e) // might not be Result<string>::Err, so we need to reconstruct it
            Err(e) => Err(e)

Now that is a bit of an ugly mess isn't it?  This is one of the biggest complaints against languages
like Whirlwind that force users to manually handle every error: your code gets long and ugly.  Even
an experience Whirlwind developer would struggle to read that all the way through the first time. 

However, there is, of course, a solution: the context manager.  We are going to first restructure
code above using one and then break down how it actually works.

    func load_data(url: string) Result<string>
        with 
            data <- fetch_data(url)
            validated_data <- validate_data(data)
            value <- extract_value(validated_data)
        do
            return Ok(value)
        else match Err(e) do
            return Err(e)

That code, while a few lines longer, is much cleaner and clearer.  So how does it work?  Well,
the `with` keyword begins our context manager.  Then, we create a series of declarations
using the extract (`<-`) operator.  Each declaration attempts to pull the value out of the `Result`
and fails (yielding the failing Result) if it can't.  The context manager chains these declarations
together so that each one progresses one after the other.  If they all succeed, the body runs and
the extracted values are made available.  On the other hand, if one of the declarations fails,
the result is piped into the `else` clause at the end of the context manager.

Don't worry if all that didn't completely make sense.  Context managers are pretty challenging and
have a lot more depth than we are covering here -- thus why it feels like some things aren't being
fully explained.  The process of chaining and accumulating values in error case is quite complex
and leverages a fearsome idea from functional programming: the monad.  Even mentioning such a term
in a guide like this is enough to turn some newer programmers away in heartbeat.  The good news is
unlike with some languages like Haskell, you don't really need to understand monads to get by in
Whirlwind.  Sure, understanding them is very helpful (and they are quite a simple and beautiful
concept when fully comprehended) but ultimately not necessary (so much so that we won't mention them
again throughout the rest of this guide).  

Simply remembering and being able to reproduce the pattern shown in the example code above should be
enough to get you started with context managers.  You can confront their intricacies at a later
date when you are more comfortable with the language and more ready to go down the functional
programming rabbit hole.
