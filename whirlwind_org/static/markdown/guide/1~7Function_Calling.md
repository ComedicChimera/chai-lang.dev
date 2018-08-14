## Function Calling
Calling functions is a pretty essential part of Whirlwind.  You will
be doing it a lot and even if you don't fully understand functions, you still
need to know how to do this.  In fact, we have already called a function
in this guide, the `Println` function.

### What is a Function?
A function simply represents a reusable, callable block of code.  You
can use functions to make your code cleaner and more efficient. Calling
a function refers to running that reusable block of code and use it functionality.
We will talk much more about functions in **Chapter 4**, but for now
we just need know how we can access and use them.

### Arguments
Function calling uses the `()` operator to call a given function.  What
we put inside the parentheses are the **arguments**. These are values
that get given to the function for it to perform some operation on.  For
example, a function that adds to numbers would take two arguments:
the two numbers you desire to add. Or a square root function would take
one argument: the number to square root. Or a function may accept no
arguments, meaning you can leave the parentheses empty.

### A Simple Call
Consider some random function `someFunc` that accepts no arguments.
We could call this function using the `()` operator like so:

    someFunc();

Function calls are considered statements and so they can stand on
their own.  Notice that we put the parentheses after the function name.
This is important as putting them before will cause a syntax error.

### Functions Return Values

Functions can also return values.  For example, an add function might
return the some of two numbers, so we can extract the value of the function
by setting the result equal to a variable or we can just ignore the return
value entirely.

    $x = anotherFunc(); // store value in x
    anotherFunc(); // ignore the value

All function return values are typed and type of the return value
does not shift from function to function.

Functions that return nothing cannot have their values store in variables,
so it is important to know when a function returns or not.

### Passing Arguments
Some functions will accept one, two, or many arguments.  We can pass
these arguments two the function by inserting them between the parentheses,
separated by commas. **Listing 1.11** shows some examples of function
calls using arguments.

#### Listing 1.11 - Calling Functions with Arguments

    /* somewhere else we declared a function add that
       takes two integer arguments. Here is how we call it */

    add(1, 2); // 3
    add(4, 5); // 9

    add('a', 'b'); // TYPE ERROR: argument types don't match
    add(4); // ARGUMENT ERROR: too few arguments
    add(5, 6, 7); // ARGUMENT ERROR: too many arguments

    $(a = 1, b = 3);
    add(a, b); // passing identifiers to function
    add(a + 2, b * 3); // you can pass expressions as well

    /* somewhere else we declared sqrt, which accepts one numeric argument
       and square roots it. */

    sqrt(4); // 2
    sqrt(16); // 4

    sqrt("number"); // TYPE ERROR
    sqrt(); // too few arguments

As you can see, functions are pretty specific as to what they accept.
However, they are pretty easy to use otherwise.

There is a lot more to talk about in terms of function calling and
other special types of arguments, but we will talk more
about that in **Chapter 4**.  This is all you need to know for now.

## Summary
That's it.  You finished Chapter 1. This was quite a large and diverse chapter.
You learned about  program structure, the main function, comments,
data types, variables, constants, constexprs, assignment, and function calling.
But, you have only begun your Whirlwind journey.
I know the information may seem daunting, but
don't worry.  After a while, you will get the hang of it.