# Lambdas and Closures

Lambdas and closures are probably some of the most powerful tools in all
of Whirlwind.  They power many of the patterns and constructs underlying Whirlwind
and are relatively simple to understand.

## Functions as Objects

Before we can understand the idea of a lambda, we must first discuss functions.
Or rather, look at functions in a unique light.  Up until now, we have exclusively
examined functions as high level, static constructs, designed only to be called.
But, that is only half the picture.

A function is in fact a type, like a struct or an interface.  And as such, they
can be made into objects (tangible items that can be stored and operated upon).
Consider the simple function that adds two numbers.

    func add(a, b: int) int
        => a + b;

What if I told you that it were possible to store this function in a variable?  Well,
as it turns out, this is possible.

    let f = add;

And even crazier, it is possible to call `f` like a function.

    f(3, 4); // => 7

Pretty cool right?  In reality, the call `()` operator is just that, an operator.
There isn't really anything special about it besides the fact that it can be used to
call functions.

> The `()` is also the only operator that can't be overloaded.

But here's where it gets even weirder.  It is possible to, in fact, pass functions
to functions.  Consider the example below:

    func applyToList(list: [int], fn: func(int)(int)) [int] {
        for (i = 0; i < list.len; i++)
            list[i] = fn(list[i]);

        return list;
    }

The function `applyToList` accepts a list a function called `fn` and applies that function
the each element in the list.  `applyToList` is what we call a **higher-order function** which
is a function that accepts another function as an argument.

> It is **not** possible to have named arguments in this context so as the names
> are not specified in the function type label.

This is also the first example of the function type label.  This label begins with the `func`
keyword and is followed by two parentheses: the first contains the types of the arguments and the
second containing its return type(s).  The function type label does distinguish between regular,
optional, and indefinite arguments by requiring you to place a `~` **after** each optional argument
and `...` **before** each indefinite argument.

> If omit the type after the `...`, the argument is assumed to be variadic.

To use our `applyToList` function, we would follow the same protocal as before.

    func increment(x: int) => x + 1;

    func main() {
        let myList = [1, 3, 5, 7];

        // after this call: myList = [2, 4, 6, 8]
        myList = applyToList(myList, increment);
    }

Here, the function `increment` is passed as the second argument and is objectified before being
passed.

> There is a method called `map` that exists on all iterables in Whirlwind that exhibits a similar
> behavior to our `applyToList`.  Look it up in the standard library docs to learn more about it.

## Lambdas

Now, you have probably already notice a flaw the above model, you have to define a function fully
before you can objectify it.  Even if you were to define `increment` locally (within the body of main
which is possible in Whirlwind), you would still run into the same problem.  And, as it happens Whirlwind
has a solution: the **lambda**.  A lambda is anonymous function, meaning it has no name by default and is
the closest thing we can get to a "function literal".  Let's see how you would fix the `applyToList` example
using a lambda.

    // applyToList up here somewhere

    func main() {
        let myList = [1, 3, 5, 7];

        myList = applyToList(myList, |x: int| => x + 1);
    }

And there it is, no increment function, no unneeded variables, just the behavior we want to occur.  You have probably
already guess what the lambda syntax is: a normal argument definition enclosed in `|` followed by a standard function
body.  Pretty easy right?  And it gets better.

Remember how earlier in the section on type classes, I mentioned that one other type has context-based inference?  Well,
here it is.  Lambdas have context based inference on the types of their arguments **and** their return type.  You do not need
to specify either.  In fact, Whirlwind does not **allow** you to specify the return type because it will always be inferable
from the function body.

So let's rewrite this scenario one more time.

    func main() {
        let myList = [1, 3, 5, 7];

        myList = applyToList(myList, |x| => x + 1);
    }

Bam.  No fuss, no mess, just plain old lambdas.  And just top it off, you can put entire function blocks inside lambdas
not just expressions.  So let's change it up and instead of adding one, calculate the factorial of each number.

    func main() {
        let myList = [1, 3, 5, 7];

        myList = applyToList(myList, |x| {
            // the right hand expression is fully executed before the loop begins
            for (n <- 2..x)
                x *= n;

            return x;
        });
    }

Now, we have a list full of factorials.  Before we move on the last concept in this chapter, we need to look at
one little thing.  No argument lambdas.  These do indeed occur, and when using them you need to be careful not
to confuse the `||` operator for the beginning of an empty lambda.  As a rule of thumb, whenever you declare a
no argument lambda, put a space in between the `|` operator.  Not only will it prevent annoying syntax errors, but
it will also make your code more readable.

## Closures

The final piece of the puzzle is the closure.  A closure is not some much its own unique data type, but rather
a type of function.  A **closure** is a function that *closes around* its external state.  That probably sounds
a bit confusing so let's look at an example.

Suppose you wanted to craft a function that everytime it was called, it would return the next consecutive integer.
You could do this using the power of closures.  Let's declare a function `consecutive` that is going to return
one of these special functions.

    func consecutive() func()(int) {

    }

Inside `consecutive` is where the real magic is going to happen.

    func consecutive() func()(int) {
        let c = 0;

        return | | => c++;
    }

Did you see what happened there?  The lambda captured `c` in its body.  Now, let's look at what calls this returned
function behave like.

    func main() {
        let cs = consecutive();

        cs(); // 0

        cs(); // 1

        cs(); // 2
    }

Most of you are probably a bit confused right now.  Why is `cs` returning a different value for `c`?  I thought it was
`0`.  Well, my friends, this the magic of the closure.  When a closure closes around its external state, it captures all
the variables in it as references by default.  This means that when we execute `c++` in the lambda, we are actually modifying
a reference to `c`.v

However, this has some interesting implications.  Let's examine the code from before more closely by adding on a few extra lines.

    func main() {
        let cs = consecutive();

        cs(); // 0

        cs(); // 1

        let cs2 = consecutive();

        cs2(); // 2

        cs(); // 3
    }

Here we see the one caveat to closures.  Because they captures their state by reference, they all modify the same state.  Effectively,
all closures generated from the same state, all modify the same **shared** state.  So, when `cs2` is created, it captures the same reference
to `c` that `cs` had and in doing so, latched on to the modified version of `c` meaning `c` had an initial value of `2`.  Moreover, when
we called `cs` again after `cs2` we see that the same principle occurs.  Because `cs2` modified not a copy of `c`, but the original version
of `c` itself, `cs` receive and incremented that modified version.  Thus, we observe the above behavior.
