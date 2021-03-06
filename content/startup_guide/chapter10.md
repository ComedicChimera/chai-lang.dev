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
internal data type, using a type cast.  

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

Now, let's say we wanted to define a function that would make a query and give us the first
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

The code above uses an is-expression to test if `next_result` is a `TextResult`.  The is expression returns
a boolean indicating if the check succeeded.

The is-expression also supports a kind of pattern matching.  For example, consider the following code:

    func display_result(sr: SearchResult) do
        if sr is TextResult do
            printf("Text Result: %s [%d words]\n", sr.get_link(), (sr as TextResult).word_count)
        elif sr is ImageResult do
            printf("Image Result: %s [%d x %d]\n", sr.get_link(), (sr as ImageResult).width, (sr as ImageResult).height)

As you can see, `display_result` takes a search result and prints a different message depending on what the
SearchResult is by checking using a standard is-expression.  You may notice that we have to downcast
three times to be able to access the unique properties of each deriving type.  This is not only wordy but,
for more complex operations, really difficult to decifer.

{{< alert theme="info" >}}More likely than not, you would just make `display_result` a method of SearchResult;
however, this is not always what makes the most sense for your situation.{{</ alert >}}

Luckily, Whirlwind's is-expression supports pattern matching.  We can rewrite the code above to use it
and then break down how it works.

    func display_result(sr: SearchResult) do
        if sr is tr: TextResult do
            printf("Text Result: %s [%d words]\n", sr.get_link(), tr.word_count)
        elif sr is ir: ImageResult do
            printf("Image Result: %s [%d x %d]\n", sr.get_link(), ir.width, ir.height)

We have added an additional part to our is-expressions to enable them to pattern match.  The name that is
defined before the colon is where an automatically downcasted version of `sr` will be stored if the is-expression
succeeds.

{{< alert theme="warning" >}}The variable will be populated with an unusable value if the test does not
succeed.  This value should not be accessed for any reason.{{</ alert >}}

This pattern matching has dramatically simplified our code and made it much more readable.  Plus, we saved ourselves
an extra downcast operation.  

The final piece of the type testing puzzle is the **type-match statement** (and its corresponding expression).  This
statement is an extension of the regular match statement that, instead of matching based on value, it matches based on
type.  Here is another version of the `display_result` function amended to use it.

    func display_result(sr: SearchResult) do
        match sr type to
            case tr: TextResult do
                printf("Text Result: %s [%d words]\n", sr.get_link(), tr.word_count)
            case ir: ImageResult do
                printf("Image Result: %s [%d x %d]\n", sr.get_link(), ir.width, ir.height)

This statement looks very similar to the regular match statement except it has the `type` keyword between the expression
being matched and the `to` beginning the body of the statement.  Moreover, instead of having value patterns in its
cases, it has type patterns.  Note that you can just use plain types as you case expressions instead of full patterns
if that makes more sense in context.

As mentioned before, this statement also has an expression form (like the value match expression) which in this case is
by far the most concise way to achieve our goal:

    func display_result(sr: SearchResult) =>
        match sr type to
            tr: TextResult => printf("Text Result: %s [%d words]\n", sr.get_link(), tr.word_count)
            ir: ImageResult => printf("Image Result: %s [%d x %d]\n", sr.get_link(), ir.width, ir.height)

{{< alert theme="info" >}}Although neither `printf` nor `display_result` explicitly return type, all functions that
return nothing implicitly return the nothing type which can be treated as a value (though it is rarely compiled as one)
so expressions like the one above are in fact valid.{{</ alert >}}
            
