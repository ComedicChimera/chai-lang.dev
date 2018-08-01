## Data Types

A **data type** is used to classify how a certain size of data in
memory should be interpreted.  Whirlwind has a number of different simple
types.

Each type has it's own specific **type designator** which is used for
marking different items as a specific type.  In **Listing 1.5**, you will
see each type designator next to a comment explaining the type is
designates.

#### Listing 1.5 - Type Designators

    int // the integer type, a positive or negative whole number

    float // the floating point type, a positive or negative decimal value

    char // the character type, a single unicode character

    str // the string type, a string of characters

    long // the long type, a larger integer that can store up to a much higher value

    double // the double type, a larger float that can store up to a much higher value

    byte // the byte type, a single byte of data

    bool // the boolean type, a true/false value

In addition to a type designator, each type also has a set of **literals**
associated with it.  A literal is just a symbol representing some
constant, immutable value such as `7` or `3.14`. The code in
**Listing 1.6** contains examples of literals with each type.

#### Listing 1.6 - Literals

    // integer, long (same literal)
    1 34 752

    // float, double
    3.14 56.0 4.1244

    // char
    'a' '\n' '😂' 'Á'

    // string
    "hi" "john" "ice cream" "\n\t" "你好，世界"

    // byte
    0b10101 /* binary literal */  0x13FFA /* hex literal */

    // boolean
    true false

There are a couple of important things to note about different literals.

The first few pertain to the character literal.  Chars can only hold
single characters.  So then what is `\n`? It turns out that there are
some characters that can be easily expressed with a character such as
tab and new line.  So we have escape sequences for those that begin
with `\`.  In this case, `\n` corresponds to a new line.

It is also important to note that chars and enclosed in single quotes
and string in double quotes.

Also, there is no such thing as a negative literal.  The `-` operator,
when used in it's unary form (operating on a single value), can be
used to change the sign of a numeric value.

Finally, booleans are case sensitive. This means that `True` or
`FALSE` are not valid literals.
