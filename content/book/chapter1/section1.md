# Hello World

The first program any programmer writes when learning a new programming language
is aptly named *Hello World*.  The task is simple: print "Hello, World!" to the
console.

Let us begin by simply showing a simple *Hello World* program:

    import println from io.std

    def main() = println("Hello, world!")

That's it: two simple lines are all we need to accomplish this task.  But, there
is quite a lot going on in them so let's break it down line by line.

## The Import Statement

Chai does not provide any default construct for printing the console.  Luckily,
such a function is provided in the standard library.  The first of our file
brings this function into scope so we can use it.

This line is called an **import statement**.  It begins with the `import`
keyword followed by the names of any symbols we want to import. In this case,
the only symbol we are importing is called `println`.  It will print its
argument to the standard out followed by a newline.

We then need to tell Chai where to find the `println` function.  This is the
second part of the statement.  It begins with the keyword `from` followed by the
**package path**.  A package path is a way of communicating to Chai where a
certain package.  We will discuss packages and package paths in more detail in
later chaptes: for now, it will suffice to say that the path `io.std` is the
path to the package defining `println`.

Because `println` is such a commonly used function, you will want to remember
this import statement as you will likely place it into your programs quite
frequently.

## The Main Function

The **main function** is called by the runtime whenever your program begins:
it acts the main "entry point" for your program.  

Functions in Chai begin with the `def` keyword followed by their name. The name
of the main function is, as you would expect, `main`.  

After the function's name, we provide a list of the arguments the function takes
enclosed in parentheses.  The main function takes no arguments so we will just
leave these parentheses empty for now.

Since our main function does not return anything, we have no need to specify a
return type and can proceed immediately onward to define the body of the
function.

There are many types of functions bodies in Chai; however, the simplest is the
**expression body** which we make use of here.  Expression bodies begin with an
`=` followed by an expression to be run as the body of the function.  Our
expression in this case is called to the `println` function we imported earlier.
We can call functions in Chai by placing parentheses after their name.

Inside these parentheses, we must provide the arguments to the function.  Our
`println` function takes one argument: the thing we want to print.  In this
case, we want to print a **string**, a programmatic representation of text, to
the console.  This string is enclosed in double quotes and contains the "Hello,
world!" we want to print.

## Compiling the Program

Now that we understand what our little program does, we need to compile it so
we can actually run it.  There are a couple of steps here, but they are all
very simple.

1. Create a new directory in a convenient location.
2. Create a file called `hello.chai` in that directory containing the Hello
   world program.
3. Open up a terminal in the directory you created.
4. Run the command `chai mod init hello` from your terminal.
5. Run the command `chai build .`.

The first few steps are self-explanatory; however, the last two warrant
a bit of discussion.

The fourth step creates a **module** in our project directory.  Modules describe
the structure of our project are necessary for compilation. We will discuss
modules in more detail later, but just know that Chai created one with sensible
default options for you when you ran the `chai mod init` command.

The fifth step actually builds our program.  We provide a path to the directory
containing the module we want to build to the Chai compiler and tell it to build
an executable out of that module. 

After performing all these steps, you should have directory that looks like this:

```text
hello/
    chai-mod.toml  <-- Chai module file
    hello.chai
    out            <-- Executable (will have a `.exe` extension on Windows)
```

We can then run our executable from the command line and see our message of
the hour appear:

```text
Hello, world!
```

Congratulations!  You have just written your first Chai program.

Before we move on, one final note: you can compile and run your program in one
command using the `chai run` command.  It works just like the regular `chai
build` command, but it builds, runs, and deletes the resulting executable for
your convenience.