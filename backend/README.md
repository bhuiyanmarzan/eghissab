## make a folder with project name
```terminal
mkdir eghissab
```
## Open the folder
```terminal
cd eghissab
```

## Initialize golan project
```terminal
go mod init github.com/bhuiyanmarzan/eghissab
```
## Open this project in VS code
```terminal
code .
```

## create a file in root
```go
// ./main.go
package main

import "fmt"

func main() {
	fmt.Println("EGHissab server is running...")
}
```

## Run this project in development mode
```terminal
go run main.go // build + execute
```

## Build go project and run it
```terminal
go build -o bin/main main.go
./bin/main
```
## Now install Makefile in system
`./Makefile`
```Makefile
.PHONY: build run
build:
	@go build -o bin/main main.go
run: build
	@./bin/main
```

## Run make file
```terminal
make run
```

## Monorepo folder structure
```terminal
cmd/api/main.go
cmd/ws/maing.go
cmd/migration/main.go
```
## Fix the Makefile
`./Makefile`
```Makefile
.PHONY: build run
build:
	@go build -o bin/api ./cmd/api/
run: build
	@./bin/api
```
## Now git initialization
```terminal
git init
```