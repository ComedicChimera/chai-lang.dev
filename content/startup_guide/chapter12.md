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

This is where the free reference comes into play.  Instead of passing all three vectors as values, we can pass
them as references which always have a fixed size of 4 or 8 bytes depending on your system.  Keeping in mind
that an `int` is 4 bytes, even for this simple example with the largest reference size, our memory usage is now
2/3 of what it original was.  And that is just for a small data structure -- imagine the benefit for a structure
with 5 or 6 fields of larger size.  

To pass by reference, we need to first change the type of the argument to a reference type.  Since we are just
passing from one function down to another, we can use free references which are stack bound to do this.

{{< alert theme="info" >}}When we explain lifetimes in the next chapter, this choice will make more sense.  Right
now, we are just trying to introduce the concepts and syntax associated with references.{{</ alert >}}

The reference type label is `&` followed by the element type.  So for a reference to an `int`, the type label is
`&int`.  For our `Vec3` above, the label is `&Vec3`.  

    func vec_add(v1, v2, v3: &Vec3) Vec3
        => Vec3{
            x=v1.x+v2.x+v3.x,
            y=v1.y+v2.y+v3.y,
            z=v1.z+v2.z+v3.z
        }

Believe it or not, that is all we have to do to refactor our code to use references.  However, two immediate questions
arise: how do we call this function and why didn't you have to change the body at all?  The first question is easier
to answer.  We can simply take references to the vector values we want to pass in.  We do this using `&` operator
which creates a free reference to whatever value we pass in.

    vec_add(&v1, &v2, &v3)

Assuming `v1`, `v2`, and `v3` are defined as vectors elsewhere, the above call would be satisfactory.

The second question of why didn't the function body change functions as a perfect segue to our next topic.

## Dereferencing and Reference Operators

Before we can talk about our `vec_add` function, we need to introduce the idea of **dereferencing**.  This is the
mechanism by which we access the internal value of a reference.  For example,

    let v = Vec3{x=12, y=-5, z=2}

    let vr = &v

`vr` stores a reference to `v` (and therefore has a type of `&Vec3`).  Now, let's say we wanted to pass `vr` to
a function: `vec_mag` that calculates the magnitude of a vector
    
    func vec_mag(v: &Vec3) double do
        return 0 // TODO

We need to be able to access the internal value of the vector reference in order to use it.  To do this, we use
the **dereference operator**.  This operator is the unary `*` and is placed before the reference.

    let v2 = *vr // accesses the value of `vr` and stores it in `v2`.

Now filling in the definition of `vec_mag`,

    import sqrt from math

    func vec_mag(v: &Vec3) double
        => sqrt((*v).x ~^ 2 + (*v).y ~^ 2 + (*v).z ~^ 2)

We first dereference `v` and then access one of its fields.  Obviously, this code is quite ugly looking.  Luckily,
Whirlwind offers a solution in the form of **reference operators**.  A reference operator is an operator that can
operate on a reference as if it were a value.  The `.` operator is one such operator -- it has a reference form
for accessing the fields and methods of references.  It is used identically to the normal `.` operator.

    func vec_mag(v: &Vec3) double
        => sqrt(v.x ~^ 2 + v.y ~^ 2 + v.z ~^ 2)

The code above is exactly equivalent to the code before -- the dereference is happening implicitly as a part of the
operator.  This finally explains our code from the previous section:

    func vec_add(v1, v2, v3: &Vec3) Vec3
        => Vec3{
            x=v1.x+v2.x+v3.x,
            y=v1.y+v2.y+v3.y,
            z=v1.z+v2.z+v3.z
        }

This code was using the reference form of the `.` operator access the vector's fields.  There are other reference operators
besides `.` -- we will cover them in later sections as they become more relevant.

## Mutability

References have two more unique and important properties that we need to dicuss.  The first is the ability to mutate
data indirectly.  Let's consider we wanted to write an *inplace* vector addition function.  The premise is that it
takes two vector references, adds them, and stores the result into the first reference.  This could be written
like so:

    func vec_add_inplace(v1, v2: &Vec3) do
        v1.x += v2.x
        v1.y += v2.y
        v1.z += v2.z

    func main() do
        let
            v1 = Vec3{x=5, y=2, z=-4},
            v2 = Vec3{x=6, y=-2, z=3}

        vec_add_inplace(&v1, &v2)

        println(v1.x, v1.y, v1.z) // prints `11 0 -1`

At face value, this seems odd.  Conventionally, when we pass values to functions, nothing is changed if that value is
mutated inside the function.  For example,

    func add_one(x: int) do
        x++

    func main() do
        let x = 4
        add_one(x)
        println(x) // still `4`

This is because the value is copied when it is passed to the function.  Now, this is true for everything in Whirlwind.
However, copying a reference simply entails copying the address value itself, not the data stored in it.  So, when we
pass vectors into `vec_add_inplace` while the value representing the address is copied, the data isn't: `v1` in `main`
points to the same value as `v1` in `vec_add_inplace`.  

We can show this more clearly by reframing our `add_one` function to actually mutate `x` as reference.

    func add_one(x: &int) do
        (*x)++ // parentheses aren't required but they help with clarity

    func main() do
        let x = 4
        add_one(&x)

        println(x) // now, it's value is `5`

We can use the deference operator on the left side of the `=` operator to mutate the value of the reference.  Since
reference operators connote an implicit dereference, our `vec_add_inplace` code is actually mutating the internal value
of the `v1` reference.  Sometimes including the implicit deferences helps make this clear:

    func vec_add_inplace(v1, v2: &Vec3) do
        // This code is *exactly* equivalent to the original implementation
        (*v1).x += v2.x
        (*v1).y += v2.y
        (*v1).z += v2.z    

## Nullability

However, dereferencing isn't always a safe operation.  Consider the code below:

    let x: &int

    println(*x)

What happens when we run this code?  That's actually a trick question, because the Whirlwind compiler will refuse to compile
that code at all.  The reason is the `x` is marked as a **null reference** because it is unitialized.  Essentially, this means
that the reference points to nothing -- dereferencing it has no meaning. 

Not all null references are as obvious at the one above.  For example,

    func scale_by(v: &Vec3, factor: double) &Vec3
        => &Vec3{x=v.x*factor, y=v.y*factor, z=v.z*factor}

There are actually two problems with the above code, both of which are related to the idea of a null reference.  The first
problem is in the actual value being returned.  See, the new struct we are creating doesn't have a well-defined place in
memory.  It does have a place but not one we can reference.  The compiler will refuse to compile the above code, stating
that you "cannot take a reference to an **r-value**".  This refers to the concept of value category -- an r-value is a
value that doesn't have a well-defined place in memory.  All references to r-values are null by definition.  That's
why we have always created variables and then referenced them -- the variables do have a well-defined place in memory.

So, let's say we fix that bug by doing the following:

    func scale_by(v: &Vec3, factor: double) &Vec3 do
        let result = Vec3{x=v.x*factor, y=v.y*factor, z=v.z*factor} 
        return &result

We still have a problem: the stack frame that `result` exists in is popped when the function `scale_by` exits -- which means
the reference is no longer valid as soon as it is returned.  Once again, Whirlwind will refuse to compile this code citing
an issue with the "lifetime" of the return value.  Lifetimes are a major topic in Whirlwind and something we will cover in
the next chapter, but suffice it to say that they help the compiler to catch errors like the one above -- it helps it to
determine when a reference will and will not be null.
