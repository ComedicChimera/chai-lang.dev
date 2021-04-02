---
name: "Arrays and Lists"
weight: 12
---

## Arrays

An **array** is a fixed-length, ordered collection of similarly typed elements.  
They are the most basic way to store multiple items together -- such as a group
of names, products, or id numbers.  

Arrays use braces to enclose their elements separated by commas.  

    let primes = {2, 3, 5, 7, 11, 13, 17}

The data type label for arrays is `[]T` where `T` is the element type.  For our
primes array, this would be `[]int`.  Note that all the elements of an array must
be **unifiable** to the same type.  This essentially means all the types can be
coerced to one, common type.

    {2, "test"}  # TYPE ERROR

We can access the elements of an array using its **index** -- an integer value
that corresponds to the position of the item in the array.  These indices start
at zero and go until one less than the length of the array.  Here is table representing
the `primes` array and its indices

| Index | Value |
| ----- | ----- |
| 0 | 2 |
| 1 | 3 |
| 2 | 5 |
| 3 | 7 |
| 4 | 11 |
| 5 | 13 |
| 6 | 17 |

We use the **subscript operator** to get elements by their index.  This operator
is formed by a pair of brackets enclosing the index positioned after the array being
indexed.

    primes[1] # => 3

We can also assign to elements of an array using the subscript operator.

    let fruits = {"orange", "mango", "apple", "kiwi"}

    fruits[2] = "grape"

    println(fruits) # {orange, mango, grape, kiwi}

As you might guess, it is also sometimes useful to get the length of an array so we can
index it from the end -- eg. to get the last element.  We can use the `len` method of
the array data type to accomplish this.

    let last_prime = primes[primes.len()-1] # 17

Arrays can also be **concatenated** -- that is you can add two arrays together to get
a new array.

    let more_primes = primes + {19, 23, 29, 31, 37}

    more_primes[more_primes.len()-2] # 31

It is important to note that an array's length is *not* part of its type.  That means that
two differently sized arrays can be stored in the same variable.

    primes = more_primes

Moreover, all types in Whirlwind act like values unless explicitly referenced (a topic to be
discussed in a later chapter).  Therefore, we can manipulate `primes` and `more_primes` as two,
completely separate arrays.

    primes[0] = 1
    println(more_primes[0]) # 2

## Lists

A **list** is a variable-length, ordered collection of similarly typed elements.  Essentially,
they are like arrays that can be resized. 

They have the type label `[T]` where `T` is the element type and use brackets instead of braces
when they are declared.

    let countries = ["Canada", "Brazil", "Switzerland", "Korea", "Egypt"]

We can manipulate lists using indices the same way would arrays.  They also have a `len` method.

    countries[countries.len()-1] = "Ethiopia"

And we can concatenate them just like arrays.

    countries += ["Japan", "Mexico", "Paraguay", "Italy"] # we can concatenate in compound form as well

{{< alert theme="info" >}}List concatenation is far more performant than array concatenation:
generally, if you need to concat to sequences, use lists.{{</ alert >}}

However, unlike arrays, we can directly add, insert, and remove elements from lists.  We do these
using several common list methods.  The first and most common list method is the `push` method which
adds an item onto the end of the list.

    countries.push("USA")

    println(countries[countries.len()-1]) # USA

Notice that `push` directly mutates the `countries` list.  It doesn't copy it or overwrite it.  Below
are a few other common list methods with demonstrated usages.

    # Remove and return the item at the end of the list
    countries.pop() # => "USA"

    # Insert an element into a list at the given index
    countries.insert(1, "Moldova") # countries = ["Canada", "Moldova", "Brazil", ...]

    # Add an item to the front of the list
    countries.push_front("China")

    # Remove and return an item at a given index in the ist
    countries.remove_at(0) # => "China"

    # Remove all items that match the given value
    countries.remove("Brazil")

    # Remove and return an element from the front of the list
    countries.pop_front() # => "Canada"

    # Insert multiple elements into a list at the given index
    countries.insert_many(3, "Tibet", "Venezuela") # countries = [..., "Korea", "Egypt", "Tibet", "Venezuela"]

Although it seems like a lot of methods to remember, you will find you use most of these quite a lot,
and that their names correspond pretty directly (and conventionally) to their function.

## Sequences and Iterables

Lists and arrays both fit into a larger category of data structures called **sequences**.  A sequence is
any collection of elements that are ordered and can be accessed by indices.

