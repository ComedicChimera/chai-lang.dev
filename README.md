# chai-lang.dev

This is the source for the [Chai programming language](https://github.com/ComedicChimera/chai) website.  
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
| `chaisite` | The main Django project |
| `chaisite/whirlsite` | Primary project configuration and common files |
| `chaisite/home` | Home page and primary subpages |
| `chaisite/docs` | Documentation |
| `chaisite/packages` | Package documentation/index |

## Guide Layout

### Unit I - Variables and Expressions

1. Hello World
2. Basic Types
3. Variables
4. Control Flow
5. Pattern Matching

### Unit II - Functions and Collections

6. Functions
7. Lists and Dictionaries
8. Loops and Sequences
9. References
10. Vectors and Matrices
11. Lambdas and Closures
12. Comprehensions
13. Sequence Functions

### Unit III - Defined Types

14. Structure Types
15. Algebraic Types
16. Method Spaces
17. Type Classes
18. Generics
19. Monads and Error Handling

### Unit IV - Packages and Modules

20. Packages and Visibility
21. Modules
22. Import Semantics
23. Build Profiles

### Unit V - Concurrency

24. Chairoutines (name WIP)
25. Futures and Events
26. Async Queues
27. Wait Groups and Task Pools




