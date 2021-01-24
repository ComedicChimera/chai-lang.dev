---
title: 'Control Flow'
weight: 4
---

## If, Else, and Elif

In Whirlwind, we can branch based on a condition using the `if` construct.
This construct will run the code in its block if the condition in its header
is true.

    let x = random_number()

    if x % 2 == 0 do
        println("`x` is even.")

The code in the block will only run if `x` is even.  Notice that if block
also begins with `do` and has its own indentation level -- much
like functions.

We can also add an `else` condition that will execute if the condition
is false.

    if x % 2 == 0 do
        println("`x` is even.")
    else
        println("`x` is odd.")

Notice that we do not need to put a `do` at the end of the else to begin the block.
This is common with many blocks that are begun by single keywords.  

{{< alert theme="info" >}}You can place a `do` at the end if you want.{{< /alert >}}

Finally, you can use the `elif` keyword to introduce a series of secondary conditions.
These conditions will be checked in order after the primary if, and the block of
the first one that is true will be executed, breaking the chain.

    let color = random_color()

    if color == "red" do
        println("Blood for the blood god!")
    elif color == "green" do
        println("A fine choice.")
    elif color == "blue" do
        println("Ocean lover, I see...")
    else
        println("Not RGB? Quirky.")

You do not need to include an `else` at the end of an `elif` chain.  However, the
statement must be constructed beginning with an `if`, followed by any `elif` blocks,
and (optionally) ending with an `else`.

## While Loops

A **while loop** is a block that executes its contents repeatedly until its condition
is no longer true.  

    let counter = 0

    while counter < 10 do
        println(counter)

        // `++` adds one to the variable counter
        counter++

The code above will print out the numbers zero through nine -- it repeats as long as
`counter` is less than `10`.  

You can exit a while loop before its condition is satisfied using the `break` keyword.

    let n: int

    while true do
        n += n

        if n % 67 == 0 do
            println("Found it:", n)
            break

Whenever the conditon in the `if` block is true, the `n` will be printed out and loop
will exit.  Notice that the condition of the loop is simply `true`.  This creates an
infinite loop that can only be exited out of using `break` (or `return`).

You can also continue onto the next iteration of a loop using the `continue` keyword.

    let word = ""
    let n = 0

    // the `.len()` method gets the length of a type such as a string
    // we will learn more about methods in a later chapter
    while word.len() < 10 do
        if n % 2 == 0 do
            word += 'a'
            n += 3
            continue

        word += 'b'
        n++

{{< alert >}}Many of these examples are trivial for sake of demonstration.{{< /alert >}}

Finally, all loops can be ended with a `nobreak` clause.  This clause will be run if the
loop is unbroken during its run -- it can be used to efficiently avoid the use of flag
variables.

    let n = random_int()

    let counter = 0
    while counter < 101 do
        if n == counter do
            println("Found a match!")
            break

        counter++
    nobreak
        println("Did not find a match for:", n)

Instead of using a flag variable, we were able to tell simply from the control flow what
happened inside the loop.  `nobreak` is an often hidden feature in certain languages
(eg. Python's else clause); however is can be quite powerful and allow for super elegant
control flow.

## Basic For Loops

**For loops** are an immensely powerful tool of control flow that we don't yet have the "meat
and potatoes" knowledge required to fully explore yet.  They apply primarily to collections
and iterables -- two concepts we will touch on in a later chapter.  However, it would be
wrong not to at least include them in this chapter.

For loops can be used to iterate over sequences of values.  One such sequence is a **range**.
A range is exactly what it says on tin -- it is a range of values, specifically integers.
Ranges are denoted by two integer values with `..` in between them.  However, unlike what
you may be used to in other languages, ranges are inclusive on both ends.  This means that
`1..10` is the numbers 1 through and including 10 not 1 through 9. 

{{< alert theme="info" >}}The primary reason for this has to do with the use cases of ranges
in Whirlwind.  For a number of reasons, it is surprisingly uncommon to generate a range based
on the bounds of a collection (ie. to create a sequence of indices) which is the primary reason
why ranges in some many languages go inclusive to exclusive.  Outside of that one big use case,
inclusive to exclusive ranges often don't make much sense and require more mental labor by the
programmer -- making them inclusive to inclusive generally makes life easier.{{< /alert >}}

To create a for loop over a range, we use the keyword `for` followed by an **iterator variable**.
This variable will store the current value in the sequence.  This is then followed by the `in`
keyword and the sequence being consumed which will be range in this case. 

So, if we wanted to use a for loop to print out the numbers 1 through 10, we could write the following:

    for n in 1..10 do
        println(n)

That is certainly a lot more concise than our while loop example above.  The loop's body will
be called once for each of the 10 numbers (ie. until the sequence is consumed).

All of the same constructs/keywords (`nobreak`, `break`, etc.) apply to for loops as well.
Thus, we can also rewrite our code to find a "match" using a for loop like so:

    let n = random_int()

    for i in 0..100 do
        if n == i do
            println("Found a match!")
            break
    nobreak
        println("Did not find a match for:", n)

Once again, a for loop has made our code much easier to read and far more concise.  As a general
rule, you want to prefer a for loop for anything involving iteration: that is going through some
arbitrary sequence of values.

Unfortunately, that is where we must leave the for loop for now; however, keep the basics in the
back of your mind as much further down the road we will be returning to explore the full power
of this construct -- there is much left to explore.