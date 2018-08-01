## Lists

A **list** is similar to an array with one small difference: it is
resizable.  Just like arrays, a list can only contain elements of
the same type, and they can be subscripted and sliced in the exact
same way.

### Declaring Lists

Unlike arrays, they are declared using the square brackets,
instead of braces.  They also have a slightly different type
specifier that looks like this: `list[type]`.  Notice that
you do not have to provide a size.

#### Listing 2.4 - Declaring a List

    $myList = [1, 2, 3, 4, 5];

    $strList: list[str] = ["apple", "orange", "dragonfruit"];

    $listOfLists: list[list[float]] = [[1.23, 3.45], [4.5, 0.4]];

Like arrays, lists must also contain at least one element as a literal.

    $emptyList: list[int] = []; // invalid

    $singleList = ['a'];

However, because lists can be resized, it is possible to declare an empty
list or reduce a list to zero elements.  Here is an example of an empty
list declaration.

    $emptyList: list[int];

Since no literal was involved, the list is valid as an empty list.
It is simply impossible to declare an empty list literal.

### Resizing Lists

A list can be resized a number of ways.  The first is by assigning
a list of a different size or a list slice to the list.

#### Listing 2.5 - Resizing a List

    $list1 = [1, 2, 3];

    list1 = [4, 5]; // valid, list resized

    list1 = [4.5]; // TYPE ERROR, list types must be the same

    list1 = [6, 7, 8];
    list1 = list1[:2]; // list slice, list resized

You can also resize a list by calling some of its **builtin methods**.
A method is simply a function that is attached to an object.
The term builtin means that the methods are builtin to the language
itself. Here is an example of calling the builtin method `push`
on a list.

    $primes = [2, 3, 5, 7, 11, 13];

    primes.push(17); // calling a builtin method with the value of 17

The `push` method is used to push an element to the back of a list.
So the primes list above now looks like this: `[2, 3, 5, 7, 11, 13, 17]`.

You can use the push method to add multiple elements to a list
at the same time like so:

    $letters = ['a', 'b', 'c'];

    letters.push('d', 'e'); // letters becomes ['a', 'b', 'c', 'd', 'e']

However, because push mutates the underlying list, it cannot
be used on constant lists so the following is invalid.

    @constList = [0.25, 0.5, 0.75, 1.0];

    constList.push(1.25); // ERROR, use of mutable method on constant object

Lists also have several other methods that can all be used like push
and all follow similar rules.

#### Listing 2.6 - Builtin Methods

    $fibonacci = [0, 1, 1, 2, 3, 5, 8, 13];

    fibonacci.pop(); // removes the last element from a list

    fibonacci.shift(); // removes the first element from a list

    fibonacci.unshift(0); // adds an element to the beginning of a list (like push)

These methods all have special additional arguments and properties. So if
you want to learn more about these methods go look at the standard library reference
and play around with them.

### Insertion and Deletion

There are two other list methods that are important to be aware of.
There are `insert` and `remove`. Each is pretty self-explanatory, but
let's show them in context so you know how to use them.

    $evens = [2, 3, 4, 6, 8, 12];

    // lets fill the hole between 8 and 12
    evens.insert(4, 10); // insert 10 at index 4

    // lets get rid of that one uneven element
    evens.remove(1); // remove the element at index 1

Just like the previous methods, these both have special arguments that
you can use the documentation to find out about.

### Method Typing
All of the method mentions above are type sensitive.  This means
that they only accept element types matching those of in the list.
For example, you cannot push at string to an integer list without
getting a type error.

This is accomplished through something called **templating**, a concept
we will discuss much later.  Just know that these methods a type sensitive
and list specific.