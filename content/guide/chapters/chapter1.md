# Hello World

The first program any programmer writes when learning new language is called a
*Hello World* program.  The task is simple: print the phrase *Hello World!* to
the console.  So, let's see how it's done.

## Setting Up the Environment

Before we begin, I am going to assume that you have already installed Chai.  If
you haven't, go check out our [install page](/install) for intrusctions on how
to do that.

Once you have Chai installed, go ahead and create a new directory.  This can be
anywhere on your computer.  However, there are some restrictions on what it can
be named.  For this tutorial, I am going to name it `hello`.  I recommend
sticking with this name for now until you know more about the language.  For
future exercises, you can either just overwrite this directory or create new
ones named something like `exercise1` which will be valid.  

Once you have created your directory, you want to go into it and open up a
command-line instance (if you haven't already been doing things from the
command-line).  Once you are in your directory, run the command:

```text
>> chai mod init hello
```

This will create a new *module* in that directory.  This module will be
described in a file called `chai-mod.toml`.  The compiler will choose some
sensible default options for you so it's best to just leave that file alone for
now.

It is worth mentioning that all Chai projects no matter how small require a
module and an isolated directory.  If you plan on keeping your exercise
solutions separate, you will need to repeat this process of creating a directory
and a module for each one.  We will talk more about why this is in later
chapters, but I promise it is worth the hassle. 

It should be mentioned that the `hello` in the command above is the module's
name. You can choose different names when you create different modules, but
again, there are some restrictions (same as with the directory).  

Next, you will need to choose an editor.  I personally recommend
[VSCode](https://code.visualstudio.com/) as that is the editor the language was
developed in.  I will not be leaning too heavily on any particular editor's
features in this tutorial so you can choose anyone you want.  It will be helpful
for you to have an editor with support for Chai so you can get access to
conveniences like syntax highlighting and compile-error squiggles.

Once you have chosen an editor, go ahead and open it in your directory and open
a new file called `hello.chai`.  If you have made it this far, then you are
ready to get started.

## Our First Program

We are going to begin by simply presenting the *Hello World* program in its
entirety and then break down how it works.

    import println from io.std

    def main() = println("Hello World!")

As you can see, it is fairly short, but there is a lot going on in the those
two lines.  Let's break it down piece by piece.

### The Import Statement

The first line is called an **import statement**.  Import statements are used to
access definitions stored in other *packages*.  In this case, we are importing a
special function called `println`.  This function will allow us to print to the
console.  It is stored in a package called `std` which is a part of the `io`
module.  Thus, we access it using the **package path** `io.std`.  This gives
us a statement of the form:

    import def_name from pkg_path

where `def_name` is `println` and `pkg_path` is `io.std`.  The words `import`
and `from` are **keywords** meaning they are reserved by the language to have
special meaning.  Keywords cannot be used as ordinary names, and there quite a
lot of them in Chai: we will cover most of them in this guide.

It is okay not to understand how this statement works completely: we will cover it
in more detail later.  The key idea is that we need to "bring in" some definitions
to allow us to print to the console, and this statement is how we do that.  

You will likely see and need to add that specific import statement at the top of
most of your early projects since printing to standard out is such a common
operation.  

### The Main Function

After our import statement, the next line of code begins with the `def` keyword.
This keyword begins a **function definition**.  As the name would imply, these
define functions.  

Those unfamiliar with the concept of a function may be a bit lost at this point,
especially since we have spent the first part of the guide talking about them.
Essentially, a **function** is a reusable piece of code.  The idea is that
instead of writing everything out explicitly, we can *name* certain blocks of
code that perform a specific task and *call* them to reuse their functionality
when we want to.  Functions can perform any task and often we will even give you
a value back.  

As an example, the `println` function we imported above contains all the code
necessary to print to the console (and calls lots of other functions to make it
happen). When we call it (which we will do in a moment), we will pass it the
thing we want it to print as an **argument**.  Arguments give us a way to
"customize" the behavior of functions when we call them. As another example, if
we had a function to multiply two numbers, those numbers would be its arguments
since they determine what it's going to multiply. We will talk a lot more about
functions in this guide, so it's important to get used to them early.

Getting back to our definition, we are going to define a special function called
`main`.  This function is the *entry point* to our program meaning that when
the program is run, the `main` function will be called, and the code in its body
will be run.  

The empty parentheses after the name would normally contain the arguments, but
since the main function doesn't take any arguments, we just leave them empty.

So far, this gives us:

    def main()

### The Print Call

After defining what is called the **signature** of `main` (its name, arguments,
and return type -- we will talk about that last one later, but it also doesn't
appear in main's signature), we now need to define its **body** which is the
code we want to run when it is called.

We are going to use an **expression body** for `main`: all the code `main` runs
can be represented in a single expression (we will talk much more about
expressions later).  We denote this by writing an `=` after the signature.

Finally, we get to the actual meat of the body.  We want to print to the console
so we are going to call our `println` function from earlier.  To do this,
we start by placing its name, `println`, followed by a pair of parentheses which
will contain the arguments we are passing to `println`. 

The only argument `println` accepts is a **string** which contains the text we
want to print.  Strings give us a way to store text inside our program. The
content of the string is enclosed in double quotes and can be any UTF-8 text
(with some exceptions we will discuss in the chapter on strings).  The text we
want to print is *Hello World!* so we put that in the string. 

This completes our main function:

    def main() = println("Hello World!")

## Running the Program

Now that we understand the *Hello World* program, we are going to compile and
run it.

The first thing we need to do is *compile* it: convert it into an executable.
This can be done by running the following command in the directory of your
*Hello World* program.

```text
>> chai build .
```

The argument to the `build` command is the directory of the module you want to
build relative to you.  In this case, the module is in our current working
directory so we use `.` to indicate that.

Once this happens, you should an executable with the same name as your module in
a folder called `out` in your working directory.  It will have a platform
appropriate extension (eg. `.exe` on Windows or no extension on many Unix
platforms).  Your working directory should now appear as follows:

```text
hello/
    out/
        hello.exe
    hello.chai
    chai-mod.toml
```

You can then run this executable, and you should see the following output:

```text
Hello World!
```

Viola!  Our program works!  Congratulations, you just wrote your first Chai
program!

> **Exercise:** Modify the *Hello World* program to print *Hello Chai!* instead
> of *Hello World!*.

There is a small shortcut to the above process: you can compile and run your
program with a single command:

```text
>> chai run .
```

This will also "clean up" the executable after it runs it, so you won't have a
lingering executable file if you just wanted to quickly test your program.
