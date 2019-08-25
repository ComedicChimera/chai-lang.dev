# Pointers and Lifetimes

This chapter will cover some of the more systems-
oriented aspect of Whirlwind including the stack,
the heap, dynamic allocation, and, of course,
pointers and references.

## The Stack

The stack is the most basic way in which memory is laid out.
It is a series of sequential memory addresses the allocated
and deallocated in order beginning at the top of the address
space.  

The stack itself is organized in terms of **stack frames**
which are specifically designated regions of memory.  For example,
each function in Whirlwind will create a new stack frame to house
its stack members whenever called.  

All variables in Whirlwind are by default allocated on the stack as are
all designated value types and pointers.

## The Heap

The heap is second way memory is organized.  Due to the innate ordinal nature
of the stack, it is impossible to allocate memory out of order or hold on
to memory for longer than its corresponding stack frame exists.  This is
where the heap comes in.  The heap is a data structure in memory used for
stack values that can be allocated and deallocated at will.  Additionally,
the heap does not need to know the precise size of the object it is allocating
at compile time which makes it very powerful.

In Whirlwind, the heap is not **garbage collected** which means that all memory you
allocate on the heap you are for the most part directly responsible for.
The compiler will in some cases insert heap deallocation calls into your code,
but at least part of your heap memory, you will have to be fully accountable for.

Now, enough with all the theoretical, lets talk about how you use the stack and
the heap.

## Pointers

A pointer is a special data type that points to a position in memory.  It can
be used to point to both stack and heap variables and one of the most essential
ways of passing data around in Whirlwind.

Pointers are light weight and many data structures in Whirlwind such as arrays
can be very large.  Since all data types are Whirlwind are passed by value, meaning
a new copy is created each time they are passed to a function, it is often inefficient
to pass the whole data structure around.  Instead, we can use the light weight nature of
pointers to pass data around.

The pointer type specifier is simply a star followed by the data type being pointed to.

    let p: *int;

The variable `p` is now designated as a pointer to an integer.  Now, we have to point
`p` at something. To do this we use the **indirection** operator.

    let v = 34;

    p = &v; // p = to the indirection of v

Now, p is set to point to the `v`.  You can access and modify the value stored by a pointer using
the **dereference** operator.

    *p = 14;

    let v2 = *p; // v2 is now 14

The value pointed to `p` which is in this case `v` is set equal to 14 and then through `p` accessed
to set the value of the variable `v2`.

> It is possible to have pointers to pointers Whirlwind, though this pattern is rarely ever used.
> However, you can chain as many dereference operators as necessary without parentheses to account
> for this edge case.  A similar pattern is followed in the type specifier as well.

Be careful when dereferencing pointers because if you dereference a pointer with no value or a deleted value,
you can cause an error.

    let p2: *char;

    let c = *p2; // ERROR

Fortunately, situations like this can actually be caught by the compiler and will often be flagged as compile
errors.  However, in more complex situations, it is not guaranteed that the compiler will be able to detect an
invalid dereference call, so be careful.

Pointers in Whirlwind also support **pointer arithmetic** which allows you to directly modify the address of the pointer
as if it were an integer.  So for example, if you had pointer that pointed to the start of a swath of memory called `ptr`,
you could update its value to point to next value in that region of memory simply be incrementing it.

    ptr++;

But, you must careful to not lose memory by moving pointers around.  If you solely desire to move a pointer to a new location
and discard the previous memory, it is often best to just use the pointer's `move()` method to change what it points to.

    // newAddr declared up here somewhere

    ptr.move(newAddr);

The move method will change the address of the pointer, but also delete any *heap memory* that it was pointing to previously.
This can save you from a lot of unnecessary pain that can result by forgetting to delete memory.

## Dynamic Allocation

Dynamic allocation is the process of allocating values on the heap.  You can dynamically allocate any value using the
`make` keyword followed by the data type you want to allocate.

    let hp = make int;

The variable `hp` is now a **heap pointer**.  This means that it is pointing to value allocated on the heap, in this case
an integer.  It is important to note that all values allocated on the heap cannot be directly stored.  They instead must be
reference through pointers.  So, all values returned by dynamic allocation calls are pointers.

There is however an important difference between this kind of pointer and the stack pointers we saw previously.  That is the
in Whirlwind, a heap pointer (or dynamic pointer) is actually a different type from the pointers we looked at previously.  They
are declared just like normal pointers except with the `dyn` prefix before them.

    let x: dyn* int;

These pointers can only be created by a heap call as mentioned before; however, we did not mention that the heap call itself can
only be performed under certain circumstances.  The most simple example of invalid heap allocation occurs when it occurs as an expression
statement.

    make int;

