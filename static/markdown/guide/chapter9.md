# Structs

The struct is the most fundamental method for stucturing data and an
essential aspect of the Whirlwind language.  Knowing how to properly, wield
the power of structs will enable you to write cleaner, simpler programs and
think of your data in a more human fashion.

## Defining a Struct

A struct itself is a collection of named values of different types. Each value is designated
by an identifier and labeled with a type akin to the way arguments are laid out.  To declare
a struct, you use the `struct` keyword followed by a name and the members of the struct enclosed
in curly braces.

    struct Point2D {
        x, y: int;
    }

The above code creates a struct called `Point2D` that contains two member variables `x` and `y` that
are both integers.  As I said before, structs can contain as many types and variables as you want.

Another example of a struct might be a struct representing a product in a store.

    struct Product {
        name: str;
        price: float;
        count: int;
    }

As you can see, this struct contains three variables of all different types.  In structs, variables are
often grouped by type, but this is not necessarily the case.  For example, `Point2D` could just as easily
be written as:

    struct Point2D {
        x: int;
        y: int;
    }

Just remember that each member or member series must be delimited by a semicolon.

## Create Struct Instances

The struct definition acts a template for the corresponding instances.  An instance represents a single unit
of data that matches the form of the struct definition.  Each struct can have numerous instances and none
of the instances have any specific influence over each other.

To create an instance of a struct, we use the `new` keyword once again, followed by the struct's name and parentheses.

    let p = new Point2D();

This will create a new instance of the Point2D struct.  You can get and set values by using the `.` operator followed
by the property you want to access.

    p.x = 4;

    let y = p.y;

The compiler will mark any invalid struct accesses (due to type or name) during compilation and you will get an error.

    p.z = 3; // ERROR, z doesn't exist

    p.x = 4.3; // ERROR, type mismatch

It is also possible to initialize struct properties during creation using an initializer list.  To do this, you
replace the parentheses with braces followed by initializers for each individual property.

    let p2 = new Point2D{ x=15, y=-7 };

The same guards apply as do with property accesses.  In addition, you don't need to initialize every property in the list,
just the ones you intend to use.

> Like variables, all struct properties are default initialized.

## Constructors

A constructor is the function that is called whenever a struct is created.  It is what is responsible for initializing the struct
and internally, used to create the instance.  You can define a custom constructor using the `constructor` keyword.

Consider our `Point2D` example.  What if we wanted all Point2D's to default to the position (1, 1) instead of (0, 0).  Using a
constructor, we could do this.

    struct Point2D {
        x, y: int;

        constructor() {
            this.x, this.y = 1, 1;
        }
    }

The constuctor looks somewhat like a function, but it cannot return anything.  This is because it implicitly returns the new struct
instance.  Also, notice that all member variables of the struct are prefaced by the `this` keyword.  All `this` is as a reference to
the parent struct.  Without it, it would be impossible to access the member variables of the struct.

Constructors can also take arguments.  For example, the Point2D constructor might take in an x and a y coordinate so as to avoid the
slightly more verbose initializer list.

    struct Point2D {
        x, y: int;

        constructor(x, y: int) {
            this.x, this.y = x, y;
        }
    }

Now, we could create an instance of this struct by using passing the values for `x` and `y` in the constructor call.

    let p = new Point2D(4, 3);

Notice that we do not need to name the members anymore since we are just passing arguments to the constructor, not manually initializing
the struct.

> Structs always retain a **default constructor** which allows their constructor to be called with no arguments, even if one is not
> defined.  The default constructor is always overridden by any user-specified no-argument constructor.

Finally, it is possible to give a struct multiple constructors.  These constructors need to be distinguishable by their arguments and are
all declared the same way.

    struct Point2D {
        x, y: int;

        constructor() {
            this.x, this.y = 1, 1;
        }

        constructor(x, y: int) {
            this.x, this.y = x, y;
        }

        constructor(both: int) {
            this.x, this.y = both, both;
        }
    }

Now, you can call any of the constructors using the new syntax and all are valid.  You can have as many constructors as you want so long as
they can be distinguished by their arguments.  For example, it is not possible to specify two no argument constructors or two constructors that
both accept a single integer.
