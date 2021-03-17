---
title: "Region and Lifetimes"
weight: 13
---

## Heap Allocation

So far, we have only dealt with memory allocated on the stack.  However, there is another kind 
of memory that we allocate in some cases even more than stack memory: heap memory.  

As mentioned in the previous chapter, the heap is the companion data structure to the stack and
grows up toward the stack.  However, unlike the stack, the heap is, by default, not bound by
any specific allocation structure.  We can allocate wherever and whenever we like, and heap
memory lasts forever until we explicitly delete it.  This affords us a number of advantages over
the stack: heap memory is

1. Persistent -- It can last as long as we need it to and be deleted whenever
2. Dynamic -- It's size can be determined completely at runtime; we don't have to know how much we
need until we need it
3. Resizeable -- It can be reallocated and adjusted as the amount of data we need to store changes

But, no boon is without its costs, and heap memory is no exception.  Particularly, the fact that
heap memory is persistent means that if we forget to deallocate it, we could cause a **memory leak**:
a situation in which program memory is never deallocated (often repeatedly) and left to float in
the proverbial "aether" of the heap, just wasting space.  

Conversely, heap memory can also be deleted at any time which means we have no guarantee that it
exists whenever we access it.  Remember those pesky null references from the previous chapter?  Those
turn from a minor inconvenience, easily debugged, into a veritable nightmare -- you never know
whether or not something exists or whether it has been moved or deleted right underneath your nose.

There are three main approaches languages take to dealing with the problem of heap memory.  The
first is the approach taken by languages like C and C++ which is **manual memory management**.
The programmer is responsible for keeping track of their own heap memory (with the help of some
data structures like smart pointers and patterns like RAII in the case of C++).  This is often the
most computationally efficient method for managing memory but comes with the major caveat of
introducing a plethora of insidious memory errors and turning debugging into a horror show for
the unseasoned developer.

The second approach is to employ a **garbage collector** to manage memory for the user at runtime.
This complex mechanism uses a variety of techniques to detect when memory is being used and when it
isn't to tactfully clean up the proverbial garbage the application leaves behind.  In contrast to
the first approach, this approach is much easier for the developer and can often save a lot of time
and code bloat that would otherwise be spent dealing with chaotic mess of the memory wildwest that
is the unmanaged heap.  The cost is in performance: even the most cleverly designed garbage collector
can slow an application down significantly and drives its memory usage through the roof.  

The third and final solution attempts to find a balance between these two approaches and has gained
popularity recently with languages like Rust championing it as their signature feature: 
**static memory analysis**.  Essentially, in these languages, the compiler is responsible for analyzing
your code to determine whether or not it is "memory safe".  To do this, the compiler and language have
a set of rules built in to dictate how the developer can manipulate memory -- rules that the compiler
can check against to determine if the code is valid.  For example, Rust uses a paradigm called ownership
from which to build its ruleset.

That long-winded history lesson finally brings us to Whirlwind.  Where does our language of study fit
into the picture?  Whirlwind is solidly in the third category of static memory analysis.  However,
Whirlwind takes a markedly different approach to static analysis encapsulated in its **region model**.

## Regions

A **region** is a small, scoped allocation arena on the heap.  Essentially, they act like variable length
stack frames for storing data on the heap.  

When we allocate data on the heap in Whirlwind, we have to specify which region it belongs to.  This
region is known as its **owner**.  All data in a region is deleted when its owner is closed.  

Regions in Whirlwind close automatically -- they have a given scope of existence and close whenever that
scope ends.  All memory in that region is then returned to the allocator to be used later.  Since regions
have no fixed length, you can allocate as much memory as you want in a region.  However, you must be
cognizant of how long you want that memory to linger since, again, that memory can only be freed once the
region is closed.  

A reference that is allocated in a region is called an **owned reference**.  It type label is similar to
that of a free reference except the label has the `own` keyword placed before it.  So an integer reference
allocated on the heap would have the type label: `own& int`.  

To actually allocate such a reference, we first need to create a region to allocate in.  This can be done
with a **region declaration** which specifies a name for a given region as well as its lifetime.  All region
declarations begin with the `region` keyword followed by the name of the region.  This is simply the local variable
name that can be used to refer to the region during allocation.  

For a **local region** declaration -- that is a region whose lifetime is confined to the stack frame of the
function it is created in -- we need to place the `local` keyword at the end of the region declaration.  So
a sample declaration of a local region `re` would look like this:

    region re local

To actually allocate in this region, we need to use a **make expression**.  All make expressions begin with
the keyword `make` followed by a **region specifier** and an **allocation parameter**.  As you might guess,
a region specifier tells the make expression what region to allocate in.  The first and most common region
specifier is the **explicit specifier** which is written as `in[r]` where `r` is the name of the region being
allocated in.  

The allocation parameter tells the make expression what we are allocating.  The first kind of allocation
parameter is simply a type label: the make expression will allocate a reference with an element type of the
type label and return it.

Putting all this together, the make expression for an `own& int` in our local region `re` would look like this:

    make in[re] int

All in all not that complicated.  Now let's put all these pieces together to allocate a `User` reference on the
heap.  The definition of `User` is:

    type User {
        id: int
        name, email: string
    }

Now for the actual allocation function,

    func create_user(name, email: string) own& User do
        region re local
        let u = make in[re] User

        // `get_id` defined elsewhere
        u.id = get_id()
        u.name = name
        u.email = email

        return u

We create a region, allocate our user, populate its fields, and return it.  However, the above code has a problem:
a big one.  In fact, this problem is so large, the compiler will refuse to compile this code at all.  Remember what
the lifetime of a local region is: its lifetime is confined to that of its enclosing function.  This means that when
`create_user` returns, all the memory inside `re` is deleted, including our `User` which means we are returning a 
null reference which is a big problem.  

To fix this, we need to change the lifetime of our region.  How do we do that?  The naive answer would be to have
some specifier that says the region is "nonlocal", allowing to exist outside the scope of the current function.
However, the obvious problem with this is that we have no way of knowing how long the caller wants that `User` reference
to exist.  

TODO: rework regions to work more like type parameters (syntax unclear...), try to remove local region declarations:
they are just tedious and annoying

We solve this by having the caller specify what region to allocate in.  We do this by having the caller pass a region
in as a value.  See, when we declare a region, we are really just declaring a special variable that stores some reference
to our region -- it isn't and can't be treated like a reference (and implementation wise it may not actually be one), but
internally, it is just data.  That means we can treat it like a value.  All regions use the special type label `region` to
denote that an argument or variable is storing a region.  Now, let's refactor our `create_user` function to take in a region
value.

    func create_user(re: region, name, email: string) own& User do
        let u = make in[re] User

        // `get_id` defined elsewhere
        u.id = get_id()
        u.name = name
        u.email = email

        return u





## Lifetimes

