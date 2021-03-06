---
title: "Duck Typing"
weight: 11
---

## Extension Binding

It is possible and common to bind a type interface onto a type without explicitly
deriving any interfaces.  This is (informally) called **extension binding** and looks
just like a regular binding without any named implementations.

    type Vec2 {
        x, y: double
    }

    interf for Vec2 of
        func add(other: Vec2) do
            this.x += other.x
            this.y += other.y

        func mul(scalar: double) do
            this.x *= scalar
            this.y *= scalar

        func dot(other: Vec2) double
            => this.x * other.x + this.y * other.y

We bind a type interface onto our `Vec2` type that provides three methods: `add` (for vector addition),
`mul` (for scalar multiplication), and `dot` (for the dot or scalar product of two vectors).  Now,
when we use our `Vec2` type in code, we can invoke these methods.

    let
        v1 = Vec2{x=3, y=4},
        v2 = Vec2{x=-7, y=2}

    println(v1.dot(v2))

    v1.add(v2)
    v1.mul(v2.dot(v1))

    println(v1.x, v1.y)

{{< alert theme="info" >}}It is always possible to call methods of a type directly regardless of whether
or not the type explicitly implements any interfaces.  For example, for our `CountingGreeter` on the
previous page, you could call `cg.greet` directly without upcasting.{{</ alert >}}

Up until this point, we have only looked at **explicit implementation** where an interface binding
states directly what interfaces it implements.  However, there is an alternative form of implemented
called **implicit implementation** where a type can implement an interface simply by matching its
**abstract methods**, that is a method of a conceptual interface that does not provide an implementation. 
For example, consider the conceptual interface `Scaleable` defined below:

    interf Scaleable of
        func mul(scalar: double)

This interface defines one abstract method: `mul`.  If you will notice, both `Vec2.mul` and `Scaleable.mul`
have the exact same signature: they take one parameter called `scalar` that is a double, return nothing,
and are called `mul`.

Because `Vec2` provides exact matches to all of the methods of `Scaleable`, `Vec2` is said to implicitly
implement `Scaleable` and thus can be treated as a deriving interface of it.

    func scale_by_2(s: Scaleable) Scaleable do
        s.mul(2)
        return s

    func main() do
        let v = Vec2{x=1, y=6}
        v = scale_by_2(v) as Vec2

        println(v.x, v.y) // prints `2 12`

This process is called **duck typing** where one type is implicitly coerced to another.  We can also use
tools like type-testing on `Scaleable` as is `Vec2` were an instance of it.

    s is Vec2 // Totally valid

This is a surprisingly common kind of implementation and is really useful for connecting different
APIs (such as integrating a package someone else has written to your own codebase).

## Virtual Methods

A **virtual method** is a method defined in a conceptual interface that has a default implementation.
These methods do not need to implemented in deriving interfaces and when called without such a
sub-implementation will run in their default form.  These methods are usable for providing common
functionality without requiring reimplementation everytime.

Consider our `Scaleable` interface again.  Let's create a virtual method called `div` that "divides"
instead of multiplying.

    interf Scaleable of
        func mul(scalar: double)

        func div(scalar: double)
            this.mul(1/scalar)

Notice that `div` has a body and that we can call other methods of `Scaleable` within the bodies of its
virtual methods.

Using this `div` method, we could add a method `half` that calculates half of a given `Scaleable`.

    func half(s: Scaleable) Scaleable do
        s.div(2)
        return s

Because `Vec2` implicitly implements `Scaleable` still, we can pass it to half and use the `div` virtual
method of `Scaleable`.

    half(Vec2{x=4, y=3}) // => Vec2{x=2, y=1.5}

{{< alert theme="info" >}}This makes logical sense since the virtual method `div` only uses other methods
of the interface which are guaranteed to have implementations.{{</ alert >}}

However, the virtual method `div` is not defined on `Vec2` itself (unlike `mul`).

    v.div(3) // ERROR

This is one of the most important distinctions between implicit and explicit implementation: virtual methods
are only accessible as methods on types that derive by explicit implementation.

For example, if we define an entirely different type,

    // 2 x 2 matrix
    type Mat2 {
        // [ a b ] 
        // [ c d ]
        a, b, c, d: double
    }

And define a binding for it that explicitly implements `Scaleable`,

    interf for Mat2 is Scaleable of
        func mul(scalar: double) do
            this.a *= scalar
            this.b *= scalar
            this.c *= scalar
            this.d *= scalar

Then, we can use the `div` method on `Mat2`:

    let m = Mat2{a=3, b=4, c=5, d=6}
    
    // implemented `mul` method
    m.mul(4)

    // virtual `div` method of `Scaleable`
    m.div(2)

The reason for this is that `Mat2` is always `Scaleable` by definition whereas something like `Vec2` is
only `Scaleable` by duck typing (ie. only in certain contexts) so it doesn't inherit the virtual methods
of `Scaleable` onto the type itself.

## Method Overriding

**Method overriding** is the mechanism by which the body of a virtual method can be superceded (or overridden to use
our current "lingo") by a different implementation in a deriving type.  Consider the following situation:

    import format from io::fmt
    
    interf Displayable of
        func to_string() string

        func display() do
            println(this.to_string())

    type User {
        id: int
        name, email: string
    }

    interf for User of
        func to_string() string
            => format("%s (%d) at %s", this.name, this.id, this.email)

        func display() do
            printf("%s: %s", this.name, this.email)

`User` implicitly implements `Displayable` meaning we can upcast it to a `Displayable` interface instance.  However,
notice that both `User` and `Displayable` provide implementations of the method `display`.  So then, what
happens when the code below is run:

    let u = User{id=12345, name="John Doe", "jdoe@email.com"}
    let d = User as Displayable

    d.display() // ???

The answer is that the `display` method of `User` is still run even though we upcast to an interface instance of
`Displayable`.  This is because `User` **overrides** the display method of `Displayable`, and the overridden method
is always called over the default implementation.  This is true of both implicit and explicit implementations of
interfaces.

    d.display() // prints `John Doe: jdoe@email.com`
