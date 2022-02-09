# Variables and Input

*Procedural code* is the basis of most modern programming languages.  In short,
it consists of a series of *blocks* which chain together different *statements*
which perform actions.  In the absence of other control flow structures (which
will consider later in this guide), execution proceeds sequentially: one
statement is executed after the other.

In this chapter, we will explore how this process works in Chai.  Although the
previous two chapters have covered material that is mostly consistent with other
languages, this chapter is where Chai starts to diverge from the norm a little
bit.

## Blocks

A **block** is an expression which sequentially executes a series of
**statements** contained within it.  In Chai, blocks begin with the `do` keyword
followed by a newline and concluded by the `end` keyword.  Each statement is
concluded by a newline.

    do
        # statement 1
        # statement 2
        # ...
    end

Importantly, blocks are considered an expression, so we can make one the body of
our main function.

    def main() = do
        # statements
    end

Because this pattern is so common, Chai allows the use of a **block body** for
functions.  Essentially, you just drop the `= do` from the code above, and Chai
will implicitly define a block as the body of the function.

    def main()
        # statements
    end

Now, it is time for a brief digression on program structure.  In Chai, all
executable code, that is code that actually does something, must be contained
within a definition, most often a function.  For now, we are going to be putting
all that code inside the main function.  However, since it is inconvenient to
copy and paste a main function definition for every example, this guide will
frequently show snippets of executable code and discuss them without enclosing
them in a function.  These snippets do *not* constitute full Chai programs. You
should also assume that `println` is always imported and available.

As mentioned before, blocks are made up of statements.  There are many different
kinds of statements, but the first one we are going to look at is the
**expression statement**.  These statements are exactly what they say on the
tin: expressions used as statements. 

    do
        5
    end

That is perfectly valid block.  However, there are some restrictions on how
expressions can be used in expression statements.  Most "complex" expressions
will need to be wrapped in parentheses.

    do
        2 + 2  # ERROR
    end 

    do
        (2 + 2)  # OK
    end

A notable exception to this rule is the function call:

    do
        println("Hello, there!")
    end

## Variable Declarations

The next kind of statement we are going to look at are 
**variable declarations**. These define **variables** which are named locations
to store values.  

Variable declarations begin with the `let` keyword followed by name and an
**initializer**. The initializer contains the initial value that will be stored
in the variable.  Let's look at an example.

    let x = 10

In that declaration, `x` is the variable name and `10` is the value that is
stored in it.

We can then access this value later by simply using the variable's name.

    let y = x * 2

All variables also have a fixed type.  For example, `x` and `y` both store
numbers, specifically numbers of a type of `i64`.  This type is *inferred*
from the type of the expression that is used to initialize them.

If you want to specify the type of a variable, you can use a **type label**.

    let pi: f64 = 3.14159264

Only one variable with a given name can be declared in a given scope.

    let x = 5.5  # ERROR

You can also declare multiple variables at once by separating them with commas.

    let z = 78, u = -12.56 * pi

Variable names can be single letters as we have already seen or they can be full
words or even combinations of words.  However, there are some restrictions.

1. Variable names cannot begin with numbers.
2. Variable names can only contain letters (upper and lower), numbers, and underscores.
3. A single underscore is not a valid variable name.
4. A variable name may not be the same as a keyword.

Here are some examples of valid variable names.

    my_var
    var123
    YOUR_VARIABLE
    eggMan
    _i_like_pancakes__
    a0b1x_12

Although many kinds of variable names are allowed, the convention in Chai is to
name variables using all lowercase letters and separate words with underscores.
This is called *snake case*.  Here are some examples:

    token
    app_name
    password
    count_of_items
    length_of_dataset

The [Style Guide](/docs/style-guide) has a more definitive list of standards for
naming conventions.

## Assignment

All variables in Chai are **mutable** meaning you can change their value after
you assign to them.  Changing a variables value is called **assignment** and is
done using the `=` operator.  Assignment constitutes its own statement.

    let name = "Bob"

    println(name)  # prints Bob

    name = "Alice"

    println(name)  # prints Alice

A variable's type cannot change during program execution: so, you can only assign
a value to variable that is the same as the type of that variable.

    let x = 5.6

    x = "Hello"  # ERROR

At this point, it is worth talking a little bit more about type inference.  Let's
consider a simple example:

    let y = 5

    y = 0xff

    y = 5.5

