# collect

[![Build Status](https://travis-ci.org/syntaqx/collect.svg?branch=master)](https://travis-ci.org/syntaqx/collect)

> `collect` provides functions to perform operations on collections of data.

We often need our programs to perform operations on collections of data, like
selecting all items that satisfy a given predicate or mapping all items to a new
collection with a custom function.

In some languages it’s idiomatic to use generic data structures and algorithms.
Go does not support generics; in Go it’s common to provide collection functions
if and when they are specifically needed for your program and data types.

This package implements some of the most common collection functions that I
like to be able to use in a pinch, and eventually I'll expand on it.
