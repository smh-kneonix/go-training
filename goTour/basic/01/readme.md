# playground

```go
package main
func main(){}
```

to run go use this command

```bash
go run goFileName.go
```

# import

group import

```go
import (
	"fmt"
	"time"
	// rand is folder inside the math folder
	"math/rand"
)
```

single import

```go
import "fmt"
```

# time

```go
import "time"
//give you the current time
fmt.Println(time.Now())
```

# print

```go
import "fmt"
fmt.Println("Hello, World!")
fmt.printf("by g you can write data in next argument as string %g", 2)
```

# math

```go
import "math"
import "math/rand"
fmt.Println(rand.Intn(100)) // give you random number between 1 to 100
fmt.Println(math.Sqrt(8))
```

# export

start with capital letter (var or func)
