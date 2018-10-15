## The Select Statement

The **select statement** is another conditional construct.  However,
it is starkly different to an if statement in that is based on checking
one value against multiple other values and choosing one that matches.
It is analogous to the switch statement from other languages.

### An Example

Let's examine an example of when a select statement would be very useful.
Consider you have an integer variable `a`. You need to print out a different
message for each value of `a` between 1 and 5.  Using an if statement,
that code would look like the following:

    if (a == 1)
        Println("Foo");
    elif (a == 2)
        Println("Bar");
    elif (a == 3)
        Println("FooBar");
    elif (a == 4)
        Println("Fizz");
    elif (a == 5)
        Println("Buzz");

We can make this code significantly more legible and concise through the use of
a select statement.

#### Listing 3.4 - A Selective Improvement

    select (a) {
        case 1:
            Println("Foo");
        case 2:
            Println("Bar");
        case 3:
            Println("FooBar");
        case 4:
            Println("Fizz");
        case 5:
            Println("Buzz");
    }

While it may appear longer at first, you will notice that there is
significantly less repetition and that this code is much more
expressive.

### Simple Select Statements
Every select statement begins with the `select` keyword followed by the
item that you will be selecting over wrapped in parentheses.  This item can be anything
from a literal to the return value of a function.

    select (item) {

    }

Next, we must add in the cases.  Each case is begun with the
`case` keyword an represents a value to matched against.

    select (name) {
        case "Bob":
            Println("Hi, Bob");
        case "Steve":
            Println("Go away Steve.");

        // -- snip --
    }

You can have as many cases as you want, but a select statement must contain
at least one case.

The value in the parentheses will be compared to each value and if there is
a match, the code in the case region will be run.

### The Case Block
The case block is a little bit special. There are no wrapping braces
around it but, it does declare an implicit sub-scope.  Furthermore,
you can place as many statements in it as you desire.

If you are coming from other languages, you might be wondering about the
absence of the `break` keyword at the end of each case region.
In Whirlwind, this keyword is not necessary and is already implicitly
there. Because of this, the following is invalid.

    select (a) {
        case 1:
            Println("blah, blah, blah");
            break; // CONTEXT ERROR

        // -- snip --
    }

It is also important to note that a case does not have to be a literal.
In fact, it can be also be any expression.

    select (a) {
        case 2 ^ 3:
            Println("Do something");
        case randomValue():
            Println("Do something else");
    }

Finally, the case value **must** be the same type or coercible to the type
of the item being matched upon. This means if you are selecting over an
integer, you cannot have a string as one of the cases.

### Matching Multiple Values
Sometimes, you want the same code to run for multiple cases.  Consider
the example in below:

    select (name) {
        case "Bob":
            Println("Greetings!");
        case "Tom":
            Println("Greetings!");
        case "Sara":
            Println("Greetings!");
        case "Steve":
            Println("Go away Steve.");
    }

That code is super-redundant and defeats the whole purpose of the
select statement.  Fortunately, Whirlwind offers a way to shorten this
code: you are able to place multiple values in the case statement.
This is demonstrated in the **Listing 3.5**.

#### Listing 3.5 - Elimination of Periphrasis

    select (name) {
        case "Bob", "Tom", "Sara":
            Println("Greetings!");
        case "Steve":
            Println("Go away Steve.");
    }

You can store as many values as you want in a single case statement.
You just must ensure that they are separated by commas.

### The Default Case
Select statements also offer their own form of an else: the default case.
This is a case that will run only if no other case matches. It
must occur at the end of the select and is totally optional.

It is constructed by the `default` keyword, followed directly by a colon
and the case block.

    select (a) {
        case 1:
            Println("Hello, there.");
        case 2:
            Println("General Kenobi!");
        default:
            Println("Why are you here?");
    }

That default block will be run on all values that are not 1 or 2.

As said before, the default block must come **after** all the case blocks
meaning restructuring the previous statement to the following would be
cause a syntax error.

    select (a) {
        // SYNTAX ERROR
        default:
            Println("Why are you here?");
        case 1:
            Println("Hello, there.");
        case 2:
            Println("General Kenobi!");
    }


