# Decorators

A decorator is one of the more niche but also powerful tools provided by Whirlwind.
It is an advanced application of all the concepts we have previously learned
distilled down with a concise, powerful syntax

## The Premise of a Decorator

The main functionality of a decorator is simply to provide common behavior between
different functions.  Consider a simple example where you wanted to have a certain
set of a functions log when they are called and when they return.  You could simple
insert your log statements at the end of each function like so:

    // assume log is defined up here somewhere

    func fn1() int {
        log("called")

        // blah blah

        log("returning")

        return 0;
    }

    func fn2() int {
        log("called")

        // blah blah blah

        if (condition) {
            // blah
            log("returning");
            return 2;
        }

        log("returning")
        return 3;
    }

    // etc.

As you can see, this behavior becomes quite tedious to implement on each function individually especially
as the functions get more complicated.  This is the problem that a decorator could solve.

A decorator by itself is merely a function that takes a function as its argument and spits out a new, wrapped
function.  Consider our logging example from earlier.  Let's create a decorator for that intended behavior.

    func logger(f: func()(int)) const func()(int) {
        func wrapper() int {
            log("called");

            let res = f();

            log("returning");

            return res;
        }

        return wrapper;
    }

As you can see, the decorator takes in the desired function and returns a wrapper
for that function.  There are two important things to notice here: first the
return type is `const` because `wrapper` is a constant and if it weren't constant
we would get an error.  Secondly, and more importantly, `wrapper` calls `f` inside
of its body and when it is returned it is not called.  This gives us the desired
behavior.

Now, we can use decorators and little bit of lambda magic to improve out setup from before.

    const fn1 = logger(| | { 
        // blah blah

        return 3;
    });

    const fn2 = logger(| | {
        // blah blah blah

        if (condition) {
            // blah
            return 2;
        }

        return 3;
    });

Now that is significantly more concise, but it can certainly be improved.  But, before we get to that,
there are, in fact, two things that can be improved here.  

## The Yield Statement

Remember our `logger` function and how we had to use a variable to store the return value of `f`?  
While it seems fairly trivial, it is a little annoying especially as your decorators get more complex.  
So how about we make that a little prettier.  Introducing, the `yield` statement.

    func logger(f: func()(int)) const func()(int) {
        func wrapper() int {
            log("called");

            yield f();

            log("returning");
        }

        return wrapper;
    } 

Much nicer.  Now we have no need for that pesky `res` variable.  For those of you wondering what yield does,
worry no longer, I have the answer.  All yield is is a deferred return statement.  In essence, yield evaluates
the expression on the right-hand side and then returns the result at the end of the function.  This can actually
be pretty handy as it can be used anywhere to shorten up your code.  But, for now, let's get back to decorators.

## Improving Decorators

As we said before, the normal Whirlwind syntax for decorators is a tad ugly and definitely be made faster to write
and far more expressive.  Well, the good news is, Whirlwind provides a special syntax just for decorators to make life
way easier.

Let's rewrite that example section using the decorator syntax.

    @[logger]
    func fn1() int {
        // blah blah

        return 3;
    }

    @[logger]
    func fn2() int {
        // blah blah blah

        if (condition) {
            // blah
            return 2;
        }

        return 3;
    }

Now we're talking.  That syntax does the exact same thing we saw before, but it is twice as idiomatic and expressive.  Better yet,
it makes adding decorators onto functions so much easier since all you need to do is copy and paste a little tag right above them
to apply the decorator.  But wait, there's more!

With this syntax, it is also possible to apply multiple decorators.  Say you wanted a function (or set thereof) to log when it is called
and returns, and you wanted it to restrict the return value to only set a group of values. No problem, just write to little decorators
and using our decorator syntax, apply them both.

    @[logger, bounded]
    func fn3() int {
        // -- snip --
    }

That's it.  The decorators are applied in the order they are specified so in this case, `logger` first then `bounded` and a decorated function
is created.

Oh, one more thing.  While we are on the topic of making decorators easy, you can also omit the `[]` if you are only adding one decorator to a function
as is the most common use case.

    @logger
    func fn1() int {
        // blah blah

        return 3;
    }

And with that, we are done looking at basic decorators, now it is time to look at some even more powerful decorators.

## Complex Decorators

A complex decorators is a decorator that in addition to accepting a function as its first argument, also accepts additional arguments that allow
it to vary its behavior.  We declare these kinds of decorators the same way we do normal decorators but with additional arguments.

    func logger2(f: func()(int), fnName: str) const func()(int) {
        func wrapper() int {
            log("called " + fnName);

            yield f();

            log("returning from " + fnName);
        }

        return wrapper;
    }

Now, our new logger we also include the function it being called and returned from in its log messages without having to create unique decorators or
worse.  Applying these kinds of decorators is also easy, simple add the additional arguments in parentheses after the function call.

    @logger2("fn1")
    func fn1() int {
        // blah blah

        return 3;
    }

    @logger2("fn2")
    func fn2() int {
        // blah blah blah

        if (condition) {
            // blah
            return 2;
        }

        return 3;
    }

As you can see, we are simply passing the `fnName` argument in parentheses in the decorator call and moving on.
Complex decorators and decorators in general are a super easy way to avoid code duplication in a concise, logical,
and reusable manner without an enormous cost the efficiency.
