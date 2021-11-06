# Pattern Matching

This chapter introduces the concept of pattern matching in Chai and one of
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

**Match expressions** are Chai's primary mechanism for performing pattern
matching.  We have already seen the basic idea of pattern matching at work in
tuple unpacking, but in order to understand the idea more generally, we need to
first discuss what a pattern actually is.

### Patterns

A **pattern** represents the semantic "shape" of a given construct.  Patterns
can match many different kinds of constructs depending on their initial shape.
In general, patterns are made up of three elements:

1. Values
2. Wildcards
3. Shape Templates

To understand how each of these elements work, let's start with the simplest
pattern imaginable:

    4

Although it may not look like it, that is, in fact, a pattern.  It matches
the value `4`.  

We can even test this in Chai using a special case of the match expression:
the **test-match** expression which is written simply:

    value match pattern

It returns a boolean if the match is successful and false if it isn't.  So, if
apply this expression to different values using the simple pattern above we can
explore Chai's pattern matching logic.

    4 match 4  # => true
    5 match 4  # => false

In the pattern `4`, the single element of that pattern, namely `4`, would fall
into the category of value element.  Value elements in patterns are matched
using standard equality comparisons.

However, this is fairly basic, lets take our patterns a step further an
introduce **wildcards** which match *any* value they are compared against.
Wildcards are denoted using the `_` symbol.
    
    4 match _               # => true
    5 match _               # => true
    "test" match _          # => true
    false match _           # => true
    (5, 4.5, 'a') match _   # => true

> You might think that the wildcard would introduce problems for the type system
> since they can match any type, but because Chai only matches patterns to
> values with known types, the actual "type" of the wildcard is always trivially
> inferrable.

Okay, so that's great and all, but it doesn't seem that useful.  In this simple
case, you would be right to think that.  But, we have to cover the building
blocks before we can cover the big idea. 

The real power of pattern matching comes from **shape templates** which when
combined with the aforementioned features can lead to some really clever control
flow.  Shape templates are essentially ways of structuring patterns so that they
match specific data structures, allowing you to concisely match over their
internals.  The simplest set of shape templates are, unsurpisingly, those for
tuples.  Let's first start with an example before diving into the weeds.

    (5, 4) match (_, 4)  # => true

The above pattern uses a shape template for a tuple with 2 elements.  It uses a
wildcard in place of one element to denote that that element can be anything and
a specific value in the other.  So, we can combine our simple pattern elements
together with a specific shape to concisely express what could otherwise be
quite an annoying computation.  Here are some more examples:

    ("test", 123, 4.5) match (_, 123, 4)  # => false

    ('a', 'b', 'c', 3) match ('a', _, 'c', 3)  # => true

    (5, 4) match (_, _)  # => true

That last pattern in the above sample might seem silly: why not just match
against `_`.  The answer leads us to the final part of our discussion of
patterns: typing.  Patterns are strongly typed like everything else in the
language.  If you try and match values that don't make sense with respect to a
given pattern, you will get a type error.

    "Hello" match (_, _)  # TYPE ERROR

The above expression matches a string to a tuple-shaped pattern which obviously
makes no sense and will produce an error at compile-time.  Thus, it is useful in
situations where a type might not be as easy to predict to use more explicit
patterns since you they will allow to produce type errors at compile-time for
unexpected types as opposed to introducing possible run-time bugs.

Note that this typing logic does apply to single values as well:

    5 match "test"  # TYPE ERROR

For values, the `match` operator acts exactly like the `==` operator.

Finally, patterns in Chai must contain constant values.  This means that you
cannot place full expressions in patterns: only the pattern elements we
discussed above.

    6 match 3 + 2  # ERROR

This allows Chai to generate more optimized code for pattern matching and to
avoid confusion in the case of pattern variables which we will discuss later.

### Simple Match Expressions

Now that you understand pattern matching, match expressions are a natural
extension of your knowledge.  Put simply, **match expressions** are a control
flow construct that allows you match against multiple pattern **cases** and run
the first matching case.

To understand match expressions, let's look at a very simple usage of them.
Consider you have some number `n` that is inputted by the user, and you
want to perform different logic depending on what that number is.  Naively,
you might choose to use an if expression: ie.

    if n == 3
        println("n is a magic number")
    elif n == 4 || n == 2 || n == 8
        println("n is a small power of 2")
    elif n == 5
        println("n is a prime number greater than 4")
    # etc...
    end

> I am aware that the logic above is entirely obtuse: it exists to demonstrate
> match expressions without introducing a bunch of complex ideas along with
> them.

As you may notice, the logic is quite repetitive and quite slow to execute (lots
of conditional branches).  This is where the match expression comes in.  It
allows us to concisely and performantly represent the above logic as a series of
cases (which more logically fits its structure anyway).  Here is the above if
tree restructured as a match expression:

    match n
        case 3
            println("n is a magic number")
        case 2, 4, 8
            println("n is a small power of 2")
        case 5
            println("n is a prime number greater than 4")
        # etc...
    end

