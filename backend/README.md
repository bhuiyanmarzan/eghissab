# Eghissab Backend

A clean and simple Go backend project for the Eghissab application. This README explains how to set up, run, and manage the backend in a structured and professional way.

## Project Overview

This backend is built with Go and follows a basic CLI-style application structure. It is designed to be easy to extend and can later include:

- REST API endpoints
- WebSocket support
- database migrations
- service and repository layers
- authentication and authorization
- environment-based configuration

## Requirements

Before starting, make sure you have the following installed:

- Go 1.21 or newer
- Git
- VS Code or another code editor
- Make (optional, but recommended)

## Monorepo Structure

```text
project/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── Makefile
│   ├── go.mod
│   └── README.md
├── frontend/
└── .git/
```

## 1. Create the Project Folder

```bash
mkdir eghissab
cd eghissab
```

## 2. Initialize the Go Module

```bash
go mod init github.com/bhuiyanmarzan/eghissab
```

If you are working inside the backend folder, the command should be run from that directory.

## 3. Open the Project in VS Code

```bash
code .
```

## 4. Create the Main Application File

```go
// ./cmd/api/main.go
package main

import "fmt"

func main() {
    fmt.Println("EGHissab server is running...")
}
```

## 5. Run the Application

### Development mode

```bash
go run ./cmd/api
```

This command compiles and executes the application directly from the source code.

### Build the binary

```bash
go build -o bin/api ./cmd/api
./bin/api
```

This generates an executable file in the `bin` directory and runs it.

## 6. Use the Makefile

The project includes a `Makefile` to simplify common tasks.

```makefile
.PHONY: build run

build:
	@go build -o bin/api ./cmd/api/

run: build
	@./bin/api
```

### Commands

```bash
make build
make run
```

`make run` automatically builds the application before launching it.

## 7. Git Initialization

```bash
git init
```

You can then add a remote repository:

```bash
git branch -M main
git remote add origin <your-repository-url>
git push origin -U main
```

## 8. Recommended Project Workflow

1. Create or update backend feature folders.
2. Write Go logic in packages or services.
3. Run the server using `make run`.
4. Build for deployment using `make build`.
5. Commit stable changes to git.

## 9. Suggested Future Structure

As the project grows, this is a good structure to follow:

```text
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   └── utils/
├── pkg/
│   └── logger/
├── Makefile
├── go.mod
└── README.md
```

## 10. add igtignore file







