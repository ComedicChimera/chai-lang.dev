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
    t.0 = 1  # COMPILE ERROR

You can, of course, mutate variables and fields holding tuples, just not the tuples
themselves.

    t = (5, 6)  # ok

## Tuple Unpacking

## Match Expressions

## Test-Match Expressions

