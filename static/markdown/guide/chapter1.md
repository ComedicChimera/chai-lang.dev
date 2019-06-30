# Introduction to Whirlwind

Whirlwind is a statically-typed, imperative, compiled language
built to bridge the gap between hardcore systems languages
such as C++ and extremely dynamic, data science languages such as Python.

This guide will:

- Teach you the basics of programming in Whirlwind.
- Familiarize you with your compiler and Blend.
- Introduce you to important packages in the Standard Library.
- Demonstrate best practices and good style in Whirlwind.

While this language certainly isn't perfect, it can be a powerful
and useful tool to help you work as efficiently and logically
as possible.  If you are intimidated by Whirlwind, don't worry:
the more you work with it, the more everything will make
sense.  The rewards of mastery are well worth the pain.

## Some Interesting Quirks

There are couple of things that make Whirlwind stand out from other
languages.  The first and most obvious example is Whirlwind's addition
of **agents**.  These are intended as a powerful tool to enable you
to have different processing kernels that logically and organically
communicate with each other, akin to neurons in a brain.  Another important
distinction are the functional aspects of Whirlwind.  While it is primarily an
imperative language as stated before, it supports a number of functional programming
techniques including closures, comprehensions, partial functions (the good kind), and
an inline case statement.  We also provide an entire package dedicated to giving you a powerful, yet
simple functional toolkit.

## Our Approach to Object Orientation

Whirlwind does support objects, however in a way slightly different than what you might expect.
Firstly, objects **cannot** inherit from each other.  However, objects can implement interfaces
both explicitly and implicitly (a topic discussed in detail later). 
Objects also support constructor overloading and private constructors.  This is in contrast to
the rest of Whirlwind which does not support function and method overloading.

## Finally...

Now that you have a good idea just what your getting into, I wish you the best of luck on
your adventures with Whirlwind and hope that you will stick around long enough to really
appreciate it.  I think you will really enjoy it.
