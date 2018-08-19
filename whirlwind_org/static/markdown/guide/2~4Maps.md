## Maps

A **map** is a data structure used to organize information in a key/
value system.  In each map instance, you have one key type and value type.
Each key corresponds to one value.  All the keys in a map must be
unique, but not all of the values have to be unique.

### Map Literals

You can declare a simple map using the braces, similar to an array.  However,
you separate each key and value with the `:` symbol.  The code
in **Listing 2.7** shows an example.

#### Listing 2.7 - Declaring a Simple Map

    $myMap = {
        'a': 1,
        'b': 2
    }

The map above, all of the keys are chars and the values are integers.
Notice that the key-value pairs are separated by commas.  Also notice
that each key is unique.

The map in **Listing 2.7** contains two keys: `'a'` and `'b'`. Key 'a' maps
to 1 and key 'b' maps to 2, as designated by the `:`.

Maps also have a specific type designator.  It is formatted like so:
`map[key_type, value_type]`. So the map type designator for the map
above is `map[char, int]`.

### Modifying Maps

You can access the keys in a map using the subscript operator with the key
as your value.  For example, if you wanted to get the value of 'b' in the
map in **Listing 2.7**, you could so like this: `myMap['b']`.

You can use this to modify maps as well and add keys to them.  For example,
if we wanted to add the key `'c'` with a value of 3 into `myMap`, we would
do the following:

    myMap['c'] = 3; // 'c' is now 3

You can change the value of map keys the same way.

    myMap['a'] = 0; // 'a' is now 0

    myMap['b']--;
    myMap['c']--;

As you would expect, all of the valid assignment operators can be applied
this way.

However, it is important to note that if you try to get a value from a map
that doesn't exists or modify using a non-standard assignment method,
you will get a **KeyException** at runtime that can crash your program if
not properly handled.

### Map Methods

Just like lists, maps also have a couple essential methods for working with
them.  The first 2 are fairly simple: `keys()` and `values()`. As you might
expect, each returns as list of the keys and values respectively.

    $keys = myMap.keys(); // ['a', 'b', 'c']

    $values = myMap.values(); // [0, 1, 2]

You can remove a key from a map by using the the `remove()` method.

    myMap.remove('c'); // remove the key 'c' and its corresponding value from the map

You can insert a key a value pair using the the `insert()` method.

    myMap.insert('c', 2); // add 'c' back into the map with a value of 2

Finally, you can use the `update()` method to change the value of a key if it
exists.

    myMap.update('d', 3); // nothing will happen since d is not in the map

    myMap.update('c', 4); // the value of 'c' is now 4

You can follow this method up with an `elseInsert()` call to add a default
value if none exists.

    myMap.update('d', 15).elseInsert(4); // if d doesn't exist set, it's value to 4
    myMap.update('c', 3).elseInsert(15); // c does exist so elseInsert does nothing

The `elseInsert()` method can also be used on just a raw map without having
to call update. It has the same behavior described above.

    myMap.elseInsert('e', 5); // e does not exist and therefore is inserted with the value 5

    myMap.elseInsert('a', 0); // a exists so nothing happens

Notice that, in this case, a key is required, unlike in the examples
with update.

### Hashability

In addition to all of the previous conditions,
all of the keys must be **hashable**, which just means
they are able to hashed for use in a map.
You can check whether or not something is hashable, by using the
builtin `hashable` function.

#### Listing 2.8 - Testing for Hashability

    hashable(34); // true
    hashable(3.4); // true

    hashable([1, 2, 3]); // false
    hashable({'a': 1, 'b': 2}); // false

All of the builtin simple types are hashable, however none of the collections
are.