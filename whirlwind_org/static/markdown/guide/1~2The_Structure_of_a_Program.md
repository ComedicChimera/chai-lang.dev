## The Structure of a Program

All Whirlwind programs follow the same basic structure.  This is to
say, every program is composed of the same basic parts.  There are
three major building blocks that will be used in almost every program:
**Expressions**, **Statements**, and **Blocks**. Now, let's break down
the purpose each and the syntax associated with them.

### Expressions
An expression is the smallest unit of code in Whirlwind.  An expression
is computable piece of code.  Here are examples of common expressions:

    3 + 4
    6 < 7
    "Hello, world!"
    fn()
    x
    6 << 5
    5 / 2 ^ 4

You'll notice than expressions can perform many actions and don't have
any particular syntax associated with them.  Expressions can
also contain many values and **operators**.  An operator is just a symbol
that designates some action like `+` or `<`.  One common thing about
expressions is that they all evaluate to some value: `3 + 4` evaluates
to 9.

An important note about expressions is that they cannot stand on their
own.  An expression must be wrapped by or be transformed into a statement.
However, not all expressions can be transformed into statements; they
can only be wrapped.

### Statements
A statement is the smallest piece of code that can stand on their own.
Here some examples of statements:

    $x = 4;
    Println("Hello, world!");
    use include stdio;
    delete x;
    break;

Just like expressions, statements can be very diversified.  However,
they are all united by the commonality that they perform something
that affects the overall state of the program or yield some output.

However, statement do all have one syntactical element in common and it
is that they end with a semicolon.  This is true for all statements
and your program will not compile if there are statements lacking
a semicolon.

One last thing, don't worry if you don't understand what any of those
statements do.  We will cover each of them in detail later on.

### Blocks
A block represents a collection of statements.  They also always declare
some sub region of your program that can have it's own state.
You can designate a block with `{}`.  Here is an example of a block:

    {
        Println("Hello, world!");
    }

*Where have we seen this before?*

### The Main Function
Now that we have covered the various elements that can be used to
construct a program, there is one more thing that is important to discuss.
This is the **Main Function**.  It represents the starting point of
execution in any program and is essential for any program to run.

As the name would imply, the Main Function is a **function**. We will talk
a lot more about these later, but for now just think of them as
**callable** blocks of code.  The code in **Listing 1.2** shows how to
declare a main function.

#### Listing 1.2 - The Main Function

    func main() {

    }

The code inside of the block will be first code to run when your program
begins.

***Note** Declarations and inclusions can be written outside
of the main function and they will be run before the main function is called.
But, for practical purposes, we say the main function is called first.*

That's it, now you know how to structure a program.  This is certainly
a lot of information to take in, but don't worry; the more you
work with Whirlwind, the more these things become easier to understand.
