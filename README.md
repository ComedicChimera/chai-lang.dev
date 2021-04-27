# whirlwind-lang.dev

This is the source for the [Whirlwind programming language](https://github.com/ComedicChimera/whirlwind) website.  
You can find it hosted here *insert link*.

## Dependencies:

This website is built with Django and Svelte.js.

Download and Install:

- Python 3.9 or later
- Node.js (and NPM)

## Building and Running

All actions in this respository can be easily run from the Powershell script `cacao.ps1`. 

Install Required Dependencies:

    cacao setup

Build Static Website Content:

    cacao build

Run Development Server:

    cacao rundev

## Repository Layout

| Directory | Purpose |
| --------- | ------- |
| `content` | Static, markdown context for the page
| `whirlsite` | The main Django project |
| `whirlsite/whirlsite` | Primary project configuration and common files |
| `whirlsite/home` | Home page and primary subpages |
| `whirlsite/docs` | Documentation |

## Guide Layout

Units:

| Unit | Title | Chapters |
| ---- | ----- | -------- |
| I | Language Basics | 1-7 |
| II | Defined Types | 8-11 |
| III | Iterables | 12-17 |
| IV | *untitled* | 18-21 |
| V | Concurrency | 22-24 |
| VI | Vectors | 25 |  

Chapters:

1. Introduction
2. Basic Types
3. Variables
4. Control Flow
5. Pattern Matching
6. Functions
7. Error Handling
8. Structured Types
9. Algebraic Types
10. Interfaces
11. Duck Typing
12. Arrays and Lists
13. Dictionaries
14. Comprehensions
15. Using Strings
16. Lambdas and Closures
17. Iterable Methods
18. References
19. Generics
20. Type Sets and Aliases
21. Packages and Modules
22. Coroutines
23. Futures and Async Iterators
24. Strand Communication
25. Vectors
26. Where to next?

## Not Covered In Guide

- Operator Definitions
- Specialization
- Partial Function Calling
- Annotations
- Metadata
- The `init` function
- Volatile Values and References
- `Copyable` and Copy Overriding
- `Contextual` (outside of error handling)
- Struct Inheritance
- Variadic Arguments
- Fallthrough to Match
- With Expressions
- Inline If Expressions
- Bitwise Operators
- `HeapCollection` and Collection Constructors
- `when` Conditions
- Custom Iterators
- The `yield` statement


