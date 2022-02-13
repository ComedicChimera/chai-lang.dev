# Conditional Logic

Until now, our programs have been entirely linear; however, such programs are
rarely very useful since they lack the ability to make decisions based on input
or external conditions.  In this chapter, we will explore how our programs can
respond and change their progression based on conditions.

## Booleans

A **boolean type** is a type which only has two possible values: True and False.
It represents the logical idea of something having a "truth value".  For
example, the statement "4 is greater than 5" is false while the statement "the
sky is blue" is true.  

Boolean types have two literal values: `true` and `false`.  The type label for
boolean is a `bool`.

We rarely use booleans all on their own: more often, we want to test some
condition and have it produce a boolean that we can act on.  This is where
**relational operators** come in.  These operators take some number of values
and produce a boolean based on the "truth" of some relation between them.  

The simplest relational operator is the *equality operator* which compares two
values to see if they are equal.  It uses the `==` symbol.

    5 == 5      # => true
    6.7 == 4.5  # => false

As you can see, it works exactly like the arithmetic operators we already know,
but it gives a boolean instead of a number.  Like arithmetic operators, you
can also combine them with parentheses.

    (5 == 6) == false  # => true

The equality operator works for all types, but it can only compare two values of
the same type.  The equality operator also has an evil twin: the inequality
operator which does the exact opposite of the equality operator.  It uses the
symbol `!=`.

    2 != 2          # => false
    "abc" != "def"  # => true

The remaining four relational operators are: greater than (`>`), less than (`<`),
greater than or equal to (`>=`), and less than or equal to (`<=`).  These work
similarly to `==` and `!=` except they only work for numbers.

    5 < 6        # => true
    7.2 >= 7.2   # => true
    2e-2 > 8.92  # => false 
    86 <= 86.0   # => true

All of the relational operators are lower precedence than all the arithmetic
operators meaning you can embed arithmetic expressions as the operands of
relational operators without parentheses.

    5 * 6 - 2 < 7 * 8  # => true
    1/(1/4) == 4.0     # => true
    2 + 2 >= 5 * 2     # => false

### Logical Operators

**Logical operators** operate specifically on boolean values.  They represent
different logical queries about and between boolean values.  

To understand what this means, let's look at the first example of a boolean
operator: the NOT operator.  This operator inverts the input boolean that is
given to it.  It uses the symbol `!`.

    !true   # => false
    !false  # => true

Notice that the `!` comes before the boolean value.  It also higher precedence
than all the other logical operators as well as the relational operators.

    !(5 < 7)  # => false

The next two logical operators are AND and OR.  Both take two boolean values and
combine them into a new boolean.  The AND operator returns true only if both its
arguments are true and uses the symbol `&&`.  The OR operator returns true if
either or both of its arguments are true and uses the symbol `||`.  OR is lower
precedence than AND.

    true && false   # => false
    true || false   # => true
    true && true    # => true
    false || false  # => false

    false && true || true    # => true
    false && (true || true)  # => false

    !true || true && !(false && true)  # => true

<guide-exercise>
{
    "label": "4.1",
    "content": "Evaluate the logical expression: true && false || true && !(true && !false)",
    "hint": "NOT before AND, AND before OR",
    "solution": {
        "type": "text",
        "text": "false"
    }
}
</guide-exercise>

Both AND and OR are also both lower precedence than the relational operators.

    5 < 7 && 8 < 9             # => true
    "test" == "abc" || 7 >= 8  # => false

### Multi-Comparison

Consider a fairly common problem in game programming: testing if a point is
inside a box.  Mathematically, this is equivalent to testing if the x-coordinate
of the point is between the x-coordinate edges of the box and if the
y-coordinate is between the y-coordinate edges of the box.  

We can try and express this programmatically like so:

    leftEdge < x && x < rightEdge && bottomEdge < y && y < topEdge

