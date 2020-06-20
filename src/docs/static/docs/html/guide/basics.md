# Basics

This section will walk through some Whirlwind fundamentals as well as building
and running your code.

## Hello World

As many of you already know, the first program any programmer most often writes
when starting out in a new language is *Hello World*.  The task of this program
is simple: print "Hello, World!" to the console.  Here's what that program
looks like in Whirlwind:

```whirlwind
import io::std

func main do
    std::println("Hello, World!")
```

The first thing to notice is the lack of any form of line or block punctuation.
Whirlwind is a whitespace aware language meaning that instead of using ugly
curly braces everywhere to signify blocks, we use indentation as well as a
variety of keywords to begin different blocks to denote the same idea. In this
case, the body of the function is contained inside the `do` block.  Also notice
the absence of semicolons: Whirlwind uses newlines to delimit statements so as
to avoid cluttering your code with needless symbols.

Next, notice that our program must begin inside of a **main function**.  This is
the function called whenever your program starts; therefore, any executable
program must have one.

With all that introductory rambling our of the way, let's build our program.
Assuming you followed the setup from the introduction and have placed the Hello
World program in your file, simply type the following command into your terminal
or console to build your code.

```terminal
whirl build .
```

Notice that we did not need to pass in a file name of any kind.  This is because
Whirlwind projects are structured in terms of packages which consist of all of
the files in a directory; therefore, the compiler will generate code in terms of
the entire directory (referenced with `.`).  After this command has run, you
should see an executable pop up in your folder with that name of your current
working directory (followed by the executable extension for your platform if one
exists).

> Your working directory name should be a valid Whirlwind identifier (beginning
> with a letter of an underscore followed by any number of letters, numbers, and
> underscores).  The compiler will still work if it isn't, but you will have
> trouble with imports later on.

If you run the executable, you should see the following appear in the terminal.

```terminal
Hello, World!
```

For the sake of concision, we can tell the compiler to build and run our
executable as a temporary so that we don't have to run it ourselves (and so that
we don't have a binary file lurking around in our directory; you can of course
change its build location to your liking: refer to the [CLI Reference](/docs/cli-reference)
for details).  To build and run all at once, simply use the following command.

```terminal
whirl run .
```

It will appear in the same form as the build command but with the word `run`
instead of a `build.`

## Variables and Constants

There are two primary kinds of named values in Whirlwind: variables and constants.
Both represent a named location in memory; however, variables are mutable and constants
are immutable.

> For the purposes of concision, the definition of a main function will be
> omitted in many code samples.  Note that variables and constants *can* be
> declared in global scope (ie. outside of a function), but that for simplicity
> we will assume that the definitions in this section will occur in the main
> function if no such function is provided (ie. just use context clues).

Variables are declared with the `let` keyword followed by a series of definitions
written like so:

```whirlwind
    // single variable definition
    let x = 10

    // multiple variable definitions
    let a = 4.5, b = 'a'
```

The variables do not have be of the same type and can be set equal to any
complex expression (on the right-hand side of the `=`).

```whirlwind
    let c = x + a * 2
```

Variables can be assigned to using the `=` operator.

```whirlwind
    b = 'b'
    x = 67 * 12
```

Several compound assignment operators also exist:

```whirlwind
    x += 10  // equivalent to x = x + 10
```

> A full list of valid compound operators can be found in the [Language Specification](/docs/lang-spec).  
> You will see a few common ones used in this guide.

It is also possible to assign to multiple values as once.  For example, the code
below swaps the values of `v1` and `v2`.

```whirlwind
    let v1 = 12, v2 = 10

    v1, v2 = v2, v1
```

Note that the right-hand side of the assignment operator is always *fully evaluated*
before the assignment it performed.

Constants are created similarly to variables but using the `const` keyword
instead of the `let` keyword.  

```whirlwind
    const pi = 3.14159265
```

Constants cannot be assigned to or mutated in any way.  

```whirlwind
    pi = 2.71828  // COMPILE ERROR!
```

## Data Types

Up until this point, we have somewhat ignored the subject of data types are
Whirlwind allows us to elide a type label whenever it can infer the type of the
variable (which you will find is most of the time).  However, if we wanted to be
more specific, we could place a type label between the name and initializer to
specify the type like so.

```whirlwind
    let my_var: int = 0
```

We have now specified that `my_var` has a type of `int` (a 32-bit integral
type).

