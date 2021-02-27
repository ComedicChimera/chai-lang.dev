---
title: "Interfaces"
weight: 10
---

## Case Study: The Greeter

As with algebraic types, the best way to understand the concept of an interface is to
look at the whole picture of how interfaces fit into the language of Whirlwind.

At a high level, an **interface** is way of grouping types by common behavior.  An
interface defines a set of **methods**, functions that are bound onto a specific type,
that must be implemented in order for the type to have **implemented** that interface.

For example, let's define a `Greeter` interface that requires that all implementing
types to implement a `greet` method that takes a name and prints a greeting.

    interf Greeter of
        func greet(name: string)

The `Greeter` interface begins with the keyword `interf` followed by the name, `Greeter`,
and the keyword `of` denoting the beginning of a block of definitions.  Then, in the body
of the interface, we define one method called `greet` that takes a single argument: `name`.

{{< alert theme="info" >}}These kinds of interfaces are called conceptual interfaces.{{</ alert >}}

To implement interface, we are going to first need to define some types that will implement
it.  Here are two such types:
    
    type CountingGreeter {
        message: string
        count: int
    }

    type NeutralGreeter {
        greeting_kind: string
    }

We will then need to provide `greet` methods for these two types.  We can do this through
an **interface binding**.  Interface bindings allow us to bind a **type interface** onto
an arbitrary and specify which interfaces that binding will implement.  A type interface
is just like a regular interface except the methods have bodies.

A binding begins with the `interf` keyword but instead of being followed by a name, we
follow it with the keyword `for` and the type we want to bind to.  Finally, we use the
`is` keyword followed by an interface to specify which interface the binding implements.

The method of this type interface be defined like normal functions except they will also
have access to a special value: `this`.  `this` points to the type being operated on.

    import println, printf from io::std
    
    interf for CountingGreeter is Greeter of
        func greet(name: string) do
            printf("%s! I have greeted you %d times.\n", this.message, this.count)
            this.count++

    interf for NeutralGreeter is Greeter of
        func greet(name: string) do
            println("Hello,", name + ".", "End of", this.greeting_kind + ".")