# Memory and the Heap

This chapter will cover some of the more systems
oriented aspect of Whirlwind including the stack,
the heap, pointers, references, and all things memory.

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

These pointers can be dereferenced like any old pointer and their type is no different from that of a normal pointer.  The
only difference is where they are pointing.

Because heap values are not always automatically deallocated when the current stack frame expires, you will often want to
delete heap values manually.  This is done with the `delete` keyword.

    delete hp;

The heap pointer we created is now null and the value it held has been deallocated.  You can use this pointer to store other
values, but you cannot access its value until it has been reallocated definitively.

The compiler will flag all uses of `hp` until it is set to point to a definite value.  This includes conditional
uses of `hp` such as those inside an `if` tree.  However, conditional uses will normally result in warnings
not a compile error, but unless you know for sure that the value is not accessed when it not allocated, I would
recommend you treat these warnings as if they were explicit errors.

## References

References are like slightly more convenient pointers. They exist for when you want to use a pointer, but don't want to
have to use the dereferencing syntax associated with them.

To create a reference, you need only use the `ref` keyword.

    let arr = {1, 2, 3, 4};

    let r = ref arr;

The variable `r` is now a reference to `arr`.  References can be treated as if they were the value they store, but in reality
they are just masked pointers.

    let second = r[1]; // perfectly valid

The reference type specifier is connoted with the `ref` keyword followed by the type referenced.

    let r2: ref [int];

References are just as light weight as pointers, but without the additional hassle.  However, references are type safe meaning you cannot
set a reference to a non-reference, but you can set a non-reference equal to a reference.

    let arr2: [3]int = r; // valid

    r = arr2; // ERROR

References are very powerful, but they do still require that you are aware of their presence.  References are for the most part completely managed
by the compiler, so you should be safe when using them.
