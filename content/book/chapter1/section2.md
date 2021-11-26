# Variables and Numbers

We will now consider several smaller programs, focused on arithmetic, to
familiarize ourselves with some of the mundane but essential features of Chai.

To begin, let us write a program to display the
[Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number).  For those
unfamiliar, the Fibonacci sequence is a sequence wherein the next number in the
sequence is defined as the sum of the previous two numbers, beginning with the
entries 1 and 1.  For example, the first few numbers of the Fibonacci sequence
appear as follows:

```text
1, 1, 2, 3, 5, 8, 13, 21, 34, ...
```

In our program, we will print the Fibonacci numbers less than 1000.  In a
similar fashion to our *Hello World* program, let's take a look at the whole
program first and then discuss what makes it work.

    import println from io.std

    def main()
        let a = 1, b = 1

        while a < 1000
            println(a)
            a, b = b, a + b
        end
    end

Our program begins with a taste of the familiar: an import statement just like
the one in our *Hello World* program.  This makes sense given that we want to
print to the console.  We also have a `main` function; however, this main
function is constructed differently.  




