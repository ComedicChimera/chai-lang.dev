# Variables

## Do Blocks

So far, we have only dealt with single expressions.  However, when you need to
use more complex statements or chain a series of expressions together, you need
to use a **do block**.  

Do blocks being with the keyword `do` and are wrapped by a new indentation
level.

    do
        ...

We can then amend our main function to use a do block instead of single
expression.

    def main() = do
        ...

For all the programs that follow this, we will consider our main function to
have a do block that way we can use complex statements in it.  Note that our
`Hello World` program with a do block is effectively identical.

    def main() = do
        println("Hello, world!")

When working with variables and assignment as we will in this chapter, do blocks
are absolutely necessary so make sure to include them.

## Defining Variables

We use the `let` keyword in Chai to create a **variable**, a named location to
store a value.  Simply place the name of the variable you want to create after
the keyword followed by an **initializer**.

    let x = 0

This variable can then be used in code by name:

    let y = x + 2

Chai implicitly infers the type of variables based on the right-hand side. You
can explicitly specify the type using a **type specifier**.

    let pi: float = 3.14

If you include a type specifier, you don't need to provide an initializer.

> This is only true for nullable types *insert link*.  Most of the types in Chai
> (inc. all the ones we currently know) are nullable so this is rarely an issue.
> Nullability is not discussed much in this guide.


    let z: string

You can declare multiple variables at once using commas.

    let name = "Bob", age = 50

You can also declare variables in an indented block.

    let
        var1 = 23
        var2 = 6.7
        var3 = 'h'

## Assignment

Variables can be **assigned** to which means changing their value.  We use the
`=` to assign to values in Chai.

    var1 = 42

You can assign to multiple variables at once. 

    name, age = "Emily", 32

Notice that each variable on the left-hand side corresponds to the value at its
position at right-hand side.  

In assignment, the right side is fully evaluated before the assignment occurs.
This means that you can swap the values of two or more variables using
multi-assignment.

    x, y = y, x

You can only assign a value of a type that is either equal to or coercible to
the type of the variable being assigned to.  If your types do not match, you
will get an error.

    name = 45 # TYPE ERROR

Chai also supports many **compound assignment** operators.  A compound
assignment operator is an operator that performs some operation between the
variable and the value being assigned and then stores the result into the
variable.

    # Expanded Form
    x = x + 2

    # Compound Form
    x += 2

Those two statements are equivalent: `+=` is the compound assignment operator.
Almost all operators in Chai have compound forms.  Here are some other examples:

    let y = 5.6

    y -= 1.2

    y /= 2

    let b = false

    b &&= true

Finally, Chai also supports **increment** and **decrement** statements which add
1 or subtract 1 from a numeric value respectively.

    x++ # increment x by 1

    y-- # decrement y by 1

> Unlike what you may used to in languages like C, `++` and `--` are statements
> not operators.