Much better!  Now, we don't have to type `n ==` over and over again.  The
statements structure itself is fairly intuitive.  The argument to the header,
begun with the `match` keyword, is the value you want to compare; each case
begins with a `case` keyword followed by a block of code and list of values that
case matches.

As you can imagine, the more cases you have, the more efficient this construct
becomes, especially if you have cases that run for multiple inputs.

However, this is just the beginning.  After all, we haven't even seen any
pattern matching yet.  Well, it turns out that those case expressions are
actually patterns!  So, we can write code like this:

    let pair = some_coord_pair()

    match pair
        case (0, 0)
            println("the origin")
        case (0, _)
            println("on the y-axis")
        case (_, 0)
            println("on the x-axis")
        case _
            println("some other point")
    end

Notice that last case at the end, that is what we call the **default case**, and
it is how we can quickly make match expressions exhaustive when we use them as
expressions.  By our pattern matching rules, we know it will match anything so
we can guarantee our match expression yields a value.

As you might expect, when we use match expressions as expressions, their branches
must always yield the same value.  Furthermore, you can use the `->` notation
to yield single expressions as you would with if expressions.

    let loc = match pair
        case (0, 0) -> "origin"
        case (_, 0) -> "y-axis"
        case (0, _) -> "x-axis"
        case _ -> "not on axis"
    end

Finally, cases must be **distinct**: this means that you cannot have multiple
cases with identical patterns.  For example:

    match n
        case 0 -> -1
        case 0 -> 1   # ERROR
    end

As you have seen, however, you can have patterns whose possible matches overlap.

    match n
        case 0 -> -1
        case _ -> 1
    end

In this case, the pattern matching algorithm simply picks the first match.



### Fallthroughs

**Fallthrough** is a control flow statement that jumps to the next case in the
sequence.  It is useful for when you have some default behavior you want to
occur for multiple cases or if you have cases that should flow in sequence.

The fallthrough statement uses the `fallthrough` keyword placed on its own line.

    match n
        case 1
            println("n is the multiplicative identity")
        case 2
            println("n is the first even natural number")
            fallthrough
        case 3, 5
            println("n is a prime number")
    end

Notice that fallthrough ignores the actual case's condition and just blindly
jumps to the next case.

Fallthrough can be used within more complex logic as well.

    match n
        case 1
            println("Do something")
        case 2
            println("Do something else")

            if n == m
                fallthrough
        case 3, 5
            println("Do something even more else")

However, using fallthrough in this way may render your match expressions
non-exhaustive.  In this example, the if expression doesn't always yield a
value: so the match expression by relation becomes non-exhaustive, and its value
can't be used.
### Case Guards

A **case guard** is a condition placed on a case the specifies when it can
match.  These are particularly useful in cases where you want the flexibility
and "scalability" of the match statement but have a few bits of more complex
logic nestled into the matching.

Case guards are placed after the full list of case patterns prefixed by the
`when` keyword.  The guard expression itself is just a condition like an if
statement.

    match n
        case 1
            println("n is the multiplicative identity")
        case 2
            println("n is the first even natural number")
            fallthrough
        case 3, 5
            println("n is a prime number")
        case _ when n >= 1
            println("n is a natural number")
    end

Note that this is different from simply placing an if expression in the case:
the case guard will prevent the case from matching in the first place. The match
expression will not even consider the case pattern if the case's guard is not
satisfied.  This can allow you to have multiple identical patterns provided
there is some case guard differentiating them.

    match n
        case 2 when m < 2
            println("m < n")
        case 2
            println("n = 2")

> I am aware the above logic is a completely unnecessary use of case guards, but
> it demonstrates the idea at hand.

## Pattern Variables

A **pattern variable** is a special kind of wildcard that gives a name to
whatever value it matches allowing it to act as variable.  

Pattern variables can be constructed by placing variable names in place of `_`
to cause whatever value matched the wildcard to be stored in that variable.

For example,

    let pair = some_tuple()

    let squared_dist = match pair
        case (0, 0) -> 0
        case (a, 0) -> a ** 2
        case (0, b) -> b ** 2
        case (a, b) -> a ** 2 + b ** 2
    end

As you can see, the variables match just like wildcards, but they can be used
the branches of case expressions.  These variables work like header variables
in if statements: they are only visible within the scope of the case or block
they are applied.

Notably, test-match expressions can also use header variables.

    if triple match (a, 0, 0)
        println("The triple at", a, "on the x-axis")
    end

Header variables *cannot* be used in cases that can be reached by fallthrough.
This is because fallthrough doesn't actually perform a pattern match so the
variable is never extracted.

    match pair
        case (0, 0)
            println("at origin")
            fallthrough
        case (0, a)  # ERROR: pattern variable after fallthrough
            println("on y-axis at", a)
        case (b, 0)  # this is ok
            println("on x-axis at", b)
    end

Similarly, pattern variables can only be used in cases that have a single
pattern since Chai won't know at compile-time which variable will be defined.

    match pair
        case (a, 0), (0, b) -> a + b  # ERROR: pattern variables in case with multiple patterns
    end


