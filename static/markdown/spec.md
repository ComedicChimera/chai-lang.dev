# The Whirlwind Language Specification

## Table of Contents

1. [Introduction](#intro)

2. [Copyright](#copyright)
    - [Language and Website](#lang-copy)
    - [Packages](#package-copy)
    - [Alternate Implementations](#alt-impl)

3. [Lexemes and Grammar](#syntax)
    - [Lexical Elements](#lexemes)
    - [Notation](#notation)
    - [Formal Grammar](#grammar)

4. Data Types
    - Primitives
    - Arrays
    - Lists
    - Dictionaries
    - Pointers
    - References
    - Tuples
    - Functions
    - Structures
    - Type Classes

5. Bases
    - Literals
    - Identifiers
    - This
    - Super
    - Null
    - Val
    - Comprehensions
    - Lambdas
    - Type Casts
    - Partial Functions

6. Expressions
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
    - Is and As Expressions
    - Extract Expressions

7. Statements
    - Variable Declarations
    - Assignment
    - Return Statements
    - Yield Statements
    - Delete Statements
    - Break Statements
    - Continue Statements
    - Expression Statements

8. Block Statements
    - If Statements
    - Select Statements
    - For Loops
    - Context Managers
    - After Clauses

9. Declarations
    - Function Declarations
    - Decorator Tags
    - Variant Declarations
    - Structure Declarations
    - Interface Declarations
    - Interface Bindings
    - Type Class Declarations

10. Scoping
    - Sub Scopes
    - Shadowing
    - Captures
    - Declaration Order

11. Memory and Ownership
    - Stack Allocation
    - Default Initialization
    - Heap Allocation
    - Heap Deallocation
    - Null
    - Ownership
    - Lifetimes

12. Functional Behaviors
    - Closures
    - Higher Order Functions
    - Composition
    - Chaining
    - Monads

13. Interface Binding
    - Type Inferfaces
    - Methods
    - Bind Labels
    - Bind Order
    - Generic Binding

14. Type Relations
    - Constancy
    - Coercibility
    - Casting
    - Is Operator
    - Classifying Interfaces

15. Generics
    - Generic Types
    - Generic Restrictors
    - Generate Construction
    - Monomorphism
    - Variance

16. Operators
    - Reference Operators
    - Increment and Decrement
    - Arithmetic Operators
    - Conditional Operators
    - Bitwise and Logical Operators
    - Functional Operators
    - Type Operators
    - Operator Overloading

17. Concurrency
    - Fibers
    - Asyncs
    - Futures
    - Threads
    - Pooling
    - Events
    - Communication

18. Packages and Inclusion
    - Include Statements
    - Export Statements
    - Packages
    - The Prelude

19. Annotations
    - File-Level Annotations
    - Struct-Specific Annotations
    - Function-Specific Annotations
    - Local Annotations

20. Optimizations
    - *insert here*

21. Runtime and Execution
    - The `main` Function
    - The Fiber and Thread Registry
    - The Heap
    - Compile-Time Intrinsics

## <a name="intro"></a> Introduction

This specification is a complete reference manual and description
of the Whirlwind Programming Language.  It describes the exact behavior
and construction of each language element as well as the relation between
language elements.

Whirlwind is a compiled, modern, and multipurpose language designed with intentionality.
It is strongly-typed, versatile, expressive, concurrent, and relatively easy to learn.  It does not
have a garbage collector and accordingly places a heavy emphasis on memory safety.
It boasts numerous new and old features and is designed to represent the needs of any software developer.

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
    "OWN" := /\bown\b/,
    "VOL" := /\bvol\b/,
    "MAKE" := /\bmake\b/,
    "WITH" := /\bwith\b/,
    "STATIC" := /\bstatic\b/
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
    "REF" := /\bref\b/,
    "VAL" := /\bval\b/,
    "AS" := /\bas\b/,
    "STRING_TYPE" := /\bstr\b/,
    "FLOAT_TYPE" := /\bu?float\b/,
    "BOOL_TYPE" := /\bbool\b/,
    "CHAR_TYPE" := /\bs?char\b/,
    "BYTE_TYPE" := /\bbyte\b/,
    "LONG_TYPE" := /\bu?long\b/,
    "DOUBLE_TYPE" := /\bu?double\b/,
    "INT_TYPE" := /\bu?int\b/,
    "VOID_TYPE" := /\bvoid\b/,
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

*insert grammar when ready*

This grammar is designed to be processed into an object by the grammar processor and subsequently passed
to the parser for efficiency's sake.  It is only loaded once per run of the compiler regardless of how
many files are being processed and remains any memory throughout compilation.
