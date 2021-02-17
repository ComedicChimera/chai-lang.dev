---
title: "Introduction"
weight: 1
---

## Installation

If you haven't already installed Whirlwind, visit our [installation page](/install).
This will walk you through the short installation process and make sure you
are all set up and ready to go for this tutorial.

## Setting Up a Workspace

In order to get started programming in Whirlwind, you will need a workspace to run
code.  

Assuming you already have the language installed, go to your development directory of
choice (wherever you want to store your Whirlwind projects) and create a new folder. 
You can call it whatever you want. 

After you have created it, navigate into the directory and create a new file.  Again,
this can be called whatever you want.  

Then, go ahead and open up a terminal.  Once you have put some code into the file, you
can use the command, `whirl run .` to build and run your code (current working package).

If you want to simply produce an executable, you can use the command `whirl build .` to
build your current package.  The executable will be created in a directory called
`bin` by default.

## Hello Whirlwind

Now, we are going to walk you through writing your first Whirlwind program.  This program
is going to print out `"Hello, world!"` to the console and is often the first program
written when learning a new language.  

Firstly, we need to import the function `println` that we are going to use to print to
the console.  To do this, add the following line to the top of your file:

    import println from io::std

`io::std` is the name of the package that contains `println`.  We will talk more about packages
and importing in a later section -- you won't need more than this for a while. 

Now, we need to create your main function.  This is the function that is called at the start
of your program.  This function is simply called `main` and should be defined as follows:

    func main() do
        ...

The `do` at the end of the function denotes the beginning of a block of code and a new
indentation level.  Whirlwind uses indentation to determine where a block begins and ends
and indentation level to determine where statements fit into the control flow of your program.

Finally, we need to call the print function.  We are going to pass in the string `"Hello, world!"`.
A string is simply a piece of text -- we will talk more about them in the next chapter.

    func main() do
        println("Hello, world!")

That's it.  We place the arguments to our function inside parentheses.  We do not need to use any
form of punctuation to end a line in Whirlwind.  The linebreak at the end denotes that the line is
over.  Whirlwind is generally whitespace sensitive and uses whitespace to discern the structure of
your program -- however, it is still fairly flexible as you will see in later chapters.

If you run the above code, you will see the text `Hello, world!` printed out in your terminal. 
Congratulations, you just wrote your first Whirlwind program and took your first step into learning
a new programming language!

