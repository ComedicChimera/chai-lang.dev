# Introduction

Chai is a cross-platform, general-purpose, compiled programming language that
takes inspiration from numerous languages and paradigms and attempts to fit them
together in a cohesive way.  It is strongly and statically typed with an
emphasis on generics and generic code.  It uses a runtime garbage collector to
avoid the hassle of manual memory management.  It places a heavy emphasis on
expression-oriented programming, first-class functions, mathematical constructs,
and concurrency.

## Principles

Chai's first and most important design principle is *pragmatism*: the idea being
that Chai is a language that can be used to write productive, real-world code
quickly.  Chai wants to minimize the mental distance between having an idea for
a system or component of an application and turning it into code.  It is a
language where you don't have to negotiate with compiler or navigate some
opinionated minefield of language rules, but rather you can simply get down to
business writing code.

The second principle is *concision*: the idea being to minimize verbosity while
maximizing readability.  Languages such as Java can be very verbose and flowery
with lots of modifiers all written out which while it does take more time to
type and adds extra lines of code also makes programs readable.  Conversely,
languages like Haskell provide an assortment of operators and syntactic sugar to
make your programs shorter while completely obfuscating the meaning of the code
to anyone not well versed in the language's "arcana".  Chai attempts to strike a
balance between these two: allowing you to write code that is easy to read and
interpret without wasting your time writing out endless boilerplate and
frivolous preludes.

The final guiding principle of Chai is *performance*.  While this principle is
self-explanatory, its implementation is not.  Many programmers, including myself
to some degree, believe that it is impossible for code that is easy to write to
run fast.  For example, Python is a beautiful language that takes almost no time
to get relatively productive in and is mostly very easy to write, but it runs
insultingly slow even with the numerous C-based workarounds and hacks.  The
grizzled developer will tell you that you have to get down to nuts and bolts in
a language like C or C++ and deal with the ugliness of manual memory management
and all the other fanfare of the systems world to produce a truly fast program.
I am not here to disagree with that assessment: there is an unavoidable truth in
it.  The more decisions the language makes for the programmer and the more
abstraction that exists between the programmer and their computer, the more
overhead their applications have to deal with.  However, this doesn't doesn't
mean that reasonably "fast" code can't be beautiful, easy, or even, with the
right guide, the most obvious option.  This is the fundamental idea behind
performance in Chai: the language makes some concessions for the sake of the
other principles such as a garbage collector but attempts to present itself in a
way that encourages the writing of more efficient code.  For example, Chai
doesn't use classes or objects in the same way as a traditional Java-like
language but rather uses an alternative paradigm based on generics such that
things like virtual method tables or boxing and unboxing are just simply not
needed.  As another simpler example, Chai puts a heavy emphasis on "powerful
expressions" (such as control flow expressions) which often end up being more
efficient than their more procedural counterparts.

That is a lot of design philosophy to open even a fairly dry formal text, but
I think it helps to inform some of the possibly odd design decisions made
and expressed throughout this specification.

## Author's Note

On a personal note, I designed Chai for me with my silly neuroses and perhaps
ill-earned preferences in mind, the idea being that others might share some of
my opinions and grievances, and they might enjoy Chai as much as I do. I hope
you find Chai to be a reliable, charming little artifice that you can always
return to.  I wish you well on your coding adventures to come.