There are many **primitive types** in Whirlwind.  Here are some of the most
common:

| Type Label | Type Info |
| ---------- | --------- |
| `int` | 32-bit signed integral type |
| `float` | 32-bit floating-point type |
| `double` | 64-bit floating-point type |
| `uint` | 32-bit unsigned integral type |
| `char` | 32-bit Unicode code point (uses single quotes - interpreted) |
| `string` | UTF-8 encoded, builtin string type (uses double quotes for interpreted strings or backticks for raw strings) |

Whirlwind is a strongly and statically typed language.  This means that all
types must be determined at compile-time and must remain constant for any given
value.  For example, the following code is invalid because the type of right-
hand side of the expression is not coercible to the type on the left-hand side
of the expression.

```whirlwind
    my_var = 5.2  // COMPILE ERROR: cannot coerce a float to an int
```

> Due to the complexity and power of Whirlwind's type system and inferencer, you
> may on occasion be confused as to why a certain type error is occurring.  In
> this case, it is often beneficial to put type labels in your code to better
> isolate the root cause of the error.

While the compiler can implicitly coerce between some types, some type conversions
will require an explicit type cast as any valid coercion could cause some unintended
behavior (eg. Whirlwind cannot coerce a float to an int).  To perform such a type
cast, we use the `as` keyword.

```whirlwind
    my_var = 5.2 as int  // NO ERROR
```

Note that casts can cause data to be lost or changed and should be used with
care (eg. the value of `my_var` is now the integer value `5` not `5.2`).

## Collections

Collections are builtin data structures used to store a finite number of values
(or key-value pairs).  The simplest collection is the array which is fixed-size
sequence of values of the same type.  They can be created like so:

```whirlwind
    let arr = {1, 2, 3, 4, 5}
```

The above code creates an array of integers called `arr`.  The type label for an
array is `[]T` where `T` is the type of arrays elements.  

We can calculate the length of the array using the array's `len` method.

```whirlwind
    let arr_length = arr.len()
```

We can access the elements of an array using the index (`[]`) operator.

```whirlwind
    let first = arr[0]
    arr[4] = 10
```

> Whirlwind also supports negative-indexing to access elements starting from the
> end of a collection (eg. `arr[-1]` = `arr[arr.len()-1]`).

All data types in Whirlwind are **pure-value types**.  This means that the user
can rely on data structures to act as values.  For example,

```whirlwind
    let arr2 = arr

    arr2[4] = 5

    arr  // {1, 2, 3, 4, 10} - no mutation
```

The full array was copied into our new `arr2` variable so that it acts like a value.

> We can circumvent this copying through the use of references and reference
> operators which will see later on in the guide.  Moreover, Whirlwind will
> avoid copying wherever posssible so that these semantics have minimal
> performance impact.

We can also use the slice operator (`[:]`) to extract a sub-array of the
existing array. Slices are inclusive to the first index and exclusive to the
last.

```whirlwind
    let arr3 = arr2[1:4]  // arr3 = [2, 3, 4]
```

The first operand to the slice operator can omitted to denote that the slice
starts at the beginning of the collection.  Similarly the last operand can be
omitted to denote that the slice goes until the end of the collection.  However,
at least one of the two must be present (unless you are performing a step-slice).

```whirlwind
    arr3 = arr3[:2] // arr3 = [2, 3]
```

The second common collection in Whirlwind is the list.  A list is like an array
but it is resizable (and optimized for resizing).  The type label for lists is
`[T]` where `T` is the type stored by the list. Lists are written like so:

```whirlwind
    let list = ['a', 'b', 'c']
```

The same operators the apply to arrays also apply to lists.  For example,
we can get the length of a list using its `len` method, and we can index and
slice them using the corresponding operators.

```whirlwind
    list[0]  // 'a'

    list[-1] = 'd'  // list = ['a', 'b', 'd']

    list = list[:2] // list = ['a', 'b']

    list.len()  // 2
```

However, lists also have several other useful behaviors.  For example, lists
can be concatenated (like strings).  

```whirlwind
    let list2 = list + ['c']  // list2 = ['a', 'b', 'c']
```

> Arrays can also be concatenated but such an operation is often significantly
> less efficient (both memory-wise and time-wise) than the equivalent with lists
> due to the differences in how lists and arrays are implemented.  Generally, if
> you want to concatenate, you should use a list.

