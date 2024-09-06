## interface

Interfaces in Go are a powerful feature that allows for abstraction and polymorphism. They define a set of method signatures that a type must implement to satisfy the interface

An interface is defined using the interface keyword:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

```

In Go, interfaces are implemented implicitly. There's no implements keyword:

```go
type File struct {
    // ...
}

func (f *File) Read(p []byte) (n int, err error) {
    // Implementation
    return len(p), nil
}
```

notice: Now \*File satisfies the Reader interface.

An interface value consists of two components: a concrete type and a value of that type.

```go
var r Reader
f := &File{}
r = f

```

var r Reader:

This declares a variable r of type Reader.
Reader is an interface type.
At this point, r is a nil interface value, meaning it doesn't hold any concrete type or value.

f := &File{}:

This creates a new File struct and assigns its address to f.
f is now a pointer to a File (\*File).

r = f:

This assigns f (which is a *File) to the interface variable r.
This assignment is valid if and only if *File implements all the methods declared in the Reader interface.
After this assignment, r is no longer nil. It now holds two pieces of information: a. The concrete type: \*File b. The value: the address of the File struct

You can extract the underlying value from an interface:

```go
f, ok := r.(*File)
if ok {
    // r was a *File
}

```

Useful for handling multiple types:

```go
switch v := i.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}

```

Interfaces can embed other interfaces:

```go
type ReadWriter interface {
    Reader
    Writer
}

```

Methods with pointer receivers implement interfaces for pointer types only:

```go
type Incrementer interface {
    Increment()
}

type Integer int

func (i *Integer) Increment() {
    *i++
}

// Only *Integer satisfies Incrementer, not Integer

```

An interface value is nil only if both its type and value are nil:

```go
var i interface{}
fmt.Println(i == nil) // true

var p *int
i = p
fmt.Println(i == nil) // false, even though p is nil

```

Single-method interfaces are often named by the method name plus 'er' (e.g., **Reader**, **Writer**).

## any type

```go
func printAnything(value interface{}){
    fmt.Println(value)
}
// or
// any is special key for alise of interface{}
func printAnything(value any){
    fmt.Println(value)
}
```

### handle any type

you can use special way in switch statement to handle enter type

```go

func main(){
    printAnything(2) // Integer: 2
    printAnything(2.6) // Float: 2.6
    printAnything("hello") //hello
    // newNote is a litral struct
    printAnything(newNote) // noting printed cuze we dont specific what to do with this special kind of type we can pass it but nothing happend
}

func printAnything(value interface{}){
    switch value.(type){
        case int:
        fmt.Println("Integer: ",value)
        case float64:
        fmt.Println("Float: ",value)
        case string:
        fmt.Println(value)
    }
}

```

### handle any type without switch

```go
func printAnything(value interface{}){

    intVal ,ok :=value.(int)
    if ok {
        fmt.Println("Integer: ",intVal)
    }
    floatVal ,ok :=value.(float64)
    if ok {
        fmt.Println("Float: ",floatVal)
    }
    stringVal ,ok :=value.(string)
    if ok {
        fmt.Println("String: ",stringVal)
    }
}
```

notice: mabey you thinking about the switch way was better but you need this way some time

## generics

```go
func main(){
    result := add(1,2)
}

func add[T int|float64|string](a,b T)T{
    return a+b
}
```
