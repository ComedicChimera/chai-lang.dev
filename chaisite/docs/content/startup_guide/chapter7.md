# Lists and Dictionaries

Up until this point, we have primarily dealt with single values: numbers,
strings, etc.  However, in most real programming, we want to deal more with
collections of values.  In this chapter, we will explore the bread and butter of
Chai's data manipulation toolkit: lists and dictionaries.

## Lists

A **list** is a variable-length, collection of similarly typed elements.  Lists
are written between `{}` and their elements are separated by commas.

    let primes = {2, 3, 5, 7, 11}

The list type label is `List[T]` where `T` is the element type of the list.

    def last_element(l: List[int]) int

You can get elements of lists using the **index operator** which is denoted with
`[]`.  An **index** is a numeric value corresponding to an elements position
within a list (or other data structure).  In Chai, indices start `0` for the
first elements and increase until one less than the length of the list.

    primes[0]  # => 2
    primes[2]  # => 5

> If you index outside the bounds of the list, you will get a fatal panic.

You can the length of a list using the `len` method.  A **method** is like a
function that is bound onto a specific type.  Methods are called using the `.`
operator placed after the value you want to call the method on followed by the
method name and its arguments.

    primes.len()  # => 5

As you will see in later chapters, Chai handles methods differently then you may
be used to in other programming languages.  Therefore, it is also to call
methods in **space form**.  This type of call would look like the following.

    List.len(primes)  # => 5

This notation is generally more verbose but occasionally useful for calling list
methods such as a `List.new` which cannot be called in **postfix form**.
    
    let empty_list: List[int] = List.new()

> The type label is necessary so Chai knows what the element type of the list
> is.  You can also use explicit type parameter specification which is much less
> repetitive; we will discuss that in the chapter on generics.

Returning to indices, we can also mutate lists individual elements directly by
placing the index operator on the left side of the assignment operator.

    primes[primes.len()-1] = 13  # set the last element to 13

As mentioned before, lists are variable length meaning we can add and remove
elements trivially using list methods.

    # remove the last element
    primes.pop()          # primes = {2, 3, 5, 7}

    # add a new element onto the end of the list
    primes.push(11)       # primes = {2, 3, 5, 7, 11}

    # remove the element at index 1
    primes.remove_at(1)   # primes = {2, 5, 7, 11}

    # insert an element at index 3
    primes.insert(3, 17)  # primes = {2, 5, 7, 17, 11}

    # remove all instances of a value within a list
    primes.remove(11)     # primes = {2, 5, 7, 17}

There are also several methods that are useful for accessing different
subsections of a list.

    # get the first element
    primes.first()  # => 2

    # get the last element
    primes.last()   # => 17

    # get all but the last element
    primes.front()  # => {2, 5, 7}

    # get all but the first element
    primes.tail()   # => {5, 7, 17}

Lists can also be concatenated using the `+` operator (like strings).

    let fruits = {"orange", "pear", "grape"}
    let vegetables = {"carrot", "broccoli", "asparagus"}

    println(fruits + vegetables)  # prints `{orange, pear, grape, carrot, broccoli, asparagus}`

You can also use the compound assignment (`+=`) operator to add a list onto another.

    fruits += {"peach", "apple"}  # fruits = {orange, pear, grape, peach, apple}

> List concatenation in Chai is very efficient in most cases: you should
> generally feel free to "concatenate at will".

## Slicing

**Slicing** is a special technique to get and set subsections of a list. It uses
the `[]` operator; however, instead of specifying just one index, you specify a
range of indices separated by `:`.  The start index is first and is inclusive;
the end index is last and is exclusive.

    let t = {3.141, 2.718, 1.414}

    println(t[0:2])  # prints `{3.141, 2.718}`

You can also elide the start index to denote that you want to start from the
beginning of list (index `0`).

    println(t[:2])  # prints `{3.141, 2.718}`

Similarly, you can elide the end index to slice until and including the end of
the list.

    println(t[1:])  # prints `{2.718, 1.414}`

You can also mutate slices of a list by putting the slice operator on the
left side of the assignment operator.  

    let letters = {'a', 'b', 'c', 'd'}

    letters[:2] = {'f', 'e'}

However, make sure that the slice you're assigning to is the same length as the
new list you are writing to the slice.  Otherwise, you will get a fatal panic.

    letters[1:3] = {'h'}  # FATAL PANIC

## Dictionaries

A **dictionary** is a variable-length collection of key-value pairs.  It is used
to organize data based on keys such as a name or a date as opposed to simply 
a position within the collection.

Dictionaries are also enclosed in `{}`; however, each element is a pair
separated where the key and value are separated by a `->` and pairs are
separated by commas.

    let ages = {
        "John" -> 42,
        "Emily" -> 23,
        "Nathan" -> 12
    }

The `ages` dictionary relates a person to their age.  We can then use the index
operator (`[]`) to access the values.  However, instead of passing in an index,
we pass in that value's key: for `ages`, this would be the person's name.

    ages["John"]  # => 42

You can also set keys in the dictionary by using the index operator on the left.

    ages["Nathan"] = 15  # set Nathan's age to 15
    ages["Emily"]++  # increase Emily's age by 1

The type label for dictionaries is `Dict[K, V]` where `K` is the key type and `V`
is the value type.

    def get_all_with_age(d: Dict[string, int], age: int) List[string]

> Note that the key type must be hashable.  This is true for all types we have
> learned thus far but will not be true (by default) for most user-defined
> types.

To add new entries to dictionaries, you can simply assign a value to the key you
want to add.

    ages["Natalie"] = 37

If you access a key that is not in the dictionary and are not assigning to it,
you will get a fatal panic.

    let age = ages["Killian"]  # FATAL PANIC

Dictionaries, like lists, have several methods that useful for manipulating
them. Here are a few of those methods.

    # get the number of pairs in a dictionary
    ages.len()  # => 4

    # check if a dictionary has a given key
    ages.has("Emily")  # => true

    # adds a new value for a key if that key is not in the dictionary
    ages.default("Barry", 67)  # adds `"Barry" -> 67`

    # updates a value for a key if that key was already in the dictionary
    ages.update("John", 44)  # updates John's entry

If you are having trouble keeping track of all the methods for lists and
dictionaries, don't worry: a) all methods are documented in the Standard Library
Reference and b) you will learn all the most useful ones through practice.  

## Copy Semantics

Unlike most languages you are probably used to, Chai implements collections as
value types meaning their contents are copied whenever they are passed to
functions, stored as unique variables, or otherwise moved around.  This means
that, for example, you can safely pass lists to functions and manipulate them
without worrying about inadvertently changing the original data structure.

    def add_one(l: List[int]) =
        l.push(1)

    def main() = do
        let l = {3, 2}
        add_one(l)  # l = {3, 2}

A consequence of this is that passing around collections by value can be quite
expensive: in a later chapter, we will study references and learn how to avoid
copying when unnecessary. 

> Chai will optimize out as many copies as it can (a technique called copy
> elision). However, it will always ensure your collections act like values so
> you don't need to worry about accidently manipulating them.