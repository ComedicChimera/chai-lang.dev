---
title: Dictionaries
weight: 13
---

## Understanding Dictionaries

A **dictionary** is a collection that maps keys to values.  It is the final fundamental 
collection in Whirlwind and is unique from the other collections in that it is ordered
by keys as opposed to indices.

To understand what we mean by this, consider the simple example of running a shop that
sells fruits.  You would need to map each fruit to a given price.  You could do this with
a match statement, but that would not be very flexible.  What if you wanted to add a new
fruit with modifying your code.  This is where a dictionary comes in.  

    let fruit_prices = {
        "orange": 3,
        "strawberry": 2,
        "dragonfruit": 7
    }

The `fruit_prices` variable holds a dictionary that maps string keys to integer values.
In this case, the keys are fruits and the values are prices.  Notice that the dictionary
uses braces to enclose its key-value pairs which are separated by commas.  The key and
the value are separated by colons.

If we wanted to get the price of an orange, we could simply index the dictionary with the
key `"orange"`.

    fruit_prices["orange"]  # => 3

Similarly, we could change the price of a strawberry by assigning to its value.

    fruit_prices["strawberry"] = 1

If we try to get a key that is not in the dictionary, we would get a runtime error.

    fruit_prices["grape"]  # KEY ERROR

However, we can use the dictionaries `get` method to get an option type instead.

    fruit_prices.get("grape")  # => None

Or, we can use the `get_default` method to return a default value instead of `None` if
the key isn't present.

    fruit_prices.get_default("tomato", 5)  # => 5

Adding keys to the dictionary is easy, just assign to the index you want to add.

    fruit_prices["tomato"] = 5

    println(fruit_prices["tomato"])  # => 5

Assigning to a value in a dictionary if that value does not exist has the effect of
adding a value to the dictionary.

Here are a couple other useful dictionary methods:

    # Remove a key (and value) for the dictionary and return the value
    fruit_prices.remove("strawberry")  # => 1

    # Check if a dictionary has a key
    fruit_prices.has("tomato")  # => true

    # Get the number of key-value pairs in the dictionary
    fruit_prices.len()  # => 3

Of course, dictionaries themselves are values and can be passed around as such. 
Dictionaries use the type label `[K: V]` where `K` is the key type and `V` is the value
type.  Our `fruit_prices` dictionary has the type label `[string: int]`.  

## Dictionaries as Iterables

Dictionaries, like arrays and lists, are iterable.  When we iterate over a dictionary
by default, we iterate over its keys.

    for fname in fruit_prices do
        println(fname)

We can iterate over the values using the `values` method of the dictionary.

    for price in fruit_prices.values() do
        println(price)

{{< alert theme="info" >}}There is an equivalent `keys` method that is occasionally
useful for isolating the keys of a dictionary.{{</ alert >}}

It is often useful to iterate over both the keys and values at the same time.  To
do this, we use the `pairs` method.  In order to understand how to use this method,
we need to introduce the idea of a **compound iterator** which allows us to iterate
over multiple items at the same time.  To use such an iterator, all you need to do is
specify multiple iterator variables separated by commas.

    for fruit, price in fruit_prices.pairs() do
        println(fruit, price)

Finally, it is worth noting that dictionaries are **ordered**.  This means that when
you iterate through the elements of the dictionary, you iterate through them in the
order you inserted them or initially specified them (and that this order is deterministic 
every program run).

{{< alert theme="info" >}}This is in contrast to a hash map which does not guarantee
order.{{</ alert >}}

## Hashability

In order to be a key in a dictionary, a key must be **hashable**.  In Whirlwind, this means
that it implements the `Hashable` interface by providing a `hash` method.  A hash is simply
a numeric value that represents the value of the data structure.  It is used to facilitate
lookups in the dictionary.  

Most builtin types are hashable; however, there are a few that aren't.  Most notably, none of
the builtin collections are hashable which means they can't be used as keys.  If you try to use
them, you will get an error.

    [[]int: int]  # ERROR

    {
        [4, 5, 3]: "test"  # ERROR
    }

You will rarely encounter hashability as an issue, but if you do, it is good to be aware of
the concept and the solution: provide a `hash` method.  

Remember that `hash` must return a unique, 64 bit unsigned integral value that represents the
data structure you are hashing (eg. the hash of `"hello"` is always the same).  If you are
unsure of how to implement a `hash` method, most of the time, Google and Stack Overflow are
your two best resources. 
