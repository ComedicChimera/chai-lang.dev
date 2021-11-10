# Lists

A **list** is a *resizeable* collection of elements of the same type: lists give
us a way to store and operate on many things together as opposed to dealing with
single discrete values as we have up until this point.

## List Literals and Indexing

You can create a new list using curly braces enclosing a list (ha, get it?) of
comma-separated elements like so:

    let my_list = {1, 2, 3}

Lists, unlike tuples, can only contain elements of the same type; thus, we can
only store numbers in `my_list` and creating lists with a mixture of types will
result in an error.

    {4.5, "test"}  # TYPE ERROR

Lists use the data type `List[T]` where `T` is the type of the lists element
making them our first *generic type*: a type that has many different forms with
similar behavior.  For `my_list`, a viable type label might be `List[i64]`.

Lists can be of any size, and as long as they are of the same element type, they
are type equivalent: eg. the following it completely legal:

    my_list = {4, 5, 6, 7, 8}

This further distinguishes them from tuples which have to be a fixed size.

If we want to access the elements of a list, we can use the **index operator**
which, as the name might suggest, retrieves an element of a list based on its
**index**.  For those unfamiliar, an index is a number identifying where
something is in a collection.  In Chai, all indices start at 0 and increase in
order with the elements until the last element is reached.

The index operator is placed after the list we want to index.  It uses `[]`
inside of which we place the value of our index.

    let first = my_list[0]  # get's the first element

You can also mutate values at specific indices by placing the index operator
on the left-hand side of the assignment.

    my_list[1] = 2
    println(my_list)  # {4, 2, 6, 7, 8}

Lists can be **concatenated** (joined together) using the `+` operator.

    let my_list2 = my_list + {9, 10}
    println(my_list2)  # {4, 2, 6, 7, 8, 9, 10}

## List Methods

A **method** is a function that is bound to a specific type.  This methods
are (generally) related to the type they operate upon.

We can access methods of a given type using the `.` operator.

> We will cover methods in much more detail in another unit: this section serves
> as an introduction to them.

The first method we will look at is the `len` method of lists which returns the
length of the list.  Here is an example usage of that method.

    let primes = {2, 3, 5, 7, 11}

    println(primes.len())  # 5

As you can see, used the `.` to access the `len` method of the `primes` list and
called it as we would a normal function.

> This syntax is actually just syntactic sugar for a more verbose method calling
> pattern.  However, there is no need to further disambiguify it here.

We can conveniently use the `len` method to access the end of the list.

    primes[primes.len()-1] = 13

Notice that we have to subtract `1` since lists are zero-indexed: the last index
is one below the actual length of the collection.

Lists have several other useful methods as well.  The first of which is the `push`
method which adds a new element to the end of the list.

    primes.push(17)

    println(primes)  # {2, 3, 5, 7, 13, 17}

Here are several other useful list methods:

    # Insert a new element into the list at the given index
    primes.insert(4, 11)  # primes = {2, 3, 5, 7, 11, 13, 17}

    # Pop an element off the end of the list and return it
    primes.pop()  # primes = {2, 3, 5, 7, 11, 13}; => 17

    # Remove an element from a specific index in the list and return it
    primes.remove(0)  # primes = {3, 5, 7, 11, 13}; => 2

    # Remove all elements of the given value from the list
    primes.prune(11)  # primes = {3, 5, 7, 13}

    # Get the last element of the list
    primes.last()  # => 13

    # Check if a list contains an item
    primes.has(2)  # => false

Don't worry about trying to memorize them all straight away: lists are probably
one of the most commonly use constructs so you will pick them up fairly quickly
from just using the language.