Why is the grouping important?  First and foremost, it is because all sequences shared a number of common
properties.  For example, they can all can be subscripted (using the `[]` operator).  

Another common property of sequences is that they are all **iterable**.  What does that mean?  At a high
level, anything that it is iterable can be *iterated over* by something like a for loop -- ie. we can loop
over each element and do something with it.

{{< alert theme="info" >}}This is not a "technical" definition; rather, it is an abstrast consequence of the
actual definition.{{</ alert >}}

We have already glimpsed for loops in a previous section; however, now that we have a more complete understanding
of Whirlwind, we can dive more deeply into them.  A for loop is made up of two parts: the iterable variable(s) and
the iterable.  They are written like this:

    for item in iterable do
        ...

`item` is the name of our iterator variable -- it is where each element of the iterable will be stored as we iterate
over it.  `iterable` can be replaced by anything that is iterable such as a list.  The code below goes through every
item in our `countries` list from the previous section and prints it out on a separate line.

    for country in countries do
        println(country)

We can use loops to do all kinds of things such as adding up the elements of an array or finding 
the maximum element of a list.

    func sum_of(l: []int) int do
        let s = 0

        for i in l do
            s += i

        return s

    # assuming list contains only positive integers
    func max_of(l: [int]) int do
        let mx = 0

        for i in l do
            if i > mx do
                mx = i

        return mx

If we want to iterate through something by its indices, we can use a method called `indices` to do that.

    let nums = [5, 6, 2, 3]

    for i in nums.indices() do
        println(i, nums[i])

{{< alert theme="info" >}}This method exists for all iterables (even those not organized by indices directly).
It is part of a larger category of methods called iterable methods we will study in a later chapter.{{</ alert >}}

You might have already guessed by now, but the groupings of iterable and sequence correspond directly to two interfaces:
`Iterable<T>` and `Sequence<T>`.  The `<T>` denotes that they are *generic interfaces* -- a topic we will cover much more
in depth in a later chapter.  For the next few chapters, you are going to see that notation pop up more and more often:
we are going to be using these chapters to gently introduce the topic before diving into fully.  

The `T` inside the angle brackets is the element type of the iterable or sequence -- in much the same way that the `[T]` in
lists denotes the element type.  Using these two interfaces, we can rewrite our `sum_of` method to work on every kind of iterable.

    func sum_of(it: Iterable<int>) int do
        let s = 0

        for num in it do
            s += num

        return s

    func main() do
        println(sum_of([1, 2, 3])) # 6

        println(sum_of({1, 1, 2, 3, 5})) # 12

While these two interfaces may not appear supremely useful at the moment, understanding the ideas that underpin them is
critical to understanding Whirlwind's model of data manipulation as a whole: in future chapters, we will build on the ideas
that we started in this chapter.

## Slicing

The final topic on our introduction to sequences is **slicing**, a technique for taking the subset of a sequence.  Slicing
using the `[]` operator just like indexing; however, we are now have to specify two bounds: a start and an end index. 
The start index tells us where we are slicing from and the end index tells us where we are slicing to.  We separate these
two indices with a `:`.

    let elements = ["chlorine", "nitrogen", "xenon", "cobalt"]

    println(elements[1:3]) # ["nitrogen", "xenon"]

Notice that our slice goes inclusive to exclusive -- it includes the element at the start index and excludes the element at
the end index.  

We can elide the start index of a slice to slice from the start of a sequence up until a specific index.

    elements[:2] # ["chlorine", "nitrogen"]

Similarly, we can elide the end index to slice from a specific index to the end of the list.

    elements[1:] # ["nitrogen", "xenon", "cobalt"]

We cannot, however, elide both parts at the same time -- this would have no meaning.

    elements[:] # SYNTAX ERROR

Using slicing, we could easily rewrite our `max_of` function to work for both positive and negative integers (without having
to involve any constants like `MIN_INT`, repeating any handling of integers, or introducing any "complex" control flow).

    func max_of(s: Sequence<int>) int do
        let mx = s[0]

        # ignore the first value
        for item in s[1:] do
            if item > mx do
                mx = item

        return mx

We can also assign to slices of data structures to overwrite the data contained in them with the contents of the slice.

    elements[1:3] = ["lithium", "argon"]

Note that if the two slices being assigned are of different sizes, the assignment will fail causing a panic.

    elements[:2] = ["barium"] # RUNTIME ERROR
