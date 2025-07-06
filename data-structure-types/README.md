# 📚 Go Data Structure Types

Welcome to the **Data Structure Types** section of the [Go-Exercises](../README.md) repository!  
This directory contains hands-on Go examples demonstrating the core data structures every Go developer should know.

---

## 📂 Contents

- [Arrays](#arrays)
- [Slices](#slices)
- [Maps](#maps)
- [Using `make` and `append` with Slices](#using-make-and-append-with-slices)
- [Deleting Elements from Maps](#deleting-elements-from-maps)
- [How to Run These Examples](#how-to-run-these-examples)
- [Resources](#resources)

---

## Arrays

- **Definition:**  
  An array in Go is a fixed-size, ordered collection of elements of the same type. The size of an array is part of its type, and once defined, it cannot be changed.

- **Why use arrays in Go?**  
  Arrays are used when you need a collection of elements with a known, fixed size. They provide fast, index-based access and are useful for low-level programming, memory management, or when working with data that does not change in size.

- **File:** [`array.go`](array.go)
- **Description:**  
  Demonstrates how to declare, initialize, and use arrays in Go. Arrays are fixed-size, ordered collections of elements of the same type.

---

## Slices

- **Definition:**  
  A slice is a dynamically-sized, flexible view into the elements of an array. Slices are much more common than arrays in Go and provide powerful ways to work with sequences of data.

- **Why use slices in Go?**  
  Slices are used for collections where the size can change at runtime. They support efficient resizing, slicing, and appending, making them ideal for most list-like data structures in Go.

- **File:** [`slice.go`](slice.go)
- **Description:**  
  Shows how to create and manipulate slices, which are dynamically-sized, flexible views into arrays.  
  Includes:
  - Creating slices
  - Slicing operations
  - Iterating with `range`
  - Printing elements and indices

---

## Maps

- **Definition:**  
  A map is an unordered collection of key-value pairs. Keys must be of a comparable type, and values can be of any type.

- **Why use maps in Go?**  
  Maps are used for fast lookups, associations, and dictionaries. They are ideal when you need to associate unique keys with values and retrieve them efficiently.

- **File:** [`map.go`](map.go)
- **Description:**  
  Explains how to use Go's built-in map type for key-value storage.  
  Includes:
  - Creating and initializing maps
  - Iterating over keys and values
  - Counting elements

---

## Using `make` and `append` with Slices

- **Definition:**  
  The `make` function is used to create slices, maps, and channels with a specified length and capacity. The `append` function adds elements to a slice, growing its length as needed.

- **Why use `make` and `append`?**  
  `make` allows you to efficiently allocate slices with a known capacity, reducing memory reallocations. `append` enables dynamic growth of slices, making them flexible for handling variable-size data.

- **File:** [`make-slice-function.go`](make-slice-function.go)
- **Description:**  
  Demonstrates the use of the `make` function to create slices with specified length and capacity, and the `append` function to add elements dynamically.  
  Key points:
  - `make([]T, length, capacity)` creates a slice with zeroed elements.
  - `append(slice, elems...)` adds elements, growing the slice as needed.
  - Capacity increases automatically when exceeded.

---

## Deleting Elements from Maps

- **Definition:**  
  The `delete` function removes a key and its value from a map. If the key does not exist, `delete` does nothing.

- **Why use `delete` with maps?**  
  Removing elements from a map is essential for managing memory and ensuring that only relevant data is stored. It allows you to update your data structures efficiently.

- **File:** [`map-element-delete.go`](map-element-delete.go)
- **Description:**  
  Shows how to remove elements from a map using the `delete` function.  
  Includes:
  - Deleting existing and non-existing keys
  - Safe access of missing keys (returns zero value)
  - Iterating over keys and values

---

## How to Run These Examples

1. **Navigate to this directory:**
    ```bash
    cd data-structure-types
    ```
2. **Run any example:**
    ```bash
    go run <filename>.go
    ```
    Example:
    ```bash
    go run slice.go
    ```

---

## Resources

- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
- [Effective Go: Arrays, Slices, and Maps](https://go.dev/doc/effective_go#arrays)
- [Go Tour: Slices](https://go.dev/tour/moretypes/7)
- [Go Tour: Maps](https://go.dev/tour/moretypes/19)

---

Happy coding with Go! 🚀