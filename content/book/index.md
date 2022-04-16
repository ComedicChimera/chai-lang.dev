# Introduction

Chai is a general-purpose programming language designed for systems programming
and enterprise applications.  It is strongly-typed, garbage-collected, and
concurrent.  Its design is influenced by both functional and imperative
programming languages. 

Chai programming languages are comprised of *packages* which are organized into
*modules*.  This architecture allows for simple and intuitive dependency
management.  Furthermore, it enables programs to grow and evolve organically
while still remaining organized and maintainable.

Chai's syntax is designed to simple, expressive, consistent, and concise.  It
attempts to maintain a friendly appearance that allows for even complex code to
appear to relatively straight forward.

Chai places a heavy emphasis on expressions and extends them in a way some
programmers may find a bit unusual.  Most notably, Chai rephrases control flow
constructs as expressions including both conditional and iterative mechanisms.
This is not to say that they can't be used in the more conventional manners that
most programmers are used, but rather that they are capable of much more than
their purely imperative counterparts.

Finally, Chai programs are written to be complete and predictable.  Chai
recontextualizes errors as alternate paths of execution that must be handled no
differently than any typical variation in program flow.  Chai programs do not
"fail" at runtime: they simply proceed differently in response to an "invalid"
input.  This paradigm enables programs to fault tolerant and ensures that
resources are properly and safely managed during program execution.

## What is this book?

This book is meant to help the reader learn how to program in Chai.  It focuses
on practical applications of the language and tries to teach the reader how to
write useful programs as quickly as possible.  This book strives to give
relevant examples and discuss the actual construction of Chai programs rather
than simply giving a list of "rules" for how to write in Chai.  Furthermore, in
addition to teaching the language itself, this book also attempts to demonstrate
how to write "good" programs not just "correct" ones: useful algorithms and
important design principles with be illustrated if not fully discussed along
with topics in the language.

This book is not intended to be an introduction to programming or Computer
Science in general: it assumes some familiarity with basic programming concepts
including variables, loops, functions as well as comfort working with code a
computer.  Knowledge of the command-line as well as a modest understanding of
what a compiler is, what an executable is, etc with be assumed.  That being
said, a novice programmer should still be able to follow along althrough they
may need to use other resources to supplement their learning.

In addition to exploring the language itself, this book will also discuss
relevant elements of the standard library with are essential or typically useful
to building real-world programs: topics like standard I/O, file I/O, directory
manipulations, concurrency, networking, and more will be studied as a part of
this book.

While it is first and foremost a learning material, this book also, at least
initially, acts as a complete reference (and specification) of the Chai
programming language.  Everything that is in the language is discussed in some
part of this book.  As such, you may find that not every section is
relevant to you.  While you are encouraged to explore each topic covered in
this book, doing so is not required to learn or even become relatively effective
in one's use of the Chai programming language: you are encouraged to skip around
as necessary.

As a final note, you may find that Chai does things a little differently than
you are used to, even and especially if you have a background in more
"conventional" programming languages.  You may be tempted to attempt to write
Chai in a more "traditional" way and fail to utilize of some of the
more unique features of the language.  While this is certainly a way to construct
programs using Chai, it is not at all the best way.  If you write Chai in
the way that is intended to be written, taking full advantage of even if the
features and patterns you find "unusual", then you will that the language shines
in a rather unique way and may even come to enjoy programming in it a great deal.

In my experience, Chai is a wonderful, expressive, and versatile programming
language that can be naturally used for a wide variety of applications.  I hope
that this book can help you to use it well and perhaps even learn to love it as
much as I do.

## Book Structure

This book is organized into chapters with are comprised of sections.  The book
begins with a study of the most basic principles of the language which are
necessary to write useful programs.  From there, it explores each of the unique
topics in the language at length in an order that should appear fairly logical
and intuitive.  That being said, Chai is designed to be cohesive: many of the
different elements of the language fit together and build on each other in such
a way that it is occasionally necessary to divide them over multiple chapters so
that they can be fully explored.

The conventions used in the book are as follows:

- **bolded text** denotes a definition of an important term
- *italicized text* denotes the usage of the first use of a commonly understood,
but still important term worth highlighting.  Such terms may receive a more
formal or "canonical" definition later on.
- `inline code` denotes a "literal" piece of source code, often a piece of an
example program, being discussed

As you read, you will find that many sections include one or more *exercises*
sprinkled throughout their content.  Here is an example of an exercise:

<guide-exercise>
{
    "label": "Example",
    "content": "This will be where the problem statement will go.  Click the
    button with the question mark to get a hint.  Click the button with the
    lightning bolt to reveal the solution.",
    "hint": "This is the hint.",
    "solution": {
        "type": "text",
        "text": "This is the solution."
    }
}
</guide-exercise>

These exercises are designed to help you practice what you are learning and to
encourage you to play around with the language and get a feel for how to use it.
I strongly recommend that you try and complete as many of the exercises as you
can because it will help cement what you are learning and make sure you don't
move to quickly over things you might not actually understand.  

Now, without further adieu, it is time to begin your journey into learning the
Chai programming language.  Let us begin...