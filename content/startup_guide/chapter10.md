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

Now that we've seen all of the necessary definitions, let's take a look at how you might use
an interface in code.

    func greet_n(g: Greeter, name: string, times: int) do
        for _ in 1..times do
            // this syntax accesses the greet method and calls it on `g`
            g.greet(name)

    func main() do
        let cg = CountingGreeter{}
        greet_n(cg, "Bob", 10)

        let ng = NeutralGreeter{greeting_kind="conversation"}
        greet_n(ng, "Allison", 5)

The function `greet_n` takes an argument of type `Greeter`.  This specifies that anything
that implements the `Greeter` interface can be passed in.  Thus, we can use both kinds of
greeter as the argument to `greet_n`.  Notice that calling `greet` on our `Greeter` instance
implicitly invokes the specific `greet` method of whatever type was passed in.

{{< mermaid-graph >}}
    stateDiagram-v2
        state "greet_n called" as s1
        state "g.greet called" as s2
        state "CountingGreeter.greet called" as s3
        state "NeutralGreeter.greet called" as s4
        s1 --> s2
        s2 --> s3 : g is CountingGreeter
        s2 --> s4 : g is NeutralGreeter 
{{</ mermaid-graph >}}

This behavior has a fancy term: **polymorphism**.  Breaking that word down into its roots,
"poly" refers to "many" and "morph" refers to "form".  So polymorphism is the idea of something
having many forms.  In this case, our `Greeter` instance `g` can either be a `NeutralGreeter` or
a `CountingGreeter`.  We, as the programmer don't have to care because we know that we want
to greet someone by a specific name -- the details of what type of greeting happens are unimportant,
at least for the purposes of writing our code.

Lot's of things in programming can be thought of in this way.  For example, if you are building
a UI library, you might have an conceptual interface representing a button in general and
then a bunch of different kinds of buttons that all implement that conceptual interface.  Then,
we you are writing your event loop and someone clicks on a button, you can just call the button's
click handler -- you don't care what that handler does just that the button was clicked.  This
is the core idea of interfaces: many data structures that all do the same thing in different ways.

## Type Testing and Downcasting

The **type-test** expression or is-expression is a way to test what type an interface is really
storing.  For example, in our `greet_n` method above, `g` could either be a `CountingGreeter` or
a `NeutralGreeter`.  Sometimes, it is useful to know what type of interface we are dealing with,
especially if we are planning to **downcast**, that is convert an interface instance into its
internal data type using a type cast.  

{{< alert theme="warning" >}}Downcasting can fail at runtime if the internal type doesn't match
the type we are downcasting to so it is really important to test and make sure the cast you are
performing is valid.{{< /alert >}}

Let's consider a simple example where we have some form of querying API that gives us search results.
There will be two kinds of search results: text and images.  We will implement this "setup" using
interfaces.

    interf SearchResult of
        func get_link() string

    type TextResult {
        link: string
        word_count: int
    }

    interf for TextResult is SearchResult of
        func get_link() string => this.link

    type ImageResult {
        link: string
        width, height: int
    }

    interf for ImageResult is SearchResult of
        func get_link() string => this.link

Now, let's see we wanted to define a function that would make a query and give us the first
result the fits some parameters.  For our example, we will say that the only parameter is whether
or not to allow image results. 

To test if an interface has a certain internal value, the is-expression is used.  It is constructed
using a value, the keyword `is`, and a type to match against.  If the value matches that type, it
returns true.

    // result retrieval function -- implementation unimportant
    func get_next_result(search_str: string) SearchResult

    // main query function
    func query(search_str: string, allow_images: bool) SearchResult do
        while true do
            let next_result = get_next_result(search_str)

            if allow_images || next_result is TextResult do
                return next_result

TODO: more explanation, is pattern matching, type match expression?

            
