# Functions

A function is a callable block of code that can accept an external state in the form of arguments.
In Whirlwind, a function is composed of three primary parts: the name, arguments, and return
types.  

## Basic Functions

All functions begin with the `func` keyword and take the following form.

    func add(x, y: int) int {
        return x + y;
    }

`add` is the name of the function and `x` and `y` are the arguments, both are integers.
The return type of the function is an integer.

As you can see from its body, this function take in two arguments, adds them, and returns
their sum.  The `return` keyword is used to return a value or values.  

Since this function only has one statement in its body which is a return statement,
you can rewrite the using an expression return.

    func add(x, y: int) int => x + y;

This is just shorthand for the first function definition, as it is common for functions to
return single, complex expressions.

Functions can also take no or as many arguments as you choose of multiple types.  For example,
if you have a function that returns the time, its signature might look like this.

    func getTimeInHours() int {
        // get time in hours
    }

Similarly, a function that sums a list, might look like the following.

    func sum(lst: list[int]) int {
        let s = 0;

        for (n <- lst)
            s += n;

        return s;
    }

Functions can take as many or as few arguments as necessary.

## Function Calling

...

