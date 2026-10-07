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
```
bin/
backend/bin/
```
## build production ready server

```
env -> CGO_ENABLED=0 GOOS=linux GOARCH=amd64
command -> go build
parameter -> -trimpath -ldglags"-s -w" -o bin/api ./cmd/api

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldglags"-s -w" -o bin/api ./cmd/api

// production on linux mechine
go build -tags netgo -ldflags "-s -w" -b bin/api ./cmd/api
```
`CGO_ENABLED=0`

এটি নির্দেশ করে যে CGO (C-Go ইন্টারফেস) ডিজেবল থাকবে। অর্থাৎ, Go কোড থেকে C লাইব্রেরি ব্যবহার করা হবে না। এর ফলে তৈরি হওয়া বাইনারি ফাইলটি সম্পূর্ণ স্ট্যাটিক (static) হয় এবং এটি অন্য Linux সিস্টেমে চালাতে কোনো অতিরিক্ত C লাইব্রেরি প্রয়োজন হয় না।

`GOOS=linux`

এটি নির্দেশ করে যে বিল্ডটি Linux অপারেটিং সিস্টেমের জন্য তৈরি করা হবে।

`GOARCH=amd64`

এটি নির্দেশ করে যে বিল্ডটি amd64 আর্কিটেকচার (অর্থাৎ 64-বিট প্রসেসর, যেমন Intel বা AMD) এর জন্য তৈরি করা হবে।

`go build`

এটি Go কম্পাইলারের মূল কমান্ড, যা সোর্স কোড থেকে এক্সিকিউটেবল ফাইল তৈরি করে।

`-trimpath`

এই ফ্ল্যাগটি কম্পাইল করার সময় ফাইল পথের (file path) তথ্য বাইনারি ফাইল থেকে সরিয়ে ফেলে। এর ফলে বাইনারি ফাইলটি ছোট হয় এবং আপনার কম্পিউটারের বা কোড রেপোজিটরিয়ের পথের গোপনীয়তা বজায় থাকে (নিরাপত্তার জন্য ভালো)।

`-ldflags"-s -w"`

এটি লিঙ্কারের (linker) কাছে কিছু অপশন পাঠায়:
- -s: ডিবাগ তথ্য (debug information) বাদ দেয়, যা ফাইলের সাইজ কমায়।
- -w: DWARF সিম্বল টেবিল বাদ দেয়, যা আরও সাইজ কমায়।
- দ্রষ্টব্য: আপনার লেখা -ldglags তে একটা ছোট টাইপো আছে, সঠিকটি হলো -ldflags।

`-o bin/api`

এটি নির্দেশ করে যে তৈরি হওয়া এক্সিকিউটেবল ফাইলটির নাম হবে api এবং এটি bin নামক ফোল্ডারে সেভ হবে।

`./cmd/api`

এটি নির্দেশ করে যে Go কোডের cmd/api ডিরেক্টরি থেকে প্রধান প্যাকেজটি (main package) কম্পাইল করতে হবে।
### সংক্ষেপে:
এই কমান্ডটি আপনার Go প্রজেক্টের cmd/api ফোল্ডার থেকে কোড নিয়ে, একটি Linux amd64 সিস্টেমের জন্য <b>একটি ছোট ও নিরাপদ এক্সিকিউটেবল ফাইল</b> (bin/api) তৈরি করে, যেখানে কোনো অতিরিক্ত C লাইব্রেরি বা ডিবাগ তথ্য থাকবে না। এটি সাধারণত প্রোডাকশন বা ডককন্টেইনারে অ্যাপ্লিকেশন চালানোর সময় ব্যবহৃত হয়।





