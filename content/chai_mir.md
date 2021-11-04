# Chai MIR

This file documents Chai's **M**iddle **I**ntermediate **R**epresentation or
MIR.  It is an abstract form that sits between Chai and LLVM to make the
back-end's job of optimization much easier and to allow Chai to target more than
just LLVM.

It does have a textual representation which will be shown here.  You may imagine
that all of these text instruction will have corresponding abstract forms.

*Note*: Chai MIR is particularly useful for developers looking to adapt Chai to
target other platforms.  Chai MIR is fully platform agnostic, easy to parse, and
guaranteed to be well formed and semantically valid given that is produced by
the compiler.  A developer can use the existing Chai compiler as a front-end to
produce Chai MIR and compile that directly to another target (without having to
compile the full LLVM dialect).

## Program Structure

Each Chai MIR file is composed of the contents of an entire Chai package. These
files begin with a series of **symbol definitions** that define external symbols
to the package.  At this point the compilation process, the concept of an import
statement no longer exists so instead these Chai MIR files are constructed much
like C program files where all symbols from outside the current Chai MIR file
(ie. translation unit) are defined as externals to be linked later.

Textually, these definitions begin with the `extern` keyword followed by the
standard definition syntax.  

Also notice that Chai MIR is not whitespace sensitive: all lines end with a
semicolon.

Immediately after these definitions come the definitions of global constants.
These are defined with the `static` keyword followed by the name of the global
constant, a type extension, the equals sign followed by the constant value.
For example, the static constant for `pi` would look like:

```chmir
static pi f32 = 3.14159265;
```

After these constants, should be the global non-static variables of the
package.  These must be defined in the following form: a `var` keyword,
an identifier, and a type label.

```chmir
var list core.types.List.[int];
```

Yes, that long prelude is a type label.  As you will see such preludes are
going to be common-place in generated Chai MIR which includes no generics
and no namespacing.  In fact, the provided sample code thus far would likely
never be produced by the compiler given that those constants would be namespaced.

This brings up another important distinction between Chai and Chai MIR: Chai MIR
identifiers can include: `.` and `[]`.  This is to allow easy translation
between the naming schemes of the two languages.

After these variables, the various functions of the program can be defined.

## Program Start Up, Shutdown, and Global Initialization

If a the Chai MIR is set to target an unsupported OS or architecture, then the
default runtime code will not be generated.  Thus, the reader must provide a
suitable entry point.  This entry point must call two functions in order to
properly initialize Chai: `__init` which calls all global initializer functions
generated with Chai packages (to initialize global variables) and
`core.runtime.init_runtime` which performs runtime initialization.

In addition, this function should also gracefully exit using whatever routines
are necessary.  Finally, in order to actually run user code, the function
defined `__main` must be called.  This function has the Chai signature:

    def __main(argv: Array[string]) i32

This signature is standardized to allow different forms of `main` to be called
identically.

## Symbol Visibility

TODO


