---
title: "Comprehensions"
weight: 14
---

## Array and List Comprehensions

A **comprehension** is a way to generate a new collection from an iterable.  Comprehensions
can be used to generate any one of the three builtin collections: arrays, lists, and dictionaries.

Let's break down the syntax for an array comprehension so we can understand how comprehensions work.

    {x for x in 1..10}

Firstly, notice that we enclose the comprehensions in the grouping symbols of the collection we want
to create: we are using braces to denote that we want to create an array.  The comprehension itself
is made up of two parts: the `x` and the `for x in 1..10`.  The `x` is the value that is put into
the array as it is being built; the "inline for-loop" is the iterable itself that is used to generate
the various `x` values.

TODO: unwrap into for-loop

## Filter Conditions

## Dictionary Comprehensions