## array

```go
func main() {
	prices := [4]float64{99.00, 9.00, 39.00, 19.00}
	fmt.Println(prices)
}
```

[4]float64{99.00, 9.00, 39.00, 19.00}:
[4]: we specific our length of array and if we dont assinge value go put null value in there
float64: is type of value in array

```go
func main() {
	prices := [4]float64{99.00}

    prices[1] = 52.99

	fmt.Println(prices[0]) // 99
	fmt.Println(prices) // {99 52.99  }

}
```

### slice

```go
func main() {
	prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
	fmt.Println(prices)

	featurePrice := prices[1:4]
	fmt.Println(featurePrice) // [36.99 41.99 24.99]
}
```

```go
func main() {
	prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
	fmt.Println(prices)

	featurePrice := prices[:4]
	fmt.Println(featurePrice) // [99 36.99 41.99 24.99]
}
```

```go
func main() {
	prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
	fmt.Println(prices)

	featurePrice := prices[1:]
	fmt.Println(featurePrice) // [36.99 41.99 24.99 12.99]
}
```

notice: you cant use -1 for get last element, and for end of slice prices[1:5] here is 5 for last element you can put last index +1 but moste of the time we use [number:] for last index

you can use slice on slice

slice is just like a window on a array so if you change a index value in a slice of an array, in original array that index will change.

```go
func main() {
	prices := [5]float64{99, 36.99, 41.99, 24.99, 12.99}
	fmt.Println(prices) // [99 36.99 41.99 24.99 12.99]

	featurePrice := prices[1:]
	featurePrice[0]=192.99

    fmt.Println(prices) // [99 192.99 41.99 24.99 12.99]
}
```

### len and cap

len and cap are built-in functions in Go that are used with arrays, slices, and some other data structures.

1. len (length):

-   For arrays: Returns the total number of elements in the array.
-   For slices: Returns the number of elements currently in the slice.

2. cap (capacity):

-   For arrays: Returns the same value as len (the total number of elements).
-   For slices: Returns the maximum number of elements the slice can hold without reallocation.

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}

    fmt.Printf("Array: %v\n", arr) // [1 2 3 4 5]
    fmt.Printf("Length: %d\n", len(arr)) // 5
    fmt.Printf("Capacity: %d\n", cap(arr)) // 5
}

```

For arrays, len and cap always return the same value, which is the total number of elements in the array.

slice:

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    slice1 := arr[1:]
    slice2 := slice1[:1]
    fmt.Println(len(slice1), cap(slice1)) // 4 4
	fmt.Println(len(slice2), cap(slice2)) // 1 4
}
```

-   Length of slice1 is 4 because it contains 4 elements [2, 3, 4, 5]
-   Capacity of slice1 is 4 because there are 4 elements from its starting point (index 1 of the original array) to the end of the underlying array

-   Length of slice2 is 1 because it contains only one element [2]
-   Capacity of slice2 is 4 because there are still 4 elements from its starting point to the end of the underlying array

and why its not 5 when thay all from arr ? cuze your select is left selection(start point), [1:] you cant access before index1 in cap, but you can access after index1 in [:1]

1. Slices are views into an underlying array. They don't copy the data.
2. The length of a slice is the number of elements it contains.
3. The capacity of a slice is the number of elements in the underlying array, counting from the first element of the slice to the end of the array.
4. When you create a new slice from an existing slice, the new slice still refers to the same underlying array.
5. Slicing doesn't create a new array, it creates a new slice value that points to the original array.

why that it is?

```go
arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:]
	slice2 := slice1[:1]
	fmt.Println(len(slice1), cap(slice1)) // 4 4
	fmt.Println(len(slice2), cap(slice2)) // 1 4

	slice2 = slice2[:4]
	fmt.Println(slice2) // [2 3 4 5]
	fmt.Println(len(slice2), cap(slice2)) // 4 4
```

### dynamic array

in go for creating an array we should pass length of array like:

```go
arr := [5]int{1, 2, 3, 4, 5}
```

but with slice feature we can make an array dynamicly:

```go
arr := []int{1, 2}
```

yeah i know maby you wandring its like create an array and this dont have any bussines about slice, but belevie me this syntax is like an slice from an original concret array that length is dynamic you know

```go
func main() {
	prices := []float64{36.99, 41.99, 1.99}

	fmt.Println(prices[1])
	prices[2] = 25.99 // ok!
	prices[3] = 44.99 // panic: runtime error: index out of range [3] with length 3
}
```

but how this dynamic and how to add more items in it ?

```go
func main() {
	prices := []float64{36.99, 41.99, 1.99}

	fmt.Println(prices[1]) // 41.99
	prices[2] = 25.99
	prices = append(prices, 78.99) // [36.99 41.99 25.99 78.99]
	fmt.Println(prices)
}
```

for remove frist item in array you can use **[1:]**

notice: you can add many items you want to a slice with append func

```go
prices = append(prices, 78.99,45.99,22.99)
```

apend slice items to a other slice:

```go
prices := []float64{36.99, 41.99, 1.99}

	discountPrices := []float64{27.99, 17.99, 38.99}

	prices = append(prices, discountPrices...)
	fmt.Println(prices)
```

this 3 dots is special and its get out items from a slice out

## map

```go
websites := map[string]string{
		"google":             "https://google.com",
		"amazon web service": "https://aws.com",
	}
	fmt.Println(websites) // map[amazon web service:https://aws.com google:https://google.com]
    fmt.Println(websites["amazon web service"]) // https://aws.com
    websites["linkedIn"] = "https://linkedIn.com"
	fmt.Println(websites) // map[amazon web service:https://aws.com google:https://google.com linkedIn:https://linkedIn.com]
    delete(websites, "google")
	fmt.Println(websites) // map[amazon web service:https://aws.com linkedIn:https://linkedIn.com]
```
