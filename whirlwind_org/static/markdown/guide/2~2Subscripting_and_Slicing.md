## Subscripting and Slicing

Arrays are pretty useless if you can access or modify their elements.
Luckily Whirlwind provides two tools for doing just that: **subscripting**
and **slicing**.

### Subscripting

You can access elements of arrays using the **subscript operator**. This
operator consists of an index put between brackets immediately following
the array.  An **index** is an integral value representing some position
in an array.  Array indices start at 0; this means that first element
of an array is at position `0`. **Listing 2.2** contains examples of subscripting.

#### Listing 2.2 - Subscripting

    $perfectSquares = {1, 2, 4, 9, 16, 25};

    // subscripting is also an expression so it cannot stand on its own
    $firstElement = perfectSquares[0];

Just keep in mind that index must be within the bounds of the array. So the
following is invalid for the array above.

    perfectSquares[10] = 100; // ERROR: Subscript outside bounds of array

However, you can use negative indices to subscript the array from the back as shown below:

    $lastElement = perfectSquares[-1];

Negative indexes start at -1, since there is no negative 0.

### Subscript Assignment

You can also use subscripting to modify elements of arrays by placing
the subscript on the left side of assignment.

#### Listing 2.3 - Subscripting to Modify Array Elements

    $names: array[str, 3] = {"Tom", "Steve", "Bob"};

    // subscript assignment
    names[0] = "Jeff";

    // swap two elements
    names[0], names[1] = names[1], names[0];

You can also use negative indexes in assignment.

    names[-1] = "Emily";

This kind of assignment also behaves like a statement.

Index values can also be expressions like so:

    $a = -2;
    names[a + 2] = "John";

### Slicing
A **slice** is similar to a subscript except it allows to select multiple elements
of an array as their own sub-array.  The slice syntax is similar to subscript
syntax, but it contains multiple values.

The slice has 3 parts: the start, the end, and the step.  Each part is separated by a
`:` and represents an index in the array.  The start is inclusive (start index included), and the
end is exclusive (end index not included).

    $arr = {1, 2, 3, 4, 5};
    arr[0:3:1] // first 3 elements

In a normal slice, the step parameter can also be
excluded completely, and will default to 1.

    arr[0:3];

If the beginning of the slice is 0, then the start can
also be excluded, **but** the delimiting colon must still
be there to allow the compiler to distinguish a slice from a
subscript.

    arr[:3]; // still first 3 elements

Similarly, the end can also be excluded if it is the end of the array,
however, the colon must still be there (for the same reason as before).

    arr[3:]; // last 2 elements of the array

As always, the step must be delimited by a colon as well.

    $nums = {0, 8, 4, 3, 0, 2, 3, 5};

    $lastBy2 = nums[2::2];

Normally, the start must always be less than the end, but when using a negative step,
the inverse is true.  Negative steps allow you to slice the array backwards.

    $first5bwds = nums[5:0:-1];

    $first5 = nums[5:0]; // invalid, start > end (step not negative)

    $first7bwds = nums[:7:-1]; // invalid start < end (step negative)

As with indexes, slice parameters are also expressions.

    $x = nums[:2 * 2:4 - 2];

    $start = 3;
    $y = nums[start:];

### Slice Assignment
You can use slices on the left hand side of assignment to set
whole regions of an array to a specific value as shown in **Listing 2.3**.

#### Listing 2.3 - Using Slices to Modify Arrays

    $fibonacci = {1, 1, 0, 2, 3, 5, 8};

    fibonacci[:3] = fibonacci[2:0:-1]; // let's fix the sequence

As you can see, both slices and subscripts are very important.  They will be some
of you most used operations and are some of the most powerful operations
in all of Whirlwind.