# Variables and Input

*Procedural code* is the basis of most modern programming languages.  In short,
it consists of a series of *blocks* which chain together different *statements*
which perform actions.  In the absence of other control flow structures (which
will consider later in this guide), execution proceeds sequentially: one
statement is executed after the other.

In this chapter, we will explore how this process works in Chai.  Although the
previous two chapters have covered material that is mostly consistent with other
languages, this chapter is where Chai starts to diverge from the norm a little
bit.

## Blocks

A **block** is an expression which sequentially executes a series of
**statements** contained within it.  In Chai, blocks begin with the `do` keyword
followed by a newline and concluded by the `end` keyword.  Each statement is
concluded by a newline.

    do
        # statement 1
        # statement 2
        # ...
    end

Importantly, blocks are considered an expression, so we can make one the body of
our main function.

    def main() = do
        # statements
    end

Because this pattern is so common, Chai allows the use of a **block body** for
functions.  Essentially, you just drop the `= do` from the code above, and Chai
will implicitly define a block as the body of the function.

    def main()
        # statements
    end

Now, it is time for a brief digression on program structure.  In Chai, all
executable code, that is code that actually does something, must be contained
within a definition, most often a function.  For now, we are going to be putting
all that code inside the main function.  However, since it is inconvenient to
copy and paste a main function definition for every example, this guide will
frequently show snippets of executable code and discuss them without enclosing
them in a function.  These snippets do *not* constitute full Chai programs. You
should also assume that `println` is always imported and available.

As mentioned before, blocks are made up of statements.  There are many different
kinds of statements, but the first one we are going to look at is the
**expression statement**.  These statements are exactly what they say on the
tin: expressions used as statements. 

    do
        5
    end

That is perfectly valid block.  However, there are some restrictions on how
expressions can be used in expression statements.  Most "complex" expressions
will need to be wrapped in parentheses.

    do
        2 + 2  # ERROR
    end 

    do
        (2 + 2)  # OK
    end

A notable exception to this rule is the function call:

    do
        println("Hello, there!")
    end

## Variable Declarations

The next kind of statement we are going to look at are 
**variable declarations**. These define **variables** which are named locations
to store values.  

Variable declarations begin with the `let` keyword followed by name and an
**initializer**. The initializer contains the initial value that will be stored
in the variable. Let's look at an example.

    let x = 10

In that declaration, `x` is the variable name and `10` is the value that is
stored in it.

We can then access this value later by simply using the variable's name.

    let y = x * 2

All variables also have a fixed type.  For example, `x` and `y` both store
numbers, specifically numbers of a type of `i64`.  

If you want to specify the type of a variable, you can use a **type label**.

    let pi: f64 = 3.14159264

TODO

## Assignment

## Command-Line Input


