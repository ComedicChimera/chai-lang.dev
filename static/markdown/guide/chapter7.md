# Tuples

A tuple is perhaps one of the most important data types in all of Whirlwind.
It represents an ordered set of values.  It can hold many different types and up to
any number of them.

## Basic Tuples

You can declare a tuple by placing a series of types in parentheses in the type specifier.

    let tuple: (int, str);

The above code creates a tuple storing an integer and a string.  It is important to note
that a tuple must contain at least two values.

Tuple literals are effectively lists wrapped in parentheses instead of brackets.  However,
tuple literals can contain many types instead of just one.  You would set the value of the above
variable like so.

    tuple = (-4, "ab");

You can access the individual members of a tuple using the following syntax.

    let first = tuple.0;
    let second = tuple.1;

Notice that you are placing the integer indices after the tuple `.` instead of brackets like
you would for a collection.  Additionally, it is important to note that tuples **cannot** be
negatively subscripted.

## Tuple Immutablility

Tuples are what we call an immutable data type.  This means that their members cannot be directly
modified.  For example, if you wanted to change just the first member of tuple, you would have to
reinitialize the entire tuple.

    tuple = (2, tuple.1); // no error

    tuple.0 = 2; // ERROR

The reason for this is that tuples are intended to be concise units of data.  They are not intended
to be thought of as a multiple pieces of data, but rather one discrete piece of data.

## Tuple Unpacking

Tuple unpacking is possibly the most important tuple operation in Whirlwind. It is the primary method by
which values are extracted for tuples and the primary context tuples are used.

Consider the code initializing the variables first and second above.

    let first = tuple.0;
    let second = tuple.1;

Using tuple unpacking we could make this operation far more concise and intuitive.

    let (first, second) = tuple;

As you can see, all tuple unpacking does is take each value in the tuple and pattern match it to
a value in an assignment or a declaration.  This is far more syntactically efficient and far easier
to understand.

As said before you can also unpack during assignment.  Consider you have already declared three variables
`x`, `y`, and `z` and a tuple representing an ordered triple called `triple`. Instead of assigning each variable
individually, you could simply unpack the tuple into the given variables and save yourself a lot of time.

    x, y, z = triple;

You can also omit a value if you don't want to extract by using `_` instead of a variable name.

    x, _, z = triple;

Now, the value of `y` will not be updated, but `x` and `z` still will receive a value.  Make sure to always provide
the correct number of slots when you are unpacking otherwise you will likely get an error.

    x, y = triple; // ERROR

Tuple unpacking is an extremely powerful tool and will soon become a staple of your Whirlwind toolkit.  Just be careful
that you don't frivolously use this tool as there are cases when a simple `.0` might suffice.