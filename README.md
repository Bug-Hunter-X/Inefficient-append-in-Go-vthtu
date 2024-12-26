# Inefficient Append in Go

This example demonstrates an inefficient way to use Go's `append` function.  The `slowAppend` function repeatedly creates a new slice and appends to it, leading to performance problems when used within a loop or when appending numerous elements.  This can lead to unnecessary memory allocations and slowdowns, especially for larger datasets.

The solution shows how to improve efficiency by appending directly to the original slice in a loop.

## How to reproduce the bug

1. Clone this repository.
2. Run `go run bug.go` to see the inefficient append.

## How to fix the bug

1. Refer to `bugSolution.go` for an improved implementation.
2. The solution appends directly to the original slice, reducing allocation overhead.