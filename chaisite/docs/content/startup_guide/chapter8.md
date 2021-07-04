# Loops and Comprehensions

A **loop** is a control flow construct that repeats the code inside its body.
However, on top of this fairly standard behavior, Chai's loops also function as
**generators**, constructs used to generate sequences of values.  

## While Loops

A **while loop** repeats the content of its body until the condition in its
header is no longer true.  They begin with the `while` keyword and have
identical body syntax to that of if and match expressions.

    while cond do
        ...

    while cond -> ()

A common pattern with while loops is create, repeat, mutate pattern which looks
like the following:

    # create
    let i = 1
    while i < 11 do
        # repeat
        println(i)

        # mutate
        i++

The code above counts up from `1` to `10` printing out the numbers as it goes.
You might notice that the code is a little bit verbose: it is spread out over
several lines, and `i` is not needed outside of the loop yet it is declared
above it.  We can fix this by combining multiple statements into the while
loop's header.

    while let i = 1; i < 11; i++ do
        println(i)

The above code is equivalent to the previous listing but way shorter. We can
also elide different pieces of the header if we only want two parts of the
pattern.

    def count_up(start, end: int) = do
        while start <= end; start++ do
            println(start)

Another common pattern in programming is the [do-while pattern](https://en.wikipedia.org/wiki/Do_while_loop).
While Chai doesn't support this pattern directly, we can use our three part
while loop to pretty easily accomplish a similar task.

    while let ok = true; ok; ok = condition do
        ...

> This is actually a pattern that Chai will recognize and optimize into a more
> efficient do-while construct during compilation.

## For Loops

A **for loop** iterates over the elements of a sequence, executing its body for
each element.  A common example of a sequence would be the lists that we studied
in the previous chapter.

For loops begin with the `for` keyword followed by an iterator pattern, the `in`
keyword, and the sequence to iterate over.  For example, if we wanted to iterate
over the elements of a list, we could do so easily using for loops.

    for item in list do
        println(item)

The `item` variable stores the value of each element as the loop iterates.  

Another common sequence is a **range**.  Ranges are sequences of numbers
spanning from a start value up until an end value.  The start and end values are
separated by a `..`.  We can use ranges to express the idea of counting upward
much more simply then we did using while loops.

    for n in 1..11 do
        println(n)

The code above prints out the numbers `1` through `10`.

> This method is generally preferable to the while loop method as its is much
> more concise and clear.

Many types have methods that give us sequences.  For example, if we wanted to
iterate through the indices of a list as opposed to its elements, we could use
the `indices` method.

    for i in list.indices() do
        ...

Dictionaries are also sequences.  By default, when we iterate over them, we
iterate over their keys.

    for key in dict do
        ...

However, we can also iterate over their values using the `values` method.

    for value in dict.values() do
        ...

So far, we have only used single-value iterator patterns: names.  However, for
loops support pattern matching meaning we can actually iterate over multiple
values at the same time.  

Consider the `pairs` method of dictionaries, this returns a sequence of tuples
containing the key value pairs of the dictionary.  Using pattern matching, we
can quickly iterate over these pairs, extracting the variables as we go.

    for (key, value) in dict.pairs() do
        println(key, "->", value)

We can also perform a similar operation on lists using the `enumerate` method to
iterate through both the indices and elements at the same time.

    for (ndx, item) in list.enumerate() do
        println(ndx, "=", item)

## Sequences



## Break and Continue

## Comprehensions