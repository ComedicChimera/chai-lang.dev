---
title: 'Variables'
weight: 3
---

## Mutable Variables

We use the `let` keyword in Whirlwind to create a **mutable** variable.
Simply place the name of the variable you want to create after the
keyword followed by an **initializer**.

    let x = 0

This variable can then be used in code by name:

    let y = x + 2

Whirlwind implicitly infers the type of variables based on the right-hand
side.  You can explicitly specify the type using a **type specifier**.

    let pi: float = 3.14

If you include a type specifier, you don't need to provide an initializer.

{{< alert theme="warning" >}}This is only true for nullable types *insert link*.
Most of the types in Whirlwind (inc. all the ones we currently know) are nullable
so this is rarely an issue. Nullability is not discussed much in this guide.
{{< /alert >}}

    let z: string

You can declare multiple variables at once using commas.

    let name = "Bob", age = 50

You can also declare variables in an indented block.

    let
        var1 = 23
        var2 = 6.7
        var3 = 'h'

## Assignment

Mutable variables can be **assigned** to -- this is what makes them mutable.
We use the `=` to assign to values in Whirlwind.  Assignment changes the
previous value of the variable.

    var1 = 42

You can assign to multiple variables at once. 

    name, age = "Emily", 32

Notice that each variable on the left-hand side corresponds to the value at
its position at right-hand side.  

In assignment, the right side is fully evaluated before the assignment occurs.
This means that you can swap the values of two or more variables using
multi-assignment.

    x, y = y, x

You can only assign a value of a type that is either equal to or coercible to
the type of the variable being assigned to.  If your types do not match, you
will get an error.

    name = 45 # TYPE ERROR

Whirlwind also supports many **compound assignment** operators.  A compound
assignment operator is an operator that performs some operation between the
variable and the value being assigned and then stores the result into the variable.

    # Expanded Form
    x = x + 2

    # Compound Form
    x += 2

Those two statements are equivalent: `+=` is the compound assignment operator.

## Constants

A **constant** is a variable that can't be mutated -- they are immutable variables.
They are declared the same as variables -- the only difference being that they use
the `const` keyword for declaration instead of `let`.

    const e = 2.718

All of the same declaration variations that apply to mutable variables also apply
to constants:

    const a = 0, b = 1

    const
        c = 5.2
        d = "test"
        v = 42

    const x: uint = 2

The only difference being that you cannot assign to constants -- once declared their
value is set for the remainder of their scope.

    e = 3.14  # ERROR
