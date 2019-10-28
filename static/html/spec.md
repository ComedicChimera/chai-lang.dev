# The Whirlwind Language Specification

## Table of Contents

1. [Introduction](#intro)
    - [Purpose](#1-purpose)
    - [Copyright](#1-copy)
    - [Notation](#1-notation)

2. [Lexical Structure](#lexical-structure)
    - [Comments](#2-comments)
    - [Punctuation](#2-punctuation)
    - [Identifiers](#2-identifiers)
    - [Keywords](#2-keywords)
    - [Operators](#2-operators)
    - [Literals](#2-literals)

3. [Primitive Types](#prim-types)
    - Byte Types
    - Character Types
    - Integral Types
    - Floating-Point Types
    - Boolean Types
    - String Types
    - Any Types
    - None Types

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
    - Constant Expressions

6. Statements
    - Variable Declarations
    - Constancy and Constexpr
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

### <a name="1-purpose"></a> Purpose

Whirlwind is a compiled, modern, and multipurpose language designed with intentionality.
It is strongly-typed, versatile, expressive, concurrent, and relatively easy to learn.  It does not
have a garbage collector and accordingly places a heavy emphasis on memory safety.
It boasts numerous new and old features and is designed to represent the needs of any software developer.

### <a name="1-copy"></a> Copyright

Whirlwind is by nature a piece of intellectual property.  That being said, it is
powered by and thrives off of its community and as such it has some rather unique
rules for usage.

### <a name="1-notation"></a> Notation

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

### <a name="2-comments"></a> Comments

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

### <a name="2-punctuation"></a> Punctuation

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
    ;   ,   (   )   {   }   
    .   =>  #   @   ... :

Many of the aforementioned punctuation elements have multiple uses and meanings and do not always denote punctuation.

### <a name="2-identifiers"></a> Identifiers

In Whirlwind, an `identifier` is anything that represents a symbol.  That is to say, it is a name.  In Whirlwind, all identifiers must follow
the following regular expression:

    ebnf
    /[^\d\W]\w*/

This regular expression indicates that an identifier begins with some letter upper or lowercase and is followed by any number of letters, numbers
or underscores. Additionally, it is important to note that identifiers must be bounded on either side by any character that is not a letter, number or underscore.

Note that Whirlwind does not allow dollar signs in variable names and that a single underscore is **not** a variable name and is used for other purposes in the
language.  However, variables may be prefixed with any number of underscores although prefixing with more than one underscore is not recommended as it could cause
conflicts with compiler or prelude declared symbols.

### <a name="2-keywords"></a> Keywords

A keyword is any word, which in this context is a combination of lowercase letters, that is reserved by the compiler for any special purpose.  Below is a list
of all keywords used in Whirlwind.

    whirlwind
    let      const  if      elif   else     for
    select   case   default break  continue when
    after    return yield   delete from     make
    with     func   async   await  variant  constructor
    operator type   struct  interf include  export
    this     super  new     null   is       then
    value    as     vol     static own      dyn

It is important to note that the last four items in this list above are most accurately referred to as modifiers not keywords, but they follow the same semantics
as all other keywords do.

Finally, it is acceptable (although in many cases not adviseable) to use keywords within identifiers: a keyword may be part of a larger identifier.
For example, the following identifiers would be considered valid:

    whirlwind
    typeX
    my_static_var
    new13
    my_value
    forThis
    orelse
    include6
    withAll

However, if the identifier is malformed then the keyword will still match separately.  Furthermore, keywords are case-sensitive so an identifier comprised
of a differently-cased form of one of the keywords would be acceptable as well.

### <a name="2-operators"> Operators

An operator is any symbol that denotes an operation.  Whirlwind contains many different kinds operators, varying both in the number of operands they accept and
the symbols (or keyword) used to represent them.  For the sake of efficiency, below is a list of all of the standard operators.

    whirlwind
    >>= :> := ++ -- -> <  ~  *  ?
    !=  !  && || ^^ |  <- >  ~/ ~^ 
    /   %  == >= <= =  &  ~* +  -

Notably, some of the operators listed can be combined with the `=` operator to form a [compound assignment](#6-assignment) operator.  

Finally, they are several operators that were not listed in the above list because they are either primarily considered some other type of lexical element.
Furthermore, all of these operators are considered non-standard operators due to the operands they accept and/or the function they perform.  Many of these
non-standard operators also do not accept operands in the traditional unary or binary manner that all of the standard operators listed above.

### <a name="2-literals"></a> Literals

A literal represents a single, discreet value used in code.  For example, the number `3` could be considered a literal because it is single, constant
value.  There are 7 different types of literals in Whirlwind and each is associated with a specific [primitive type](#prim-types).

#### Integral Literals

An integral literal represents a non-negative whole number that corresponds to one of the [integral types](#3-int-types) and a specific signage.
All integral literals will take the following form.

    ebnf
    /\d+[ul]*/

Notice that these literals can end with a suffix denoted which specific integral type they correspond to.  The suffix `u` marks the integral
literal as unsigned and the `l` suffix marks it as a long type.  By default, integral literals default to signed integers.

If the value held by the integral value is outside the range of allowed values for a signed integer literal, it will be interpreted as the smallest
type that can hold its value (smallest meaning smallest possible range).  Moreover, since the suffixes only increase the maximum range, they can
never contradict with this upcasting pattern; rather, they merely provide a base size to cast up from if necessary.  

Finally, if the value is too large to be stored in the largest possible integral type, then the program will fail to compile.

#### Floating-Point Literals

A floating-point literal represents a decimal number that corresponds to one of the [floating-point types](#3-float-types).  All floating-point literals
will take the following form.

    ebnf
    /\d+\.\d+d?/

Akin to integer literals, floating-point types can also end with a suffix `d` that denotes it as a `double` instead of a `float`.  Additionally,
if the value it stores has a greater precision than the base float type allows or has a greater value that the float type allows, the literal
will automatically be upcast to a double type.  If this upcast is still insufficient, the program will fail to compile.

#### Character Literals

A character literal represents a single unicode point that corresponds to the [character type](#3-char-type).  This literal is
enclosed in single quotes and must contain at least one value.

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
