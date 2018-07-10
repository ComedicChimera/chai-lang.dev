## Variables and Constants
In programming, we need ways of storing and accessing data in memory, whether it
be an integer, a string, or a massive, complex object.  In Whirlwind,
we use something called a **variable** to accomplish this.

### Variables
A variable is an identifier that references some point in the computer's
memory.  The code below shows how to declare them.

    $x = 3;

This variable declaration has 2 parts: the **declaration** and the
**initialization**. The declaration is doe with the `$` operator followed
by some name, in this case `x`. The second part, initialization, is
where the value of a variable is set.  This performed with the
'=' operator, followed by an expression.

Once x is declared, it can be accessed from anywhere as shown in **Listing 1.7**.

#### Listing 1.7 - Variables in Action
    // notice that x is accessible from anywhere
    $y = x * 3; // 9
    $z = x + 4; // 7

    $a = z + x + y; // all 3 variables are now usable

You can also change the value of a variable through **assignment**.
This is done with the `=` operator, it uses the same syntax as initialization.

    x = 5; // x's value is now 5
    z = x * y; // y is still 9, but x is now 5, so z now equals 45

Variable is are **type immutable**.  This means that their type cannot
change.  Since x, y, z, and a are all integers by **type inference**,
you cannot set z equal to say a string.

    z = "String"; // TYPE ERROR; this will not compile

If you want to specify a variable's type at declaration, you can use a
**type extension**. Type extensions consist of the `:` operator, followed
by a type designator.  The code below shows how to utilize type
extensions for variables.

    $b: int = 4; // : int is the type extension

    $c: float = 5; // 5 will automatically be coerced into a float (5 -> 5.0)

    $d: bool = 5; // TYPE ERROR, types do not match

As you can see, the type extension always goes in between the
identifier and the initializer.  Additionally, the
type extension always take precedence over the type of the
initializer and so if they do not agree, as shown with `d`, you will
get a type error during compilation.

Finally, you can declare a variable without initializing it as long
as you specify a type. This is called **default initialization** and
The code below shows an example of default initialization.

    $var: int;

Unlike other languages, Whirlwind allows variables to be used without
initialization.  It automatically assumes a null (or default) value
for the variable.  In this case, `var` has a null value of 0 because it
an integer.  We will talk more about null and null values in future chapters.

    a = var; // a now equals 0, because var is 0

It is important to note that you must always give a type during
a declaration.  This type can be inferred from
the initializer or explicitly stated with a type extension, but
it must always be there.  Whirlwind will not allow you to declare
variables with no type.  For example, the code below is invalid.

    $invalidVar; // ERROR: unable to infer type of 'invalidVar'


### Constants
Constants are just like variables, except their value cannot be
changed.  They are declared using the `@` operator as opposed to
the `$` operator.

    @a = 5;

Unlike variables, constants must always be initialized.  You can provide
a type extension should you desire to, but there must always
be an initializer.

    @b: int; // ERROR: all constants must be initialized

Because constants cannot be changed, there are not valid on the left hand side
of assignment.

    $x: int;
    @a = 45;

    x = a; // valid, no constant is being mutated
    a = 20: // ERROR: constant value is changed

Constants cannot be assigned to under any context, even if the
value they are being assigned is the same as their current value.

    @constant = 15;
    $variable = 15;

    constant = variable; // invalid

### Constexprs
There is a third kind of constant in Whirlwind called a **constexpr**.
Constexpr stands for constant expression.  These are constants whose
default value is predictable at compile time.  They are initialized
using the `:=` operator as opposed to the `=` operator.

    @constexpr := 5 + 4;

Constexpr values are significantly faster at run time as their
value can be statically determined by the compiler.  The docs page
contains a document that outlines all valid constant expressions.

The reason constexpr's exist is not just because they are faster at
runtime, but also for other things that need to be compile time constant,
such as array bounds (which we will talk about in Chapter 2).

You will likely not use constexpr's very much, but it is important
to know of their existence, so you can avoid confusing errors later
down the line.

### Compound Declaration

Compound declaration refers to declaring multiple variables or constants
at once.  **Listing 1.8** show examples of compound declaration.

#### Listing 1.8 - Compound Declaration

    $(a, b, c): int; // all 3 variables are given a type of int

    $(x = 4, y, z = 3.2): float; // some values are initialized

    @(s = 'a', r = 'b'); // both values must be initalized since it is a constant

As you can see, compound declaration consists of using one of the two
declaration operators, followed by a set of identifiers wrapped in parentheses.

These identifiers can provide initializers or be given the overall type of
the declarative group.

In addition, one declarative group can contain variables (or constants) of multiple
types.

    @(a = 'a', pi = 3.14);

Similarly, variables (or constants) can specify their own type within the group as
well.

    $(x: int, y: float);

You can even provide an overriding type that all uninitialized values will
default to.

    $(b = 4, c: char, d, e): float; // b is int, c is char, and d & e are floats

### Compound Assignment

Just like declaration, you can also assign to multiple values at once.
This is done with **compound assignment** as looks like so:

    $(a = 4, b = 3);

    a, b = b, a; // a and b are now swapped

Assignment works by matching each side in order.
So in the example above, the first variable `a` matches
to the first expression that happens to contain the value
of `b` and so on.

You can do this with up to as many values are you want as long as
there is an equal number of a values on both sides.  Because of this,
the code below is invalid.

    $(x = 'a', y = 'b', z = 'c');

    x, y, = z, y, x; // invalid; too many values on right hand side

## Summary
That's it.  You finished Chapter 1. This was quite a large and diverse chapter.
You learned about  program structure, the main function, comments,
data types, variables, constants, constexprs, and assignment.
But, you have only begun your Whirlwind journey.
I know the information may seem daunting, but
don't worry.  After a while, you will get the hang of it.

Now let's move on to Chapter 2, where will cover some new data types and
cooler ways to assign values.