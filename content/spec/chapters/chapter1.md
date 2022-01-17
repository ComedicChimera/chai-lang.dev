# Lexical Elements

Chai source code is Unicode text encoded in UTF-8.  This document will use the
term character to refer to a Unicode code point in user source text.

## Comments

A **comment** is a part of the source code that is ignored by the compiler.
Comments primarily serve as code documentation.

Comments come in two varieties:

1.  A *line comment* begins with a `#` character and continues until the end of
    the line.
2.  A *block comment* begins with the the sequence `#!` and continues until a
    matching `!#` is found.  Block comments can span multiple lines.

Comments can not begin inside a rune literal, a string literal, or another
comment.

Line comments count as a newline; block comments even if they span multiple
lines, do not count as a newline.

## Whitespace

Chai considers the following characters to be **whitespace**: `\x09`, `\x0B`,
`\x0C`, \x0D`.  

Unless inserted within a string or rune literal, such characters have no
semantic meaning within Chai source code.  They are primarily used to facilitate
good code style.

## Newlines and Split-Joins

Chai considers **newlines** (`\x0A`) to be significant: they are used to delimit
statements and definitions among other purposes.  They may generally *not* be
placed arbitrarily within user source text.

However, if one wants to continue an expression or definition over multiple lines
without it counting as a newline, a **split-join** may be used.  Split-joins begin
with a backslash (`\`) followed immediately by a newline.

Additionally, newlines may be inserted after commas, semicolons, opening
parentheses, opening braces (`{`), opening brackets (`[`), and `->` *without*
being counted as a newline.  Furthermore, they may be inserted before closing
group symbols (`)`, `}`, and `]`) without being counted as a newline.

Multiple newlines placed in immediate succession will only be counted as a
single newline.  Furthermore, a file may begin with an arbitrary number of
newlines.  Finally, an end-of-file may also be counted as a newline.

## Identifiers

An **identifier** is a lexical token generally used for referring two named
values such as functions or variables.

Identifiers must match the regular expression: `\b[a-zA-Z_]\w*`.  Below are
some examples of valid identifiers:

    a
    b10
    HEll0
    _my_func
    UserError
    left4ded
    pi

Some identifiers are reserved as *keywords* meaning they have special syntactic
significance.  Keywords cannot be used as regular identifiers.  A listing of all
the keywords in Chai is provided below:

    def     union  type    class   space  for  
    oper    let    const   import  from   pub 
    with    async  while   if      elif   closed
    else    match  case    do      break  continue 
    return  after  when    end     await  fallthrough
    as      is     async   await   fn     then
    catch   null   sizeof  i8      u8     i16
    u16     i32    u32     i64     u64    nothing
    f32     f64    string  bool    super

In addition to these reserved keywords, the `_` symbol also has special meaning
and may not generally be used as an identifier.  Its relevancy to pattern
matching will be discussed in more detail in later sections.

### Intrinsic Constants

Chai also reserves several identifiers as **intrinsic constants**: constants
which refer to a specific value within the language.

The identifiers `true` and `false` refer to the Boolean values of true and
false respectively and always correspond to a boolean type.  These identifiers
are often thought of as "boolean literals".

The identifier `null` is used to refer to a type's *null value*.  

## Number Literals

Chai provides four different kinds of number literals.  These literals
correspond to a different set of possible types the literal can have.

1. *Numeric Literals*: can be any numeric type
2. *Float Literals*: can be any floating-point type
3. *Integer Literals*: can be any integral type
4. *Imaginary Literals*: can be any complex type

> The specifics of these types will be discussed in later chapters.

### Numeric Literals

Numeric literals consist of any series of digits representing a base 10 integer.
They may also contain arbitrary underscores separating digits.  These
underscores do not have any effect on the value of the number; rather, they
exist to make large literals easier to read.

This can be represented exactly by the regular expression: `\b\d(_?\d)*`.

Below are some examples of these literals:

    45
    7
    1240
    1_000_000

### Float Literals

Float literals are used to represent floating-point numbers.  Such numbers must
either consist of a series of digits separated by a decimal point or by an
exponent (scientific notation). A literal may also contain a decimal followed by
an exponent. Decimals are denoted with the `.` character and exponents with
either an `e` or an `E`.  Exponents can be negative; however, they must be
integral.  Similar to numeric literals, underscores may be added in between
digits before and after the decimal to improve clarity.

This can be represented exactly by the regular expression:
`\b\d(_?\d)*(\.\d(_?\d)*([eE]\-?\d+)|[eE]\-?\d+)`.

Below are some examples of these literals:

    3.141592
    6.626e-34
    81e9
    6.022E23
    10_000.123_456

### Integer Literals

Integer literals are used to represent integers non-base 10 integer numbers.
They begin with a `0` followed by a **base prefix**: `b` for binary, `o` for
octal, `x` for hexadecimal.  The "digits" of the literal are placed after this
prefix.  These digits may be separated by underscores.

Regular expressions for each literal are given below:

- Binary: `0b[01](_?[01])*`
- Octal: `0o[0-7](_?[0-7])*`
- Hexadecimal: `0x[0-9a-fA-F](_?[0-9a-fA-F])*`

Below are some examples of these literals:

    0b1010
    0xff
    0o125
    0xab_01_7E

In addition, integer literals may also represent special kinds of base 10
numbers: those which are explicitly unsigned or "long" (64 bits).  Such integers
are represented by standard numeric literals followed by a suffix: `u` for
unsigned or `l` for long. Integer literals may have both suffixes applied in any
order.

Below are examples of these literals:

    1u
    238l
    67ul

### Imaginary Literals

Imaginary literals are used to represent the imaginary components of complex
numbers. These literals are comprised of either a numeric or float literal
followed immediately by `j` representing the imaginary number.

Below are examples of these literals:

    12j
    21.5j
    6.45e10j

## Rune Literals

Rune literals are used to represent a single unicode character.  They consist of
an arbitrary unicode character enclosed in single quotes.  Examples include:

    'a'
    ' '
    'θ'

Rune literals may not contain newlines or unescaped single quotes and backslashes. 

With the exception of escape codes, rune literals may only contain a single
character.

### Escape Sequences

To denote special characters such as newlines, single quotes, or backslashes,
escape sequences must be used.  An escape sequence begins with a backslash
followed by an **escape code**.  Below is a table of escape codes:

| Code | Character |
| ---- | --------- |
| `a` | alert |
| `b` | backspace |
| `f` | form feed |
| `n` | newline |
| `r` | carriage return |
| `t` | tab |
| `v` | vertical tab |
| `0` | null terminator |
| `'` | single quote |
| `"` | double quote |

