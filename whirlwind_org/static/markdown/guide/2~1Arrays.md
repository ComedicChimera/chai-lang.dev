## Arrays

An **array** represents a continuous set of sequential data organized by
indices.  You can think of it as a simple data set with a set size.  All
collections, including arrays, are strongly typed; this means that
a collection can only contain one type.

***Note** It is possible to have multi-type collections through the use
of type unions, but that will be covered later.*

### Array Literals

Like simple types, collections also have literals.  Array literals consist
of a set of element separated by commas and wrapped in braces.  Here is
an example of string array literal containing 3 elements.

    {"Jeff", "Bob", "Steve"}

You can put as many elements are you want into an array literal, but they
must all be of the same type.  Here are some examples:

    {'a', 'b', 'c', 'd'} // an array of 4 chars

    // space is not important in Whirlwind, so you can condense elements like so
    {1,2,3,4,5,6} // an array of 6 integers

    {3.14, 6.28, 9.42} // an array of 3 floats

    {2 + 3, 4 * 5} // array values are expressions, so this is valid

    $x = 45;
    {x ^ 2, x * 2, x} // you can even use variables

Like all other literals, array literals are expressions, so the above code
is invalid from a normal standpoint.

You can also have arrays the contain a single element, though you will not
use them very often.

    {45 ^ 2} // that is valid as an array literal

### Array Type Designators

In addition to contains multiple pieces of data, all collections have
unique to type designators.  Here is the type designator for an
array of 5 integers.

    array[int, 5]

There are a couple important parts to the type designator.  First of all,
it begins with the keyword `array` followed by two values separated
by commas and wrapped in brackets.  The first argument is the type of the
array elements, in this case, `int`.  The second argument is the size
of the array.

As you may have guessed, arrays have a static size that cannot be changed.
In Section 2.3, we will look at **lists** which do have a variable
size.

You can use these type designators just like you use all other normal
designators, as shown by **Listing 2.1**.

#### Listing 2.1 - Arrays

    $arr: array[int, 4] = {1,2,3,4};

    $arr2: array[float, 3] = { 1.2, 4.5, 5}; // 5 is valid since type coercion applies to arrays

