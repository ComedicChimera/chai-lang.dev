# Hello World

We will begin our study of Chai by writing a simple *Hello World* program.  The
task itself is simple: print the phase "Hello World!" to the console.

However, there is a lot involved in actually making this happen.  First, we need
to learn how to actually construct a basic Chai program, how to compile it, and
how to run it.

## Running Your First Chai Program

The *Hello World* program in Chai appears as follows:

    import println from io.std

    def main() = println("Hello World!")

Before we consider how this program actually works, let's first discuss how to
compile it as there is a bit more machinery involved than you may expect.

First, we need to create a file to store our program.  In Chai, all program
files have the extension `.chai` and are created as UTF-8 text files.  Go ahead
and create a file named `hello.chai` and enter the code above into it.  Make
sure that this file is in its own directory, separate from other source files.

Then, open up your command-line and navigate to the directory containing the
source file you just created.  Once there, run the following command:

```text
chai mod init hello
```

You will see a new file appear in your directory called `chai-mod.toml`.  This
is called a **module file**: all Chai projects, however small, must have one. We
will discuss them in much more detail when talk about *modules* in a later
chapter.

That done, you can then run the command:

```text
chai build .
```

to build the program into executable.  After a few moments, you should see a new
directory appear inside your working directory called `out` which contains an
executable called `hello` (with a `.exe` extension if you are on Windows).  You
can then run the executable, and you should see the following output:

```text
Hello World!
```

## The Import Statement

Let us now discuss the construction of the program you just executed beginning
the first line:

    import println from io.std

This line is called an **import statement**.  By default, Chai does not provide any
"built-in" way to print to the console.  So, we have to import that functionality
from the **standard library**, a collection of modules containing common utilities
for working on any system.

The statement begins with the `import` **keyword**.  A keyword is special word
reserved by the language for a specific meaning.  

Next, we specify the name of the thing we want to import.  In this case, we are
importing a *function* called `println`.  This function provides us with the
ability to print a line to standard out, ie. the console.  

Finally, we conclude our statement with the `from` keyword followed by the
**package path**. The package path tells Chai where to look for the function we
just imported.  In this case, we are telling Chai that the function is defined
in the `std` *package* of the `io` module.  As you might guess, modules are
comprised of packages which contain the actual source code, the code defining
`println` in our example.  We will take a much in depth look at the structures
of modules and packages in later chapters, but it is important to familiarize
yourself with them now as they are an essential part of the language.

As an aside, the `io` module contains all the common functionality used for I/O
in general as well as a standard interface for console I/O and, as we will see
later, File I/O.  This module is provided by the standard library and will
always be accessible to you.  

## The Main Function

The next line of our program:
    
    def main() = println("Hello World!")
    
is called a **function definition**.  As the name would imply, it defines a
function which for our program is called `main`. `main` is a special function
which acts as the entry point for all our programs. Every executable program
must have a main function.

Our function definition begins with the `def` keyword followed by the name of
the function which as discussed is `main`.  The parentheses following the name
contain the *arguments* to the function.  The main function does not take any
arguments so we leave those parentheses empty.  

After this header, we move on to the **body** of the function.  The body
contains the code that will be executed when the function is called.  

For our function, the body begins with an `=` sign.  This denotes that body
of the function is an *expression*.  

## The Println Function

The expression contained in the body of our main function is a *call* to the
`println` function we imported earlier.  We call a function by first referencing
its name followed by a pair of parentheses containing the arguments we want to
pass to the function.  

The `println` function prints the text passed to it followed by a newline.  So,
the only argument we want to pass to it is simply the text we want to print.  

We represent text in Chai using **strings**.  A string is an array of bytes
representing UTF-8 encoded text.  We can input text directly into our programs
using a **string literal**. String literals are enclosed in double quotes. Since
the text we want to print is "Hello World!", we pass the string literal `"Hello
World!"` as the first argument to our `println` function.

That's it: all the pieces of the Chai *Hello World* program.  Congratulations on
writing your first Chai program -- you have taken an important step into
learning the Chai programming language!






