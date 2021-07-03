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
8. Loops and Comprehsnions
9. References
10. Vectors and Matrices
11. Lambdas and Closures
12. Sequence Functions

### Unit III - Defined Types

12. Structure Types
13. Algebraic Types
14. Method Spaces
15. Type Classes
16. Generics
17. Monads and Error Handling

### Unit IV - Packages and Modules

18. Packages and Visibility
19. Modules
20. Import Semantics
21. Build Profiles

### Unit V - Concurrency

22. Chairoutines (name WIP)
23. Futures and Events
24. Async Queues
25. Task Pools




