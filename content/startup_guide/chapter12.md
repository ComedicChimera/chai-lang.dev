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





