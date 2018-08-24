## Constants
Constants are just like variables, except their value cannot be
changed.  This means you can use them to represent unchanging values
like `pi` or to ensure that some value isn't modified.

### Constant Syntax

They are declared using the `@` operator as opposed to
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

There is a another kind of constant in Whirlwind called a **constexpr**.
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

### Enumerated Constants

Sometimes you want initialize a bunch of related constant values,
like colors for example.  In this case, it is common
to simply give them a series of incremental values
moving upward. For example, green is 0, red is 1, ect.
Constants that follow this pattern are called
**enumerated constants**.  

Since it is would be extremely tedious initialize the constants
with values like this, and even more precarious to change them
using the default constant syntax, Whirlwind offers a simpler
syntax to do this quickly.

    @{
        GREEN
        RED
        BLUE
    };

That above is an enumerated constant.  Each value has
an incremental value starting at 0 and going up to 2.
Notice that the enumerated constant syntax uses braces
instead of parentheses and that they are **no commas**
between values.

It is also common to want to have the enumerated constants
start from a given value an move up. Whirlwind enables this
by enabling to specify a **start value**.  You do this by
setting the first variable equal to value
you want to start at.

    @{
        THREE = 3
        FOUR
        FIVE
        SIX
    };

This will make sure the counter starts at 3.

However, the ability to specify a start value
creates a unique problem. The value
at the beginning of the enumeration
is not necessarily compile-time determinable.

Because of this, the compiler can't auto initialize
values when generating your code. So the result code
will end up equivalent to the second block as opposed to
the first.

    // first enum const
    @(
        GREEN = 0,
        RED = 1,
        BLUE = 2
    );

    // second version
    @(
        THREE = someVal, // someVal represents a random value
        FOUR = someVal + 1,
        FIVE = someVal + 2,
        SIX = someVal + 3
    );

The second block is noticably more verbose and computationally
intensive than the first block.  However, this can be avoided
by using the constexpr initializer in the first block.

    @{
        THREE := 3
        FOUR
        FIVE
        SIX
    };

Now the resulting code looks how you would expect:

    @(
        THREE = 3,
        FOUR = 4,
        FIVE = 5,
        SIX = 6
    );

Whenever possible, try to use the constepxr initializer when
working with enumerated constants.