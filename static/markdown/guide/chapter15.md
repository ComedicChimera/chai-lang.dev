# Captures

Captures are a small, but critical part of Whirlwind.  They are how we control scope and what
exists in it.

## Capture Basics

Captures define how the higher scopes are represented in the current scope.  They do this by excluding
and redefining the variables and types that exposed in the upper scope.

A capture begins with the `with` keyword and is followed by the list of variables that are intended to be
visible in the current scope.

    let (a = 10, b = 4, c = 3);

    with [a, b] {
        a = b + 2;

        // ERROR: c is not defined
        b = a * c;
    }

By default, captures take in variables by reference, but you can change it to take in variables by value
using the `val` prefix.

    // a, b, and c are defined up here

    with [val a, b] {
        a = b + 2;

        b = a * 4;
    }

    let d = a + b; // d = 34 (10 + 24);

Because `a` is taken in by value instead of by reference, its value outside the sub-scope is not effected.

> When we say taken by reference here, we just mean that the compiler exposes the variables normally:
> no pointer or reference is created.  Similarly, taken by value just means that the compiler creates
> a local copy of the variable in the enclosed scope.

You can also redefine a variable as constant within a sub-scope or take ownership of it using the corresponding keywords.

    let p = make int;

    let a = 5;
    *p = 14;

    with [own p, const a] {
        a = 3; // ERROR: a is constant

        delete p; // this sub scope owns "p" so ok
    }

    // a is no longer constant
    a = 15;

Captures can also be used to exclude a variable by means of the `!` prefix.

    with [!a] {
        b = c * 4;

        a += b; // ERROR: a does not exist
    }

However, when the capture is excluding values, it accepts all non-excluded values unless another value is specified.

    with [!a, c] {
        // neither b nor a exist
        c = b * a;

        // A, ok
        c = 4;
    }

Using these two patterns together is almost always redundant so we don't normally recommend doing it.

Finally, if a capture is not told to exclude or include anything, it excludes everything by default.  This
will isolate any scope wrapped by an empty capture from accessing any outside information.

## Captures on Functions

In addition to being bound to ordinary sub-scopes, captures can be attached to the end of any function.

    let var = 45;

    func f() with [!var] {
        var = 4; // ERROR: var is not defined
    }

But, more importantly, captures can be used to regulate how a closure accesses and manages information.

    func getFn() func()(str) {
        let privString = "abc";

        return | | with [val privString] {
            privString += "d";

            return privString;
        }
    }

    func main() {
        let f1 = getFn();

        f1(); // abcd

        let f2 = getFn();

        f2(); // abcd

        f2(); // abcd
    }

Because we use a capture to specify that the lambda should take in `privString` by value and not by reference,
our changes to `privString` are only preserved within the temporary scope of the closure.  So even when the same
closure is called twice, `privString` still keeps it initial value because all changes made to it actually only
affect a copy of `privString`, not `privString` itself.