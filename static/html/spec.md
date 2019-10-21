# The Whirlwind Language Specification

## Table of Contents

1. [Introduction](#intro)
    - [Purpose](#purpose)
    - [Notation](#notation)

2. [Copyright](#copyright)
    - [Language and Website](#lang-copy)
    - [Packages](#package-copy)
    - [Alternate Implementations](#alt-impl)

3. [Lexical Structure](#lexical-structure)
    - [Comments](#comments)
    - [Punctuation](#punctuation)
    - [Identifiers](#identifiers)
    - [Keywords](#keywords)
    - [Operators](#operators)
    - [Literals](#literals)

4. [Collections](#collections)
    - [Arrays](#arrays)
    - [Lists](#lists)
    - [Dictionaries](#dicts)
    - [Tuples](#tuples)

5. [Expressions](#expressions)
    - Nesting
    - Precedence
    - Atoms
    - Unary Expressions
    - Arithmetic Expressions
    - Shift Expressions
    - Comparison Expressions
    - Logical Expressions
    - Functional Expressions
    - Expression Local Variables
    - Ternary Expressions
    - Case Expressions
    - Range Expressions
    - Then Expressions
    - Is Expressions
    - Cast Expressions
    - Extract Expressions

6. Statements
    - Variable Declarations
    - Assignment
    - Simple Statements
    - If Statements
    - Select Statements
    - For Loops
    - Context Managers
    - After Clauses

7. Functions
    - Function Declarations
    - Function Groups
    - First Class Functions
    - Lambdas and Closures
    - Composition
    - Partial Functions
    - Monads

8. User-Defined Types
    - Structure Types
    - Interface Types
    - Opaque Types
    - Type Aliases
    - Enumerated Types
    - Algebraic Types

9. Interface Binding
    - Methods
    - Classification
    - Default Implementation
    - Overriding
    - Operator Overloading
    - Special Methods

10. Generics
    - Generic Types
    - Generic Restrictors
    - Generate Construction
    - Monomorphism
    - Variance

11. Symbols and Scoping
    - Declaration Order
    - Compound Definitions
    - Incomplete Definitions
    - Scoping
    - Shadowing
    - Captures

12. Memory Model
    - Pointer Types
    - Heap Allocation
    - Heap Deallocation
    - Moving and Copying
    - Nullable Operators
    - Implicit Deallocation
    - Ownership
    - Lifetimes

13. Concurrency
    - Fibers
    - Futures
    - Asynchronous Functions
    - Await and Shields
    - Volatility
    - Race Conditions

14. Packages
    - Package Assembly
    - Export Status
    - Inclusion
    - Data Organization

15. Absolute Types
    - Any Types
    - Any Pointers
    - None Types

16. Annotations
    - File-Level Annotations
    - Block Annotations
    - Annotation Interpretation

17. Runtime and Execution
    - The Main Function
    - Stack Allocation
    - Forced Copying
    - Error Model
    - Stack Tracing
    - Dynamic Allocation
    - Concurrency
    - Compile-Time Intrinsics

## <a name="intro"></a> Introduction

This specification is a complete reference manual and description
of the Whirlwind Programming Language.  It describes the exact behavior
and construction of each language element as well as the relation between
language elements.

### <a name="purpose"></a> Purpose

Whirlwind is a compiled, modern, and multipurpose language designed with intentionality.
It is strongly-typed, versatile, expressive, concurrent, and relatively easy to learn.  It does not
have a garbage collector and accordingly places a heavy emphasis on memory safety.
It boasts numerous new and old features and is designed to represent the needs of any software developer.

### <a name="notation"></a> Notation

yada yada yada

## <a name="copyright"></a> Copyright

Whirlwind is by nature a piece of intellectual property.  That being said, it is
powered by and thrives off of its community and as such it has some rather unique
rules for usage.

### <a name="lang-copy"></a> Language and Website Copyright

*insert legal bs here*

### <a name="package-copy"></a> Packages

*insert more legal bs here*

### <a name="alt-impl"></a> Alternate Implementations

It is perfectly acceptable to create an alternate implementation of the Whirlwind compiler
and standard library.  Such implementations are not considered infringement on Whirlwind's
copyright so long as they give credit to original language creators and acknowledge that they
are indeed an alternate implementation of an existing programming language.  Additionally,
in order to be considered a valid implementation, one must completely and totally conform to
specification only allowing for deviation in the actual implementation details of the compiler
and standard library elements.  Effectively, this specification needs to be a correct description
of an alternate implementation in order for it to be considered valid.

## <a name="syntax"></a> Lexemes and Grammar

This section will describe the syntactic and lexical structure of the Whirlwind programming language
in detail.

### <a name="lexemes"></a> Lexical Elements

The list below describes the exact and complete list of token names and their corresponding regular expressions
as they are read by the compiler.  

    "STRING_LITERAL" := /"(?:[^"\\']|\\.)*"/,
    "CHAR_LITERAL" := /'(?:[^"\\']|\\.)*'/,
    "FLOAT_LITERAL" := /\d+\.\d+/,
    ">>=" := />>=/,
    ":>" := /\:>/,

    ":=" := /\:=/,
    "++" := /\+\+/,
    "--" := /\-\-/,
    "#" := /#/,
    "@" := /@/,

    "..." := /\.{3}/,
    "~*" := /~\*/,
    "~/" := /~\//,
    "~^" := /~\^/,
    "?" := /\?/,
    "." := /\./,
    ":" := /\:/,
    "," := /,/,
    "!=" := /!=/,
    "!" := /!/,
    "AND" := /&&/,
    "OR" := /\|\|/,
    "XOR" := /\^\^/,
    "|" := /\|/,
    "<-" := /<-/,
    "->" := /->/,
    "+" := /\+/,
    "-" := /-/,
    "*" := /\*/,
    "/" := /\//,
    "%" := /%/,
    "==" := /==/,
    ">=" := />=/,
    "<=" := /<=/,
    "=>" := /=>/,
    "=" := /=/,
    ";" := /;/,
    ">" := />/,
    "<" := /</,
    "&" := /&/,
    "~" := /~/,

    "(" := /\(/,
    ")" := /\)/,
    "{" := /\{/,
    "}" := /\}/,
    "[" := /\[/,
    "]" := /\]/,

    "LET" := /\blet\b/,
    "CONST" := /\bconst\b/,
    "_" := /\b_\b/,

    "IF" := /\bif\b/,
    "ELIF" := /\belif\b/,
    "ELSE" := /\belse\b/,
    "FOR" := /\bfor\b/,
    "SELECT" := /\bselect\b/,
    "CASE" := /\bcase\b/,
    "DEFAULT" := /\bdefault\b/,
    "BREAK" := /\bbreak\b/,
    "CONTINUE" := /\bcontinue\b/,
    "WHEN" := /\bwhen\b/,
    "AFTER" := /\bafter\b/,

    "RETURN" := /\breturn\b/,
    "YIELD" := /\byield\b/,

    "DELETE" := /\bdelete\b/,
    "FROM" := /\bfrom\b/,
    "VOL" := /\bvol\b/,
    "MAKE" := /\bmake\b/,
    "WITH" := /\bwith\b/,
    "STATIC" := /\bstatic\b/
    "DYN" := /\bdyn\b/,

    "FUNC" := /\bfunc\b/,
    "ASYNC" := /\basync\b/,
    "AWAIT" := /\bawait\b/,
    "VARIANT" := /\bvariant\b/,
    "CONSTRUCTOR" := /\bconstructor\b/,
    "OPERATOR" := /\boperator\b/,

    "TYPE" := /\btype\b/,
    "STRUCT" := /\bstruct\b/,
    "INTERF" := /\binterf\b/,

    "INCLUDE" := /\binclude\b/,
    "EXPORT" := /\bexport\b/,

    "THIS" := /\bthis\b/,
    "NEW" := /\bnew\b/,
    "NULL" := /\bnull\b/,
    "IS" := /\bis\b/,
    "THEN" := /\bthen\b/,
    "VALUE" := /\bvalue\b/,
    "AS" := /\bas\b/,

    "STRING_TYPE" := /\bstr\b/,
    "FLOAT_TYPE" := /\bu?float\b/,
    "BOOL_TYPE" := /\bbool\b/,
    "CHAR_TYPE" := /\bs?char\b/,
    "BYTE_TYPE" := /\bbyte\b/,
    "LONG_TYPE" := /\bu?long\b/,
    "DOUBLE_TYPE" := /\bu?double\b/,
    "INT_TYPE" := /\bu?int\b/,
    "ANY_TYPE" := /\bany\b/,

    "BOOL_LITERAL" := /\b(true|false)\b/,
    "HEX_LITERAL" := /0x[0-9A-F]+/,
    "BINARY_LITERAL" := /0b[10]+/,
    "IDENTIFIER" := /[^\d\W]\w*/,
    "INTEGER_LITERAL" := /\d+/

The name is the value enclosed in quotes on the left hand side of the `:=` and the regular expression is the value on the right
hand side of the `:=` and is enclosed in `/`.

In this context, a token refers to the distinct element of any given program file that matches a given regular expression.  
Each token is comprised of a name, a value, and a position. The name is listed above, the value is whatever element matched the
regular expression specified for the given token type (name).  The position is where in the file the token was found and the match's
length. The tokens are extracted from the program file by the scanner according to the rules listed and are passed to the parser
in the order that they appear in the program file.

### <a name="notation"></a> Notation

Our context-free grammar uses a modified form of EBNF (Extended Backus-Naur Form) that allows for comments and does not include a `?` operator
or token literals.  Additionally, it uses a different production declaration operator.

The below code block outlines the syntactic notation used in our custom EBNF notation.

    // This is a comment

    /* This is a multiline comment */

    // Tokens (Terminals) are encased in single quotes, and all productions are ended by semicolons.
    production_name: 'TOKEN' ;

    // Our format allows for alternators and parentheses
    alternated_production: ( production | 'TOKEN' ) | production ;
    
    // and optional blocks
    optional_production: [ production ] ;

    // star and plus operators
    operator_production: production* production+ ;

This simple notation is all that used to define the Whirlwind grammar.  However, several conventions are
followed in the Whirlwind Language Grammar itself.

- All sections are prefixed by titles in all caps.
- Related productions are grouped together.
- Each production group has one new-line on either side of it and is labeled with comment.
- No production contains capital letters.
- Any alternator which requires multiple lines should follow Haskell style.

Any other patterns that appear in the grammar are not convention and the last convention is not
always respected.

### <a name="grammar"></a> Grammar

The Whirlwind Language Grammar is partially ambiguous: it allows for left-recursion and productions
with multiple beginning symbols in common.  The parsing algorithm was custom-written for Whirlwind and
is designed to deal with this complex grammar with ease.  The start symbol for the grammar itself is
`whirlwind`.

Below is the complete grammar for Whirlwind exactly as it is read by the compiler.

    // grammar goes here

This grammar is designed to be processed into an object by the grammar processor and subsequently passed
to the parser for efficiency's sake.  It is only loaded once per run of the compiler regardless of how
many files are being processed and remains in memory throughout compilation.

## <a name="data-types"></a> Data Types

This section describes the data types used Whirlwind, and their behavior.  It will also outline what
and how they compile.  However, this section does not clarify which operators are valid on which types
nor how to declare them.

### <a name="primitives"></a> Primitives

In Whirlwind, a primitive is the simplest type.  It is comprised of no sub-types and translate directly to
LLVM.  Whirlwind divides the primitives into four categories: boolean types, integral types, and floating point types.

#### The Boolean Type

The boolean type is a type with two states: true (1) and false (0).  It is a classic boolean type akin to those
found in languages like C++, Rust, and Go.  It designated with the type label `bool` and compiles to an `i1` in
LLVM.  

#### The Integral Types

There are four integral types in Whirlwind: the byte, the char, the int, and the long.  An integral type is a whole
number that spans a specific range and can be either signed and unsigned both depending on the type.  Integral types
are considered numbers, but two have additional connotations.

The first integral type is the byte.  It occupies one byte of memory and translates as an `i8` in LLVM.  If it is unsigned,
it holds a value of exactly 8 bits and has a type label of `byte`.  If it is signed, it holds a value of 7 bits with
the first bit is used as a sign bit and has a type label of `sbyte`.  This type, although it is an integral type, also
can represent raw binary data such as in a data stream.  No change need occur for this to be the case; rather, it one
must simply use the type differently.

The second integral type is the char. It occupies two bytes of memory and translates as an `i16` in LLVM.  If it is unsigned,
it holds a value of exactly 16 bits and has a type label of `char`.  In this state, it is also used as Whirlwind's primary
character type that is it holds a single Unicode (UTF-16, excluding the top 1792 and all private-use characters) character point.
If it is signed, it holds a value of 15 bits with the first bit being used as a sign bit.  This version uses the type label `schar`.

The last two integral types are relatively similar differing only in their size.  Both are, by default, a signed state, can be converted
to unsigned by prefixing their type label with a `u`, and represent a plain, old integer.  The first is the int type which is designated
with the type label `int` and occupies 4 bytes of memory.  It translates as an `i32` in LLVM.  The second is the long type which is
designated with the type label `long` and occupies 8 bytes of memory.  It translates as an `i64` in LLVM.  Both types, in their signed forms, 
will hold up to one bit less than the number of bits they occupy in memory with their first bit being used as a sign bit.  
