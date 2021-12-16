# Lexical Elements

Chai source code is Unicode text encoded in UTF-8.  This document will use the
term character to refer to a Unicode code point in user source text.

## Comments

A **comment** is a part of the source code that is ignored by the compiler.
Comments primarily serve as code documentation.

Comments come in two varieties:

1. A *line comment* begins with a `#` character and continues until the end of
   the line.
2. A *block comment* begins with the the sequence `#!` and continues until a
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

TODO

## Number Literals

TODO

## String Literals

TODO

## Rune Literals

TODO