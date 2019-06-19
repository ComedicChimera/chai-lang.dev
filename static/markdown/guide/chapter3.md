# Variables and Data Types

Variables are used to store information under a given name.  Variables
in Whirlwind come in two distinct flavors **variables** and **constants**.
Variables are declared with `let` and constants with `const`.

    let x = 20;

This code declares a variable `x` with a value of 20.  The main difference
between variables and constants is that variables can be **assigned** to
and mutated after initialization and constants cannot.  

    let y = 45;
    const pi = 3.141592;

    y = 3; // ok
    pi *= 2; // ERROR

You'll notice that you can combine normal mathematical operators such as the `*`
(multiplication) operator seen below with the assignment operator as shorthand for the
equivalent code.

    pi = pi * 2;

It is also important to note that variables in Whirlwind can be shadowed such that
a variable can be declared in one, higher scope and its name can be reused and redeclared in
the lower scope.

    let t = 0;

    // subscope
    {
        let t = "asd"; // t is shadowed
    }

This can be very useful in many situations, but it cannot also be harmful as once a variable is
shadowed, it's value cannot be accessed by that name in the scope
in which it is shadowed.

## Type Safety

Whirlwind is a type safe language.  This means that the **data type** of the variables cannot
be mutated, regardless of which type they are.  For example, our variable `y` is a type of
`uint` (unsigned integer) and that cannot change.  You can specify the data type of the variable
by using a type extension.

    let a: int = 5;

Even though 5 is an unsigned integer, because a is marked as a signed integer and unsigned types 
can be coerced to signed types, a has a type of signed integer.  They are several types
central types in Whirlwind and all of them have **type aliases** associated with them as shown
below.

    int, uint // signed and unsigned integer (32 bits)

    long, ulong // signed and unsigned double size integer (64 bits)

    float, ufloat // signed and unsigned floating point number (32 bits)

    double, udouble // signed and unsigned double floating point (64 bits)

    str // string of characters (unicode)

    char // single unicode code point (16 bits)

    byte // single byte of data (uses byte literals)

    bool // boolean type (8 bits)

You'll notice that all of the numeric types have both a signed and unsigned form.  This is always
designated by placing a `u` in front.

> `char` is special in that it is technically numeric, but it defaults to being unsigned as this is it's normal use.
> To designate it as signed, you have to place an `s` in front.

## Null Initialization

All variables in Whirlwind are initialized on declaration.  However, what they are initialized as
is determined by the declaration.  So far, we have been initializing variables to definite values.
But, if a type extension is provided a variable can be declared without an initializer.  This is called
**null initialization**.

    let x: int;

Because of null initialization, x is initialized to a value of 0, which is its **null value**.  No variable
is ever left uninitialized and because of this, you can declare variables and use them without specifying a value.

It is important to note that a variable must have a data type, whether it is given by a type extension or inferred from
the type of an initializer.  Be careful to keep in mind that a type extension always trumps an intializer and that they must
agree.  Because of this, the following code is invalid.

    let y; // ERROR - unable to infer data type of y

These rules apply equally for variables and constants.

## Constancy Propagation

We have spent a lot of time talking about variables, but we need to quickly discuss constants.  Constancy
is unique in Whirlwind because it tied both to name and to value.  In Whirlwind, constancy is permanent
and inviolable.  No matter but what means you mutate a constant value, whether it is by reference, pointer or
any other means.

    const p = 4; 

    let r = &p; // create a reference to p

    *r = 2; // ERROR - unable to mutate a constant value

    let a = 42;
    r = &a; // r is not constant so this is ok

    *r = 3; // passes because a is non-constant

I know we have not covered much of the concepts used above, but it is simply to illustrate the point
of just how long constancy lasts.  For experienced programmers, this should be a wake up call
to you to be careful with constants.

## Multi Declaration

Multi-declaration is just a way of declaring multiple variables at once.  It looks like so.

    let (a = 0, b = 'a', c = 0.3);

Notice that each variable is separated by commas and that the variables have different data types.
You can also just use a type extension instead of an initializer if you want to.

    let (a: int, b = 'a', c: double);

You can also define an overarching type extension to fill in any unspecified variables.

    const (a: int, b = 'a', c, d): double;

The variables `c` and `d` are now both doubles.
