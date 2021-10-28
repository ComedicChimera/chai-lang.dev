# Hello World

Every programmer learning a new language should start by learning how to print
"Hello world!" to the program's console.  This is a basic task and the first,
albeit simple, stepping stone to learning the language.

Note that this chapter assumes you have already installed Chai.

## Setting Up

Before we begin writing our program, we need to create a simple project.
Navigate to whatever directory you find most convenient, and open up your
terminal of choice in that directory.  Then, run the following command:

    chai mod new --empty hello

In your directory, you should see that the command created a directory called
`hello`. Navigate into this directory, and you should see the following:

```language-text
hello/
    chai-mod.toml
```

What is that `chai-mod.toml` file doing there?  The answer: it is the **module
file** for our new project.  **Modules** are the backbone of Chai, and you
cannot build or run Chai code without one. 

As might guess, Chai uses [TOML](https://toml.io) to mark up its module files.
If you were to open this file, you would actually see quite a lot of information
which, again, we will discuss in more detail later.  The good news is that you
won't have to deal with them directly: the command you just ran initialized a
brand new module for you with some sensible default options selected.  This may
seem tedious, but Chai is designed with larger projects in mind where this idea
of a module makes more sense and will likely save you quite a lot of headaches
down the road.

Now that we have our module all set up, go ahead and create a file with the
`.chai` extension in the `hello` directory.  We will call our file `hello.chai`
but as you will see, it doesn't really matter what you name it as long as it has
the `.chai` extension.

Go ahead and open this file in the editor of your choice and with that tiny
prelude, we are ready to start coding.

## The Hello World Program

A simple *Hello world* program in Chai looks like so:

    import println from io.std

    def main() = println("Hello, world!")

There are a couple key parts to break down here:

1. The file begins with an **import statement** that retrieves a function
   `println` from a package at the path `io.std`.  `std` is a sub-package of the
   `io` module -- we will study what that really means in a later chapter.
2. The program begins with a **main function**, named `main` and defined using
   the `def` keyword.  This function must be included in every executable
   program and invoked at the beginning of program execution.
3. The body of the main function is just an expression; we use the `=` operator
   to bind an expression the result of a function.  In this case, it is an
   expression that returns `nothing`, which is a special type in Chai denoting
   "no value".
4. The `println` function call itself is fairly simple: we can `println` by its
   name using parentheses and pass it a single string argument.  Strings are
   collections of UTF-8 encoded text in Chai and are denoted with double quotes.

That is all the code necessary to say "Hello, world!" -- pretty neat right?

## Running the Program

To actually build our program to an executable, we need only run the following
command:

    chai build .

You should see you directory structure update like so:

```language-text
hello/
    out/
        hello.exe  <-- our binary (will be called just `hello` on Unix platforms)
        hello.pdb  <-- a debugging file for GDB
    chai-mod.toml
    hello.chai
```

You can then run this executable from the terminal, and you should see the
following output:

    Hello, world!

Congratulations, you just took your first step into becoming a Chai programmer!

It is worth mentioning that we can streamline our build process a bit if we just
want to run the program.  We can just use the `run` command.

```language-text
> chai run .
Hello, world!
```

This will compile our binary, run it, and then remove any trace of the binary's
existence for your convenience.

### That `.` Argument

However, the astute among you may have one lingering question: why does the
compiler accept a directory and not a file when building?  The answer goes back
to that idea of modules: we are passing the compiler the path to the module it
should build not a specific file.  When building any given module, Chai will
automatically compile all Chai source files in the directory as one **package**.
If this seems a bit confusing, don't worry.  We will explore packages and
modules much more in later chapters, but for now, just know you have to pass the
module path not the file path to the compiler to build.  

In line with this, if you want to follow along with this tutorial on your own,
just replace/update the code in the Chai file you have created and run the
commands as we did above.

## Aside: Comments

**Comments** are text that is ignored by the compiler.  In Chai, comments do
not have ANY special behavior outside of just being comments.  

Line comments begin with hash symbol and continue until the end of the line.

    # I am a line comment

Block comments begin with `#!` and end with `!#`.  They can go over multi-lines
or be nested within a single line.

    #! I
    am a
    block
    comment !#

    some_code #! I am an nested block comment !# some_more_code

We will use comments frequently throughout the guide, and you are encouraged to
use them in your own source code to improve readability.