If you try and compile the code above, you will get a compilation error on the
third line.  It will look like this:

```text
no type overload of `{int}` matches type `{float}`
```

This is a bit of odd error considering the circumstances: what is the compiler
trying to tell you?  In order to understand why this happened, we need to
understanding how Chai determines the type of something.

Essentially, it starts with a set of possible types for something and slowly
prunes those possible types down as it learns more about your program.  Let's
start with line 1.

    let y = 5

Here, Chai doesn't know much about the type of `y`.  Since `5` is a number
literal, Chai only knows that `y` must be a numeric type, but it doesn't know
which of the many types that are available.

Proceeding to line 2:

    y = 0xff

Because `0xff` is an integer literal, Chai now knows that `y` must be an integer:
all the float types are no longer considered viable types for `y`.  So when Chai
reaches line 3:

    y = 5.5

It sees that `y` is being assigned to a float type, but it also knows that `y`
can only be an integer.  Thus, it encounters a type mismatch between the set of
possible types of `y` and the set of possible types of `5.5`: there is no overlap.
This is why Chai produces an error.  

The language of the error is also relevant.  Chai refers to something called a
**type overload**.  As you may be able to guess, a type overload is the formal
name from one of the possible type possibilities.  All the possible types of `y`
are called the **type overload set** of `y`.  

Now, we can actually interpret what the compiler is trying to tell us.  

```text
no type overload of `{int}` matches type `{float}`
```

Reading it directly, it saying that no type overload of `{int}`, meaning no
possible type of integer, matches (or is in) the set of float types.

### Compound Assignment

Now let's take a break from type mumbo jumbo and take a look at a cool bit of
syntactic sugar.  Suppose you have a variable `a` as defined below:

    let a = 10

Now, suppose you want to double `a` (multiply it by 2).  You could write:

    a = a * 2

However, this line is fairly repetitive and as you will discover a common task.
So, Chai gives us a shorthand.

    a *= 2

This is exactly equivalent to the previous line, just with few characters.  This
is called **compound assignment**.  You can do this with all the arithmetic
operators:

    a += 6  # a = a + 6

    a /= 2  # a = a / 2

By far the most common task out of all compound assignments is adding and
subtracting one, also known as incrementing and decrementing.  Using compound
assignment, we can shorten these to:

    a += 1  # increment
    a -= 1  # decrement

However, for this special case, Chai offers yet another little bit of shorthand.

    a++  # increment
    a--  # decrement

Once again, this code is exactly equivalent to the previous pair of lines, just
shorter.

### Multi-Assignment

You can also assign to multiple values at once.

    let a, b: i32

    a, b = 0, 1

The position of the variable on the left corresponds to the position of the
value that is assigned to it on the right.  For example, in the above code, `a`
is assigned to `0`, and `b` is assigned to `1`.

Intuitively, one might think that you can simply decompose multi-assignments
into two separate assignment statements, but this is in fact not the case.
Let's consider an example to understand why:

    a, b = b, a

This code does exactly what you would expect it to do: it swaps the values of
`a` and `b`.  But, if you decompose it into two separate statements, the
result actually changes:

    a = b
    b = a

Following the execution of that code, you see that it really doesn't have the
intended behavior: first `a` is assigned to `b` and then `b` is assigned to `a`
which now holds the value of `b` so `b` is assigned to itself. 

This is ultimately because the multi-assignment is *not* equivalent to the two
separate statements.  In the case of multi-assignment, Chai fully evaluates
the right-hand side of the assignment statement before it stores the results
into the variables.  This means that swapping code is actually equivalent to:

    let a_temp = a
    let b_temp = b

    a = b_temp
    b = a_temp

Obviously, the Chai compiler will generate code that is a bit more optimized
than the above, but the premise still holds.  

You can also use compound assignment operators with multi-assignment.

    let x = 5.6, y = 7.2

    x, y *= 2, 3  # x = 11.2, y = 21.6

Multi-assignment is an amazing tool that you are encouraged the use plentifully.
It helps make code cleaner and more concise without losing meaning.

## Command-Line Input and Output

In Chapter 1, we introduced the `println` function to print strings to the
console. Now that we know a little bit more about Chai, it is time that we
discuss some other means of interacting with the command-line.

First things first, the `println` function doesn't only work with strings.
You can also print numbers to the command-line just as easily.

    println(12)  # prints 12

