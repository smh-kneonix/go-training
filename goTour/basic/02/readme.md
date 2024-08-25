# create function

for creating function you can

```go
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 3))
}
```

### share type in function

if you have multi arg with same type you can do this

```go
func add(x, y float64) float64 {
	return x + y
}
```

### return multiple result

```go
func add(x, y int, a,b float32) (int,float32) {
	return x + y, b + a
}

func main() {
	a, b := add(1,2,2.5, 3.2)
	// var a, b = add(1,2,2.5, 3.2) //or use this
	fmt.Println(a,b)
}
```

# named return value

only use in short function

```go
func add(x, y int, a,b float32) (l int,n float32) {
	l = x+y
	n = a+b
	return
}
```

# var

var can make list of variable like

and like an argument you can specify type as last

```go
var hi,bye bool
```

### with initializers

```go
func main() {
	var a,b int = 1,2
	// var a,b = 1,2 //or you can use this
	fmt.Println(a,b)
}
```

## short assignment
its only can use inside function
> its no need to type by this short assignment

```go
k:=3
```
## group assign
```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```
# basic type

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

>HINT: `uint` only can get positive number

>HINT: `uintptr` use fore address of data

>HINT: `rune` its for using for unicode character

>HINT: `complex` ذخیره کردن اعداد مختلط


# zero values
if you have var and not initialize we treated as zero values

0 for numeric types,

false for the boolean type, and

"" (the empty string) for strings.