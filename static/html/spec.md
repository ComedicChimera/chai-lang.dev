# The Whirlwind Language Specification

## Table of Contents

1. [Introduction](#intro)
    - [Purpose](#purpose)
    - [Copyright](#copy)
    - [Notation](#notation)
    
2. [Lexical Structure](#lexical-structure)
    - [Comments](#comments)
    - [Punctuation](#punctuation)
    - [Identifiers](#identifiers)
    - [Keywords](#keywords)
    - [Operators](#operators)
    - [Literals](#literals)

3. [Collections](#collections)
    - [Arrays](#arrays)
    - [Lists](#lists)
    - [Dictionaries](#dicts)
    - [Tuples](#tuples)

4. [Expressions](#expressions)
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
    - Constant Expressions

5. Statements
    - Variable Declarations
    - Constancy and Constexpr
    - Assignment
    - Simple Statements
    - If Statements
    - Select Statements
    - For Loops
    - Context Managers
    - After Clauses

6. Functions
    - Function Declarations
    - Function Groups
    - First Class Functions
    - Lambdas and Closures
    - Composition
    - Partial Functions
    - Monads

7. User-Defined Types
    - Structure Types
    - Interface Types
    - Opaque Types
    - Type Aliases
    - Enumerated Types
    - Algebraic Types

8. Interface Binding
    - Methods
    - Classification
    - Default Implementation
    - Overriding
    - Operator Overloading
    - Special Methods

9. Generics
    - Generic Types
    - Generic Restrictors
    - Generate Construction
    - Monomorphism
    - Variance

10. Symbols and Scoping
    - Declaration Order
    - Compound Definitions
    - Incomplete Definitions
    - Scoping
    - Shadowing
    - Captures

11. Memory Model
    - Pointer Types
    - Heap Allocation
    - Heap Deallocation
    - Moving and Copying
    - Nullable Operators
    - Implicit Deallocation
    - Ownership
    - Lifetimes

12. Concurrency
    - Fibers
    - Futures
    - Asynchronous Functions
    - Await and Shields
    - Volatility
    - Race Conditions

13. Packages
    - Package Assembly
    - Export Status
    - Inclusion
    - Data Organization

14. Absolute Types
    - Any Types
    - Any Pointers
    - None Types

15. Annotations
    - File-Level Annotations
    - Block Annotations
    - Annotation Interpretation

16. Runtime and Execution
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

### <a name="copy"></a> Copyright

Whirlwind is by nature a piece of intellectual property.  That being said, it is
powered by and thrives off of its community and as such it has some rather unique
rules for usage.

### <a name="notation"></a> Notation

Our grammatical notation uses a modified form of EBNF (Extended Backus-Naur Form) that allows for comments and does not include a `?` operator
or token literals.  Additionally, it uses a different production declaration operator.

The below code block outlines the syntactic notation used in our custom EBNF notation.

    ebnf
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

This simple notation is all that used to define the Whirlwind grammar and will be used in this specification.  

Additionally, regular expressions are infrequently used throughout this specification and will be marked with
forward slashes on either side when used.

## <a name="lexical-structure"></a> Lexical Structure

This section will describe the basic lexical structure as well as some of the simpler code fragments (eg. literals)
in the Whirlwind programming language.

### <a name="comments"></a> Comments

A comment represents any piece of code that will not be processed by the compiler.  Comments take two forms in
Whirlwind: single-line and multi-line.  They take the following form:

    whirlwind
    // single-line comment

    /* multi-line comment */

As the name would imply, the former type of comment occupies one line, and the latter type of comment can be spread over multiple lines.  
Additionally, multi-line comments must always be bounded on both sides the appropriate symbols (`/*` for the left, `*/` for the right).

    whirlwind
    /* multi-
    line
    comment */

Finally, it is also legal (although not advised) to embed single-line comments within multi-line comments

    whirlwind
    /*
    // nested comment
    */

Because the content of all comments is completely ignored, it is of no consequence to the compiler what you place inside the comment,
excluding the symbols used to end multi-line comments which will cause the compiler to think the comment has ended.

### <a name="punctuation"></a> Punctuation

Whirlwind requires that several different pieces of punctuation be used in different scenarios.  The most notable of which is the
semicolon which primarily denotes the end of block-less [statement](#statements).  Another notable piece of punctuation is the colon which
primarily denotes a type label (or type extension; both names are acceptable) all of which take the form `: type` where type is replaced
with the desired type for the label.

Another important piece of punctuation is the brace: its two forms (left and right) are used to begin and end a code block.

    whirlwind
    {
        // code here
    }

There are several other pieces of punctuation used in Whirlwind, and they are, including those already mentioned, listed here.

    whirlwind
    ;   :   ,   .   (   )   {   }   ?


Many of the aforementioned punctuation elements have multiple uses and meanings and so no are not stictly considered punctuation.  Moreover,
there are several other [operators](#operators) that could be considered punctuation.

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
