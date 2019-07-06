# Overloading

Overloading is the process by which a function or method is given additional functionality
that it did not possess before.  Overloading comes in two forms in Whirlwind: **group overloading**
(otherwise just known as overloading) and **operator overloading**.

> Note: In this chapter and in all future chapters, when we refer to function, assume that methods
> are also included unless otherwise stated.

## Group Overloading

Group overloading is the simplest form of overloading.  It involves declaring multiple functions
with the same name and different sets of parameters.  Then, when the function is called
Whirlwind will determine which function to call.

    func alert(name: str) {
        // -- snip --
    }

    func alert(name: str, msg: str) {
        // -- snip --
    }

    func main() {
        alert("Tom"); // calls first method
        alert("Tom", "Hi."); // calls second method
    }

As you can see, Whirlwind will map each call to its appropriate signature to produce the desired behavior.
The only important thing to remember is that the functions must be distinguishable by their parameters alone.
So the following would be considered invalid.

    func f(x: int) int => x * 4;

    func f(x: int) float => x * 3.14;

Because the two functions have the same parameters, Whirlwind cannot distinguish them solely based on their call.
Additionally, if you have functions with optional, indefinite, or variadic arguments, Whirlwind may not be
able to as easily distinguish function calls between them and so you may end up being unable to effectively
overload them at all.

    func g(x: int) int => x ~^ 2;

    func g(x: int, y = 0) int => (x + y) * 4;

The above overload also fails since a call to the former could just as easily be a call to the latter.

> The language specification contains the exact rules by which an overload is evaluated.  You can read
> that to get the details of how Whirlwind distinguishes calls in more complex scenarios.

The final piece of the puzzle for this kind of overloading is the idea of a **function group**.  A function
group is merely a function and all of its overloads taken together as one entity.  In the first example,
both variations of the function `alert` when referred to together would be a function group.  This is where
the term group overloading comes from.

## Operator Overloading

Operator overloading is a little more complex.  First and foremost, it only exists on methods and employs a
different syntax from functions.  Operator overloading is how you would define the behavior of a contrived type
under a certain operator.  

For example, if you wanted to allow a special type you defined like say a complex number to be operated on by the
`+` operator, this would be how you would do that.

The syntax for operator overloads is fairly simple.  It begins with the `operator` keyword followed by the operator
you are overloading, followed by a syntax similar to how you would define a function's parameters and its return type and body.

For the complex number type mentioned, an overload for the `+` operator would look like this.

    struct Complex {
        r, i: double;
    }

    interf for Complex {
        operator+(other: Complex) Complex
            => new Complex(this.r + other.r, this.i + other.i);
    }

That is really all there is to operator overloading, you can define an overload for almost any operator with only one
restriction: it must take the same number of arguments as the original operator.

> You can find a list of the operator symbols for overloading in the language specification.

Because operator overloads are really just methods internally, they can also undergo group overloading like a function.
