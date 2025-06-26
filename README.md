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

# Happy coding with Go! 💻💙