You will find that most types that ship with Chai either as part of the core
language or standard library are acceptable as arguments to `println`.  In a
later chapter, we will learn how to make our own types "Showable".

Often, however, we don't just want to print things to the command-line, we also
want to read things in.  After all, what good is a program that can't respond to
input!

Command-line input is a bit more complicated than output.  We will introduce
some basic tools in this chapter and elaborate more upon them in later chapters.

Let's start with the basics.  First things first, we need to bring in some more
functions from `io.std`.  The function we want to bring in now is called
`scanln`. We can do this by simply adding another name to the import statement.  

    import println, scanln from io.std

`scanln` is the exact opposite of `println`: it reads a line from the console
and drops the newline.  It returns the line it reads in as a string.  To get the
return value of a function, we simply store the "result" of calling it as a
value.  To understand what this means, let's look at example.  Below is a simple
program which inputs a line and then prints it right back out.

    let line <- scanln()

    println(line)

If you run this program, you should see behavior like the following:

```text
>> Hello!
Hello!
```

You may notice something a bit unusual about the variable declaration: I used an
`<-` instead of an `=`.  This is because `scanln` actually returns something
called a `Result` type that conditionally contains the string value of the line
if Chai succeeds in reading from the command-line.  We will learn much more
about `Result` and that special `<-` in later chapters, but for now, just know
that whenever you use `scanln`, you need to use the `<-` instead of an `=`. Note
that this is also true in assignment:

    let line: string

    line <- scanln()

    println(line)

Often, we also want to read in numbers.  To do this, we are going to need a
function called `scanf`.  We can amend our import statement to bring it in
instead of `scanln`.

    import println, scanf from io.std

`scanf` is a special function that performs something called *formatted I/O*.
Basically, it allows us to extract values from user input by providing a pattern
we want that input to fit. 

Additionally, you will note that `scanf` doesn't actually return the value it
reads in.  This is because `scanf` can actually read many values from the
command-line instead of just one like `scanln`.  So, we instead use something
called a *reference*.  Essentially, references give us a way to store into
variables without knowing the name or location of the variable.  References are
massive topic in Chai that will take up a whole chapter later on.  But, because
they are so integral to Chai, it is good to exposed to them early. 

As you might be able to tell, there is a lot more going on here than I have
space to explain to you in this chapter.  One might argue that I am jumping
quite a bit ahead, but considering as reading in numbers from the command-line
is such an important skill, I am going to, for now, just give you a pattern to
copy whenever you need to read in a number using `scanf`.  

Here is a sample call to `scanf`.  

    let d: i64
    scanf("{}\n", &d)

As you can see, `scanf` is taking two arguments separated by commas.  The first
argument is a special string that essentially tells `scanf` how we want our value
to be read in.  In this case, we only want to read one value on its own line.
The second argument is a reference to the variable where we want to store the scanned
value.  That `&` is the symbol we use to create a reference. 

As an example of this special usage of `scanf`, here is a program that read in a
number, squares it, and prints the result of the squaring.

    let input: f64

    scanf("{}\n", &input)

    println(input ** 2)

Here is an example of running this program:

```text
>> 5
25
```

There are two things worth noting here.  

Firstly, we didn't need to change that `scanf` line at all to account that we
are scanning in a float instead of an integer. This is because `scanf` knows
from the type of `input` what kind of value to scan in.  

Secondly, if it was not already obvious, functions accept full expressions as
their arguments so we can perform the squaring inside the call to `println`: the
expression is evaluated, and the result passed to `println`.

One final note on `scanf`: what happens if the user enters something you don't
expect?  While we don't yet know how to deal with errors, you can already learn
how to identify the behavior.  For numbers, all that will happen if we run the
program and input something that Chai can't convert into a number is that nothing
will be stored into `input` meaning it will have its default value of `0`.  For
example,

```text
>> hello
0
```

`hello` can't be converted into a number.  So nothing gets stored into `input`.
In later chapters, we will learn how to more effectively deal with errors like
this.  However, this knowledge should at least help you eliminate bugs in your
programs.

<guide-exercise>
{
    "label": "3.1",
    "content": "Write a program that prompts the user to enter two numbers and
    prints their product.",
    "hint": "Use scanf twice.",
    "solution": {
        "type": "text",
        "text": "TODO"
    }
}
</guide-exercise>

