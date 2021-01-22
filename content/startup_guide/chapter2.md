---
title: "Basic Types"
weight: 2
---

{{< alert theme="info" >}}All code used in this chapter and beyond
is assumed to have the appropriate enclosing constructs (`println`
import, main function, etc.) unless otherwise stated.{{< /alert >}}

## Numeric Types

Whirlwind supports two kinds of numeric types: **integral** types and
**floating-point** types.  They have the following type labels:

    // Integral Types
    short   // 16-bit, signed integer
    ushort  // 16-bit, unsigned integer
    int     // 32-bit, signed integer
    uint    // 32-bit, unsigned integer
    long    // 64-bit, signed integer
    ulong   // 64-bit, unsigned integer

    // Floating Types
    float   // 32-bit, floating-point number
    double  // 64-bit, floating-point number

{{< alert theme="info" >}}The strings of text prefixed by `//`
are comments -- text that is ignored by the compiler.{{< /alert >}}

Here are some examples of various different numeric literals:

    45   // integer
    5.6  // floating-point
    5e7  // floating-point

All numeric types support the following basic arithmetic operations:

| Operator | Behavior |
| -------- | -------- |
| `+` | Adds two numbers |
| `-` | Subtracts two numbers |
| `*` | Multiplies two numbers |
| `/` | Divides two numbers and yields a floating-point number |
| `~/` | Divides two numbers and yields a integer number |
| `%` | Calculates the remainder of a division operation |
| `~*` | Raises a number to an integer power |

All of these operators are binary: they take two numbers and return
a single value.  Note that Whirlwind supports two division operators:
one that performs floating division and one that performs integer division.

Here are some examples of some of the operators:

    45 + 2   // => 47
    4 - .86  // => 3.14
    7 % 2    // => 1
    0.5 ~* 2 // => 0.25
    5 / 2    // => 2.5
    7 ~/ 2   // => 3

Finally, you can also use the unary `-` operator to negate a numeric value.

    -45 
    -(2 + 4)
    (3 + 4) * (7 ~^ 2)

## String and Rune Types

A **string** is a group of characters.  All strings in Whirlwind are UTF-8 encoded
and enclosed in double quotes.  Strings have the type label `string`.  

A **rune** is a single Unicode character, represented as a 32-bit integral value.
Runes have the type label `rune` and are enclosed in single quotes.  

Both runes and strings support the use of escape codes.  Here are some fairly
common and/or useful ones.

| Code | Character |
| ---- | --------- |
| `\n` | New-line |
| `\t` | Tab |
| `\b` | Backspace |
| `\\` | Backslash |
| `\"` | Escaped double quote |
| `\'` | Escaped single quote |

{{< alert theme="info" >}}Check the Language Specification *insert link* for a
more definitive list.{{< /alert >}}

Both string and runes support the use of the `+` operator.  For both strings
and combinations of strings and runes, the `+` operator will concatenate (join)
the items.  However, for combinations of exclusively runes, `+` will add their
numeric values.

    "Hello" + " there" // => "Hello there"
    "Yo" + '!'         // => "Yo!"
    'a' + '\n'         // => 'k'

## Boolean Types

A **boolean** is a true/false (Boolean) values.  There have the type label `bool`
and two literal values: `true` and `false`.  

There are two logical operators for comparing booleans: `&&` and `||`.  The
former performs logical AND and the latter performs logical OR.  Booleans
also support logical NOT with the `!` operator.

    true && false  // => false
    true || false  // => true
    false || false // => false
    true && true   // => true
    !true          // => false

Additionally, the numeric types support four ordinal comparison operators:
`>`, `<`, `>=`, `<=`.  The two operators with equal signs attached translate
as *greater than or equal to* and *less than or equal to* respectively.  Here
are some examples of them:

    5 > 4             // => true
    8 < 5.2           // => false
    4 + 5 >= 4 - 3.42 // => true
    7 <= 7            // => true

Finally, most types support direct comparison using the `==` and `!=` operators.
The former tests for equality and the latter tests for inequality.

    5 == 4         // => false
    'a' != 'b'     // => true
    "abc" == "abc" // => true

## Type Errors and Casting

If a certain operator is applied between types for which such an application
is undefined (ie. it makes no sense), then a type error is generated.  Specifically,
most operators have a small defined set of valid types of operands.  Here are
some examples of invalid combinations of operators, all of which will generate
an error at compile-time.

    5 > "abc" 
    !'a'
    5 == "bc" // can't compare items of different types
    54 and 76

It sometimes convenient or necessary to convert one type into another.  Some
conversions happen implicitly: from ints to floats for example.  These kinds
of conversions are called **coercions**.  However, some more complex conversions
such as from signed integers to unsigned integers, must be converted explicitly
using a **type cast**.  Type casts use the keyword `as` followed by the type
to convert to.  Here are some example casts:

    5.5 as int  // => 5
    "a" as rune // => 'a'
    1 as bool   // => true

Not all casts are valid and some will fail at compile-time or at run-time (causing
a runtime panic).  


