# Numbers and Types

We now arrive at the ever-boring and often confusing discussion of the
representation of numbers.  Given that computers were spawned forth from
calculators, it only makes sense that numbers would be a critical part of any
programming language.  Unfortunately, numbers in many programming languages can
be a bit arcane and while Chai tries to avoid this confusion as much as
possible, there are still some common pitfalls worth mentioning.

However, before we begin with numbers, there is one small topic, we need to
discuss first.

## Comments

A **comment** is piece of source text that will be ignored by the compiler. They
are primarily useful for documentation, although we will be using them more as
annotations in this guide.  There is a whole lot of debate and "programmer
philosophizing" about how to write good comments and when they should be used:
if you want to hear my take on it, go read the
[Chai Style Guide](/docs/style-guide).

From a practical standpoint, comments come in two flavors: **line comments**
and **block comments**.  Line comments begin with a `#` and span until the
end of the line.

    # This is a line comment

The other variety of comment is enclosed between a `#{` and a `}#`.  All
the text between those two pairs of characters will be ignored.  These comments
can span multiple lines or be squeezed into the middle of line.

    #{ This
    is a
    block
    comment }#

    code #{ This is also a block comment }# more code

Notably, comments don't "nest": a comment including inside another comment has
no meaning beyond simply being text inside that comment.  
## The Type System

Before we can discuss numbers, we need to introduce the concept of a **type**.
In short, a type (also called a *data type*) is way of classifying data. See, to
computers, all data is simply stored as binary, 1s and 0s.  Types give us a way
to give those binary digits special meaning based on how we want to use them.
Formally, a type is a set of possible values and operations that can be applied
to and between those values.

Chai is a *strongly* and *statically* typed programming language.  **Strong typing** 
means that generally types don't mix: if a value is expected to be of a specific
type, then it must be exactly that type.  It also means that the compiler will
avoid doing *type coercion*: implicitly converting a value of one type to
another.  **Static typing** means that everything must have a well-defined,
known type at compile-time.  Furthermore, any value can only have *one* type
that cannot change during program execution.  This is distinct from interpreted
languages like Python or Javascript where types change all the time, and it can
be impossible to know the type of something until you run the program.

This discussion of typing brings us to the set of **numeric types**: types used
for numbers.  Generally, these types split into two groups: integral (used for
representing positive and negative integers) and floating-point (used for
representing all real numbers).  Each type has a **type label** which is a
special keyword or expression used to label values as being of that type.  Each
type also has a fixed bit size: a value of that type will occupy precisely that
number of bits. 

## Integral Types

For integers, there are a total of 8 types: a signed and an unsigned variant for
each of four sizes.  **Signedness** refers to whether or not the type can be
negative.  Furthermore, in combination with the bit size, it helps to determine
the maximum value the type can store.  To understand what this means, let's
consider a decimal number like `336`.  The number of bits in a binary number is
analagous to the number of digits in a decimal number.  So a 3-digit "decimal
type" could store a maximum value of `999`.  A similar logic is used for binary
integers: an 8 bit, unsigned binary integer can store a maximum value of 2^8 - 1
= 255 (mathematically analagous to the 3 digit decimal number storing a maximum
value of 10^3 - 1).  Finally, if the number is signed, then a bit used to store
the sign (0 for positive, 1 for negative).  So the maximum value of the number
is based on a maximum one bit less (2^7 - 1 = 127) and has a minimum value of
-128 (since 0 is considered positive so the negatives have a little extra room).
With all that annoying garbage out of the way, here are the types:

    # unsigned: 
    u8  # 8 bit
    u16 # 16 bit
    u32 # 32 bit
    u64 # 64 bit

    # signed
    i8  # 8 bit
    i16 # 16 bit
    i32 # 32 bit
    i64 # 64 bit

The keywords on the left are the type labels for those integers.

## Floating-Point Types

Chai provides only two floating-point types whose labels are `f32` and `f64`
respectively.  As with integers, the number at the end corresponds to the size:
32-bits (single-precision) and 64-bits (double-precision).  

The meaning of that size is a bit more complex in the case of floating-point
types.  Floating-point types and floating-point math in general is compliant
with and specified by a standard known as IEEE-754.  If you have a whole lot of
time to kill and perhaps take delight in reading dull, dense engineering jargon,
then you are welcome to read the standard yourself.  If not, then I will try to
hit the highlights here for you.  I should also note that I am going to use the
term "float" to refer to floating-point values from here on out.  I mention this
because to many C, C++, Java programmers, the term float corresponds to
specifically the 32-bit variety of floating-point numbers.  However, since Chai
simply uses sizes to denote types, I will consider myself free to use the term
"float" to mean all kinds of floating-point values.

