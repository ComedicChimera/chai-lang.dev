# Conditional Logic

Until now, our programs have been entirely linear; however, such programs are
rarely very useful since they lack the ability to make decisions based on input
or external conditions.  In this chapter, we will explore how our programs can
respond and change their progression based on conditions.

## Booleans

A **boolean type** is a type which only has two possible values: True and False.
It represents the logical idea of something having a "truth value".  For
example, the statement "4 is greater than 5" is false while the statement "the
sky is blue" is true.  

Boolean types have two literal values: `true` and `false`.  The type label for
boolean is a `bool`.

We rarely use booleans all on their own: more often, we want to test some
condition and have it produce a boolean that we can act on.  This is where
**relational operators** come in.  These operators take some number of values
and produce a boolean based on the "truth" of some relation between them.  

The simplest relational operator is the *equality operator* which compares two
values to see if they are equal.  It uses the `==` symbol.

    5 == 5      # => true
    6.7 == 4.5  # => false

As you can see, it works exactly like the arithmetic operators we already know,
but it gives a boolean instead of a number.  Like arithmetic operators, you
can also combine them with parentheses.

    (5 == 6) == false  # => true

The equality operator works for all types, but it can only compare two values of
the same type.  The equality operator also has an evil twin: the inequality
operator which does the exact opposite of the equality operator.  It uses the
symbol `!=`.

    2 != 2          # => false
    "abc" != "def"  # => true

The remaining four relational operators are: greater than (`>`), less than (`<`),
greater than or equal to (`>=`), and less than or equal to (`<=`).  These work
similarly to `==` and `!=` except the only work for numbers.

    5 < 6        # => true
    7.2 >= 7.2   # => true
    2e-2 > 8.92  # => false 
    86 <= 86.0   # => true

All of the relational operators are lower precedence than all the arithmetic
operators meaning you can embed arithmetic expressions as the operands of
relational operators without parentheses.

    5 * 6 - 2 < 7 * 8  # => true
    1/(1/4) == 4.0     # => true
    2 + 2 >= 5 * 2     # => false

### Multi-Comparison

TODO

### Logical Operators

TODO

## If, Elif, and Else

TODO

## Block Expressions

TODO