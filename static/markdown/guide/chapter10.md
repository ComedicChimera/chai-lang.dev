# Interfaces

Interfaces are the primary method of classifying types in Whirlwind.  They
are also by far one of the most versatile tools available in Whirlwind.

There are two primary kinds of interfaces: **type interfaces** and **classifying interfaces**.
The latter is normally just referred to as an interface as it is the standard version of
interface available in other languages.

## Methods

Methods are functions that are bound onto a type.  You can think of them as type specific
behaviors.  You have already been exposed to various methods before such as those present
on lists and dictionaries (has, get, etc.).  Methods are called by first placing a `.` operator
after the variable followed by the method name and standard call syntex.

    var.method();

The above code calls a method of the type of the variable `var` by the name `method`.  The variable
itself can be accessed inside the method through the use of the `this` pointer akin to the constructors
we saw before.

    // -- snip --

    // method definition of var
    func method() int {
        return this.x ~^ 2; // return the property of var called "x" squared
    }

    // -- snip

In order to access any property of call any other method of `var`, you must use the `this` pointer.

## Interface Binding

You can specify the methods of a type in a type interface which is a collection of methods
specific to one type.  You must then bind this type interface
to whatever type you want to receive it.  A difference between Whirlwind and almost every other language
is that Whirlwind allows you to bind methods onto **any** type.  You can attach methods to integers
or to lists or the structs.  In fact, all the default methods in the standard library were attached
this way.

To understand interface binding, we will use an example.  Consider you have a type `Point` which is
defined as follows.

    struct Point {
        x, y: int;
    }

Let's say you wanted to bind a method called `distance` onto `Point` that would calculate the distance between
two points.  You would first need to declare a type interface for (bound to) Point.  You can do this using the
following syntax.

    interf for Point {

    }

As you can see the bind and the declaration occur simultaneously; it is impossible to declare a type interface without
binding it.

You can then simply write the methods for the given in type inside of the type interface declaration.  In this case,
we will implement our distance method here as follows:

    interf for Point {
        func distance(other: ref Point) double {
            return sqrt((this.x - other.x) ~^ 2 + (this.y - other.y) ~^ 2);
        }
    }

Then, given any two instances of `Point` it will be possibly to find their distance.

    let p1 = new Point{x = 4, y = 2};

    let p2 = new Point{x = -7, y = 6};

    let d = p1.distance(ref p2); // distance between p1 and p2

This definition of distance will work anywhere in which `Point` is also defined and can be called
on any instance of Point.

> It is technically impossible to bind to the static version of a struct or a type class.

## Classifying Interfaces

A classifying interface is an interface used to classify types based on their methods and provide
functionality based on that shared classification.  To declare a classifying interface, you follow
a similar procedure to that of declaring a type interface, but replace the binding statement
with the name of the interface.

    interf MyInterf {
        // methods go here
    }

As mentioned before, classifying interface fulfill two roles: functionality classifiers and providers.
As such, there are two ways to define methods within them: via incomplete declaration or via complete
declaration.  

Incomplete declaration defines a method that must be defined in any type that implements
the interface.  This is how a classifying interface classifies.  To perform an incomplete declaration
simply replace the body of the function with a semicolon.

    interf MyInterf {
        func mandatoryMethod() int;
    }

Now, all derived types must implement mandatoryMethod in order to be considered a "MyInterf".  These
methods must be implemented exactly as their signature defines with no variation.  

To specify that a type implements an interface, you use the `is` syntax following the bind syntax.

    interf for SomeType is MyInterf {
        func mandatoryMethod() int {
            // -- snip --
        }
    }

As you can see, the implements of any given type must be listed after `is`.  It is important
to note that a type **CAN** implement multiple interfaces.

> You cannot bind type interfaces to classifying interfaces as classifying interfaces have another
> mechanism for defining functionality.  This all prevent interfaces for implementing other interfaces.

Now that `SomeType` implements `MyInterf`, a `SomeType` may be passed in place of a `MyInterf` like so.

    func doSomething(a: MyInterf) int
        => a.mandatoryMethod() * 3 + 7;

    func main() {
        let t: SomeType; 

        doSomething(t);
    }

Using the interface as the type means that **all** implementees of `MyInterf` may be passed as a in the above
method.  As you can see above, the parameter `a` may only access the `mandatoryMethod` as that is the only method
that is guaranteed to present on any given `MyInterf`.  So, even if `SomeType` were defined with additional methods
and/or properties, because it is passed as a `MyInterf` only `MyInterf` properties may be accessed.

> It is impossible to pass a raw classifying interface as an argument to any method that accepts a classifying interface
> as a type because interfaces do not have instances and the type designation of a classifying interface technically
> refers to the implementees not the interface itself.

## Implicit Implementation

Because of the nature of interfaces in Whirlwind, it is possible to pass a type as a classifying interface without the
type explicitly declaring itself a member of that interface.  Consider our example of `Point` from before.  Say there was
an interface that resembled the following.

    // placing "I" before the name of an interface is a common convention
    interf IPoint {
        // note: the type here is "Point" not "IPoint"
        func distance(other: ref Point) int;
    }

It would be possible to pass `Point` in place of an `IPoint` because the definition of distance within `IPoint` is congruent
with said definition of distance within `Point`.  So the following code, would in fact be valid without modifying the
type interface definition for `Point`.

    func doSomethingElse(p1, p2: ref IPoint) double
        => p1.distance(p2) * 0.l7;

    func main() {
        let p1 = new Point{x = 3, y = 2};
        let p2 = new Point{x = -1, y = 5};

        doSomethingElse(p1, p2);
    }

The demonstates **implicit implementation** a very important concept in Whirlwind, but it has one caveat that will soon
become apparent.

## Defining Functionality

In our discussion of classifying interfaces, I mentioned that there were two kinds of definition within an interface:
complete and incomplete.  So far, we have only examined the ramifications of the latter.  Complete declaration entails
defining a method with a body.  When you do this, any type that implements the interface will also inherit that method
**WITHOUT** having the implement it.

    interf IThing {
        func method1() int;

        func method2(x: IThing) int
            => this.method1() + x.method1();
    }

    struct Thing {
        a: float;
        b: str;
    }

    interf for Thing is IThing {
        func method1() int
            => floor(a) * b.len();
    }

    func doThing(a: IThing) float
        => cbrt(a.method1());

    func main() {
        let t = new Thing {a = 3.14, b = "beef stew"};
        
        t.method1(); // still works

        doThing(t); // still valid

        let t2 = new Thing {a = 1.141, b = "stew beef"};
        
        t1.method2(t2); // this is valid as well
    }

As you can see in this length example, completely defined methods are applied to any type to explicitly implements an interface.
They are also valid in general interface cases such as in `doThing` although that is not shown above.

If you recall from earlier, I mentioned that there is a caveat to implicit implementation.  Well, here is where it lies.  Implictly
implemented types do not receive completely defined methods because they do not technically implement the interface.  This does not
mean that when passed as an interface they will not receive the methods, but rather that it will not be globally defined on the actual
type itself.
