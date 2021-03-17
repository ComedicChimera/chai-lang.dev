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

## Lifetimes

