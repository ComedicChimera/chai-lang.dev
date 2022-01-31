# The Basics

This chapter is going to cover an assortment of basic topics in the language
that are just essential to do anything remotely useful: this will be our
"lightning round" so to speak.

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

## Numbers and Types

We now arrive at the ever-boring and often confusing discussion of the
representation of numbers.  Given that computers were spawned forth from
calculators, it only makes sense that numbers would be a critical part of any
programming language.  Unfortunately, numbers in many programming languages can
be a bit arcane and while Chai tries to avoid this confusion as much as
possible, there are still some common pitfalls worth mentioning.

Before we can discuss numbers, we need to introduce the concept of a **type**.
In short, a type (also called a *data type*) is way of classifying data. See, to
computers, all data is simply stored as binary, 1s and 0s.  Types give us a way
to give those binary digits special meaning based on how we want to use them.
Formally, a type is a set of possible values and operations that can be applied
to and between those values.

Chai is a *strongly* and *statically* typed programming language.  
**Strong typing** means that generally types don't mix: if a value is expected
to be of a specific type, then it must be exactly that type.  It also means that
the compiler will avoid doing *type coercion*: implicitly converting a value of
one type to another.  **Static typing** means that everything must have a
well-defined, known type at compile-time.  Furthermore, any value can only have
*one* type that cannot change during program execution.  This is distinct from
interpreted languages like Python or Javascript where types change all the time,
and it can be impossible to know the type of something until you run the
program.

This discussion of typing brings us to the set of **numeric types**: types used
for numbers.  Generally, these types split into two groups: integral (used for
representing positive and negative integers) and floating-point (used for
representing all real numbers).  Each type has a **type label** which is a
special keyword or expression used to label values as being of that type.  Each
type also has a fixed bit size: a value of that type will occupy precisely that
number of bits. 

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

TODO: floats, literals, and operators

## Variables

## Command-Line Input

