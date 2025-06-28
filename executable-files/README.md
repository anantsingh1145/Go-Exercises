# 🚀 Building Executables with Go (Golang)

This guide helps you compile your Go source code into executable binaries for different platforms.

---

## ✅ Prerequisites

- [Install Go](https://go.dev/dl/)
- Set up your environment (`GOPATH`, `GOROOT`, and `PATH`)
- Confirm installation:

```bash
go version
````

---

## 🛠️ Step-by-Step: Create Executable from Go Code

### 1. Write Your Go Program

Example: `hello.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go Executable!")
}
```

---

### 2. Build the Executable (for Current OS)

```bash
go build hello.go
```

This will generate an executable:

* `hello` on Linux/macOS
* `hello.exe` on Windows

---

### 3. Create a Executable file in other directory

Make sure you are in the directory where you have written the GO code

```go
go build -o ../executable-files/hello ./hello.go
```

### 4. Run the Executable

```bash
./hello        # Linux/macOS
hello.exe      # Windows
```

---

## 🌍 Build for Other Platforms (Cross Compilation)

### Linux (from macOS or Windows):

```bash
GOOS=linux GOARCH=amd64 go build -o hello-linux hello.go
```

### Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o hello.exe hello.go
```

### macOS (from Linux/Windows):

```bash
GOOS=darwin GOARCH=amd64 go build -o hello-mac hello.go
```

---

## 🔍 Useful Go Commands Cheat Sheet

| Command               | Description                                     |
| --------------------- | ----------------------------------------------- |
| `go run hello.go`      | Run Go program directly without building binary |
| `go build`            | Compile code into executable                    |
| `go env GOARCH GOOS`  | Show system architecture and OS info            |
| `GOOS=linux go build` | Cross-compile for Linux                         |
| `go fmt`              | Format Go source code                           |
| `go mod tidy`         | Clean up unused modules                         |
| `go help <command>`   | Get help on any Go command                      |

---

## 📚 Resources

* [Official Go Documentation](https://golang.org/doc/)
* [Tour of Go](https://tour.golang.org/)
* [Effective Go](https://go.dev/doc/effective_go)

---

**Happy Coding! 💻🚀**
