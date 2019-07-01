# Hello Whirlwind

In this chapter, we will learn how to write a **hello world** program in Whirlwind.
This program is the first step in learning any program language and hopefully will
introduce you to some of Whirlwind's syntax and layout.

For this tutorial, I am going to assume you already have Whirlwind installed.  If you
don't, you can visit our [install](/install) page to get your version and a handy installation guide.

## Creating the Program File

Since Whirlwind does not have a REPL, we will need a file to work with.  Just create a file called
`main.wrl` in the directory of choosing.  Make sure that the it has the `.wrl` extension as this
designates it as a program file.  Many features of Whirlwind will not work if it is not named that way.

Then simply open this file in the editor of your choosing and you are ready to go.  It is important to
note that Whirlwind files are just text files with a special extension.  There is nothing magical about
them, just the designation that they contain code.

## Writing the Main Function

Whirlwind, like many other compiled languages, requires the implementation of a **main function**.  This
is the starting point for your program and where we will start writing our code.  Declaring the main
function is easy.  It looks like this:

    func main() {

    }

Thats it.  No fuss, no mess, just a simple declaration.  There are however a couple things to notice about
this declaration.  Firstly, notice the use of the `func` keyword.  This will be used everywhere you want to
declare a function and in the same form.  Next is the name, `main`.  In Whirlwind, the function's name always
follows the `func` keyword.  The parentheses the come next are where arguments are designated.  We will get to
those as well as the noticably missing return type in our section on functions, later.  Finally, the braces
designate a code block.  Inside them is where we will place our *hello world* instruction and any other code
we want to run from the main function.

## Importing the "stdio" Package

Before we can write our print instruction however, we need to import the **stdio** (standard I/O) package.  This package contains
several important definitions used when reading from and writing to the console.  Packages are an essential
piece of Whirlwind and a tool you will be using quite frequently.  To import a package, we need to use the
**include statement**.  This is a statement that goes at the top of your file and describes what package to import
and what definitions you want from it.

> We are actually including a sub-package of the io package which handles all forms of I/O in Whirlwind
> called the `std` sub-package which is where the standard I/O definition are stored.
> This sub-package is often referred to as the stdio package even though that is not its name.

A simple include statement for the stdio package will look like this:

    include io::std;

This brings the `std` sub-package of the `io` package into scope so that we can use all of its definitions.  
This will be all we use for now.  However, later on we will learn how to use the full power of the include statement to export
and import specific definitions.

## Calling the Println Function

The `Println` function from the stdio package is what we be using to print to the console.  It writes whatever you tell it
to to the console, followed by a new line.  To call the Println function, we will use the following syntax.

    std::Println();

Notice that the work Println is followed by 2 parentheses, one open and one closed.  These parentheses are where we will put our
**arguments**. The `std` prefix before `Println` is simply an instruction telling the compiler what package (or sub-package in this case) to find the function in.  
Finally, the semicolon at the end marks where our instruction ends.  These tell the compiler where one instruction ends and another begins.

Now we need to tell Println what to print.  In order to do this we have to pass it an argument to print.  In this case, we will be
passing in a **string**.  A string is just series of characters that all go together.  In this case, the contents of our string will be
`Hello, World!`. To designate a string we simply take the content we want to be stored by the string and wrap it in double quotes.  We then
need to pass our string value to Println, so we will be put it inside the parentheses we designated for arguments earlier.

    std::Println("Hello, World!");

That's it.  Our finished print instruction.  You can see where and how we have inserted our string literal and how we have called our function.
Get familiar with this syntactic structure as it will likely be one of your most used in Whirlwind.

## Putting It All Together

Now that we have each part of our program, we need assemble them together and build it. Putting it together is simple.  We first need to add
in our include statement and then put our print statement inside the braces of the main function definition.  After doing that, our program finished program
is only 5 lines long and very expressive.

    include io::std;

    func main() {
        std::Println("Hello, World!");
    }

Next step is to build it.  Go ahead and open a command window and navigate to the directory your file is stored in.  Once there, enter
the following command into your prompt, replacing my file name with yours.

<div class="command-window">
    whirl build hello.wrl
</div>

Whirlwind will build that file and output an executable with the same name as your file.  Since I am on Windows, mine is called `hello.exe`.
Now, you can go ahead and run your executable from command prompt.  And you should see the following printed out.

<div class="command-window">
    Hello, World!
</div>

And that's it. You have just written your first Whirlwind program.  Welcome to the family.  The basics you learned here will be useful for the
rest of your Whirlwind programming career.
