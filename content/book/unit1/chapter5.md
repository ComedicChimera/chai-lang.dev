# Pattern Matching

This section introduces the concept of pattern matching in Chai and one of
Chai's most fundamental data structures that happens to be the basis for most
pattern matching.

## Tuples

A **tuple** is an ordered set of typed values.  They are written enclosed
in parentheses and separated by commas like so:

    let my_tuple = (5, "demo")

The tuple type label is simply a list of the element types enclosed in
parentheses.  For example, the type label for `my_tuple` would be: 
`(i32, string)`.  

We can access the elements of a tuple using the `.` operator followed by the
**index** of the element.  The index is a numbered position of the element
within the tuple starting at 0 for the first element.

    let el1 = my_tuple.0  # first element
    let el2 = my_tuple.1  # second element

Note that the index must be an integer constant.

> The indexing scheme used for tuples is unique since the indices have to be
> determined at compile-time.  As you will see later, most collections in Chai
> allow for variable expressions and use a different operator for indexing.  In
> this respect, tuples behave more like structures (which we will study later),
> than collections.

Tuples can contain any number of elements of disparate types.  For example, all
of the below are also valid tuples.

    (5, 6.6, "yo", 'a')

    ('7', 42, 3.14)

Finally, in Chai, tuples are **immutable**.  This means their elements cannot be
individually mutated.

    let t = (2, 3)
    t.0 = 1  # ERROR

You can, of course, mutate variables and fields holding tuples, just not the tuples
themselves.

    t = (5, 6)  # ok

## Tuple Unpacking

**Tuple unpacking** is the first and simplest kind of pattern matching: it
allows us to conveniently extract the values of tuples into variables.

For example, consider you had the following situation: you have two triples of
numbers, and you want to extract the elements to perform some arithmetic on them
(say a [dot product](https://en.wikipedia.org/wiki/Dot_product) for example).
Using what we know so far, here is what our code looks like:

    let triple1 = (2, 3, 5),
        triple2 = (5, -7, 8)

    # -- some code here --

    let dot = triple1.0 * triple2.0 + triple1.1 * triple2.1 + triple1.2 * triple2.2

Not only is that code very long and repetitive, but it is hard to read and
write. When writing this documentation, I accidentally mistyped the tuple
indices twice! 

This is a perfect use case for tuple unpacking!  How does it work?  The idea is
that we, at once, bind several variables to the elements of each tuple and work
with them directly.  Let's start by seeing this in action and then explain how
it works.

    # -- triple definitions above --

    let x1, y1, z1 = triple1
    let x2, y2, z2 = triple2

    let dot = x1 * x2 + y1 * y2 + z1 * z2

Much better!  Now we can actually read what our code is doing.  As you can see,
the mechanism is basically what it says on the tin: we set a variable each to
value at each position in the tuple using their position in the declaration to
tell what variable corresponds to which value.

You can do this in assignment as well.  For example, if we wanted to quickly add
a tuple to three values we already know, we could just use a compound assignment
operator and some tuple unpacking.

    # a, b, c defined somewhere up here

    a, b, c += triple1

Isn't that neat?  We just did an operation that would normally take three lines
of code, the word `triple1` typed three times, and a bunch of tuple indexes and
simplified it into one line of code that is easy to read and understand. 

However, what if you don't want to use every value of a tuple?  This is where
the `_` symbol comes in.  Whenever you use `_`, that communicates to the
compiler that you want to ignore whatever value would be stored into it.

For example, if we just wanted the first and third values of the triple, we
could use `_` in place of a variable name for the middle value.

    let first, _, third = triple1

Note that `_` cannot be used as a value anywhere.  It has no type and no value:
it is just a placeholder to tell the compiler you don't care about something.

    let y = _ + 5  # ERROR

You can also use `_` is assignment.

    _, b, c = triple2

And, you can work with tuples of any size when using unpacking.

    let t1 = (2, "test"), t2 = (6.6, 23, 'd', 812)

    let a, _, _, b = t2

    a, _ *= t1

    let d = a * 2

    _, b, _, d %= t2

    let _, s = t1
    let _, x, r, _ = t2

    println(s, x, r)

Tuple unpacking is one of the most commonly used tools in Chai: as you will see
later on, many functions that want to return multiple values return tuples that
you can then be conveniently unpacked to get their values.
## Match Expressions

## Test-Match Expressions

