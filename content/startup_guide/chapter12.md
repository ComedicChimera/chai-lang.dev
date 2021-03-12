---
title: "References"
weight: 12
---

## What is a reference?

A **reference** is value that points to a single value or block located somewhere in memory.
But what does the phrase "points to" mean?  Essentially, all memory in a computer has an
**address** which is a unique number that identifies its location within your programs
memory.  A reference simply stores the value of that address that can be then used to access
this memory.

For most if not all programs you will write, the memory layout of your program will look
something like this:

![](/images/memory_diagram.png?width=265.5px&height=485.5px)

with the highest addresses being at the top and lowest addresses being at the bottom.  The two
regions of memory we are going to be most interested in are the **stack** and the **heap**.  These
are where the majority of our memory operations are going to take place, and as the arrows would
imply, the stack grows downward and the heap grows upward as more memory is allocated.  We will learn
more about each of those two regions as we go further in this chapter and the next chapter, but for
now, we are going to focus on the structure of memory itself.

As mentioned before, a reference is simply address.  So if we were to take a random selection of
memory from the stack, we could represent it in the following table (these addresses are arbitrary
and not necessarily indicative of where memory would actually be located in the stack):

| Address | Data |
| ------- | ---- |
| 0xffbb11ff | 0x12 |
| 0xffbb11fe | 0x45 |
| 0xffbb11fd | 0x32 |
| 0xffbb11fc | 0x54 |

The table represents 4 bytes of memory -- each address corresponds to a single byte.  Both the addresses
and the data are represented in hexadecimal.  Those four bytes could be the bytes of an integer or a
float or part of another data structure entirely.  For sake of example, let's say that they are the bytes
of an `int`.

A reference to that integer would then be the memory address denoting the start of that memory block.
This will almost always be the lowest address (`0xffbb11fc` in this case).  

However, two questions immediately arise: how does the computer know how large the reference is and
how does the computer know what type of data the reference stores?  Both of these are answered by
the fact that references have an **element type** that fulfills the dual purpose of indicating the
type of the data being referenced and its size (the size of the data type is the size of the
reference data).

## Free References

The first and simplest kind of reference in Whirlwind is the **free reference**.  These are references to
values on that stack.  For those unfamiliar, the stack is a contiguous block of memory that grows
downward from the top of the address space.  It is defined in terms of **stack frames** where each function
(or method, closure, etc.) pushes a frame onto the stack when it is entered and pops one off when it exits.

{{< alert theme="info" >}}This is a simplification and more of a logical model than a description of how
things actually work down at the assembly level -- it will be a good enough understanding for our purposes.
{{</ alert >}}

The memory for variables, arguments, expression values, and more is primarily stored on the stack -- that
`int` we looked at in the previous section could very well be the data for a variable.  The key point is
the stack is where we store most of our stuff.  The pushing and popping mechanism helps us to easily free
memory once we no longer need it -- it would be incredibly impractical to just leave all that stuff sitting
around on the stack.

However, the biggest downside to the way we have being handling memory thus far is that we have only been
using pure values which means that every time we move them around, be it passing them to a function or
assigning them to a variable, we were copying all that data.  

{{< alert theme="info" >}}This again is a simplification: the Whirlwind compiler performs a common
optimization at a part of regular program compilation called copy elision where the number of copies
performed is dramatically reduced through a variety of techniques.{{</ alert >}}

Let's consider a simple example.  

    type Vec3 {
        x, y, z: int
    }

    func vec_add(v1, v2, v3: Vec3) Vec3
        => Vec3{
            x=v1.x+v2.x+v3.x,
            y=v1.y+v2.y+v3.y,
            z=v1.z+v2.z+v3.z
        }

In absence of any compiler optimization, we are copying 3 vectors when we passed them into `vec_add`.  While
a small sum of memory in the short-run, if you called this function numerous times with a large number of
vectors, your memory usage would grow dramatically.  And this is assuming you aren't writing some recursive
function involving vectors.

{{< alert theme="info" >}}The compiler actually optimizes the function above to involve zero copies, but that
is only because this function is fairly simple for the sake of demonstration.{{</ alert >}}

TODO: introduce free references, reference operators, and dereferencing

