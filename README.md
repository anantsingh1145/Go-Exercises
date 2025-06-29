# Go-Exercises
Go-Exercises

# Go (Golang) Setup and Initialization Guide

This guide will walk you through the steps to install Go and set up your first Go project.

---

## 🛠️ Step 1: Install Go

### 🔗 Download Go

Visit the official Go downloads page:

👉 https://go.dev/dl/

Download the installer for your OS and follow the instructions.
For Mac, Linux and windows download the respective version.

After the executable file get downloaded, Just install it.

### ✅ Verify Installation

After installation, verify that Go is installed correctly:

Open your terminal make yourself as a root user and go to directory

```bash
cd /usr/local/go  #this is for Mac

#then run the below command
go version
```

You will be able to see the output like:

```bash
go version go1.24.4 darwin/arm64
```

## 📁 Step 2: Initialize a Go Project
### 1. Create a new directory:

``` bash
mkdir Go-Exercises
cd Go-Exercises
```

### 2. Initialize a Go module (this creates go.mod):

```bash
go mod init github.com/yourusername/Go-Exercises
```

## 👨‍💻 Step 3: Create Your First Go File
### Create a hello.go file:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### Run the program:

```bash
go run hello.go
```

## 🧾 Go Commands Cheat Sheet

### ⚙️ Common Go Commands

| Command                  | Description                                      |
|--------------------------|--------------------------------------------------|
| `go run main.go`         | Runs the Go program without building a binary.  |
| `go build`               | Compiles your Go code into a binary executable. |
| `go run`                 | Like `go run main.go` but can take multiple files. |
| `go env GOARCH GOOS`     | Displays Go’s target architecture and OS.       |
| `GOOS=linux go build`    | Cross-compiles your program for Linux.          |
| `go help`                | Displays help for Go tooling (e.g., `go help build`). |
| `go env`                 | Displays all the go environments. |

---

### 🔧 Bonus Tools

| Command                | What it does                                |
|------------------------|----------------------------------------------|
| `go fmt main.go`       | Auto-formats your Go code.                  |
| `go mod tidy`          | Removes unused dependencies.                |
| `go get <package>`     | Installs a Go package.                      |

---

### 📚 Resources

