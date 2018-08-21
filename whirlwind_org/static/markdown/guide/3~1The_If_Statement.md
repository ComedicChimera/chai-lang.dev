## The If Statement

The **if statement** is a conditional contruct that allows
you to regulate which code runs on the basis of a boolean value.

### The If Syntax

The if statement has a fairly simple syntax.  It consists of 3 parts:
the **if**, the **condition**, and the **body**.  The first part, the if,
is simply the keyword `if`.  The condition is a boolean expression wrapped
in parentheses following directly after the `if.  Finally the body is
a block or single statement that will only be executed if the condition
is true.  The code in **Listing 3.1** demonstrates this.

#### Listing 3.1 - The If Statement

    if (5 > 3) {
        // do something
    }

In this case, the body will always execute since the condition is always true.
But, if the condition were not true, the body would be ignored and the program
would simply keep going.

    include random;
    include stdio;

    $randVal = random.Randint(0, 10);

    // will only run if randVal is even
    if (randVal % 2 == 0)
        stdio.Println("Random value is even.");

Using the if statement, you can determine how your program runs very quickly.
But, sometimes you want a bit more power, that is where the secondary clauses
come in.

### The Else Clause

If statements are often followed by what is call the **else clause**. Consider the
example below.

    if (randVal % 2 == 0)
        Println("Randval is even.");
    if (randVal % 2 != 0)
        Println("Randval is not even.");

That code seems a little a bit redundant right? Effectively, you are checking
the same condition two times! Well, you can use the else clause the shorten this
statement significantly.

The else clause is basically the opposite of an if statement. It will only run if
the previous if clause was **false**. So you can use the do something if a condition
was not true.  The code in **Listing 3.2** demonstrates the same block as above, but
using an else clause instead.

#### Listing 3.2 - Else in Action

    if (randVal % 2 == 0)
        Println("Randval is even.");
    else
        Println("Randval is not even.");

See how much shorter that was? As you can see, the else clause is extremely useful for shortening up
your code.  But, it extremely important to note that you can only use the else statement when it is
paired with an if statement.  So the following is invalid on it's own.

    else
        Println("Randval is not even.");

As stated before, the else clause is optional so you don't have to use it on every if statement.

### The Elif Clause

There is another type of if clause called the **elif clause**. This clause
is like a merger of if and else. It will only run if the if statement
preceding is false *and* its condition is true.  Consider the following
example.

    if (randVal % 2 == 0)
        Println("Randval is even.");
    else {
        if (randVal % 3 == 0)
            Println("Randval is divible by 3.");
        else
            Println("Randval is not even or divisible by 3.");
    }

Seems a little verbose right? Well, using an elif clause you can tighten that up
to look like the following.

    if (randVal % 2 == 0)
        Println("Randval is even.");
    elif (randVal % 3 == 0)
        Println("Randval is divible by 3.");
    else
        Println("Randval is not even or divisible by 3.");

The elif clause doesn't have to always be paired with an else.
It is just as possible to omit the else entirely if you don't
need it.  **Listing 3.3** showcases this nicely.

#### Listing 3.3 - No Else Required

    if (randVal % 2 == 0)
        Println("Randval is even.");
    elif (randVal % 3 == 0)
        Println("Randval is divible by 3.");

It is important to note that just like else, elif must be paired with an if statement.
Furthermore, elif must also come before the else.

    if (randVal % 2 == 0)
        Println("Randval is even.");
    else
        Println("Randval is not even or divisible by 3.");
    // invalid, will cause an error
    elif (randVal % 3 == 0)
        Println("Randval is divible by 3.");

You can also stack elif clauses if you want to check multiple cases at
once.

    if (randVal % 2 == 0)
        Println("Randval is even.");
    elif (randVal % 3 == 0)
        Println("Randval is divible by 3.");
    elif (randVal % 5 == 0)
        Println("Ranval is divisible by 5.");

You can stack as many elifs as you want and they will all follow similar logic.
It is important to note that each elif in the chain will only execute if the if
statement *and all previous elifs* have failed.