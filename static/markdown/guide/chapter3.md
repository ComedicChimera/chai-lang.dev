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

    string // string of characters (unicode)

    char // single unicode code point (16 bits)

    byte // single byte of data (uses byte literals)

    bool // boolean type (8 bits)

You'll notice that all of the numeric types have both a signed and unsigned form.  This is always
designated by placing a `u` in front. `char` is special in that it is technically numeric, but
it defaults to being unsigned as this is it's normal use.  The designate it as signed, you have to place
an `s` in front.
