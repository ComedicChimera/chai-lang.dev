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
have access to a special value: `this`. `this` points to the type being operated on.

    import printf from io::std
    
    interf for CountingGreeter is Greeter of
        func greet(name: string) do
            printf("%s! I have greeted you %d times.\n", this.message, this.count)
            this.count++

    interf for NeutralGreeter is Greeter of
        func greet(name: string) do
            printf("Hello, %s. End of %s.\n", name, this.greeting_kind)
        
Now, we have created two different kinds of `Greeter`, one that provides you with a
"neutral greeting" and another that counts how many times it has greeted you.  Both
of these greeters can be interacted with in the same way while having different specific
behaviors: they both greet you but in different ways.  

{{< mermaid-graph >}}
    classDiagram
    Greeter --|> CountingGreeter : Impl
    Greeter --|> NeutralGreeter : Impl
    Greeter : greet(name string)
    CountingGreeter : greet(name string)
    NeutralGreeter : greet(name string)
{{< /mermaid-graph >}}

This is the heart of how interfaces work in Whirlwind.  Conceptual interfaces like `Greeter`
define a general pattern of interaction and a type interface implements that general pattern
on a specific type.  

Lot's of things in programming can be thought of in this way.  For example, if you are building
a UI library, you might have an conceptual interface representing a button in general and
then a bunch of different kinds of buttons that all implement that conceptual interface.
