# The Whirlwind Language Specification

PAGE-BREAK

## Table of Contents

1. Introduction
    - Usage of Specification
    - Design of Whirlwind
    - End-Game Vision

2. Copyright
    - Language and Website
    - Packages
    - Alternate Implementations

3. Lexemes and Grammar
    - Lexical Elements
    - Modified EBNF Syntax
    - Formal Grammar

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
    - Extended Expressions

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
    - Coercibility
    - Constancy
    - Casting
    - Classifying Interfaces
    - Is and As

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

19. Optimizations
    - *insert here*

20. Runtime and Execution
    - *insert here*