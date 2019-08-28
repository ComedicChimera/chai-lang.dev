# Comprehensions

A comprehension is a construct that is used to create
a new collection from a pre-existing collection based
on a set of parameters.

## Array and List Comprehensions

The most basic form of a comprehension is a single-element no-filter comprehension. Such comprehensions take the form of a either an array or a list comprehension. The type of comprehension simply determines which collection will be generated: an array comprehension will produce and unsized array as its output.

To create one of these comprehensions begin by typing either `{}` for arrays
and `[]` for lists to create the corresponding comprehension.

Next, we must describe the desired behavior for our comprehension.  A basic
comprehension has two parts: the action and the iterator, and they are described
in that order.  Let's look at simple comprehension that takes a range of numbers
and multiplies them by 2.

    [x * 2 | x <- 1..10];

The first thing to notice is that this is a list comprehension.  The second is
the action.  This is the `x * 2` component of the comprehension.  The result of
this expression will be what appears in the resulting list.  The third
thing to notice is the iterator.  This is the part of the comprehension where
you describe from what you desire to create a new collection and what variable
you want to use in the action.  This iterator syntax should already be familiar to you as you have seen it before: in the for-each loop.  The final and
simplest component of the comprehension is the `|` which separates the action
from the iterator.

The result of the above comprehension is a new list that takes the form
`[2, 4, 6, 8, 10, 12, 14, 16, 18, 20]`.

## Adding in a Filter

A filter is a clause that can be added to a comprehension which specifies what
condition a given element must satisify in order to appear in the resulting collection.  It comes at the end of the collection and preceded by the `when`
keyword.

    let list = [1, 2, 3, 5, 6, 9, 10];

    let newList = [x ~^ 2 | x <- list when x % 2 == 0];
    // newList = [4, 36, 100]

The above collection takes every item in `list` is puts its square in the resulting collection if it is divisible by 2.  

## Dictionary Comprehensions

A dictionary comprehension is the final form of a comprehension.  It is
similar to a normal comprehension with one small tweak.  Instead of having
one action expression it has two: one for the key and one for the value.

You declare it with `{}` and separate the action expressions with a `:`.

    let numbersAndChars = {x: cast<char>(x + 65) | x <- 1..10};

This creates a new dictionary with the numbers 1 through 10 as its keys and
the characters as its values starting at "A" and going until "J".
