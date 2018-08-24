## Variables and Constants

In programming, we need ways of storing and accessing data in memory, whether it
be an integer, a string, or a massive, complex object.  In Whirlwind,
we use something called a **variable** to accomplish this.

### Variables

A variable is an identifier that references some point in the computer's
memory.  The code below shows how to declare them.

    $x = 3;

This variable declaration has 2 parts: the **declaration** and the
**initialization**. The declaration is doe with the `$` operator followed
by some name, in this case `x`. The second part, initialization, is
where the value of a variable is set.  This performed with the
'=' operator, followed by an expression.

Once x is declared, it can be accessed from anywhere as shown in **Listing 1.7**.

#### Listing 1.7 - Variables in Action

    // notice that x is accessible from anywhere
    $y = x * 3; // 9
    $z = x + 4; // 7

    $a = z + x + y; // all 3 variables are now usable

You can also change the value of a variable through **assignment**.
This is done with the `=` operator, it uses the same syntax as initialization.

    x = 5; // x's value is now 5
    z = x * y; // y is still 9, but x is now 5, so z now equals 45

Variable is are **type immutable**.  This means that their type cannot
change.  Since x, y, z, and a are all integers by **type inference**,
you cannot set z equal to say a string.

    z = "String"; // TYPE ERROR; this will not compile

If you want to specify a variable's type at declaration, you can use a
**type extension**. Type extensions consist of the `:` operator, followed
by a type designator.  The code below shows how to utilize type
extensions for variables.

    $b: int = 4; // : int is the type extension

    $c: float = 5; // 5 will automatically be coerced into a float (5 -> 5.0)

    $d: bool = 5; // TYPE ERROR, types do not match

As you can see, the type extension always goes in between the
identifier and the initializer.  Additionally, the
type extension always take precedence over the type of the
initializer and so if they do not agree, as shown with `d`, you will
get a type error during compilation.

Finally, you can declare a variable without initializing it as long
as you specify a type. This is called **default initialization** and
The code below shows an example of default initialization.

    $var: int;

Unlike other languages, Whirlwind allows variables to be used without
initialization.  It automatically assumes a null (or default) value
for the variable.  In this case, `var` has a null value of 0 because it
an integer.  We will talk more about null and null values in future chapters.

    a = var; // a now equals 0, because var is 0

It is important to note that you must always give a type during
a declaration.  This type can be inferred from
the initializer or explicitly stated with a type extension, but
it must always be there.  Whirlwind will not allow you to declare
variables with no type.  For example, the code below is invalid.

    $invalidVar; // ERROR: unable to infer type of 'invalidVar'

### Multi Declaration

multi declaration refers to declaring multiple variables or constants
at once.  **Listing 1.8** show examples of multi declaration.

#### Listing 1.8 - Multi Declaration

    $(a, b, c): int; // all 3 variables are given a type of int

    $(x = 4, y, z = 3.2): float; // some values are initialized

    @(s = 'a', r = 'b'); // both values must be initalized since it is a constant

As you can see, multi declaration consists of using one of the two
declaration operators, followed by a set of identifiers wrapped in parentheses.

These identifiers can provide initializers or be given the overall type of
the declarative group.

In addition, one declarative group can contain variables (or constants) of multiple
types.

    @(a = 'a', pi = 3.14);

Similarly, variables (or constants) can specify their own type within the group as
well.

    $(x: int, y: float);

You can even provide an overriding type that all uninitialized values will
default to.

    $(b = 4, c: char, d, e): float; // b is int, c is char, and d & e are floats

### Multi Assignment

Just like declaration, you can also assign to multiple values at once.
This is done with **Multi assignment** as looks like so:

    $(a = 4, b = 3);

    a, b = b, a; // a and b are now swapped

Assignment works by matching each side in order.
So in the example above, the first variable `a` matches
to the first expression that happens to contain the value
of `b` and so on.

You can do this with up to as many values are you want as long as
there is an equal number of a values on both sides.  Because of this,
the code below is invalid.

    $(x = 'a', y = 'b', z = 'c');

    x, y, = z, y, x; // invalid; too many values on right hand side

### Shadowing

Shadowing is an extremely useful concept in Whirlwind that allows you to
reuse variable names from higher scopes.

***Note*** *A scope refers a variable context, like a block.*

The principle of shadowing is that when looking up a variable, Whirlwind starts
in the nearest scope. This means that you can use the variable name twice as long as
they are in different scopes.  **Listing 1.9** demonstrates this.

#### Listing 1.9 - Shadowing

    $x = 2;

    // sub scope
    {
        // shadow occurs
        $x = 5;

        Println(x); // x is 5
    }

    Println(x); // x is 2

Just be careful that the variables aren't in the same scope as that will cause a
redefinition error.

    $x = 2;

    // ERROR
    $x = 5;
    Println(x);