This will warrant an exception that looks something like "unable to perform heap allocation without possibility of an owner".  Owner, in this
context, means a deletable name for the heap memory.

Speaking of, deletion works a little bit differently for heap pointers in Whirlwind.  First of all, any heap pointer will be automatically deallocated
via a delete call inserted by the compiler when their current stack frame expires unless they are being returned or elevated to a higher scope.  Whirlwind
will makes its best guess on whether or not it can delete something, and if it is not sure, it will opt not to delete it in order to avoid problems later on.

Given that, you may want to manually delete memory under certain circumstances to prevent memory leaks.  This is done with the `delete` keyword.

    delete hp;

The heap pointer we created now points to nothing and the memory under it has been freed.  This is important to be aware of as if you try to dereference
a heap (or stack pointer) that is null, you will get a null pointer exception by default.

## Move Semantics and Nullable Dereferencing

Dynamic allocation, as useful as it is for holding onto memory for long periods of times, introduces some issues.  Consider the code below.

    func main() {
        let p = make int;

        // do something with p

        p = make int;

        // do something else with new p
    }

The code, by itself, looks innocent enough, except for one small problem.  Because values on the heap are not deallocated automatically unless they are named
(ie. via inserting a delete statement), when p's value is changed, the memory it pointed too is never deallocated.  It is no longer named and so Whirlwind has
no way of statically determining whether or not to delete it.  And so now, it is just sitting on the heap not doing anything, and it won't be deleted until the
entire heap is deleted at the end of program execution.  This creates what we call, a memory leak, and these can be a huge problem if you are not careful.

The simple way to avoid this problem is simply by deleting the pointer everytime you change its value.

    func main() {
        let p = make int;

        // -- snip --

        delete p;
        p = make int;

        // -- snip --
    }

Now, p's underlying memory is freed, and we are safe to change its value.  This works most of the time, provided that all of this is happening in the same scope
and their is no in between.  For example, if `p` were passed to a function and that function deleted it, you could be in a bit of pickle.  Luckily, Whirlwind provides
several avenues to avoid the problems of manual deletion.

The first is the addition of move semantics.  All dynamic pointers have a builtin method called `move()` which takes in a new memory address as an argument.  Whenever it
is called, it first frees the underlying memory, handling the memory leak problem, and then assigned the pointer to a new value.

    func main() {
        let p = make int;

        // -- snip --

        p.move(make int);

        // -- snip --
    }

Now, whenever p's value is changed, its memory is safely freed, and it is immediately provided with a new value.

> Whirlwind's standard library implement to special kinds of pointers: `UniquePointer`, which automatically implements move semantics
> on its assignment operator (via a special method), and `SharedPointer` which keeps count of how many times it's value is copied and
> only applies move semantics when there is only one reference to it to avoid deleting things where not appropriate.

However, move semantics are not the universal antidote: sometimes you just want to delete something, but still need to try to access it
again.  You can employ a null test before you dereference any heap pointer you are uncertain about:

    func main() {
        let p = make int;

        // `p` may be deleted somewhere

        // null check
        let value = *p if p != null else 0;
        
        // use p elsewhere
    }

This solution of course works, but it is a little verbose.  Luckily, Whirlwind provides a **nullable dereference operator** which simply
returns null when the deference fails as opposed to throwing an error.

    func main() {
        let p = make int;

        // -- snip --

        // *? is nullable dereference
        let value = *?p;

        // -- snip --
    }

Now, we can safely dereference `p` without worrying about null pointer errors.

## Lifetimes

There is one final piece of the puzzle to Whirlwind's memory model, and it is lifetimes.  A lifetime defines, as the name would imply,
how long a piece of memory exists for.  By default, all variables have what is called a **standard life-time**.  
This means the variable (or other value) exists until it goes out of scope (with the exception of heap memory sometimes).  If it is in a function, it is
recreated everytime that function is called.  However, this can sometimes be inconvenient if for example you need data to persist between
function calls.  This is why Whirlwind implements what are called **static life-times**.

A static life-time is a life-time that extends indefinitely.  This means that a variable's value and memory are retained forever.  You can
designate a variable as static via the use of the `static` modifier in the declaration.

    func counter() int {
        let static inc = 0;

        return inc++;
    }

    func main() {
        counter(); // 0
        counter(); // 1
    }

As you can see, because `inc` is marked as static, its memory is never deallocated, and its value persists.  If you allocate any thing statically,
you should never really need to worry about deleting unless it is a heap pointer and you want to change its value safely.  These values are automatically
disposed of when the program exits so there is no reason to worry about them.
