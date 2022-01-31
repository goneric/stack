# Stack

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/goneric/stack)
[![Build Status](https://github.com/goneric/stack/actions/workflows/build.yaml/badge.svg)](https://github.com/goneric/stack/actions)
[![codecov](https://codecov.io/gh/goneric/stack/branch/main/graph/badge.svg)](https://codecov.io/gh/goneric/stack)

Lightweight, simple, fast, thread-safe Go stack using generics.

## Features
- ✅ Simple, lightweight without any external dependencies
- ✅ Type-safe with Go generics (required Go 1.18)
- ✅ Use with any Go data types, `bool`, `int`, `string`, structs or pointers
- ✅ Thread-safe, protected by `sync.RWMutex`
- ✅ Have unsafe version (without locking)

## Usage
### Installation
```shell
go get github.com/gonerics/stack
```

### Examples
#### Primitive types
You could use Stack with any primitive data types: `bool`, `int`, `string`, for examples:
```go
package main

import (
    "fmt"
    "github.com/gonerics/stack"
)

func main() struct {
    intStack := stack.New[int]()
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    i := intStack.Pop()
    fmt.Printf("Int value = %d\n", i)

    strStack := stack.New[string]()
    strStack.Push("Hello")
    strStack.Push("世界")
    str := strStack.Pop()
    fmt.Printf("String value = %s\n", str)
}
```

```shell
$ go run main.go
Int value = 3
Str value = 世界
```


#### Structs and pointers
```go
package main 

import (
    "fmt"
    "github.com/gonerics/stack"
)

type Task struct {
    Name string 
    Amount int
}


func main() {
    s := stack.New[Task]()
    s.Push(Task{Name: "First task", Amount: 100})
    task, ok := s.Pop()
    if ok {
        fmt.Printf("Pop: %#v", task)
    }

    s2 := stack.New[&Task]()
    s2.Push(&Task{Name: "First task", Amount: 100})
    s2.Push(&Task{Name: "Second task", Amount: 200})
    task, ok := s.Pop()
    if ok {
        fmt.Printf("Pop: %#v", task)
    }
}
```