As mentioned, lists are resizable and thus provide a number of methods for
manipulating their elements.  Here are a few of the most common:

```whirlwind
    // adds 'd' to the end of the list
    list2.push('d')

    // removes the last element from the list and returns it
    let end = list2.pop()

    // inserts an element at the given index
    list2.insert(1, 'e')  // list2 = ['a', 'e', 'b', 'c']

    // removes an element at a given index
    list2.remove_at(1)

    // removes all occurences of a given element from a list
    list2.remove('c')

    let list3 = [1, 1, 2, 3, 5, 8, 13]

    // remove one occurrence from a list
    list3.remove(1, n=1)  // could also be written as `list3.remove(1, 1)`
```

For a full list of list methods as well as more details about how these methods
work, refer to the [Standard Library Reference](/docs/std-lib).

The final collection worth of note is the dictionary.  The dictionary is an
ordered, compact key-value collection.  In essence, every value has an
associated key and no key may be repeated more than once.

Dictionaries have the type label `[K: V]` where `K` is the type of the key and
`V` is the type of the value.  Dictionary literals are written like so:

```whirlwind
    let names = {
        "Bob": 45,
        "Joe": 23,
        "Emma": 32
    }
```

As we can see, this dictionary pairs people's names with some numeric value (an
age for instance). We can access and mutate the values of a dictionary by using
the index operator but with a key instead of a index.

```whirlwind
    names["Bob"]  // 45

    names["Emma"] = 33
```

There are two ways of adding keys to a dictionary: assignment and manual
insertion.  The first involves simply setting a given key not already in the
dictionary equal to a value.

```whirlwind
    names["Steve"] = 52  // `"Steve": 52` is now a key-value pair in the dictionary.
```

The second is to use the dictionary's `insert` method.

```whirlwind
    names.insert("Jane", 41)  // `"Jane": 41` is now a key-value pair in the dictionary.
```

These two approaches are equivalent as should be clear for our examples.

> For those wondering, the reason for the existence of the `insert` method is
> that in Whirlwind, functions (and methods) are first-class citizens (as we
> will see later), and we are able to do more with a method than with an
> operator.

Here are some other fairly common dictionary methods.  We will see more of them later on
(when we cover loops and algebraic types).

```whirlwind
    // gets the keys in order (as a KeyView iterable)
    names.keys()  

    // gets the values in order (as a ValueView iterable)
    names.values()  

    // number of key-value pairs in the dictionary
    names.len()  // 5

    // tests if a dictionary has a given key
    names.has("Emma")  // true

    // tests if a dictionary has a given value
    names.has_value(12)  // false

    // tests if a given key-value pair exists in a dictionary
    names.has_pair("Tom", 19) // false
```

It is once again worth reiterating that dictionaries are ordered which means
that key-value pairs are stored in the order they were inserted in which we
will see more of in the next section on control flow.

## Tuples

The tuple represents fixed-size group of things where each type of thing has a
specific position within the data structure.  They are notated like so:

```whirlwind
    let tuple = (1, "abc", 1.414)
```

Notice that the elements can be of different types.  However, the tuple data
type is much more specific than that of a collection.  For example, the type
label for the above tuple is `(int, string, float)`.  

We can access the elements of a tuple using the `.` syntax.

```whirlwind
    tuple.0  // 1

    tuple.2  // 1.414
```

Tuples (like strings) are immutable and so we cannot assign to the elements of a
tuple.

```whirlwind
    tuple.0 = 2  // ERROR
```

We can also use some basic pattern matching (which we will discuss in more
detail later) to unpack the elements of a tuple into variables like so:

```whirlwind
    let (x, y, z) = tuple  // x = 1, y = "abc", z = 1.414
```

We can also use some simple pattern matching to unpack a tuple in assignment.

```whirlwind
    let tuple2 = (43, 67.12)

    x, z = tuple2  // x = 43, z = 67.12

    let tuple3 = ('a', 8.8, "eg")

    // `_` is used to mean that a value should be discard/ignored
    let (a, _, _) = tuple3  // a = 'a'

    // watch out for type mismatches!
    _, z, y = tuple3
```

As mentioned, we will explore more pattern matching when we discuss the match
statement.  However, since pattern matching is one of the unique and more
powerful capabilities of the tuple data type, it is worth mentioning here even
to introduce them.
