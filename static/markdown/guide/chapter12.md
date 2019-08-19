# Type Classes

Type classes are the most complex, but powerful tool available to you in
Whirlwind.  They are the most variable and advanced type provided by
the type system and offer numerous different options on how to approach
them.

## The Basics

A type class is at its core no more than a name.  It allows you to define
a type without any properties, methods, or really anything.  The syntax
for declaring a type class in this form (called an **empty type class**)
looks like this:

    type MyType;

This creates a new type class by the name `MyType`.  It currently has no values and
no instance but it can still be used.  A type class in this form can still have methods
bound to it and can be used like a type, it simple contains no values.

The next form a type class can take is called a **type alias**.  This is a type class that creates a
new name for an existing type.  For example, say you have situation where you wanted to store an
integer, but you wanted it to be treated like its own specific type.  You could easily accomplish this
by using a type alias.

    type Alias int;

Now, the type `Alias` can hold the value of an integer while still being its own distinct type.  The implications
of an alias can best summarized by the statement "an integer is an `Alias`, but an `Alias` is not an integer".  This means
that you can assign an integer value to `Alias`, but you cannot assign or pass an `Alias` as an integer type.

## Enumerated Type Classes

A type class can also take the form of an **enumerated type class** wherein the type class defines the literal
values it can take.  Enumerated type classes begin the same as a normal type class, but their values are prefixed
by the `|` operator.

    type Color
        | Red
        | Blue
        | Green
        ;

The `Color` type is now an enumerated type class whose four primary values are `Red`, `Green`, and `Blue`.  A usage
of the Color class might look like the following.

    func colorToString(color: Color) str
        => color case {
            Red => "Red",
            Blue => "Blue",
            _ => "Green"
        };

    func main() {
        let myColor = Color::Red;

        let strColor = colorToString(myColor);

        strColor = colorToString(Green);
    }

As you can see in the example below, enumerated type classes behave a lot like a primitive type which can accept and translate
between values.  The biggest difference is those values are specified.  

> If your enumerated type class only has one value, you can omit the preceding `|`.

There is another important thing to point out, and it is the use of the `Color` prefix.  You will notice that within the
above code the `Color` prefix is only used once: in the variable declaration.  The prefix must be used here because
the compiler cannot infer based on context what the type is supposed to be.
The variable has no type extension and therefore, no assumable type.  It is entirely possible for other `Red` variables or types
to exist simultaneously and so the compiler requires context to determine which one to use in the given scenario.  If it does
not have context, you will need to specify what you want the context to be.

    // this is ok: the compiler knows the type of the expression
    let color: Color = Green;

    // this is not ok: the compiler does not know the type of the expression
    let color2 = Red;

This pattern is called **context-based inferencing** and is applicable to another type we will see later on.

## Value Type Classes

It is also possible for an enumerated type class to have its enumerated types hold values.  This is called a
**value type class**.  The syntax for declaring one of these is very simple.  You declare a normal enumerated
type class followed by the type you want to hold in parentheses.

    type Number
        | Int(int)
        | Decimal(double)
        ;

This type class contains two value-holders `Int` and `Decimal` each of which hold the corresponding type.

> It is also possible for a value type class to have value members that contain themselves multiple values.  You simply delimit
> the values with commas inside the parentheses.

The usage of this type class is similar to the usage of the basic enumerated type class above, the only difference
being you specify the value you want each instance to hold.

    let num = Number::Int(3);

    func main() {
        let t: int = from (num as Int)

        t++;
    }

Value type classes also introduce two new operators: `from` and `as`.  The former of these operators is used elsewhere
in the language but in this context, unpacks the value of a known type class value.  In the specific case above, it
extracts the integer value from the `Int` member of the type class.  The `as` operator converts between the different forms
of a type class so that value extraction works properly.  This operator is necessary as otherwise the compiler would not
know what value to extract when calling `from` (from only works on type class members, not the type class itself).  Additionally,
as allows you to store the specific value of a given type class to be used later, while still providing you with the ability to
unpack it.

    func takeNumber(num: Number) {
        // -- snip --
    }

    func main() {
        let num: Number = Int(3); // context based inferencing applies here too

        let numInt = num as Int;

        takeNumber(numInt); // still works as type class

        let numSquared = from numInt ~^ 2; // able to use from
    }

Note above that `from` is a high-precedence operator so you don't need to wrap from calls in parentheses.

## Value Restrictors

Value type classes also have an additional power in the form of a value restrictor (or just restrictor for short).  Adding
a restrictor allows you to introduce additional logic defining what values may be contained within a value type class.
Using this logic, we could expand our `Number` definition to include a positive type.

    type Number
        | Positive(v: int) when v > 0
        | Int(int)
        | Decimal(double)
        ;

As you can see, to add a value restrictor, you need only add on an additional when expression and a give the type a local
name within the type class.

## Multi-Variant Type Classes

This is the final and most powerful form of type class.  It is a type class in which all the elements are used together.
A great example of such a type class would be an optional type class.

    type OptionalFloat
        | Some(float)
        | None
        ;

Here we are combining the traits of an enumerated type class and a value type class.  This type of combination can be done
easily and is often used within Whirlwind.  We call type classes which combine multiple types of values **multi-variant
type classes**. A usage of the above class could look like the following:

    include { Println } from std::io;
    
    // sqrt may fail if the integer is negative
    func sqrt(x: int) OptionalFloat {
        // -- snip --
    } 

    func main() {
        let res = sqrt(3);

        if (res is Some) {
            Println(from (res as Some));
        } else {
            Println("Failed to find result.");
        }
    }

The code above introduces the final operator we will cover in this chapter that is most pertinent to multi-variant and value
type classes: the `is` operator.  This operator allows you to determine to true type of any value.  In this case, we used it
to test if `res` is a `Some` or not.  You can also use `is` on any other type if you want to know its actual type (such as
with classifying interfaces).

Whew!  You did it.  That was almost certainly one of the most complex chapters in this guide.  As you might imagine, it isn't
completely comprehensive to type classes, but it is pretty close.  Just remember that the best way to truly understand these
constructs is to actually use them.  We have recently covered a lot of heavy stuff from interfaces to type classes. The good
news is, you get a break for the next four chapters which are a lot lighter and easier to understand.  Just remember to keep
going at it because eventually it will all make sense, especially once you can see the big picture.
