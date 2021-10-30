# Types and Operators

If we want to do a bit more than just print "Hello, world!", we will need to
engage in a brief discussion of Chai's type system.  Namely, Chai provides us
with several basic types:

| Label | Meaning |
| ----- | ------- |
| `i8` | 8-bit signed integer |
| `i16` | 16-bit signed integer |
| `i32` | 32-bit signed integer |
| `i64` | 64-bit signed integer |
| `u8` | 8-bit unsigned integer |
| `u16` | 16-bit unsigned integer |
| `u32` | 32-bit unsigned integer |
| `u64` | 64-bit unsigned integer |
| `f32` | 32-bit float |
| `f64` | 64-bit float |
| `string` | UTF-8 encoded string |
| `bool` | boolean |
| `rune` | 32-bit, unicode code point |
| `nothing` | empty/no value |

You will notice that Chai considers strings to be a primitive type and that it
operates on Unicode by default.  We will talk in more detail about strings in
a later section, but for now, let's focus on the numbers.

All integers and floats in Chai have fixed sizes and their size is apart of
their name.  On every platform, `i32` will be a 32 bit integer.  This tends to
ensure that your API has more standard behavior on all systems.

## Number Literals

In Chai, standard "integer" numbers can be used for *any* of the numeric types.
For example, if you write the number `42` in your program, it can be inferred to
be an integer or a float of any size depending on how it is used.  Note that
this does not mean that it doesn't have a single, concrete numeric type; rather,
the type inferencer applies no initial constraints outside of it being a number
to the literal.  This is important because Chai does NOT perform or support ANY
"type coercion" (ie. where the compiler implicitly converts between two types
without an explicit cast).  Thus, the type inferencer has to make up for that
loss by being more general with its initial constraints.

However, some literals will have more specific constraints applied to them.  For
example, the following are all floats:

    4.5
    0.141
    6.626e-34
    3E8

however, their size is not constrained.

Similarly, the following are all guaranteed to be some form of integer:

    5u   # unsigned
    67l  # long (64 bit)
    87ul # unsigned + long => `u64`

Note that if a type for a literal cannot be determined from context (which
happens surpisingly often), then the compiler will pick a sensible default:
generally one of the 32 or 64 bit forms.  If you want a specific type, then you
should state it explicitly.

## Arithmetic

Arithmetic in Chai works similarly to any other programming language.  Below are
the full list of builtin arithmetic operators.  These work for all numbers.

| Operator | Operation |
| -------- | --------- |
| `+` | Add two numbers |
| `-` | Subtract two numbers OR negate one number |
| `*` | Multiply two numbers |
| `/` | Divide two numbers and produce a floating point result |
| `//` | Divide two numbers and produce a floored, integer result |
| `%` | Find the remainder of a division operation |
| `**` | Raise a number to a non-negative, integer power |

Notice that there are two division operators in Chai: one for floating point
division and one for integer division.  This is to avoid random casts being
littered all over your program.  Note that both operators work for both kinds of
input (integers and floats).

Chai also supports parentheses and applies standard operator precedence rules
for arithmetic (ie. exponents, multiplication and division, addition and
subtraction -- performed left-to-right for ties of precedence).

    4 + 5              # => 9
    (65 * 0.8) // 2    # => 26
    0.5 ** 2           # => 0.25
    (3.14 + 2.72) * 64 # => 375.04
    5 % 3              # => 2
    -10 / (6 - 3)      # => -3.333...

## Type Casting

Type casting allows for the conversion of one type into another.  These are
performed using the `as` keyword.  All casts will fail at compile-time if the
cast is invalid.

    5.4 as i32      # => 5
    12  as f64      # => 12.0
    "string" as f64 # COMPILE ERROR

Note that casts can cause data to be lost during conversion (eg. `5.4` to `i32`
essentially floors it). 

## Booleans

**Booleans** are a common fundamental type in Chai that used to represent a
true/false value.  Their literals are `true` and `false`.

Several operators are defined on booleans, called **logical operators**:

| Operator | Operation |
| -------- | --------- |
| `&&` | Logical AND |
| `\|\|` | Logical OR |
| `!` | Logical NOT |

> Both logical AND and logical OR support short-circuit evaluation.

These operators behave as standard 
[boolean logic](https://en.wikipedia.org/wiki/Boolean_algebra) operators.

Several other operators are used to produce boolean values, called **comparison
operators**:

| Operator | Operation |
| -------- | --------- |
| `==` | True if both values are equal |
| `!=` | True if both values are not equal |
| `<` | True if the LHS is less than the RHS |
| `>` | True if the LHS is greater than the RHS |
| `<=` | True if the LHS is less than or equal to the RHS |
| `>=` | True if the LHS is greater than or equal to the RHS |

Both `==` and `!=` are defined for all values of the same type.  However, the
other comparison operators are only defined for numbers and runes by default.

Here are some examples of these operators:

    5 > 3           # => true
    "hi" == "hello" # => false
    7.6 <= -8.1     # => false      

## Rune Literals

**Runes** represent single unicode character.  They are used commonly in Chai are
represented with a single character (such as a letter) enclosed in single quotes.
They use the type label `rune`.

For example, a rune literal for the character `a` would be: `'a'`. 

Rune literals can contain multibyte unicode literals as well.

    'φ'
    '♣'

Runes also support several escape codes, which allow the inputting of special
characters such as newlines and carriage returns into literals.  They begin
with a backslash followed by one of the following codes:

| Code | Character |
| ---- | --------- |
| `a` | alert (legacy on most systems) |
| `b` | backspace |
| `f` | form feed |
| `n` | newline |
| `r` | carriage return |
| `t` | tab |
| `v` | vertical tab |
| `0` | null terminator |
| `"` | double quote |
| `'` | single quote |
| `\` | backslash

For example, to write a newline, one would use the following literal: `\n`.
