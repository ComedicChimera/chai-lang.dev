---
title: "Struct Types"
weight: 8
---

## Defining a Struct Type

A **struct type** is a data type that stores multiple typed values in named fields.
It is a **defined type** meaning that it is created by the user using a type definition.
All structs have a specific name associated with them.  

To declare a struct, we begin with the `type` keyword.  This keyword begins any type
definition.  We follow it with the name of the type being defined.  For this example,
we are going to define a struct called `User` which is going to represent someone who
is using our application.

Next, we begin by opening a pair of curly braces (`{}`) and describing our fields.
Each field has a name and a type.  Fields are separated by newlines, but fields of
the same type can be grouped together under the same type label much like arguments.

    type User {
        name, email: string
        age: uint
    }

The `User` struct has three fields: `name`, `email`, and `age`.  The first two are strings
and the latter is a `uint`.

We can provide default initializations for the fields in the struct in much the same way
we can for function arguments.

    type User {
        name, email: string
        age: uint
        admin: bool = false
    }

The type label for a struct is its name.

    func promote_to_admin(user: User) do
        ...

## Instantiating Structs

The definition of a struct type describes the general structure of the struct in much the
same way that the type label `(int, string)` defines the general structure of that tuple.
Much like tuples, we need to be able to create instances of a struct types that we can use
and manipulate.

Creating an instance of a struct (or other defined type) is referred to as **instantiation**.
To instantiate a struct, we can simple use its name followed by a pair of curly braces.

    let u1 = User{}

The instance `u1` has fields with default values only -- all the fields are blank.  If we want
to specify some specific data about our user, we can provide **field initializers** in the
curly braces of a our struct instantiation.  The initializers can be given in any order and
look very similar to the syntax for named arguments.

    let u2 = User{name="Brenda", age=27}

Our new user now has a specified name and age.  All the fields we didn't specify will still
have their default values.

Finally, it is fairly common to want to initialize one struct based on the contents of the
other.  For example, say we wanted to create a copy of `u2` that was an admin user instead
of a normal one.  We could simply refill out the fields has we did above, but for structs
with many fields, this can not only be time consuming but also error prone.  Instead,
Whirlwind provides an additional kind of initialization called **spread initialization**.
All this does is initialize our struct with values contained struct being "spread" and
then allows us to override the struct's values as we desire.

We use a special `...` syntax to perform spread initialization as shown below:

    let u3 = User{...u2, admin=true}

`u3` has the same fields as `u2` but `admin` is now marked `true`.

## Accessing Fields

Once your struct instance is created, you will want to access and manipulate its fields.  This
is done using the `.` operator followed by the name of the field you want to access.

    println(u3.name) // prints `Brenda`

We can change the value of that field by assigned to it like so:

    u3.name = "Emma"

    println(u3.name) // prints `Emma`

This operator also respects pattern matching on assignment:

    let name: string, age: uint

    // -- snip --

    u3.name, u3.age = name, age

Note that some fields can be marked as constants (by placing a `const` before the field name).
These fields cannot be assigned to after they are initialized.

    type MyStruct {
        const field: uint
    }

    func main do
        let s = MyStruct{field=3}

        println(s.field) // prints `3`

        s.field += 2 // ERROR