Additionally, Chai supports the use of Unicode escape codes to denote
unicode characters numerically.  These escape codes begin with a prefix
followed by a hexadecimal number corresponding to the escape code.

| Code | Max Value | Length |
| ---- | --------- | ------ |
| `x` | 255 | 2 |
| `u` | 65535 | 4 |
| `U` | 4294967295 | 8 |

Each hexadecimal number must be exactly the length corresponding to the prefix:
zeroes may be added before as padding.

Here are some examples of escape sequences:

    '\n' # newline
    '\'' # single-quote rune

    '\u03A9' # upper case omega: Ω
    '\xB0' # degree symbol: °

## String Literals

String literals represent strings of unicode, UTF-8 encoded, text.  These
literals come in two flavors with different syntax and properties.

### Standard String Literals

Standard string literals are enclosed in double quotes and can contain any
unicode characters except: newlines, unescaped backslashes, and unescaped double
quotes.

Standard string literals can contain escape sequences: these sequences behave
identically to those in runes.

    "Hello, world!"
    "I\nam\non\nmany\nlines with a \t tab."

    "\u03B8\xB0"

    "Hello, 世界"

### Raw String Literals

Raw string literals are enclosed in backticks and may contain *any* character
except an unescaped backtick.  These strings may span multiple lines and contain
unescaped backslashes.  

A backtick in a raw string literal may be escaped to prevent it from ending the
literal by placing a backslash before it.  This is the only context in a raw
string literal in which a backslash has any special functionality.

Here are some examples of raw string literals:

    `Hello, world!`

    `I am
    a multiline
    string!`

    `\d+(\.\d+)?`

    `¡Hola!`

    `My name is \`Bob\`!`
