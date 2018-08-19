## Tuples

A tuple is an immutable set of types used to represent a couple of
related items like a point or a name and age.

### The Tuple Type

Tuples have a special data type since they are composed of multiple types
instead of one.  The tuple type literal is just a set of data types
separated by commas and wrapped in parenthesis.

    $myTuple: (int, double);

The tuple type must be comprised of at least 2 types.

### Tuple Literals

Tuple literals are expressed as sets of expressions wrapped in 
parentheses and separated by commas.  Notice that they
are structures similarly to the their data type.

    myTuple = (2, 3.4);

Tuples are also type sensitive. This means that assigning
a tuple to a tuple comprised of different types
or different numbers of element will cause a type error.

    myTuple = ('a', "bc"); // TYPE ERROR
    myTuple = (1, 1.5, 2); // TYPE ERROR

Tuples are stored as contingent values in memory,
similar to arrays so internally, these are
equivalent.

    (1, 2, 3)
    {1, 2, 3} // stored the same

### Tuple Unpacking

Tuples can be subscripted to get individual elements.
However, you do not use the normal subscript
operator. Instead you must use **tuple unpacking**.

This is done by accessing each element to a variable.
Consider the tuple `(2, -1)`. We could get the values
from this tuple by setting each value equal to
a variable.  **Listing 2.9** shows how to do this.

#### Listing 2.9 - Tuple Unpacking

    $point = (2, -1);

    $(x, y): int;
    x, y = point; // tuple is unpacked into x and y

This is a quick and easy method to get values from
a tuple. However, it can be simplified using
**tuple initializers**. These let you set
unpack the values on declaration as opposed
to through normal assignment.

So the code in **Listing 2.9** could be
rewritten like so:

    $point = (2, -1);

    $(x, y) = point;

However, this can be problematic if the
variable type extensions to not
match with the types of the values
they are being assigned. So it is important to ensure
that the typles match up.

### Ignoring Tuple Values

Sometimes, you don't want to extract all of a
tuples values and instead only want one or two
of them.  In this case, you can use the
`_` variable.

This is a special variable that signifies an
ignored value. This variable cannot be used normally or
assigned to or declared. In fact, it is already implicitly in
scope when the program begins.

This variable can be used in both declaration and
assignment and follows the basic rules of
**pattern matching**. **Listing 2.10** demonstrates
how to use this ignored variable and pattern matching.

#### Listing 2.10 - The _ Variable

    $vec3 = (0, 1, 2);

    // get the second value
    $(_, second, _) = vec3;

    $first: int;
    // get the first value
    first, _, _ = vec3;

    vec3 = (3, 4, 5);
    // get both first and second values
    first, second, _ = vec3;

Notice that the number of elements must always match the
number of variables on the left side.  This naturally
means the creating tuples of more than 3 elements can
be very unwieldy.  

For this reason, it is important to
keep your tuples small and not over-use them.

### Tuple Immutabilty

Notice that I have not mentioned how to assign
to individual values of a tuple.

This is because tuples are **immutable**.
They're constituent values cannot be
modified.  If you want to change the value
of a tuple, you must change all of
its values.

This has to do with design and implementation of tuples.
They are designed to be represent single, one off structures.
Also, making them immutable saves on compile time costs and
prevents our syntax from being too obnoxious.

Tuples are intentionally designed to not be very reusable.
Later on, you will be introduced to structs which are
much easier to manipulate and are very reusable.

### Summary

That is it for the chapter on Collections.  You have learned a
lot and been introduced to a lot in this chapter.  This is
one of the most important chapters because the concepts taught in it
will be reused and expanded upon as we progress through the book.

Next chapter we will begin learning about control flow and start building
in program logic.