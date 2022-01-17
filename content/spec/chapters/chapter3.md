# Type System

A **type** is a set of values together with the methods bound to those values.  
All types may have a *type label* which is the syntactic construct used to
reference them.

Chai is strongly and statically typed meaning all values have a fixed, static
type, and all types are determined at compile-time.  Furthermore, Chai does not
support type coercion of any form: all type conversions in Chai must be
performed explicitly.  To combat this rigidity, Chai uses a powerful type
inferencer to reduce the amount of conversions and type labels necessary in user
code.

The Chai type system consists of the following types:

- Primitive Types
- Function Types
- Tuple Types
- Reference Types
- Structure Types
- Algebraic Types
- Hybrid Types

This chapter will focus on primitive types, type properties, casting, and
inference.  Discussions of **composite types** such as tuple or structure types
will be reserved for their own chapters.

## Primitive Types

**Primitive types** are a set of types built-in to the language itself.  Their
implementations are kept completely opaque and may be system dependent.

### Integral Types

Chai provides 8 **integral types** which each represent a specific set of
integers. There are four sizes of integer: 8 bit, 16 bit, 32 bit, and 64 bit.
These sizes correspond to the width of the integer's representation. For each
size, there is a signed and an unsigned variant.  Integers are represented as
standard binary numbers using two's complement arithmetic to encoded signedness.

The type labels for integral types begin with an `i` for signed integers and a
`u` for unsigned integers followed by the bit size of the integer.  Below
is a listing of all the integral types:

    i8   # 8 bit,  signed integer
    i16  # 16 bit, signed integer
    i32  # 32 bit, signed integer
    i64  # 64 bit, signed integer

    u8   # 8 bit,  unsigned integer
    u16  # 16 bit, unsigned integer
    u32  # 32 bit, unsigned integer
    u64  # 64 bit, unsigned integer

### Floating-Point Types

Chai provides two **floating-point types** which are represented in accordance
with the IEEE-754 standard for floating-point numbers.  

The first type is a single-precision floating-point type with a 32 bit width
denoted as `f32`.  This type consists of the set of all 32-bit IEEE-754
floating-point numbers.

The second type is a double-precision floating-point type with a 64 bit width
denoted as `f64`.  This type consists of the set of all 64-bit IEEE-754
floating-point numbers.

These types include the special floating-point values for not-a-number (NaN),
negative zero, and infinity as outlined in the IEEE-754 standard.

### Complex Types

TODO

### Rational Types

TODO

### Strings

The **string type** stores a finite string of Unicode text.  Strings are
represented as a contiguous array of bytes and a 32 bit wide, unsigned integral
length.  The byte array is a finite sequence of UTF-8 encoded text encoding the
value of the string.

Strings do not require a null terminator nor does Chai treat null terminator
characters as ending or interrupting strings.

Strings use the type label `string`.

### Booleans

The **boolean type** represents a true/false (Boolean) value as an 8 bit wide
integral value.  The false value is encoded as a `0` and the true value is
encoded as a `1`.  The remaining bits are not used.

Booleans use the type label `bool`.

### Nothing

The **nothing type** is a primarily semantic type denoting a unit value: many
expressions may not yield any explicit value and so such expressions are said to
return "nothing" or "a unit value".

The nothing type has no well-defined machine level representation as no value
with a type of nothing exists at the machine level.  During compilation, all
values with a type of nothing must be *pruned* from output program: any value,
field, parameter, variable, or other storage mechanism types with nothing should
be completely removed from the user program.  This pruning is done under the
principle that there is only ever one possible value of a nothing type. and
therefore all nothing values are identical: all usages of the nothing type are
by definition completely execution invariant.

## Equality and Equivalency

**Type equality** represents the notion of two types being semantically
identical.  **Type equivalency** represents the notion of two types being
interchangable: a value of type A which equivalent to type B can be used as a
value of type B without any conversion.  Equality implies equivalency.

Two types are equal if they are exactly identical: they correspond to the same
definition.  By contrast, additional conditions for equivalency beyond equality
are defined for specific types.  Most types are only equivalent if they are
equal, but this is not always true.

## Nullability

A type is **nullable** if it defines a **null value**.  The null value of
a type refers to the "zero" or "default" state of a type.  

TODO

## Type Casting

## Type Inference

### Constraints and Contexts

### Assertions