This code is fairly straight forward, but it does have some redundancies: `x`
and `y` both occur two times in two different tests.  It turns out that this
kind of check, that is a value being between two bounding values, is fairly
common. This is why Chai provides us with a little syntactic sugar to make our
lives easier.  

Instead of writing two separate comparisons, we can merge the two into one:

    leftEdge < x < rightEdge && bottomEdge < y < topEdge

This is called **multi-comparison**: Chai automatically combines adjacent
relational operators into a single relation.  This will work with all the
relational operators and more than two relations can be combined together:

    a <= b > c == d

If you want Chai to consider the operations separately, you will need to use
parentheses:

    (a <= b) == false

## If, Elif, and Else

Now that we understand conditions, we can use them to change the behavior of our
program.  The concept of using boolean to direct program progression is called
**control flow**. 

The simplest kind of control flow is the **if statement**.  This statement begins
with the keyword `if` and executes the code in its body only if its condition
is true.  The general form looks like so:

    if condition
        body
    end

The condition is always a boolean expression.  The body, in this case, is just a
block: ie. a series of statements.  Notice that the if body ends using the `end`
keyword just like the blocks we studied in Chapter 3.  You will see a pattern of
`end` being used to terminate logical elements of code.

Here is an example of some code that prints out a greeting if the user enters
"Hello!" using if statements.

    let input = scanln()

    if input == "Hello!"
        println("Hi!")
    end

If you run that program, you will see it only responds to your input if you
enter the exact phrase `Hello!`.  

Often, we also want to specifically run code in the case that the condition is
not true. This is called an **else clause**.  We use the keyword `else` to
denote it.

    let favorite_number = 0
    scanf("{}\n", &favorite_number)

    if -10 < favorite_number < 10
        println("Your favorite number is one digit.")
    else
        println("Your favorite number is two or more digits.")
    end

Notice that the `end` is only used once for the whole statement despite it being
made up of separate blocks.  This is because the `end` is used to close the
entire if-else statement not the individual blocks.

<guide-exercise>
{
    "label": "4.2", 
    "content": "Write a program that prompts the user to enter their favorite
    number and then displays whether or not that number is even.",
    "hint": "You can test if a number is even using the modulo operator (see
    Chapter 2).",
    "solution": {
        "type": "text",
        "text": "TODO"
    }
}
</guide-exercise>

TODO: elif

## Scoping and Shadowing

TODO

### Header Variables

TODO

## Block Expressions

TODO

## Short Circuit Evaluation

Now that we have seen some actual control flow in Chai, let's revisit those
logical operators from earlier since there is one more special behavior we need
to discuss: **short circuit evaluation**.

To understand what this behavior is, let's consider a very simple example:

    false && some_bool

Looking at the first value only, we already know that the result is always false
because both operands can't be true.  This means that the value of `some_bool`
is completely irrelevant.  Chai might as well not evaluate it at all since it
already know the answer just from looking at the first argument.  Indeed, this
is the crux of short circuit evaluation: if Chai already knows the result of a
logical operation by considering only the first operand, then it won't evaluate
the second operand at all.

Let's consider another example:

    some_bool || do_something_that_returns_a_boolean()

If `some_bool` is `true`, Chai knows that the result of `||` will also be `true`
without considering the second argument.  So, in the case that `some_bool` is
`true`, Chai will not even call `do_something_that_returns_a_boolean`: the code
will never run.

This behavior occurs for both AND and OR.  For AND, if the first operand evaluates
to `false`, then the second operand will not be evaluated, and the expression will
evaluate to `false`.  Similarly, for OR, if the first operand evaluates to `true`,
then the second operand will not be evaluated, and the expression will return `true`.

To give you a more concrete sense for this logic, here is what AND looks like
with if expressions.

    if op1 => op2 else => false end

<guide-exercise>
{
    "label": "4.4",
    "content": "Express the short circuit evaluation logic for OR using if expressions.",
    "hint": "Model your answer after the expression for AND above.",
    "solution": {
        "type": "snippet",
        "code": "if op1 => true else => op2 end"
    }
}
</guide-exercise>
