# type conversion
to change type form one to another you can easily use `type(value)`

```go
	var i int = 42
	fmt.Println(float32(i))
	fmt.Println(float64(i))
	fmt.Println(uint(i))
	fmt.Println(string(i))
```

you should always set type in using methods like this
```go
	fmt.Println(math.Sqrt(float64(i)))
```

# type interface

if you create var with type and assign it to another var its get same type even its empty

```go
    var i int
    j := i // j is an int
```

# constant
like other language you have constant var which you can NOT change them
```go
	// const
	const hi = "hi"
	// hi = "bye" //error
```

# numeric constant

by this feature you can make big value
```go
	// numeric constant
	const big = 1 << 100 // make 1 with 100 zero after that in binary
	const small = big >> 99 // delete 99 number from right in binary
	// fmt.Println(small) // you can not log this because its too much big
	fmt.Println(small) // 2
```

# for
you only have one loop which is `for`


```go
	for i := 0; i < 10; i++ {
		sum += i
	}
```
>HINT: if you make variable inside `for` its only available inside the for block
>HINT: you can use `for` without any condition like `for ; sum < 100; {`

```go
	// now its like while in other language	
	for sum < 100 {
		sum += sum
	}
```

### infinite loop

```go
	for {
		sum += sum
	}
```

# if

```go
	num := 2	
	if num==1 {
		fmt.Println(1)
	} else if num > 2 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
```


you can also define variable especially for if statement

and also you can access this var in else
```go
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
```
