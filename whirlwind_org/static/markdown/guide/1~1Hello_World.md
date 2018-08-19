## Hello World
Hello World is first program written when learning a new language.
It provides an introduction to language's syntax and some basic concepts
of both programming and the language itself.

The **Hello, World** program's purpose is simple.  It's job is print
`Hello, World` to the console.  The **Console** is just a window
to read and write information.  Every program starts with a console.

### Creating a Source File
To begin writing our Hello World program, we first need to create file.
Just create a new file called `hello.wrl` to store our code in. Notice
that `.wrl` is the extension used for Whirlwind program files.

### Writing and Running our Program
Open the file you just created and type the code in
**Listing 1.1** into it.

#### Listing 1.1 - An Example of Hello World Program

    use include stdio;

    func main() {
        Println("Hello, world!");
    }

Next, open up a command window and type the following command into it.

`whirl build hello.wrl`

Once that command has finished running, you should see an
executable file called `hello` right next to the `hello.wrl` file.

Go ahead and run that executable and you should see the following:

<div class="console">
    Hello, world!
</div>

Congratulations!  You just wrote your first Whirlwind program.  Welcome
to Whirlwind programming community!

If you're confused by the Hello World program, don't worry! We will spend
the rest of this chapter breaking down how this program works and
what we can learn from it.



