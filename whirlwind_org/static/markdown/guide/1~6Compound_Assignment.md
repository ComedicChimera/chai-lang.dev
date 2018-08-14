## Compound Assignment

In the previous section, you learned about assignment: an essential
piece of Whirlwind.
However, while normal assignment is extremely powerful, it can sometimes
be inefficient.  Consider you want to add `2` to a variable. Using
normal assignment, you would have to type the following:

    x = x + 2;

This code seems a little redundant as it requires you type to
`x` two times.  By constast, we can use compound assignment to
condense this operation significantly.  Below is an example of
code that does the exact same thing as the previous code, but it
utilizes compound assignment.

    x += 2;

The `+=` is called a compound assignment operator.  Whirlwind supports
a number of compound assignment operators, including arithmetic operators
and bitwise operators. The code in **Listing 1.9** shows examples
of compund assignment with other operators.

#### Listing 1.9 - Compound Assignment Operators

    $x = 2;

    x += 2;
    x -= 3;
    x *= 10;

    x ^^= 2;
    x &&= 2;

    x >>= 3;

Compound assignment operators can even be used in
multi assignment.

    x, y += 2, 3;

As you can see, compound assignment can be very useful.
However, there is one very common case that can be made
even more concise through the use of special operators.

### Increment and Decrement

Whirlwind provides two special compound assignment operators, called
**increment** and **decrement**. Their purpose is faily simple,
to add or subtract 1 from a numeric value.

    x++; // increment
    x--; // decrement

The `++` symbol signifies the increment operator, which adds 1 to a value.
Conversly, the `--` signifies the decrement operator, which subtracts 1 from
a value. 

Unlike the normal compound operators, these operators cannot be applied to
multiple values and are restricted to only operating on a single value at a time.

### The Expression Forms

Another key different between the increment and decrement operators
and the regular compound operators is that the the increment and decrement
operators also have an expression form.  This allows them to be used in expressions
as well as statements. 

When used in expressions, these operators have two forms: **prefix** and **postfix**.
The only different between the forms is the order of operation. The prefix
variants first add one to the value and then give you back the modified value.
Whereas, the postfix version gives you the original value and then adds
1 to it.

These two versions have a key syntactical different: the placement of the operator.
In the prefix version, you place the operator **before** the value.  But, in the
postifx version, you place the operator after the value.  **Listing 1.10** contains examples
of each.

#### Listing 1.10

    $(x = 4, y = 10);

    x = y++; // x is still 4, but y is now 11
    y += --x; // x is now 9 and y is 20

A final note regarding these operators is that they all follow normal
type checking rules. This means that you can't just apply the operators
normally, but instead must make sure the types match up.