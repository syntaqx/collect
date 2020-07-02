# collect

[![Test](https://github.com/syntaqx/collect/workflows/Test/badge.svg)](https://github.com/syntaqx/collect/actions?query=workflow%3ATest)
[![golangci-lint](https://github.com/syntaqx/collect/workflows/golangci-lint/badge.svg)](https://github.com/syntaqx/collect/actions?query=workflow%3Agolangci-lint)
[![codecov](https://codecov.io/gh/syntaqx/collect/branch/master/graph/badge.svg)](https://codecov.io/gh/syntaqx/collect)
[![Go Report Card](https://goreportcard.com/badge/github.com/syntaqx/protokit)](https://goreportcard.com/report/github.com/syntaqx/protokit)
[![GoDoc](https://godoc.org/github.com/syntaqx/collect?status.svg)](https://godoc.org/github.com/syntaqx/collect)

> `collect` provides functions to perform operations on collections of data.

We often need our programs to perform operations on collections of data, like
selecting all items that satisfy a given predicate or mapping all items to a new
collection with a custom function.

In some languages it’s idiomatic to use generic data structures and algorithms.
Go does not support generics; in Go it’s common to provide collection functions
if and when they are specifically needed for your program and data types.

This package implements some of the most common collection functions that I
like to be able to use in a pinch, and eventually I'll expand on it.

## License

`collect` is open source software released under the [MIT license][MIT].

[MIT]: https://opensource.org/licenses/MIT
