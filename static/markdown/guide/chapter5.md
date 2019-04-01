# Control Flow

Controlling the flow of your program is one of the most important parts of programming.
Whirlwind offers you numerous control flow structures: both at the block and at
the expression level.

## If Statements

The basic **if statement** is the simplest control flow statement.  It runs the
code in its body if its condition is true.

    // assume a and b are declared as integers somewhere

    if (a < b) {
        a = b; // only runs if a < b
    }

The if block above can be rewritten without the braces designating a block
because it only contains one statement.  If it contained more,
we would have to use braces.

    if (a < b)
        a = b;

This will be true of all control flow statements with one exception.

> In Whirlwind, there is no concept of "truthiness" so any type other than
> a boolean **cannot** be used as a condition.

If statements can also be combined with **elif** and **else** clauses to
create combined trees.

    if (a < b)
        a = b;
    elif (b < a)
        b = a;
    else {
        a = 0;
        b = 0;
    }

The elif statement runs only if all previous statements before it has failed and
its condition is true.  The else statement runs if all of the tree has failed no matter
what.

    if (a < b)
        a = b;
    // still only runs if a is not less than b
    elif (true)
        b = a;

You can add as many elif clauses as you want in between the if and the else;
however, elif cannot be used anywhere else. In addition, you do not need an elif
between the if and the else, nor do you need an else at the end of an elif tree.

## Select Statements

Select statements are value branch statements.  They work off a case system
where each case represents one or more values and their appropriate cases.

    select (a) {
        case 2:
            a = 4;
        case 4, 6:
            a = 8
        case 8:
            a = 10;
    }

As you can see, the value that is being selected over is surrounded in parentheses
after the `select` keyword.  Each individual case is designated by the `case` keyword,
followed by the values delimited by commas and closed by a colon.

Each case is a block, does not require braces, and can contain unlimited amount of
statements.  If all other cases fail, the default case will execute.  You can specify this
case using the `default` keyword.

    select (a) {
        case 2:
            a = 4;
            b = 6;
        case 4, 6, 8:
            a = b;
        default:
            // will only run if call other cases fail
            b = a;
    }

Select statements do require braces around them, regardless of how many cases they contain.
Additionally, any type can be selected over, but the types of the case values must be coercible
to the type of the head value.

## Loops

There are for kinds of loops in Whirlwind: infinite, while, for, and for each.  All loops
begin with the `for` keyword and take a different form depending on what type they are.

The first kind of loop is an infinite loop.  This loop executes the code in its body repeatedly
forever.  It takes the following form.

    for {
        // code to be ran indefinitely
    }

All loop bodies are the same as if bodies in the sense that the braces can be omitted if the for body
contains only one statement.

The next kind of loop is the while loop.  This is a loop the iterate as long as a condition is true.
These kinds of loops have an additional component in their header, which the condition.

    for (a < b * 3)
        a++;

This loop will continue incrementing a as long as it is less than three times b.

The third loop is the classic for loop.  This loop is based off of the
c-style for loop and works in a similar way.  It is composed of an iterator variable, a condition,
and a cycle clause.

    for (i = 0; i < 10; i++) {
        b *= 2 + i;
    }

As you can see, our iterator variable is `i`.  `i` will be incremented as long as it is less than 10.
This loop multiplies by every number between 2 and 11.

> The iterator variable is always accessible and mutable within the for body.

You can also omit parts of the for loop if you either don't need them or don't want them.

    let p = 0.761;

    for (; p < 3.141; p += 0.3) {
        // do something
    }

In this case, our iterator variable was already declared so we can omit it.

The final loop, the for each loop, iterates through every value in a collection.  You simply
specify an iterator variable (or variables under certain conditions) and the collection you
want to iterator over.

    let col = range(1, 11); // returns a list of the numbers 1-10

    let d: dict[int, int];

    for (n <- col) {
        d[n] = n ~^ 2;
    }

This code creates a dictionary of all the numbers between 1 and 10
and their respective perfect squares.  In this loop, `n` is our iterator variable
and `col` is our collection.

## Break and Continue

Break and continue are two simple keywords used for controlling the flow
of loops.  The `break` keyword exits the nearest enclosing loop immediately
and the `continue` keyword move the loop onto its next cycle.

    for (i = 0; ; i++) {
        if (i % 2 == 1)
            continue;

        if (i % 5 == 0)
            break;
    }

This loop continues every time `i` is an odd number and breaks whenever `i` is
divisible by 5.

## Inline Control Flow

There are two inline control flow operators in Whirlwind.  These operators
allow you the determine the value of some expression base on conditions without
having to use a full block.

The first operator is the inline condition operator.  It takes the following form.

    let x = condition ? result1 : result2;

As you can see, this operator returns the first value is condition is true and the
second value if the branch is false.

> Both inline control flow operators as low precedence operators.  This means that
> all other expressions will by default nest within them as opposed to going around them.

The second control flow operator is the inline case operator.  This operator works like the
select expression, but for inline values.

    let y = value case {
        1, 2 => 4,
        3 => 6,
        _ => 8
    };

The value at the top of the case expression is what will be selected over.  Each of the cases are
separated by commas and the possible values and the result are separated by the `=>` operator.

Notice that the last case just has the `_` symbol in it.  In the context of the inline case statement,
this designates the default case.  Because the inline case operator is an expression, it must always yield
a value and thus must always define a default case.