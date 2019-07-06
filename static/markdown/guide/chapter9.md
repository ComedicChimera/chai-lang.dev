# Ownership and Lifetimes

As a system language, managing how long resources exist and where they are available
is incredibly important to Whirlwind.  In this chapter, we will discuss how this is done.

## The Basics of Ownership

Ownership in Whirlwind is different from how you might think of it in other languages.
First and foremost, ownership primarily applies only to **heap pointers**.  It is the
method by which we manage dynamic memory.  

Ownership in Whirlwind is defined as the right to delete something.  Effectively,
only the scope which owns a given pointer may use `delete` on it.  Now, we must
discuss how Whirlwind determines ownership of a pointer.  By default, the scope
that creates a given pointer is the owning scope.  

    // this scope owns p
    let p = make int;

    {
        // this scope owns p2
        let p2 = make float;
    }

If a scope is a sub-scope of the scope that owns the given type, then that scope also
has ownership.

    let p = make int;

    {
        // this scope also owns `p`
        delete p;
    }

The only exception to this rule is when the sub-scope is a function.

    func main() {
        let p = make int;

        func sub() {
            // ERROR: ownership violation
            delete p;
        }
    }

If Whirlwind detects a specific ownership violation, it will fail compilation.

If Whirlwind determines that a pointer could be null via an owned deletion,
it will by default warn you during compilation, but it will not fail compilation.  
However, you can change this setting by specifying it in the `whirl build` call.

<div class="command-window">
    whirl build filename --unsafealert=ERROR
</div>

The other values for this argument are `WARN`, the default; `ERROR` which tells the
compiler to throw an error; and `SILENT` which disables this type of checking.

> We only recommend setting this to silent if you are confident that you're code is
> safe and/or are looking for a faster build as turning this setting off will speed
> up compilation.

Ownership is not only bound to the scope, but also to the variable itself.  Consider the
below example.

    func main() {
        let p = make int;

        let p2 = p;

        // FAIL: "p2" doesn't own the memory at "p"
        delete p2;
    }

Because `p` is the variable that owns that memory, `p2` cannot delete it.

## Transferring Ownership

There are two ways to transfer ownership of a heap pointer.  The first method occurs
implicitly when a heap pointer is returned from the scope that created it.  Consider
the below example.

    func fn() *int {
        let p = make int;

        return p;
    }

    func main() {
        let p = fn();

        // valid because p is now owned by main
        delete p;
    }

In this example, the function `fn` creates `p` initially and is thus its initial owner.
However, when `p` is returned from `fn`, its ownership is transferred to main.

The second method of tranferring ownership is explicit.  It is connoted by the `own` keyword.

    func fn(p: *own int) {
        // ok
        delete p;
    }

    func main() {
        let p = make int;
        fn(p);
    }

In this example, the variable `p` is created in main and passed to `fn`.  In `fn`, p is explicitly
specified as owned and therefore ownership of `p` is transferred to `fn`.

Because `own` is a type modifier, you can attach to anything, even a variable.
Using `own`, we can revise the example from the first section to work as intended.

    func main() {
        let p = make int;

        let p2: own *int = p;

        delete p2;
    }

You can also add this modifier to a struct to allow the struct to delete it.

    struct S {
        p: own *int;
    }

    func makeStruct() S {
        let p = make int;

        let s = new S{p = p};

        // ordinarily, this would cause an error since s doesn't own `p`
        // but because we specified that it did, this is ok
        return s;
    }

    func main() {
        let s = makeStruct();

        // ERROR: main doesn't own p, only S does
        delete s.p;
    }

If a bit of memory is owned by a given struct or type, it can only be deleted in the **finalizer**
of that struct.  A **finalizer** is simply a special method that, if it exists, will be called during
the deletion of a type.  The method looks like the following:

    interf for S {
        func __finalize__() {
            // ok because in the finalizer
            delete p;
        }
    }

By its very definition, finalizers are unable to called explicitly and instead are called implicitly
by the compiler.

## Ownership as a Responsibility

Ownership is more than just a right: it is a responsibility.  When you take ownership of a piece of
memory, you must free it.  This means two things: one, it is impossible to change where an owning
pointer points without explicitly deleting the memory, and two, it is (almost) impossible
to create memory leaks.

Because of these rules, the following code would be considered invalid by the compiler.

    func main() {
        let p = make int;

        // ERROR: the initial memory you allocated was never freed
        p = make int;
    }

There are however, a number of ways to handle deletion.  The first is by simply deleting the pointer.

    func main() {
        let p = make int;

        delete p;

        // because "p" is the l-value in the assignment, this is ok
        p = make int;
    }

This can be a little tedious because it: a) requires an extra step, b) is not very functional, and c) mandates
that a pointer be assigned to before it can be used as an r-value (not being assigned to) again.

All these reasons means that there is a probably a better way of handling this.  The first solution solves
problem b because it is a function.  This function is `free()` which is actually a method of pointers.

    func main() {
        let p = make int;

        p.free();

        p = make int;
    }

Now, we could use `free` as just a normal function which, as you will see later, has its benefits.  However, this is
not the best solution.  The best solution is the `move()` method, which solves all of our problems.

    func main() {
        let p = make int;

        p.move(make int);
    }

In this scenario, the memory `p` points to is safely freed and its value is updated all in own step.  

> Using make expressions like this is often not recommended as it creates unique ownership problems unless
> the function it is being passed to takes ownership of it (which move() does).  For safety reasons,
> the compiler only allows such use of make expressions when the function takes ownership.  Otherwise, it throws
> an error.

Now, all of this seems rather tedious.  Having to explicitly delete everything you allocate on the heap would be a huge pain.  Luckily,
Whirlwind's compiler takes care of this.  Whenever the owner of a piece of memory goes out of scope, Whirlwind implicitly frees that underlying
memory.  For this reason, things like finalizers are reserved primarily for managing things like file-handles and for handling destruction logic
like incrementing a counter.

## Lifetimes

There is one final piece of the puzzle in terms of ownership, and that is lifetimes.  By default, all variables have what is called
a **standard life-time**.  This means the variable (or other value) exists until it goes out of scope.  If it is in a function, it is
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

As you can see, because `inc` is marked as static, its memory is never deallocated, and its value persists.  However, this has some interesting
implications for ownership.  First and foremost, it means that static values are **perpetual owners** aka they can never be deleted.  Once a
static variable takes ownership of a value, it holds that ownership forever.  It is thereby impossible to transfer ownership from a static owner,
to delete a static owner, or to adjust where a static owner points to.

    func takeOwnership(a: own *int) {
        // -- snip --
    }

    func main() {
        let static v = make int;

        // ERROR
        delete v;

        // ERROR
        takeOwnership(v);
    }

The reason for this is that is impossible to determine the exact parameters of ownership on any specific static variable.  So, we simply say that once
ownership is taken, it cannot be revoked.  Now, this does not mean you cannot mutate static variables generally, just not static owners.  And, it also
does not preclude you from modifying the memory under a static variable or passing its value or even its address around.  But, any transfer or adjustment
or ownership on a static variable will be considered invalid.
