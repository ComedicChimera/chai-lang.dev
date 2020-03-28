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
    - [Byte Types](#3-byte-types)
    - [Integral Types](#3-int-types)
    - [Floating-Point Types](#3-float-types)
    - [Boolean Types](#3-bool-types)
    - [Character Types](#3-char-types)
    - [String Types](#3-string-types)
    - [Any Types](#3-any-types)
    - [None Types](#3-none-types)

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
    - Select Expressions
    - Range Expressions
    - Then Expressions
    - Is Expressions
    - Cast Expressions
    - Extract Expressions

6. Statements
    - Variable Declarations
    - Constancy
    - Assignment
    - Simple Statements
    - If Statements
    - Select Statements
    - For Loops
    - After Clauses
    - Context Managers

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
    - Generalized Boxing

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
    - Value Semantics
    - Pointer Types
    - Heap Allocation
    - Heap Deallocation
    - Moving and Copying
    - Nullable Operators
    - Ownership
    - Lifetimes

13. Concurrency
    - Fibers
    - Futures
    - Asynchronous Functions
    - Await and Shields
    - Mutexes
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
    - Fiber Management
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
    let        const      if         elif       else       for
    select     case       default    break      continue   when
    after      return     yield      delete     from       make
    with       func       async      await      variant    constructor
    operator   type       struct     interf     include    export
    this       super      new        null       is         then
    value      as         vol        static     own        dyn

It is important to note that the last four items in this list above are most accurately referred to as modifiers not keywords, but they follow the same semantics
as all other keywords do.

Finally, it is acceptable (although in many cases not adviseable) to use keywords within identifiers: a keyword may be part of a larger identifier.
For example, the following identifiers would be considered valid:

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

### <a name="2-operators"></a> Operators

An operator is any symbol that denotes an operation.  Whirlwind contains many different kinds operators, varying both in the number of operands they accept and
the symbols (or keyword) used to represent them.  For the sake of efficiency, below is a list of all of the standard operators.

    whirlwind
    >>=   :>   :=   ++   --   ->   <    ~    *    ?
    !=    !    &&   ||   ^^   |    <-   >    ~/   ~^ 
    /     %    ==   >=   <=   =    &    ~*   +    -

Notably, some of the operators listed can be combined with the `=` operator to form a [compound assignment](#6-assignment) operator.  

Finally, they are several operators that were not listed in the above list because they are either primarily considered some other type of lexical element.
Furthermore, all of these operators are considered non-standard operators due to the operands they accept and/or the function they perform.  Many of these
non-standard operators also do not accept operands in the traditional unary or binary manner that all of the standard operators listed above do.

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
integral type starting from the signed integer type that can hold its value.  Moreover, since the suffixes only increase the maximum range, they can
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

A character literal represents a single UTF-8 encoded, unicode point that corresponds to the [character type](#3-char-types).  This literal is
enclosed in single quotes and must contain exactly one value.

    ebnf
    /'(?:[^\"\\']|\\.)'/

The character literal can contain any unicode character or one of the following legal escape codes (each followed by its function in parentheses):

    \a (alert)
    \b (backspace)
    \f (form feed)
    \n (new line)
    \r (carriage return)
    \t (horizontal tab)
    \v (vertical tab)
    \0 (null terminator)
    \s (space)
    \" (escaped double quote)
    \' (escaped single quote)
    \\ (escaped backslash)

In addition, the prefix `\u` may be used to signal partial unicode character code (16 bits) and the prefix `\U` may be used to signal a full unicode
character code (32 bits) where the appropriate number of hexadecimal characters follows the the escape prefix.

#### Boolean Literals

A boolean literals represents a single boolean (true/false) value that corresponds to the [boolean type](#3-bool-types).  This literal can hold
two values: `true` and `false` as expressed below:

    ebnf
    /\b(true|false)\b/

In this literal, `true` corresponds to the boolean value true (1) and `false` corresponds to the boolean value false (0).  

#### String Literals

A string literal represents a string of unicode characters that corresponds to the [string type](#3-string-types).  This literal
matches the regular expression below.

    ebnf
    /\"(?:[^\"\\']|\\.)*\"/

This literal may contain any escape code or unicode character code that would be valid as a character literal.  Moreover, these literals
may contain as many or as few characters as necessary.  

String literals (and strings by extension) may also contain no values.  These strings are termed as empty strings and have a length of 0.
Notably, these strings do still occupy memory and can be manipulated as if they were a proper string.  However, any attempt to read a value
from this string (though a subscript) will result in a runtime error.

#### Binary and Hexadecimal Literals

Binary and hexadecimal literals represent an arbitrary piece of data.  All binary and hexadecimal literals correspond to either a [byte type](#3-byte-types)
or an [integral type](#3-int-types) depending exclusively on the type of data they hold.  Furthermore, the correspondent type is always
unsigned.

Binary literals take the following form:

    ebnf
    /0b[10]{1,64}/

Notice that these literals are prefixed by `0b` and may only contain up to 64 characters (corresponding to the `ulong` type) and must contain at least
one character.  

As noted previously, the size of the resulting type is determined by the size of the literal.  In this case, the number of bits in the literal
corresponds directly with the type is will evaluate to.  That is to say, a binary literal will be inferred as the smallest type that occupies
a greater or equal number of bits than the number contained within the binary literal.

Hexadecimal literals take the following form:

    ebnf
    /0x[0-9A-F]{1,16}/

All hexadecimal literals are prefixed by `0x` and must contain at least one character but no more than 16 characters (corresponding to the `ulong` type).
Furthermore, only capital letters may be used as hexadecimal digits within a hexadecimal literal.  

Hexadecimal literals follow a similar type progression to binary literals; however, the space that the hexadecimal literal occupies is determined
by the space required to store the binary value of the hexadecimal literal.  By contrast, with respect to the determined size of the literal,
the type determination process is identical to that of a binary literal.

Finally, in both literals, 0s still count as a additional space in the literal.  That is to say, the size of the inferred type is determined not by the
actual used memory of literal but rather by the number of characters specified in the literal.  For example, `0x000` will evaluate to a larger type than `0x00`.

## <a name="prim-types"></a> Primitive Types

Primitive types are considered to be the most simple types offered by Whirlwind (with the exception of the string type to a degree).  They are often
the building blocks for more complex types.

All primitive types that can only hold a finite set of a values can experience underflow, overflow, and/or loss of precision.  These types will **not**
throw any form of error when said events occur; it is the responsibility of the programmer to check for, prevent, and handle these edge cases where applicable.

All references to coercion and casting in this section refer exclusively to coercion and casting between primitive types.  More complex types such as
interfaces and type classes do **not** apply here.  Additionally, the any type, due to all types ability to be coerced to it, is also not included in
the descriptions of coercion and casting rules.

> Note: The phrase "a *T* type" that occurs in each type description that follows and many that occur later in this specification technically
> refers to an element of the set described by *T* type.  In essence, this shorthand is akin to saying "a value of type *T*".

### <a name="3-byte-types"></a> Byte Types

A byte type is simplest type in Whirlwind.  It represents a single byte of data with no particular type.  

It has two different variants: the signed byte and the unsigned byte.  Each has its own type label.

    whirlwind
    byte  // unsigned byte
    sbyte // signed byte

The only literal forms of a byte type are sufficiently small hexadecimal and binary literals both of which have
a type of unsigned byte (if small enough).  

Although byte types compile to an `i8` in LLVM IR, they are **not** considered integral types, but they can
be coerced to any of the integral types without an explicit type cast.  They have no additional possible explicit
casts other than to the types to which they are coercible.

### <a name="3-int-types"></a> Integral Types

An integral type represents a whole number that can be negative if the type is marked as signed.  There are several
different sizes of the integral type as listed below with their corresponding type label:

    whirlwind
    short // 16 bit, signed integral type
    int   // 32 bit, signed integral type
    long  // 64 bit, signed integral type

All integral types have absolute sizes that do not vary by platform.  All integral types are represented with
two's complement arithmetic. Moreover, each integral type has an unsigned variant that can be designated with the prefix `u`.

    whirlwind
    ushort // 16 bit, unsigned integral type
    uint   // 32 bit, unsigned integral type
    ulong  // 64 bit, unsigned integral type

All of the integral types follow the same casting and coercion rules with respect to each other.  Any integral type
can be coerced to any integral type that has a size larger than that of itself.  Moreover, any integral type
will automatically coerce from its signed form to its unsigned form; however, integral types must be explicitly cast
from their unsigned to their signed forms.  Finally, all integral types must be explicitly cast in order to transform
them into a smaller integral type or to the byte type.

Both short types and integer types (where the integer type is another way of saying the 32-bit integral type) are
coercible to both types of [floating-point types](#3-float-types), regardless of signage.  By contrast, long types
(both signed and unsigned) are only coercible to double types.  However, all integral types can be cast to either
one of the floating point types explicitly.

Standing apart from the other integral types, the signed and unsigned 32 bit integer type is also capable of being cast
to the [character type](#3-char-types).  

### <a name="3-float-types"></a> Floating-Point Types

A floating-point type represents a finite precision decimal number.  There are two different types of floating-point
types, and their type labels are listed below with their sizes:

    whirlwind
    float  // 32 bit floating-point type
    double // 64 bit floating-point type

Floating-point types make no distinction between signage at a typing level: all floating-point types have a sign.
Both floating-point types conform to the IEEE-754-2008 specifications for binary32 and binary64.

With regards to coercion and casting within the floating-point types, a `float` can coerce to a `double`, but a `double`
must be explicitly downcast to `float`.  As it pertains to other types, floating-point types must be explicitly cast
to integral types and cannot be cast to any other type.

### <a name="3-bool-types"></a> Boolean Types

A boolean type represents a Boolean truth value and is a psuedo-integral data type.  
It occupies a single byte and is denoted with the following type label:

    whirlwind
    bool

There is no sign associated with the boolean type.  It cannot be implicitly coerced to any other type, but it can
be cast into any integral type.

Notably, the boolean type only stores a single bit of information and is padded to 8 bits.  Semantically, it is only
considered to be 1 bit in size; however, it still technically occupies a byte thus the description above.  For this
reason, when taking the size of the data type, the language always returns that it is 1 byte in size.

### <a name="3-char-types"></a> Character Types

A character type represents a single UTF-8 encoded character.  It occupies 32 bits and has no sign.  Its type label
is listed below:

    whirlwind
    char

Due to its nature as a UTF-8 encoded value, any valid unicode code point can be stored in a character type.

Character types can only be coerced to one other primitive: the string.  However, they can be cast to the
integer (32 bit integral) type.

### <a name="3-string-types"></a> String Types

String types are unique among the primitive types in that represent not one value but many.  A string is
defined to be a set of characters with a finite length.  Notably, that length is not noted in the
data type and thus is not considered at a typing level: two string types of different length are
equivalent from a typing standpoint.

A string does not occupy any definite amount of space; however, the construct itself includes a pointer
to an array of characters and a length represented as a 32 bit unsigned integer.  Therefore, it can be
said that the base string type occupies the equivalent size of a struct of that contains those two members.
This definition excludes that underlying memory that a string also contains via its pointer.

The string type label is as follows:

    whirlwind
    str

Strings are unique from almost all other types in Whirlwind in two capacities.  They are fundamentally immutable
which means that elements of a string which are defined to be character types can only be accessed; never set.
Therefore, an operation such as the one shown below is considered to be invalid.

    whirlwind
    let s = "gap";
    s[0] = 'n'; // COMPILATION ERROR

There are also unique in that the underlying data of a string is only copied implicitly upon being returned from
a function as opposed to being copied in any other situation where a normal implicit reference type would be copied
(as outlined in section 12). This is due to the fact that the underlying data of a string will never be mutated and
so it is feasible for multiple instances of the same string to share the same underlying data without violating
the value semantics of Whirlwind (as outlined in section 12).

Finally, because strings implicitly include a length in their definition, they are not null terminated.

### <a name="3-any-types"></a> Any Types

The any type can be cast to any other type and all types can be coerced to it.  The any type has no definite
size as it can store any value.  However, similar to the string type, the any type is also implemented as a
struct; therefore, its size is considered to be the size of that struct.

It has the following type label:

    whirlwind
    any
