# Variables and Numbers

We will now consider several smaller programs, focused on arithmetic, to
familiarize ourselves with some of the mundane but essential features of Chai.

To begin, let us write a program to display the
[Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number).  For those
unfamiliar, the Fibonacci sequence is a sequence wherein the next number in the
sequence is defined as the sum of the previous two numbers, beginning with the
entries 1 and 1.  For example, the first few numbers of the Fibonacci sequence
appear as follows:

```text
1, 1, 2, 3, 5, 8, 13, 21, 34, ...
```

In our program, we will print the Fibonacci numbers less than 1000.  In a
similar fashion to our *Hello World* program, let's take a look at the whole
program first and then discuss what makes it work.

    import println from io.std

    def main()
        let a = 1, b = 1

        while a < 1000
            println(a)
            a, b = b, a + b
        end
    end

Our program begins with a taste of the familiar: an import statement just like
the one in our *Hello World* program.  This makes sense given that we want to
print to the console.  We also have a `main` function; however, this main
function is constructed differently.  Instead of having an `=` to denote its
body, it instead has a newline and a matching `end`.  This denotes that the
function has a **block body**: ie. it is made up of a series of statements to be
executed in sequential order.

## Variable Declarations

Our function begins with a statement called a **variable declaration**.  The
statement above declared two new variables: `a` and `b`.  These variables are
going to be used to store the entries as the fibonacci sequence: we need the
previous two values to calculate the next one.  These variables are both
initialized with a value of `1` as denoted by the `= 1` following each of their
names.

However, a natural question arises when declaring these variables: what is their
type?  Chai is a strongly and statically typed language meaning every value must
be assigned a specific data type at compile time and that type cannot change. To
the question at hand, the type is determined via **type inference**.  In
essence, Chai uses the value assigned to the variable to decide which type to
assign it.

Chai provides us with a number of built-in data types to represent numbers.  In
this case, the compiler will select the `i64` type: denoting a 64-bit signed
integer.  We can specify this type explicitly using a **type label** like so:

    let a: i64 = 1, b: i64 = 1

Of course, in this case it adds nothing for us to do so, but it is a point worth
mentioning nonetheless.  We will see a myriad of other types as we look at more
examples: a definitive list of built-in types is provided in a later chapter.

## The While Loop

## Assignment





