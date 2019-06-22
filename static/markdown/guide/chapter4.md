# Collections

A **collection** is something that is used to store multiple similarly typed values in a
sequence.  Collections are iterable and they are subscriptable.

## Arrays

An array is a fixed length collection that can store items of the same type.  You declare
an array using braces and items delimited by commas.

    let arr = { 2, 3, 4 };

That is an array of unsigned integers.  The part in braces is called an array literal.
Array literals must contain at least one item in order to determine their type.

The array data type is declared using the following type alias syntax.

    let arr: [3]uint;

The first item in the brackets is the data type of the elements and the second value is
the length of the array.  Unlike in many other languages, arrays can only have a fixed length
and it must be included in the data type.

> It is also possible for an array data type to be unsized.  This will allow any array of
> the element type to be considered equivalent, regardless of its size.  Unsized arrays are
> designated with the `[]T` type specifier.

You can get the length of an array (or any other collection) by using the `len()` method.

    let l = arr.len(); // l = 3

All a method is a function, like the one's we used in Chapter 1, that is bound to an object.
The `len()` method takes no arguments and so can be called with just parentheses.  Notice that
instead of using the `::` operator, we use the `.` operator to call methods.

## Subscripting and Slicing

Subscripting is select a single element from a collection and slicing is selecting a set of
elements from a collection.  Subscripting is done with the bracket syntax.

    let el1 = arr[0]; // first element of the collection

Subscripting for all collections starts at 0 and can accept negative indices.

    let last = arr[-1]; // last element of the collection

Finally, you can only subscript within the bounds a collection, this goes for both negative and positive
indices.

    arr[3]; // ERROR: bounds exception

Slicing is similar to subscripting except it allows to select multiple elements instead of just one.
Basic slices accept a start and an end index and all values starting from the first index and going up
until the second index.  The indices must be separated by a colon within the slice.

    let a = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

    a[1:5]; // returns the array {2, 3, 4, 5}

You can also omit the first bound if you want the slice to start at the beginning of the array and omit
the ending if you want the slice to go to the end of the array.

    a[:3]; // first 3 elements
    a[5:]; // last 5 elements

As with indexing, you can also use negative indices, but be careful.  Slicing will by default assume the first
index is the start, meaning you can very often end up with empty slices if you aren't careful with
negative indices.

    a[-7:-1]; // correct negative slice (7th from last to last)

    a[-1:-7]; // incorrect negative slice

Slices are also bounds checked and will throw a similar bounds exception if either index is out of bounds.

The final argument to a slice is the step.  The step defines how many steps should be taken between
collecting elements and in which direction the collection should go.

    a[::2]; // get every 2 elements in the list

    a[7:1:-1]; // get all elements between 7 and 1 in reverse

When stepping, positive values mean collect forwards and negative values mean collect backwards.  Notice that
when using negative values, the start index is **greater** than the end index.  This makes sense given
the fact the slice is collecting in reverse.  Finally, if a step is too great, it will be considered an
out of bounds exception.

## Ranges

Ranges are special type of array that allow you to create ordered series of numbers.
By default, they generate as an unsized array and are treated just like one. To create
a range, you must use the `..` operator, bounded by two integers.  

    let r = 1..10;

The above range creates a list of numbers 1 through 9.  Notably, Whirlwind ranges generate
including the leading bound and excluding the ending bound.  Ranges only go in one direction:
least to greatest and will generate an empty range if any other pairing occurs.

> Unlike some languages, Whirlwind only allows the creation of ranges of integers.

Ranges are high level expressions meaning in order to access the range as a whole, you must
enclose it in parentheses.

    let r = 1..9[2]; // ERROR: you can't subscript integers

    r = (1..9)[0]; // A OK

This can be somewhat inconvenient, but it allows you to put complete expressions on each side of the
range as opposed to only simple literals.

    let (a = 10, b = 17);

    let r = a * b..a ~^ b; // me me big range

This format makes editing the bounds of ranges far easier, but it also introduces some ambiguity which can
make code confusing to read so just make sure to not make the range expressions too complex

## Lists

A list is a like an array that can be resized.  Similarly, the size of a list is not factored into
the data type.  You declare a list a like so.

    let lst = ['a', 'b', 'c'];

Notice that lists use brackets instead of braces.  The list data type is declared similar to that of
an array, but with the type inside the brackets instead of the size.

    let lst: [char];

Just like arrays, it is impossible to declare an empty list literal, but you can use the pure data type
form to declare an empty list.

Lists, as collections, can be subscripted and sliced like all other collections.  As said previously,
lists can be resized.  There are several methods used to do this.  

    lst.push('d'); // adds new item onto end

    lst.remove(3); // remove item at index 3 (bounds checked)

    lst.pop(); // remove item from end

    lst.insert(0, 'e'); // insert item at front of list

Lists also define a `len()` method, akin to that of the array's `len()` method.

## Dictionaries

The final core collection is a key-value collection.  Instead of using indices, it uses
keys as element addresses.  Dictionaries are declared like so.

    let d = {
        'a': 1,
        'c': 2,
        'f': 3
    };

Dictionaries use braces like arrays, but have colons separating key from value and commas
in between the pairs.

The dictionary data type looks like so and is the only way to declare an empty dictionary

    let d: [char: int];

Dictionaries work a little bit differently in terms of subscripting.  Firstly
the subscript is done with the key name instead of the index.  Additionally, you can
index values that don't currently exist yet in the dictionary as long as you are setting
the value and not getting it.

    let k = d['a']; // returns 1

    d['j'] = 4; // ok, because the value being set

    k = d['o']; // ERROR: key doesn't exist

Dictionaries cannot be sliced, but they are ordered.  Meaning if you iterate through the keys,
they will always be in the same order they were inserted in.

Dictionaries also cannot have duplicated keys, so declaring a dictionary in the following form
is invalid.

    let d2 = {
        "pi": 3.14
        "e": 2.71,
        "root2": 1.41,
        "pi": 3.141   // ERROR: pi declared twice in dictionary
    };

Dictionaries also have several methods used for inserting and managing keys.
The first method is the `get()` method.  It works similarly to the get form
of the subscript, but it returns `null` of the value doesn't exist instead
of throwing an error.

    let j = d.get('j'); // j = 4

    let o = d.get('o'); // o = 0 (null)

The `get()` method also has a second optional argument that specifies what the
value will be if none is found.  As seen above, this defaults to null.

    let o = d.get('o'); // o = 0

    o = d.get('o', 5); // o = 5 

Dictionaries do have a `len()` method that simply returns the number
of keys in the dictionary.

Dictionaries also have a method called `has()` that tells if a dictionary contains
a given key.

    d.has('a'); // true

    d.has('o'); // false

> With regards to `has()` and `get()` methods, other collections have them,
> but they have different behavior.

Finally, dictionaries have two additional methods, `keys()` and `values()` that
give a list a of the keys and values respectively.

    d.keys(); // ['a', 'c', 'f', 'j']

    d.values(); // [1, 2, 3, 4]

The keys and values will be in the insertion order because
dictionaries are ordered.
