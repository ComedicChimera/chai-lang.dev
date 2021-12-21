# Program Structure

The smallest, compileable unit of code in Chai is a **module**.  Modules provide
structure to any given translation unit and are essential to effective
dependency management.  Modules are made up of *packages*.

## Module Files

All modules are defined by a special file named `chai-mod.toml`.  These files
use the [TOML](https://toml.io/en/) markup language to describe the
configuration of the module.

These files are placed in the root directory of a module: once a valid module
file is present, the directory is automatically elevated to a module.  The
directory containing the module file is referred to as the **module root**.

There are several "top-level" configuration options that are required for all
modules.  They are as follows:

| Field | Type | Purpose | Restrictions |
| ----- | ---- | ------- | ------------ |
| `name` | string | the name of the module | must be a valid identifier |
| `version` | string | the version number of the module | of the form: `major#.minor#.patch#` |
| `chai-version` | string | the minimum Chai version required to build the module | must be a valid Chai version |
| `temp-dir` | string | the path Chai can use as a temporary/storage directory | must be a valid path |

We will leave discussion of the specifics of these properties to their relevant
sections of the specification: most options are related to dependency
management.

## Packages and Namespacing

A **package** is a collection of Chai source files in the same directory that
share a global namespace.  The name of the package is inferred from the name of
the directory: the directory name must also be a valid identifier.

Packages are constructed *implicitly*: placing source files together in a
directory automatically makes that directory a package.  Furthermore, no
"package definition" is required in any of the source files.

Packages fall into one of two categories based on where they are placed relative
to the module that contains them.

1. *Root packages*: packages in a module root
2. *Sub-packages*: packages in sub-directories of a module root

Building a module corresponds to building that module's root package.  All
modules must have a root package and can have any arbitrary number and depth of
sub-packages.

It should be noted that Chai does not provide any explicit support for
"sub-modules" that is module inside of the sub-directories of other modules.
However, it doesn't "prohibit" them; rather, it doesn't really acknowledge them
as being different from the sub-packages of a module.

## File Structure

Individual source files are composed at the top level of *import statements*,
*definitions*, and *meta directives*.  

Import statements are used to bring other packages into the file's *local
namespace*.  A more rigorous discussion of import statements will occur in
later chapters; for now, it is sufficient to say that they must be placed 
at the start of source files.

**Definitions** are primary content of any Chai source file.  Definitions
define symbols in the package's global namespace.  There are several types
of definitions which can be enumerated as follows:

1. Functions
2. Defined Types
3. Type Unions
4. Operators
5. Global Variables
6. Method Spaces
7. Type Classes

Each of these kinds of definitions will have their own dedicated chapter which
discusses them in detail: this listing acts to specify which language constructs
are considered definitions.

Many definitions will have some number of **predicates**.  Predicates are
*expressions* which are associated with the definition and used to determine
some value(s) related to it.  These predicates are where the actual executable
contents of the Chai source file are contained.

An **expression**, which can compromise a predicate, is any group of executable
code that yields a value (which may be `()` ie. nothing/unit).  There are many
varieties of expression and ways to construct them which will be discussed over
several chapters.  Chai considers almost all executable code to be expressions
including control flow constructs which allows them to be treated as yielding
values.
