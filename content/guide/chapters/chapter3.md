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
in the variable.  Let's look at an example.

    let x = 10

In that declaration, `x` is the variable name and `10` is the value that is
stored in it.

We can then access this value later by simply using the variable's name.

    let y = x * 2

All variables also have a fixed type.  For example, `x` and `y` both store
numbers, specifically numbers of a type of `i64`.  This type is *inferred*
from the type of the expression that is used to initialize them.

If you want to specify the type of a variable, you can use a **type label**.

    let pi: f64 = 3.14159264

Only one variable with a given name can be declared in a given scope.

    let x = 5.5  # ERROR

You can also declare multiple variables at once by separating them with commas.

    let z = 78, u = -12.56 * pi

## Assignment

All variables in Chai are **mutable** meaning you can change their value after
you assign to them.  Changing a variables value is called **assignment** and is
done using the `=` operator.  Assignment constitutes its own statement.

    let name = "Bob"

    println(name)  # prints Bob

    name = "Alice"

    println(name)  # prints Alice

A variable's type cannot change during program execution: so, you can only assign
a value to variable that is the same as the type of that variable.

    let x = 5.6

    x = "Hello"  # ERROR

At this point, it is worth talking a little bit more about type inference.  Let's
consider a simple example:

    let y = 5

    y = 0xff

    y = 5.5

If you try and compile the code above, you will get a compilation error on the
third line.  It will look like this:

```text
no type overload of `{int}` matches type `{float}`
```

This is a bit of odd error considering the circumstances: what is the compiler
trying to tell you?  In order to understand why this happened, we need to
understanding how Chai determines the type of something.

Essentially, it starts with a set of possible types for something and slowly
prunes those possible types down as it learns more about your program.  Let's
start with line 1.

    let y = 5

Here, Chai doesn't know much about the type of `y`.  Since `5` is a number
literal, Chai only knows that `y` must be a numeric type, but it doesn't know
which of the many types that are available.

Proceeding to line 2:

    y = 0xff

Because `0xff` is an integer literal, Chai now knows that `y` must be an integer:
all the float types are no longer considered viable types for `y`.  So when Chai
reaches line 3:

    y = 5.5

It sees that `y` is being assigned to a float type, but it also knows that `y`
can only be an integer.  Thus, it encounters a type mismatch between the set of
possible types of `y` and the set of possible types of `5.5`: there is no overlap.
This is why Chai produces an error.  

The language of the error is also relevant.  Chai refers to something called a
**type overload**.  As you may be able to guess, a type overload is the formal
name from one of the possible type possibilities.  All the possible types of `y`
are called the **type overload set** of `y`.  

Now, we can actually interpret what the compiler is trying to tell us.  

```text
no type overload of `{int}` matches type `{float}`
```

Reading it directly, it saying that no type overload of `{int}`, meaning no
possible type of integer, matches (or is in) the set of float types.

### Compound Assignment

Now let's take a break from type mumbo jumbo and take a look at a cool bit of
syntactic sugar.  Suppose you have a variable `a` as defined below:

    let a = 10

Now, suppose you want to double `a` (multiply it by 2).  You could write:

    a = a * 2

However, this line is fairly repetitive and as you will discover a common task.
So, Chai gives us a shorthand.

    a *= 2

This is exactly equivalent to the previous line, just with few characters.  This
is called **compound assignment**.  You can do this with all the arithmetic
operators:

    a += 6  # a = a + 6

    a /= 2  # a = a / 2

By far the most common task out of all compound assignments is adding and
subtracting one, also known as incrementing and decrementing.  Using compound
assignment, we can shorten these to:

    a += 1  # increment
    a -= 1  # decrement

However, for this special case, Chai offers yet another little bit of shorthand.

    a++  # increment
    a--  # decrement

Once again, this code is exactly equivalent to the previous pair of lines, just
shorter.

### Multi-Assignment

TODO

## Command-Line Input and Output

TODO

### Reading Integers

TODO