The most basic principle is the, unlike integers, while floats do have maximum
and minimum possible values, these values are so astronomically large and small
that there are very rarely worth consideration.  Instead, they have a
*precision*: it determines how accurately they can represent a value.  In simple
terms, a float is essentially represented as
[scientific notation](https://en.wikipedia.org/wiki/Scientific_notation).  There
is an exponent which occupies a finite number of bits and a base (or as us
Computer Science nerds refer to it, the "mantissa") which also occupies a finite
number of bits.  As your values get larger or smaller, that exponent gets larger
and smaller and the number of actual "digits" in number itself increases (eg.
10^10 has 11 digits). However, floats can only represent so many of those digits
as the numbers grow larger and smaller.  For example, a float might only be to
represent less than a 20 of the digits of PI because it runs out digits to store
into.  This is what we mean by precision.

Another important detail about floats is that they are always signed: there is
no such thing as an unsigned float.

The most notable and dangerous quirk about floats in day to day use has to do
with arithmetic which we will get to shortly.

## Literals

A **literal** is a lexical (textual) representation of a specific value.  For
example, the text `12` is a literal representation of the number twelve.

Chai provides several different kinds of literals for representing numbers each
of which has an association with a certain subset of numeric types. What this
means is that some literals will only be usable as certain types of numbers.
For example, the literal `1.2` can only be a float since integers can't have
decimals. 

Let's start with the simplest kind of literals: *number literals*.  There
are simply whole numbers:

    12
    42
    8
    0
    12765
    100_000  
    2_345_990_000_123

These literals can be used for any numeric type but default to integers.  What
this means is that although they can be used to indicate both floats and
integers, in the absence of other information, the compiler will assume they are
integers, specifically an `i64`.  (TODO: update with final decision on this
topic).

Notice that you can insert underscores into the literals to help make them
easier to read.  This is true with all numeric literals.

The next kind of literals are *float literals*.  There is a bit more variety
with these so let's just see some examples:

    # Decimal Notation
    3.141592
    567.89
    100_000.567_897
    0.42

    # Scientific Notation
    1e9
    6.626e-34
    12.2E3
    1.602E-19
    10.024e12
    100E4
    11_234e-1

All these literals can only be floats and will default to the type `f64`.

Finally, there are *integral literals* which only correspond to integer types.
These literals are primarily used for notating numbers in other bases.  Chai
supports base 2 (binary), base 8 (octal), and base 16 (hexadecimal).  Each of
these literals begins with a specific prefix and has a specific set of digits
that can be used in them.

| Base | Prefix | Digits |
| ---- | ------ | ------ |
| 2 | 0b | `0`, `1` |
| 8 | 0o | `0` - `7` |
| 16 | 0x | `0` - `9`, `a` - `f`, `A` - `F` |

Here are some examples of literals in these bases:

    0b101010
    0xff
    0o172
    0xAe23
    0b110
    0o23456

These will default to the type `i64`.

<guide-exercise label="2.1"></guide-exercise>

## Arithmetic

Chai supports several arithmetic operations by default including:

| Operator | Operation |
| -------- | --------- |
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `%` | Modulo (remainder of a division operation) |
| `**` | Raise to a Power |

All these operations work on all numeric types but expect both their operands to
be of the same type.  They also return the same type as their operands.  The operators
also respect precedence (eg. multiplication before addition).  When two operations
are of equal precedence, they are executed left to right.  Here is brief listing
of operators in order from highest to lowest precedence.

1. `**`
2. `*`, `/`, `%`
3. `+`, `-`

Here are some examples of this arithmetic operators in action:

    1 + 1       # => 2
    2 - 5 * 6   # => -28
    4.5 / 2     # => 2.25
    5 % 2       # => 1
    78 + 23 - 4 # => 97

You can also use parentheses to group terms and elevate the precedence of
sub-expressions.

    (4 + 5) * (7 - 2)       # => 45
    3 * (72 / 6)            # => 36
    (3 + 6 * 8) * (33 - 27) # => 306

You can use a `-` sign as a prefix to an expression to negate the result.

    -4 * 5          # => -20
    78 + -5         # => 73
    -(7 + 8) * 2    # => 30
    -4 - -6.2       # => 2.2
    -7 * -(3 + 2.5) # => 38.5

<guide-exercise label="2.2"></guide-exercise>

There are some notable caveats worth mentioning.  Firstly, the division operator
works in two modes depending on the types of its operands: integer division and
floating-point division.  Integer division returns the largest whole number
quotient of the two numbers.  Floating-point division performs fully accurate
decimal division. 

    7 / 2   # => 3   (integer division)
    7.0 / 2 # => 3.5 (float division)

    8 / 12   # => 0    (integer division)
    8 / 12.0 # => 0.75 (float division)

Secondly, the power operator only works for non-negative integer powers by
default. 

    5 ** 2   # => 25
    0.5 ** 3 # => 0.125
    3 ** 0.5 # ERROR

You can however import the additional functionality for fractional and negative
exponents by placing the following line at the top of your file:

    import (**) from math

This is called an *operator import*: we will take a look at them in a much later
chapter.  With this import, you can now do operations like the following:

    2 ** 0.5 # => 1.41421356
    3 ** -1  # => 0.33333333    

Finally, it is worth noting that arithmetic involving floating-point numbers can
often be a bit strange.  For example, consider the simple addition below:

    0.1 + 0.2

Intuitively, we would expect this to yield `0.3`.  But, as any seasoned programmer
will tell you, this is not the case.  Instead, when you perform this addition,
the following is produced.

    0.1 + 0.2  # => 0.30000000000000004

This is called a **floating-point rounding error**.  It is a systemic problem
with how computers represent decimal numbers according to the floating-point
standard. There really isn't anyway to prevent it other than to be aware that it
exists and design your code with these kinds of errors in mind.  Feel free to
read up on floating-point errors if you would like to know more, but for now, we
will conclude our discussion of arithmetic.