- 🔗 [Official Go Docs](https://golang.org/doc/)
- 🧠 [Effective Go](https://golang.org/doc/effective_go.html)
- 📘 [A Tour of Go](https://tour.golang.org/)

---

# Go-lang Basics: Sequential & Control Flow

Welcome to this repository! This project serves as a practical exploration and explanation of fundamental programming concepts in Go (Golang), specifically focusing on **sequential execution** and various **control flow structures**.

If you're just starting with Go, understanding how your code executes step-by-step and how you can direct its path based on conditions or repetitions is crucial. This README, along with the code examples in this repository, aims to solidify these core concepts.

## Table of Contents

1.  [Introduction to Go Basics](#introduction-to-go-basics)
2.  [Sequential Execution](#sequential-execution)
3.  [Control Flow Structures](#control-flow-structures)
    * [Conditional Statements (`if`, `else if`, `else`)](#conditional-statements-if-else-if-else)
    * [Switch Statements (`switch`)](#switch-statements-switch)
    * [Looping Constructs (`for`)](#looping-constructs-for)
    * [`break` and `continue`](#break-and-continue)
    * [`goto` (Optional)](#goto-optional)
4.  [Functions Revisited](#functions-revisited)
5.  [Getting Started (How to Run Examples)](#getting-started-how-to-run-examples)
6.  [Official Go Documentation](#official-go-documentation)

---

## 1. Introduction to Go Basics

Go is a statically typed, compiled programming language designed at Google. It's known for its simplicity, efficiency, and strong support for concurrent programming. Before diving into control flow, it's good to have a grasp of:

* **Variables:** How to declare and assign values.
* **Data Types:** Understanding `int`, `string`, `bool`, etc.
* **Functions:** How to define and call blocks of reusable code.

## 2. Sequential Execution

In its most basic form, Go code executes **sequentially**. This means that statements are processed one after another, from top to bottom, exactly in the order they appear in your program.

Consider this simple example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Step 1: Initialize program") // Executed first
    var x int = 10                         // Executed second
    fmt.Println("Step 2: Value of x is", x) // Executed third
}
````

Unless a control flow statement or a function call changes the execution path, Go will always follow this linear progression.

## 3\. Control Flow Structures

Control flow statements allow you to alter the default sequential execution of your program. They enable you to make decisions, repeat actions, and jump to different parts of your code.

### Conditional Statements (`if`, `else if`, `else`)

Conditional statements execute different blocks of code based on whether a specified condition evaluates to `true` or `false`.

  * **Official Doc:** [Go Tour: If and Else](https://go.dev/tour/flowcontrol/3)

**Example:**

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }

    temperature := 25
    if temperature > 30 {
        fmt.Println("It's very hot!")
    } else if temperature > 20 {
        fmt.Println("It's pleasant.")
    } else {
        fmt.Println("It's a bit chilly.")
    }

    // Short statement (variable declaration within if)
    if num := 9; num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}
```

### Switch Statements (`switch`)

The `switch` statement provides a concise way to handle multiple conditional branches. It's more readable than a long `if-else if` chain when dealing with many conditions based on a single expression.

  * **Official Doc:** [Go Tour: Switch](https://go.dev/tour/flowcontrol/8)

**Example:**

```go
package main

import "fmt"
import "time" // Required for time.Now()

func main() {
    dayOfWeek := time.Now().Weekday()

    switch dayOfWeek {
    case time.Saturday, time.Sunday: // Multiple values in a case
        fmt.Println("It's the weekend!")
    case time.Monday:
        fmt.Println("It's Monday, time to work.")
    default: // Default case if no other matches
        fmt.Println("It's a weekday.")
    }

    // Switch without an expression (acts like if-else if)
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }
}
```

### Looping Constructs (`for`)

Unlike many other languages, Go only has one looping construct: `for`. However, it can be used in several ways, mimicking `while` loops and infinite loops.

  * **Official Doc:** [Go Tour: For loop](https://go.dev/tour/flowcontrol/1)

**Example:**

```go
package main

import "fmt"

func main() {
    // 1. Classic for loop (with initial statement, condition, and post-statement)
    for i := 0; i < 5; i++ {
        fmt.Println("Classic loop iteration:", i)
    }

    // 2. For loop as a 'while' loop (only a condition)
    sum := 1
    for sum < 100 {
        sum += sum
    }
    fmt.Println("Sum (while-like loop):", sum)

    // 3. Infinite loop (no condition - use with break!)
    count := 0
    for {
        fmt.Println("Infinite loop iteration:", count)
        count++
        if count >= 3 {
            break // Exit the loop
        }
    }

    // 4. For-range loop (iterating over collections like arrays, slices, maps, strings)
    numbers := []int{10, 20, 30, 40}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Iterating over a string (gives unicode code points)
    for i, r := range "Hello, 世界" {
        fmt.Printf("Char at index %d is %q (rune: %d)\n", i, r, r)
    }
}
```

### `break` and `continue`

These statements are used within loops to control their execution flow:

  * **`break`**: Terminates the innermost `for`, `switch`, or `select` statement. Execution then continues at the statement immediately following the terminated construct.

  * **`continue`**: Skips the rest of the current iteration of the innermost `for` loop and proceeds to the next iteration.

  * **Official Doc:** [Go Language Specification: For statements](https://www.google.com/search?q=https://go.dev/ref/spec%23For_statements) (See "Break statements" and "Continue statements" sections)

**Example:**

```go
package main

import "fmt"

func main() {
    fmt.Println("--- Using continue ---")
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue // Skip the rest of this iteration when i is 2
        }
        fmt.Println("Current i (continue example):", i)
    }

    fmt.Println("\n--- Using break ---")
    for i := 0; i < 5; i++ {
        if i == 3 {
            break // Exit the loop entirely when i is 3
        }
        fmt.Println("Current i (break example):", i)
    }
}
```

### `goto` (Optional)

The `goto` statement allows you to jump unconditionally to a labeled statement within the same function. While Go supports `goto`, its use is generally discouraged in modern programming practice as it can lead to code that is difficult to read and maintain ("spaghetti code"). It's included here for completeness, but you will rarely, if ever, need to use it.

  * **Official Doc:** [Go Language Specification: Goto statements](https://www.google.com/search?q=https://go.dev/ref/spec%23Goto_statements)

**Example (Discouraged Practice):**

```go
package main

import "fmt"

func main() {
    i := 0
Start: // This is a label
    fmt.Println("Current i (goto example):", i)
    i++
    if i < 3 {
        goto Start // Jump back to the 'Start' label
    }
    fmt.Println("End of goto example.")
}
```

## 4\. Functions Revisited

Functions in Go are fundamental for organizing code. They accept zero or more arguments and return zero or more values. They are crucial for implementing sequential steps and encapsulating logic that can be reused within various control flows.

  * **Official Doc:** [Go Tour: Functions](https://go.dev/tour/basics/4)

**Example:**

```go
package main

import "fmt"

// A function that returns a single value
func greet(name string) string {
    return "Hello, " + name + "!"
}

// A function that returns multiple values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    message := greet("Go Enthusiast")
    fmt.Println(message)

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division result:", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division result:", result)
    }
}
```

## 5\. Getting Started (How to Run Examples)

To run the examples in this repository:

1.  **Install Go:** If you don't have Go installed, follow the instructions on the [official Go website](https://go.dev/doc/install).
2.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git](https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git)
    cd YOUR_REPO_NAME
    ```
3.  **Navigate to an example file:** Each concept usually has its own `.go` file (e.g., `sequential_flow.go`, `conditionals.go`).
4.  **Run the file:**
    ```bash
    go run <filename>.go
    ```
    For example:
    ```bash
    go run conditionals.go
    ```
5.  **Build (optional):** To compile an executable:
    ```bash
    go build <filename>.go
    ./<filename> # On Linux/macOS
    <filename>.exe # On Windows
    ```

## 6\. Official Go Documentation

For the most accurate and in-depth information, always refer to the official Go documentation:

  * **The Go Tour:** [https://go.dev/tour/](https://go.dev/tour/) (Interactive way to learn Go basics)
  * **Go Language Specification:** [https://go.dev/ref/spec](https://go.dev/ref/spec) (Detailed language reference)
  * **Effective Go:** [https://go.dev/doc/effective\_go](https://go.dev/doc/effective_go) (Tips for writing clear, idiomatic Go code)
  * **Go Blog:** [https://go.dev/blog/](https://go.dev/blog/) (Articles and updates from the Go team)

-----

# Happy coding with Go! 💻💙
